package main

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/config"
	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/logger"
	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/service"
	"github.com/ZemtsovMaxim/gRPC_TestTask/pkg/api"
	"google.golang.org/grpc"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := logger.SetUpLogger(cfg.LogLevel)

	log.Info("starting application", slog.Any("config", cfg))

	server := grpc.NewServer()

	srv := &service.NetVulnService{log}

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
