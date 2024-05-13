package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/config"
	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/logger"
	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/service"
	"github.com/ZemtsovMaxim/gRPC_TestTask/pkg/api"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.MustLoad()

	log := logger.SetUpLogger(cfg.LogLevel)

	log.Info("starting application", slog.Any("config", cfg))

	startServer(cfg, log)

}

func startServer(cfg *config.Config, log *slog.Logger) {

	server := grpc.NewServer()

	srv := service.NewNetVulnService(log)

	api.RegisterNetVulnServiceServer(server, srv)

	listener, err := net.Listen("tcp", cfg.Addres)
	if err != nil {
		log.Error(fmt.Sprintf("Ошибка при создании TCP-соединения %v", err))
		os.Exit(1)
	}

	log.Info("server listening", slog.Any("address", cfg.Addres))

	if err := server.Serve(listener); err != nil {
		log.Error(fmt.Sprintf("Ошибка при запуске сервера %v", err))
		os.Exit(1)
	}
}
