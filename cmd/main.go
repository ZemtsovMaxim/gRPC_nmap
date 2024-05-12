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

	log := setUpLogger(cfg.LogLevel)

	log.Info("starting application", slog.Any("config", cfg))

	// Создаем новый gRPC сервер
	server := grpc.NewServer()

	// Регистрируем ваш сервис на сервере
	api.RegisterNetVulnServiceServer(server, &service.NetVulnService{})

	// Создаем TCP-прослушиватель на указанном адресе и порте
	listener, err := net.Listen("tcp", cfg.Addres)
	if err != nil {
		log.Info("Ошибка при создании TCP-соединения: %v", err)
	}

	log.Info("server listening", slog.Any("address", cfg.Addres))

	// Запускаем сервер и начинаем обслуживание запросов
	if err := server.Serve(listener); err != nil {
		log.Info("Ошибка при запуске сервера: %v", err)
	}

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
