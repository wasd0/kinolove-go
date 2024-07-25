package app

import (
	"context"
	"errors"
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
	"kinolove/pkg/utils/app"
	"net/http"
	"os/signal"
	"syscall"
)

func Startup() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT,
		syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	defer stop()
	runServer(ctx)
}

func runServer(ctx context.Context) {
	cfg := config.MustRead()
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log, loggerCallback := zerolog.MustSetUp(cfg)
	pg, storageCallback := storage.MustOpenPostgres(log)
	httpLog := logger.LogFormatterImpl{Logger: log}
	mux := setUpRouter(cfg, &httpLog)
	closer := &app.Closer{}

	closer.Add(loggerCallback)
	closer.Add(storageCallback)

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
	mux.MethodNotAllowed(defaultApi.MethodNotAllowed)

	server := &http.Server{Addr: addr, Handler: mux}
	closer.Add(server.Shutdown)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err, "http server start failed")
		}
	}()

	printStartMessage(log, cfg)

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.Timeout)
	defer cancel()

	if err := closer.Close(shutdownCtx, log); err != nil {
		log.Fatal(err, "Server close failed")
	}
}

func setUpRouter(cfg *config.Config, log *logger.LogFormatterImpl) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestLogger(log))
	router.Use(middleware.Timeout(cfg.Server.IdleTimeout))
	return router
}

func printStartMessage(log logger.Common, cfg *config.Config) {
	log.Info("Server started")
	log.Infof("Host: %s", cfg.Server.Host)
	log.Infof("Port: %s", cfg.Server.Port)
	log.Infof("ENV: %s", cfg.Env)
}
