package main

import (
	"github.com/rs/zerolog"
	"kinolove/internal/config"
	"kinolove/internal/logger"
)

func main() {
	cfg := config.MustRead()
	log := logger.MustSetUp(cfg)

	printStartMessage(log, cfg)
}

func printStartMessage(log *zerolog.Logger, cfg *config.Config) {
	log.Info().Msg("Server started")
	log.Info().Msgf("Host: %s", cfg.Server.Host)
	log.Info().Msgf("Port: %s", cfg.Server.Port)
	log.Info().Msgf("ENV: %s", cfg.Env)
	log.Info().Msgf("DB PATH: %s", cfg.DbPath)
}
