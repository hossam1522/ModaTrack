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

func TestCargarConfiguracionFichero(t *testing.T) {
	envContent := "LOG_FILE=./logs/test.log\n"
	tempFile := "./.env"
	err := os.WriteFile(tempFile, []byte(envContent), 0644)
	if err != nil {
		t.Fatalf("No se pudo crear el archivo temporal .env.test: %v", err)
	}
	defer os.Remove(tempFile)

	cfg, err := LoadConfigFromFile(tempFile)
	if err != nil {
		t.Errorf("LoadConfigFromFile() error = %v", err)
		return
	}
	if cfg.LogFile != "./logs/test.log" {
		t.Errorf("LoadConfigFromFile() LogFile = %v, debería ser ./logs/test.log", cfg.LogFile)
	}
}
