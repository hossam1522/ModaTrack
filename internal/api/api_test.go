package api

import (
	"ModaTrack/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
)

var dependencias = Server{
	router: chi.NewRouter(),
	bd:     models.NewBD(models.WithJSON("../models/bd_test.json")),
}

func TestGetPrendas(t *testing.T) {
	server := httptest.NewServer(dependencias.GetRouter())
	defer server.Close()

	url := server.URL + "/prendas/camisa"
	resp, err := server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un status %d, se obtuvo %d", http.StatusOK, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		t.Error(string(body))
	}

	var prendas []models.Ropa
	err = json.NewDecoder(resp.Body).Decode(&prendas)

	if err != nil {
		t.Error(err)
	}

	prendas_en_bd, _ := models.ObtenerPrenda(dependencias.bd, "camisa")

	if len(prendas) != len(prendas_en_bd) {
		t.Errorf("Se esperaban %d prendas, se obtuvieron %d", len(prendas_en_bd), len(prendas))
	}
}

func TestGetPrendasNoExistentes(t *testing.T) {
	server := httptest.NewServer(dependencias.GetRouter())
	defer server.Close()

	url := server.URL + "/prendas/chaqueta"
	resp, err := server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Se esperaba un status %d, se obtuvo %d", http.StatusNotFound, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		t.Error(string(body))
	}
}

func TestGetPrendasTalla(t *testing.T) {
	server := httptest.NewServer(dependencias.GetRouter())
	defer server.Close()

	url := server.URL + "/prendas/camisa/M"
	resp, err := server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un status %d, se obtuvo %d", http.StatusOK, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		t.Error(string(body))
	}

	var prenda models.Ropa
	err = json.NewDecoder(resp.Body).Decode(&prenda)

	if err != nil {
		t.Error(err)
	}

	if prenda.GetTalla() != models.M || prenda.GetNombre() != "camisa" {
		t.Errorf("Se esperaba una talla M, se obtuvo %s", prenda.GetTalla())
	}
}

func TestPostPrenda(t *testing.T) {
	server := httptest.NewServer(dependencias.GetRouter())
	defer server.Close()

	url := server.URL + "/prendas/chaqueta/XL/1"
	resp, err := server.Client().Post(url, "application/json", nil)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Se esperaba un status %d, se obtuvo %d", http.StatusCreated, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		t.Error(string(body))
	}

	prenda_en_bd, _ := models.ObtenerPrendaTalla(dependencias.bd, "chaqueta", models.XL)

	if prenda_en_bd.GetTalla() != models.XL || prenda_en_bd.GetNombre() != "chaqueta" {
		t.Errorf("Se esperaba una talla XL, se obtuvo %s", prenda_en_bd.GetTalla())
	}
}

func TestGetVentasPrenda(t *testing.T) {
	server := httptest.NewServer(dependencias.GetRouter())
	defer server.Close()

	url := server.URL + "/prendas/camisa/L/ventas"
	resp, err := server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un status %d, se obtuvo %d", http.StatusOK, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		t.Error(string(body))
	}

	var ventas []models.Venta
	err = json.NewDecoder(resp.Body).Decode(&ventas)

	if err != nil {
		t.Error(err)
	}

	ventas_en_bd, _ := models.ObtenerVentas(dependencias.bd, "camisa", models.L)

	if len(ventas) != len(ventas_en_bd) {
		t.Errorf("Se esperaban %d ventas, se obtuvieron %d", len(ventas_en_bd), len(ventas))
	}
}

func TestGetVentasPrendaFecha(t *testing.T) {
	server := httptest.NewServer(dependencias.GetRouter())
	defer server.Close()

	url := server.URL + "/prendas/pantalon/M/ventas/2024-06-12T15:30:45Z"
	resp, err := server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un status %d, se obtuvo %d", http.StatusOK, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		t.Error(string(body))
	}

	var ventas []models.Venta
	err = json.NewDecoder(resp.Body).Decode(&ventas)

	if err != nil {
		t.Error(err)
	}

	ventas_en_bd, _ := models.ObtenerVentas(dependencias.bd, "pantalon", models.M, time.Date(2024, 6, 12, 15, 30, 45, 0, time.UTC))
	if len(ventas) != len(ventas_en_bd) {
		t.Errorf("Se esperaban %d venta, se obtuvieron %d", len(ventas_en_bd), len(ventas))
	}
}

func TestPutVenta(t *testing.T) {
	server := httptest.NewServer(dependencias.GetRouter())
	defer server.Close()

	url := server.URL + "/prendas/pantalon/L/ventas/2024-06-12T15:30:45Z"
	req, err := http.NewRequest(http.MethodPut, url, nil)

	if err != nil {
		t.Error(err)
	}

	resp, err := server.Client().Do(req)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Se esperaba un status %d, se obtuvo %d", http.StatusCreated, resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		t.Error(string(body))
	}
}
