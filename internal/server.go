package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/ZemtsovMaxim/gRPC_TestTask/api/netvuln"

	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/service"
)

// RunServer запускает gRPC-сервер для обслуживания запросов.
func RunServer(address string) error {
	// Создание нового gRPC-сервера
	server := grpc.NewServer()

	// Регистрация вашего сервиса на сервере
	netvuln.RegisterNetVulnServiceServer(server, &service.NetVulnService{})

	// Создание TCP-соединения на указанном адресе и порте
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Ошибка при создании TCP-соединения: %v", err)
	}

	// Запуск сервера и начало обслуживания запросов
	log.Printf("Сервер запущен на адресе %s", address)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

	return nil
}
