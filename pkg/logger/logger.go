package logger

import (
	"log/slog"
	"os"
)

type Level string

const (
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
	DEBUG Level = "DEBUG"
)

var logLevel = map[Level]slog.Level{
	INFO:  slog.LevelInfo,
	WARN:  slog.LevelWarn,
	ERROR: slog.LevelError,
	DEBUG: slog.LevelDebug,
}

func Configure(l Level) {
	level, ok := logLevel[l]

	if !ok {
		level = slog.LevelInfo
	}

	log := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: level,
		},
	))

	slog.SetDefault(log)
	slog.Info("logger denined to " + level.String())
}
