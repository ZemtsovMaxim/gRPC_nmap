package service

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"strconv"

	"github.com/ZemtsovMaxim/gRPC_TestTask/pkg/api"
)

type NetVulnService struct {
	Logger *slog.Logger
}

// Конструктор для NetVulnService
func NewNetVulnService(log *slog.Logger) *NetVulnService {
	return &NetVulnService{
		Logger: log,
	}
}
func (s *NetVulnService) CheckVuln(ctx context.Context, req *api.CheckVulnRequest) (*api.CheckVulnResponse, error) {

	s.Logger.Info("Получен запрос на проверку уязвимостей:", slog.Any("incoming request", "method"))

	cmdArgs := append([]string{"--script=vulners.nse", "-p", strconv.Itoa(int(req.TcpPort))}, req.Targets...)

	cmd := exec.Command("nmap", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		s.Logger.Error(fmt.Sprintf("Ошибка при выполнении сканирования: %v", err))
		return nil, err
	}

	s.Logger.Info("Результат сканирования:")
	s.Logger.Info(string(output))

	return &api.CheckVulnResponse{}, nil
}
