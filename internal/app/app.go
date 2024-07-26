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
	log, loggerCallback := zerolog.MustSetUp(cfg)
	pg, storageCallback := storage.MustOpenPostgres(log)
	logFormatter := logger.LogFormatterImpl{Logger: log}

	var (
		repos    = repoProvider.InitRepos(pg.Db, log)
		services = serviceProvider.InitServices(repos)
		apis     = apiProvider.InitApi(services, log)
	)

	api := chi.SetupServer(cfg, log, apis, &logFormatter)
	callback := api.MustRun()

	closer.Add(loggerCallback)
	closer.Add(storageCallback)
	closer.Add(callback)

	printStartMessage(log, cfg)

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.Timeout)
	defer cancel()

	if err := closer.Close(shutdownCtx, log); err != nil {
		log.Fatal(err, "Server close failed")
	}
}

func printStartMessage(log logger.Common, cfg *config.Config) {
	log.Info("Server started")
	log.Infof("Host: %s", cfg.Server.Host)
	log.Infof("Port: %s", cfg.Server.Port)
	log.Infof("ENV: %s", cfg.Env)
}
