package app

import (
	"kinolove/internal/config"
	"kinolove/internal/repository"
	"kinolove/internal/service"
	"kinolove/internal/storage"
	"kinolove/pkg/logger"
	"kinolove/pkg/logger/zerolog"
)

func Startup() func() {

	cfg := config.MustRead()
	log, loggerCallback := zerolog.MustSetUp(cfg)
	pg, storageCallback := storage.MustOpenPostgres(log)

	var (
		userRepo     repository.UserRepository  = repository.NewUserRepository(pg.Db)
		userService  service.UserService        = service.NewUserService(userRepo)
		movieRepo    repository.MovieRepository = repository.NewMoviesRepository(pg.Db)
		movieService service.MovieService       = service.NewMovieService(movieRepo)
	)

	printStartMessage(log, cfg)

	_ = userService
	_ = movieService

	//todo implement graceful shutdown

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
