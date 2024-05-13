package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/config"
	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/service"
	"github.com/ZemtsovMaxim/gRPC_TestTask/pkg/api"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := logger.setUpLogger(cfg.LogLevel)

	log.Info("starting application", slog.Any("config", cfg))

	server := grpc.NewServer()

	srv := &service.NetVulnService{}

	api.RegisterNetVulnServiceServer(server, srv)

	listener, err := net.Listen("tcp", cfg.Addres)
	if err != nil {
		log.Error("Ошибка при создании TCP-соединения: %v", err)
	}

	log.Info("server listening", slog.Any("address", cfg.Addres))

	if err := server.Serve(listener); err != nil {
		log.Error("Ошибка при запуске сервера: %v", err)
	}

}

// Функция для логгера
func setUpLogger(env string) *slog.Logger {
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
