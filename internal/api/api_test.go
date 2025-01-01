package api

import (
	"net/http/httptest"
	"testing"

	"ModaTrack/internal/models"

	"github.com/go-chi/chi/v5"
)

func TestGetPrendas(t *testing.T) {
	router := chi.NewRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	stock := models.NewStock()
	stock.inventario[models.Ropa{Nombre: "camisa", Talla: models.M}] = 1

}
