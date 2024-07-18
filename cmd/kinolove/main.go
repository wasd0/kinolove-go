package main

import (
	"kinolove/internal/config"
	"kinolove/internal/logger"
	"kinolove/internal/logger/zerolog"
	"kinolove/internal/storage/postgresql"
)

func main() {
	cfg := config.MustRead()
	log, loggerCallback := zerolog.MustSetUp(cfg)
	defer loggerCallback()

	printStartMessage(log, cfg)

	db, callback := postgresql.MustOpenConnect(log)
	defer callback()

	type res struct {
		Val interface{} `ksql:"res"`
	}

	response := res{}

	//todo remove this
	if err := db.SelectOne(&response, "select 'hello world' as res"); err != nil {
		log.Fatal(err, "Failed to execute query")
	}

	log.Infof("Result: %v", response)
}

func printStartMessage(log logger.Common, cfg *config.Config) {
	log.Info("Server started")
	log.Infof("Host: %s", cfg.Server.Host)
	log.Infof("Port: %s", cfg.Server.Port)
	log.Infof("ENV: %s", cfg.Env)
}
