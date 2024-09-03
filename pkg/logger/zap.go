package logger

import (
	"github.com/TheZeroSlave/zapsentry"
	"github.com/getsentry/sentry-go"
	"github.com/studentkickoff/gobp/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.Logger {
	var logger *zap.Logger
	env := config.GetDefaultString("app.env", "development")
	if env == "development" {
		logger = zap.Must(zap.NewDevelopment(zap.AddStacktrace(zap.WarnLevel)))
	} else {
		logger = zap.Must(zap.NewProduction())
	}
	logger = logger.With(zap.String("env", env))

	if dsn := config.GetString("app.dsn"); dsn != "" {
		cfg := zapsentry.Configuration{
			EnableBreadcrumbs: true,
			BreadcrumbLevel:   zapcore.InfoLevel,
			Level:             zapcore.ErrorLevel,
			Tags: map[string]string{
				"component": config.GetString("app.name"),
			},
		}
		core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromClient(sentry.CurrentHub().Client()))
		if err != nil {
			logger.Fatal("initializing zap core", zap.Error(err))
		}

		logger = zapsentry.AttachCoreToLogger(core, logger)
	}

	return logger
}
