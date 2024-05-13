package tests

import (
	"os"
	"testing"

	"github.com/ZemtsovMaxim/gRPC_TestTask/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMustLoad(t *testing.T) {

	// Создаем временный файл
	file, err := os.Create("config_test.yaml")
	require.NoError(t, err, "failed to create temp file")
	defer os.Remove("config_test.yaml")

	_, err = file.WriteString("address: localhost:8080\nlog_level: debug\n")
	require.NoError(t, err, "failed to write to temp file")

	err = file.Close()
	require.NoError(t, err, "failed to close temp file")

	os.Setenv("CONFIG_PATH", file.Name())
	defer os.Unsetenv("CONFIG_PATH")

	cfg := config.MustLoad()

	assert.Equal(t, "localhost:8080", cfg.Addres)
	assert.Equal(t, "debug", cfg.LogLevel)
}
