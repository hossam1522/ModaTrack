package api

import (
	"ModaTrack/internal/models"
	"encoding/json"
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
}
