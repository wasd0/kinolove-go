package logger

import (
	"kinolove/internal/config"
	"log/slog"
	"os"
)

func MustSetUp(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case config.EnvDev, config.EnvStage:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
