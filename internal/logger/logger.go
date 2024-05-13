package logger

import (
	"log/slog"
	"os"
)

func SetUpLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "info":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case "debug":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}),
		)
	case "error":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError, AddSource: true}),
		)
	case "warn":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn, AddSource: true}),
		)
	}

	return log
}
