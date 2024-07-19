package main

import (
	"fmt"
	"github.com/go-jet/jet/v2/postgres"
	"kinolove/.gen/kinolove/public/model"
	"kinolove/.gen/kinolove/public/table"
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

	storage, callback := postgresql.MustOpenConnect(log)
	defer callback()

	//todo remove this
	stmt := postgres.SELECT(table.Users.AllColumns).FROM(table.Users)

	var dest []struct {
		model.Users
	}

	err := stmt.Query(storage.Db, &dest)
	if err != nil {
		log.Fatal(err, "Failed to execute query")
	}

	for _, user := range dest {
		if user.IsActive != nil {
			fmt.Println(*user.IsActive)
		}
	}
}

func printStartMessage(log logger.Common, cfg *config.Config) {
	log.Info("Server started")
	log.Infof("Host: %s", cfg.Server.Host)
	log.Infof("Port: %s", cfg.Server.Port)
	log.Infof("ENV: %s", cfg.Env)
}
