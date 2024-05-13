package service

import (
	"context"
	"log"
	"os/exec"
	"strconv"

	"github.com/ZemtsovMaxim/gRPC_TestTask/pkg/api"
)

type NetVulnService struct{}

func (s *NetVulnService) CheckVuln(ctx context.Context, req *api.CheckVulnRequest) (*api.CheckVulnResponse, error) {

	log.Printf("Получен запрос на проверку уязвимостей: %v", req)

	cmdArgs := append([]string{"--script=vulners.nse", "-p", strconv.Itoa(int(req.TcpPort))}, req.Targets...)

	cmd := exec.Command("nmap", cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Ошибка при выполнении сканирования: %v", err)
		return nil, err
	}

	// Пример обработки вывода команды, можно присвоить результатам каким-либо полям в структуре ответа
	log.Printf("Результат сканирования: %s", string(output))

	// Возвращаем фиктивный ответ в качестве примера.
	return &api.CheckVulnResponse{}, nil
}
