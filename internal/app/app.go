package app

import (
	"kinolove/internal/config"
	"kinolove/internal/repository"
	"kinolove/internal/storage"
	"kinolove/pkg/logger"
	"kinolove/pkg/logger/zerolog"
)

func Startup() func() {

	cfg := config.MustRead()
	log, loggerCallback := zerolog.MustSetUp(cfg)
	pg, storageCallback := storage.MustOpenPostgres(log)

	userRepo := repository.NewUserRepository(pg.Db)

	_ = userRepo

	printStartMessage(log, cfg)

	return func() {
		loggerCallback()
		storageCallback()
	}
}

func printStartMessage(log logger.Common, cfg *config.Config) {
	log.Info("Server started")
	log.Infof("Host: %s", cfg.Server.Host)
	log.Infof("Port: %s", cfg.Server.Port)
	log.Infof("ENV: %s", cfg.Env)
}
