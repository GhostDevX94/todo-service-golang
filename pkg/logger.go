package pkg

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

func InitLogger() {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "debug" || ginMode == "" {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		Logger = zerolog.New(output).
			Level(level).
			With().
			Timestamp().
			Caller().
			Logger()
	} else {
		Logger = zerolog.New(os.Stdout).
			Level(level).
			With().
			Timestamp().
			Caller().
			Logger()
	}

	log.Logger = Logger
}

func GetLogger() *zerolog.Logger {
	return &Logger
}
