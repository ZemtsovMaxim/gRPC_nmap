package tests

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/service"
	"github.com/ZemtsovMaxim/gRPC_TestTask/pkg/api"
)

func TestNetVulnServiceCheckVulnSuccess(t *testing.T) {
	logger := setupLogger()
	srv := service.NewNetVulnService(logger)

	req := &api.CheckVulnRequest{
		Targets: []string{"example.com"},
		TcpPort: 8080,
	}

	resp, err := srv.CheckVuln(context.Background(), req)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resp == nil {
		t.Error("Response is nil")
	}
}

func TestNetVulnServiceCheckVulnFailure(t *testing.T) {

	logger := setupLogger()
	srv := service.NewNetVulnService(logger)

	req := &api.CheckVulnRequest{
		Targets: []string{"example.com"},
		TcpPort: -12341234, // Некорректный порт
	}

	resp, err := srv.CheckVuln(context.Background(), req)

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if resp != nil {
		t.Error("Expected nil response, but got non-nil")
	}
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
