package logger

import (
	config "ModaTrack/internal/config"

	"github.com/phuslu/log"
)

func InitLogger() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config")
	}

	logFile := cfg.LogFile

	log.DefaultLogger = log.Logger{
		Level: log.InfoLevel,
		Writer: &log.MultiEntryWriter{
			&log.FileWriter{Filename: logFile,
				FileMode:     0600,
				MaxSize:      100 * 1024 * 1024,
				MaxBackups:   7,
				EnsureFolder: true,
				LocalTime:    true},
			&log.ConsoleWriter{ColorOutput: true},
		},
		Caller:     1,
		TimeField:  "date",
		TimeFormat: "2006-01-02",
	}
}

func GetLogger() *log.Logger {
	InitLogger()
	return &log.DefaultLogger
}
