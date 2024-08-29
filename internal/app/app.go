package app

import (
	"context"
	"kinolove/internal/api/chi"
	"kinolove/internal/app/apiProvider"
	"kinolove/internal/app/repoProvider"
	"kinolove/internal/app/serviceProvider"
	"kinolove/internal/storage"
	"kinolove/pkg/config"
	"kinolove/pkg/logger"
	"kinolove/pkg/logger/zerolog"
	"kinolove/pkg/utils/app"
	"kinolove/pkg/utils/jwtUtils"
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
	closer := &app.Closer{}
	_, loggerCallback := zerolog.MustSetUp(cfg)
	pg, storageCallback := storage.MustOpenPostgres()
	logFormatter := logger.LogFormatterImpl{}
	auth := jwtUtils.NewJwtAuth()

	var (
		repos    = repoProvider.InitRepos(pg.Db)
		services = serviceProvider.InitServices(repos, auth)
		apis     = apiProvider.InitApi(services)
	)

	api := chi.SetupServer(cfg, apis, &logFormatter, auth)
	callback := api.MustRun()

	closer.Add(loggerCallback)
	closer.Add(storageCallback)
	closer.Add(callback)

	printStartMessage(cfg)

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.Timeout)
	defer cancel()

	if err := closer.Close(shutdownCtx); err != nil {
		logger.Log().Fatal(err, "Server close failed")
	}
}

func printStartMessage(cfg *config.Config) {
	logger.Log().Info("Server started")
	logger.Log().Infof("Host: %s", cfg.Server.Host)
	logger.Log().Infof("Port: %s", cfg.Server.Port)
	logger.Log().Infof("ENV: %s", cfg.Env)
}
