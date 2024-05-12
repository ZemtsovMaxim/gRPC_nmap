package nmap

import (
	"fmt"
	"log"
	"os/exec"
)

func Scan(targets []string, tcpPort int32) (string, error) {
	cmd := exec.Command("nmap", append([]string{"--script=vulners.nse", fmt.Sprintf("-p %d", tcpPort)}, targets...)...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Ошибка при выполнении сканирования: %v", err)
		return "", err
	}

	return string(output), nil
}
