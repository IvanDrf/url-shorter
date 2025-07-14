package logger

import (
	"log"
	"log/slog"
	"os"
	"url-shorter/config"
)

const (
	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
)

func InitLogger(cfg *config.Config) *slog.Logger {
	var level slog.Leveler
	switch cfg.LoggerLevel {
	case levelDebug:
		level = slog.LevelDebug

	case levelInfo:
		level = slog.LevelInfo

	case levelWarn:
		level = slog.LevelWarn

	case levelError:
		level = slog.LevelError

	default:
		log.Fatal("cant set logger level")
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}))
}
