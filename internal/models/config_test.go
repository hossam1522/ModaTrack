package models

import (
	"os"
	"testing"
)

func TestCargarConfiguracion(t *testing.T) {
	os.Setenv("LOG_FILE", "./logs/test.log")

	cfg, err := LoadConfig()
	defer os.Clearenv()

	if err != nil {
		t.Errorf("LoadConfig() error = %v", err)
		return
	}
	if cfg.LogFile != "./logs/test.log" {
		t.Errorf("LoadConfig() LogFile = %v, debería ser ./logs/test.log", cfg.LogFile)
	}
}

func TestCargarConfiguracionPorDefecto(t *testing.T) {
	os.Clearenv()

	cfg, err := LoadConfig()
	if err != nil {
		t.Errorf("LoadConfig() error = %v", err)
		return
	}
	if cfg.LogFile != "./logs/app.log" {
		t.Errorf("LoadConfig() LogFile = %v, debería ser ./logs/app.log", cfg.LogFile)
	}
}

func TestCargarConfiguracionError(t *testing.T) {
	os.Setenv("LOG_FILE", "3")

	defer os.Clearenv()

	_, err := LoadConfig()
	if err == nil {
		t.Errorf("LoadConfig() error = %v, debería ser error", err)
		return
	}
}
