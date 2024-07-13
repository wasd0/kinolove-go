package main

import (
	"kinolove/internal/config"
	"kinolove/internal/logger"
	"kinolove/internal/logger/zerolog"
)

func main() {
	cfg := config.MustRead()
	log, callback := zerolog.MustSetUp(cfg)
	defer callback()

	printStartMessage(log, cfg)
}

func printStartMessage(log logger.Common, cfg *config.Config) {
	log.Info("Server started")
	log.Infof("Host: %s", cfg.Server.Host)
	log.Infof("Port: %s", cfg.Server.Port)
	log.Infof("ENV: %s", cfg.Env)
	log.Infof("DB PATH: %s", cfg.DbPath)
}
