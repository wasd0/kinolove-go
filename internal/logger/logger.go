package logger

import (
	"github.com/rs/zerolog"
	"io"
	"kinolove/internal/common/utils"
	"kinolove/internal/config"
	"os"
	"time"
)

func MustSetUp(cfg *config.Config) *zerolog.Logger {
	var output io.Writer
	var level zerolog.Level

	switch cfg.Env {
	case config.EnvDev, config.EnvStage:
		level = zerolog.TraceLevel
		output = zerolog.ConsoleWriter{
			Out:          os.Stdout,
			NoColor:      false,
			TimeFormat:   time.RFC850,
			TimeLocation: time.Local,
		}
	case config.EnvProd:
		level = zerolog.ErrorLevel
		output = initOutputFile(cfg)
	}

	zerolog.SetGlobalLevel(level)

	logger := zerolog.New(output).With().Timestamp().Caller().Logger()
	return &logger
}

func initOutputFile(cfg *config.Config) *os.File {
	if len(cfg.LogPath) == 0 {
		panic("No log file path")
	}

	logFileName := time.Now().Format("2006-01-02") + ".log"
	logFilePath := utils.CreateDirectoriesIfNotExists(cfg.LogPath)
	logFilePath.WriteString(utils.Separator)
	logFilePath.WriteString(logFileName)

	file, err := os.OpenFile(
		logFilePath.String(),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)

	if err != nil {
		panic(err)
	}

	return file
}
