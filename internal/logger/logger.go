package logger

import (
	config "ModaTrack/internal/config"

	"github.com/phuslu/log"
)

func InitLogger() {
	cfg, err := config.LoadConfigFromFile("../../env.test")

	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config")
	}

	logFile := cfg.LogFile

	log.DefaultLogger = log.Logger{
		Level:      log.InfoLevel,
		Writer:     &log.FileWriter{Filename: logFile},
		Caller:     1,
		TimeField:  "date",
		TimeFormat: "2006-01-02",
	}
}

func GetLogger() *log.Logger {
	return &log.DefaultLogger
}
