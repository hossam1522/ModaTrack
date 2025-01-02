package api

import (
	"ModaTrack/internal/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func getRouter() *chi.Mux {
	router.Get("/prendas/{nombre}", getPrendas)
	router.Get("/prendas/{nombre}/{talla}", getPrendaTalla)
	router.Post("/prendas/{nombre}/{talla}/{cantidad}", putPrendaTalla)
	router.Delete("/prendas/{nombre}/{talla}", deletePrendaTalla)
	router.Get("/prendas/{nombre}/{talla}/ventas", getVentasNombreTalla)
	router.Get("/prendas/{nombre}/{talla}/ventas/{fecha}", getVentaFecha)
	router.Put("/prendas/{nombre}/{talla}/ventas/{fecha}", postVenta)
	router.Delete("/prendas/{nombre}/{talla}/ventas/{fecha}", deleteVenta)

	return router
}

func getPrendas(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	bd := models.BDPrueba()
	prendas, err := bd.ObtenerPrenda(nombre)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prendas)
}

func getPrendaTalla(w http.ResponseWriter, r *http.Request) {}

func putPrendaTalla(w http.ResponseWriter, r *http.Request) {}

func deletePrendaTalla(w http.ResponseWriter, r *http.Request) {}

func postVenta(w http.ResponseWriter, r *http.Request) {}

func getVentaFecha(w http.ResponseWriter, r *http.Request) {}

func getVentasNombreTalla(w http.ResponseWriter, r *http.Request) {}

func deleteVenta(w http.ResponseWriter, r *http.Request) {}
