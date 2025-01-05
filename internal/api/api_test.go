package api

import (
	"ModaTrack/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetPrendas(t *testing.T) {
	router := getRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/prendas/camisa"
	resp, err := server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Se esperaba un status 200, se obtuvo %d", resp.StatusCode)
	}
	var prendas []models.Ropa
	err = json.NewDecoder(resp.Body).Decode(&prendas)
	if err != nil {
		t.Error(err)
	}
	prendas_en_bd, _ := models.GetBDPrueba().ObtenerPrenda("camisa")
	if len(prendas) != len(prendas_en_bd) {
		t.Errorf("Se esperaban %d prendas, se obtuvieron %d", len(prendas_en_bd), len(prendas))
	}
	if prendas[0].GetNombre() != "camisa" || prendas[1].GetNombre() != "camisa" {
		t.Error("Las prendas no son las esperadas")
	}
}

func TestGetPrendasNoExistentes(t *testing.T) {
	router := getRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/prendas/chaqueta"
	resp, err := server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 404 {
		t.Errorf("Se esperaba un status 404, se obtuvo %d", resp.StatusCode)
	}
}

func TestGetPrendasTalla(t *testing.T) {
	router := getRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/prendas/camisa/M"
	resp, err := server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Se esperaba un status 200, se obtuvo %d", resp.StatusCode)
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
	router := getRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/prendas/chaqueta/XL/1"
	resp, err := server.Client().Post(url, "application/json", nil)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 201 {
		t.Errorf("Se esperaba un status 201, se obtuvo %d", resp.StatusCode)
	}

	url = server.URL + "/prendas/chaqueta/XL"
	resp, err = server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Se esperaba un status 200, se obtuvo %d", resp.StatusCode)
	}
	var prenda models.Ropa
	err = json.NewDecoder(resp.Body).Decode(&prenda)
	if err != nil {
		t.Error(err)
	}
	if prenda.GetTalla() != models.XL || prenda.GetNombre() != "chaqueta" {
		t.Errorf("Se esperaba una talla XL, se obtuvo %s", prenda.GetTalla())
	}
}

func TestGetVentasPrenda(t *testing.T) {
	router := getRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/prendas/camisa/L/ventas"
	resp, err := server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Se esperaba un status 200, se obtuvo %d", resp.StatusCode)
	}
	var ventas []models.Venta
	err = json.NewDecoder(resp.Body).Decode(&ventas)
	if err != nil {
		t.Error(err)
	}
	ventas_en_bd, _ := models.GetBDPrueba().ObtenerVentas("camisa", models.L)
	if len(ventas) != len(ventas_en_bd) {
		t.Errorf("Se esperaban %d ventas, se obtuvieron %d", len(ventas_en_bd), len(ventas))
	}

	url = server.URL + "/prendas/camisa/L"
	resp, _ = server.Client().Get(url)
	prenda := models.Ropa{}
	json.NewDecoder(resp.Body).Decode(&prenda)

	if ventas[0].GetItemsVendidos()[prenda] != 1 || ventas[1].GetItemsVendidos()[prenda] != 1 {
		t.Errorf("Las ventas no son las esperadas")
	}
}

func TestGetVentasPrendaFecha(t *testing.T) {
	router := getRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/prendas/pantalon/M/ventas/2024-06-12T15:30:45Z"
	resp, err := server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Se esperaba un status 200, se obtuvo %d", resp.StatusCode)
	}
	var ventas []models.Venta
	err = json.NewDecoder(resp.Body).Decode(&ventas)
	if err != nil {
		t.Error(err)
	}
	ventas_en_bd, _ := models.GetBDPrueba().ObtenerVentas("pantalon", models.M, time.Date(2024, 6, 12, 15, 30, 45, 0, time.UTC))
	if len(ventas) != len(ventas_en_bd) {
		t.Errorf("Se esperaban %d venta, se obtuvieron %d", len(ventas_en_bd), len(ventas))
	}

	url = server.URL + "/prendas/pantalon/M"
	resp, _ = server.Client().Get(url)
	prenda := models.Ropa{}
	json.NewDecoder(resp.Body).Decode(&prenda)

	if ventas[0].GetItemsVendidos()[prenda] != 1 {
		t.Errorf("Las ventas no son las esperadas")
		t.Error(ventas)
	}
}

func TestPutVenta(t *testing.T) {
	router := getRouter()
	server := httptest.NewServer(router)
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

	if resp.StatusCode != 201 {
		t.Errorf("Se esperaba un status 201, se obtuvo %d", resp.StatusCode)
	}

	url = server.URL + "/prendas/pantalon/L/ventas/2024-06-12T15:30:45Z"
	resp, err = server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Se esperaba un status 200, se obtuvo %d", resp.StatusCode)
	}
	var ventas []models.Venta
	err = json.NewDecoder(resp.Body).Decode(&ventas)
	if err != nil {
		t.Error(err)
	}
	if ventas[0].GetItemsVendidos()[models.Ropa{}] != 0 {
		t.Errorf("La venta no es la esperada")
	}

	url = server.URL + "/prendas/pantalon/L"
	resp, err = server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 404 {
		t.Errorf("Se esperaba un status 404, se obtuvo %d", resp.StatusCode)
	}
}

func TestDeleteVenta(t *testing.T) {
	router := getRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/prendas/pantalon/L/ventas/2024-06-12T15:30:45Z"
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		t.Error(err)
	}

	resp, err := server.Client().Do(req)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 202 {
		t.Errorf("Se esperaba un status 202, se obtuvo %d", resp.StatusCode)
	}

	url = server.URL + "/prendas/pantalon/L/ventas/2024-06-12T15:30:45Z"
	resp, err = server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 404 {
		t.Errorf("Se esperaba un status 404, se obtuvo %d", resp.StatusCode)
	}
}
