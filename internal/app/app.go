package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"kinolove/api"
	"kinolove/internal/repository"
	"kinolove/internal/service"
	"kinolove/internal/storage"
	"kinolove/pkg/config"
	"kinolove/pkg/logger"
	"kinolove/pkg/logger/zerolog"
	"net/http"
)

func Startup() func() {

	cfg := config.MustRead()
	log, loggerCallback := zerolog.MustSetUp(cfg)
	pg, storageCallback := storage.MustOpenPostgres(log)
	mux := setUpRouter(cfg)

	var (
		userRepo  repository.UserRepository  = repository.NewUserRepository(pg.Db)
		movieRepo repository.MovieRepository = repository.NewMoviesRepository(pg.Db)
	)

	var (
		userService  service.UserService  = service.NewUserService(userRepo)
		movieService service.MovieService = service.NewMovieService(movieRepo)
	)

	var (
		defaultApi            = api.NewDefaultApi(log)
		userApi    api.ChiApi = api.NewUserApi(log, userService)
		movieApi   api.ChiApi = api.NewMovieApi(log, movieService)
	)

	mux.Route(userApi.Register())
	mux.Route(movieApi.Register())
	mux.NotFound(defaultApi.NotFound)

	printStartMessage(log, cfg)

	//todo implement graceful shutdown

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	err := http.ListenAndServe(addr, mux)

	if err != nil {
		log.Fatal(err, "error while starting server")
	}

	return func() {
		log.Info("Shutting down")
		loggerCallback()
		storageCallback()
	}
}

func setUpRouter(cfg *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Timeout(cfg.Server.IdleTimeout))
	return router
}

func printStartMessage(log logger.Common, cfg *config.Config) {
	log.Info("Server started")
	log.Infof("Host: %s", cfg.Server.Host)
	log.Infof("Port: %s", cfg.Server.Port)
	log.Infof("ENV: %s", cfg.Env)
}
