package main

import (
	"fmt"

	"github.com/studentkickoff/gobp/internal/api"
	"github.com/studentkickoff/gobp/pkg/config"
	"github.com/studentkickoff/gobp/pkg/logger"
	"github.com/studentkickoff/gobp/pkg/mjml"
	"go.uber.org/zap"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	zapLogger := logger.New()
	zap.ReplaceGlobals(zapLogger)

	go func() {
		if err := mjml.Init(); err != nil {
			panic(fmt.Sprintf("Failed to setup mjml binary: %+v", err))
		}
	}()

	server, err := api.NewServer()
	if err != nil {
		zap.L().Fatal("Failed to create server", zap.Error(err))
	}

	zap.L().Info(fmt.Sprintf("Server is running on %s", server.Addr))
	if err := server.Listen(server.Addr); err != nil {
		zap.L().Fatal("Failure while running the server", zap.Error(err))
	}
}
