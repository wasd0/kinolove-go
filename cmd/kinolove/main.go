package main

import (
	"kinolove/internal/config"
	"kinolove/internal/logger"
	"log/slog"
)

func main() {
	cfg := config.MustRead()
	log := logger.MustSetUp(cfg.Env)

	printStartMessage(log, cfg)
}

func printStartMessage(log *slog.Logger, cfg *config.Config) {
	log.Info("Server started",
		slog.String("host", cfg.Server.Host),
		slog.String("port", cfg.Server.Port),
		slog.String("env", cfg.Env),
		slog.String("db_path", cfg.DbPath))
}
