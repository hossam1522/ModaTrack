package models

import (
	"os"
	"testing"
)

func TestCargarConfiguracion(t *testing.T) {
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("LOG_FILE", "./logs/test.log")

	defer os.Clearenv()

	cfg, err := LoadConfig()
	if err != nil {
		t.Errorf("LoadConfig() error = %v", err)
		return
	}
	if cfg.LogLevel != "debug" {
		t.Errorf("LoadConfig() LogLevel = %v, deberia ser debug", cfg.LogLevel)
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
	if cfg.LogLevel != "info" {
		t.Errorf("LoadConfig() LogLevel = %v, deberia ser info", cfg.LogLevel)
	}
	if cfg.LogFile != "./logs/app.log" {
		t.Errorf("LoadConfig() LogFile = %v, debería ser ./logs/app.log", cfg.LogFile)
	}
}
