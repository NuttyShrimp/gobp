package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/studentkickoff/gobp/internal/api"
	"github.com/studentkickoff/gobp/pkg/config"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	env := config.GetDefaultString("app.env", "development")

	if env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	log.Logger = log.Logger.With().Caller().Str("env", env).Logger()

	server, err := api.NewServer()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	log.Info().Msgf("Server is running on %s", server.Addr)
	if err := server.Listen(server.Addr); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
