package api

import (
	"ModaTrack/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
	if len(prendas) != 2 {
		t.Errorf("Se esperaban 2 prendas, se obtuvieron %d", len(prendas))
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

func TestDeletePrenda(t *testing.T) {
	router := getRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/prendas/camisa/M"
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

	url = server.URL + "/prendas/camisa/M"
	resp, err = server.Client().Get(url)

	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 404 {
		t.Errorf("Se esperaba un status 404, se obtuvo %d", resp.StatusCode)
	}
}
