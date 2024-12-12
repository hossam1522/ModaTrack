package models

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	LogFile string `env:"LOG_FILE" envDefault:"./logs/app.log"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
