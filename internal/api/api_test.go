package api

import (
	"ModaTrack/internal/models"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func BD_prueba() *models.BD {
	bd := models.NewBD()

	bd.InsertarRopa("camisa", models.M, 1)
	bd.InsertarRopa("camisa", models.L, 1)
	bd.InsertarRopa("pantalon", models.M, 1)
	bd.InsertarRopa("pantalon", models.L, 1)

	return bd
}

func TestGetPrendas(t *testing.T) {
	router := chi.NewRouter()
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
}
