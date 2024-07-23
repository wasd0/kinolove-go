package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"kinolove/internal/config"
	"kinolove/internal/repository"
	"kinolove/internal/service"
	"kinolove/internal/storage"
	"kinolove/pkg/logger"
	"kinolove/pkg/logger/zerolog"
	"net/http"
)

func Startup() func() {

	cfg := config.MustRead()
	log, loggerCallback := zerolog.MustSetUp(cfg)
	pg, storageCallback := storage.MustOpenPostgres(log)
	router := setUpRouter()

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

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	err := http.ListenAndServe(addr, router)

	if err != nil {
		log.Fatal(err, "error while starting server")
	}

	return func() {
		log.Info("Shutting down")
		loggerCallback()
		storageCallback()
	}
}

func setUpRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	return router
}

func printStartMessage(log logger.Common, cfg *config.Config) {
	log.Info("Server started")
	log.Infof("Host: %s", cfg.Server.Host)
	log.Infof("Port: %s", cfg.Server.Port)
	log.Infof("ENV: %s", cfg.Env)
}
