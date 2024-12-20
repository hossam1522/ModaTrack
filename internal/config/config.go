package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	LogFile string `env:"LOG_FILE" envDefault:"../../logs/test.log"`
}

func LoadConfig() (*Config, error) {
	if _, err := os.Stat(".env"); err == nil {
		err := readConfigFromFile()
		if err != nil {
			return nil, err
		}
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func readConfigFromFile() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("línea inválida en .env: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
