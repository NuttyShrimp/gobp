package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/studentkickoff/gobp/internal/api"
)

func main() {
	// TODO: Init viper config settings
	viper.AutomaticEnv()
	viper.SetDefault("app_env", "development")
	env := viper.GetString("app_env")

	if env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Print pwd
	fmt.Println(os.Getenv("PWD"))

	viper.SetConfigName(fmt.Sprintf("%s.toml", env))
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Fatal().Err(err).Msg("Error while reading config file")
	}

	server, err := api.NewServer()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	log.Info().Msgf("Server is running on %s", server.Addr)
	if err := server.Listen(server.Addr); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
