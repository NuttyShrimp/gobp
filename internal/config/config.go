package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func bindEnv(key string) {
	envName := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	viper.BindEnv(key, envName)
}

func Init() error {
	viper.AutomaticEnv()
	env := GetDefaultString("app.env", "development")

	viper.SetConfigName(fmt.Sprintf("%s.toml", env))
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	return viper.ReadInConfig()
}

func GetString(key string) string {
	bindEnv(key)
	return viper.GetString(key)
}

func GetDefaultString(key, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return GetString(key)
}

func GetInt(key string) int {
	bindEnv(key)
	return viper.GetInt(key)
}

func GetDefaultInt(key string, defaultVal int) int {
	viper.SetDefault(key, defaultVal)
	return GetInt(key)
}

func GetUint16(key string) uint16 {
	bindEnv(key)
	return viper.GetUint16(key)
}

func GetDefaultUint16(key string, defaultVal uint16) uint16 {
	viper.SetDefault(key, defaultVal)
	return GetUint16(key)
}
