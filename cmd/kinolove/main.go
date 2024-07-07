package main

import (
	"github.com/rs/zerolog"
	"kinolove/internal/config"
	"kinolove/internal/logger"
	"os"
)

func main() {
	cfg := config.MustRead()
	log, logFile := logger.MustSetUp(cfg)

	if logFile != nil {
		defer func(logFile *os.File) {
			err := logFile.Close()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to close log file")
			}
		}(logFile)
	}

	printStartMessage(log, cfg)
}

func printStartMessage(log *zerolog.Logger, cfg *config.Config) {
	log.Info().Msg("Server started")
	log.Info().Msgf("Host: %s", cfg.Server.Host)
	log.Info().Msgf("Port: %s", cfg.Server.Port)
	log.Info().Msgf("ENV: %s", cfg.Env)
	log.Info().Msgf("DB PATH: %s", cfg.DbPath)
}
