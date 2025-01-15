package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aidenwallis/go-utils/utils"
	"github.com/studentkickoff/gobp/internal/redis"
	"github.com/studentkickoff/gobp/pkg/config"
)

type authPermissions struct {
	User        string              `json:"user"`
	Permissions map[string][]string `json:"permissions"`
}

// TODO: not fully ideal we need to fiber ctx here
func Can(ctx context.Context, userId, permission string) (bool, error) {
	var permissions []string
	redisKey := fmt.Sprintf("%s:auth:permissions", userId)
	permissionStr, err := redis.Client.Get(ctx, redisKey).Result()
	if err == nil {
		permissions = strings.Split(permissionStr, ",")
	} else if err == redis.Nil {
		// Fetch perms from auth
		newPerms, err := getPermsFromAuth(ctx, userId)
		if err != nil {
			return false, err
		}
		_, err = redis.Client.SetEx(ctx, redisKey, strings.Join(newPerms, ","), 3600*time.Second).Result()
		if err != nil {
			return false, err
		}
		return Can(ctx, userId, permission)
	} else {
		return false, err
	}

	_, found := utils.SliceFind(permissions, func(s string) bool { return s == permission })

	return found, nil
}

func getPermsFromAuth(ctx context.Context, userId string) ([]string, error) {
	authBaseUrl := fmt.Sprintf("%s/api/v1/scopes", config.GetDefaultString("auth.sko.url", "http://localhost:3004"))
	authUrl, err := url.Parse(authBaseUrl)
	if err != nil {
		return nil, err
	}
	authUrl.RawQuery = url.Values{
		"user": {userId},
	}.Encode()

	req, err := http.NewRequest("GET", authUrl.String(), http.NoBody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	authSecret := config.GetString("auth.sko.client_secret")
	if authSecret == "" {
		return nil, fmt.Errorf("auth.sko.client_secret is not set")
	}
	req.Header.Add("X-Application", authSecret)
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// nolint:errcheck // ignore error of closing body
	defer resp.Body.Close()

	var body authPermissions
	bodyDecoder := json.NewDecoder(resp.Body)
	err = bodyDecoder.Decode(&body)
	if err != nil {
		return nil, err
	}

	var permissions = make([]string, 0)
	for _, perms := range body.Permissions {
		permissions = append(permissions, perms...)
	}

	return utils.UniqueSlice(permissions), nil
}
