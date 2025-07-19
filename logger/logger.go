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
	source := false

	switch cfg.LoggerLevel {
	case levelDebug:
		level = slog.LevelDebug
		source = true

	case levelInfo:
		level = slog.LevelInfo

	case levelWarn:
		level = slog.LevelWarn

	case levelError:
		level = slog.LevelError

	default:
		log.Fatal()
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: source,
		Level:     level,
	}))
}
