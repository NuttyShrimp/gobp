package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/studentkickoff/gobp/pkg/config"
)

var Client *redis.Client
var Nil = redis.Nil

func New() error {
	cfgUrl := config.GetDefaultString("redis.url", "")
	if cfgUrl != "" {
		options, err := redis.ParseURL(cfgUrl)
		if err != nil {
			return err
		}

		Client = redis.NewClient(options)
	} else {
		Client = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", config.GetDefaultString("redis.host", "localhost"), config.GetDefaultInt("redis.port", 6379)),
		})
	}
	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	return err
}
