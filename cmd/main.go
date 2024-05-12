package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := setUpLogger(cfg.LogLevel)

	log.Info("starting application", slog.Any("config", cfg))

}

func setUpLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {

	case "info":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case "debug":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "error":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
		)
	case "warn":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}),
		)
	}

	return log
}
