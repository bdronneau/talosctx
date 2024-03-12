package talosconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lederniermetre/talosctx/internal/env"
)

var getHomeDir = os.UserHomeDir

func GetTalosConfigContent() ([]byte, string, error) {
	if configPath := os.Getenv(env.TALOS_ENV_CONFIG_NAME); configPath != "" {
		yamlData, err := os.ReadFile(configPath)
		if err != nil {
			return nil, "", err
		}

		return yamlData, configPath, nil
	}

	configPath, err := getTalosConfigHome()
	if err != nil {
		return nil, "", fmt.Errorf("get config from user home: %w", err)
	}

	yamlData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, "", fmt.Errorf("read config from user home: %w", err)
	}

	return yamlData, configPath, nil
}

func getTalosConfigHome() (string, error) {
	homeDir, err := getHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".talos", "config"), nil
}

func WritetalosConfig(path string, data []byte) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, info.Mode())
	if err != nil {
		return err
	}

	return nil
}
