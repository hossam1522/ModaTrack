package api

import (
	"ModaTrack/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func getRouter() *chi.Mux {
	router.Get("/prendas/{nombre}", getPrendas)
	router.Get("/prendas/{nombre}/{talla}", getPrendaTalla)
	router.Post("/prendas/{nombre}/{talla}/{cantidad}", postPrendaTalla)
	router.Delete("/prendas/{nombre}/{talla}", deletePrendaTalla)
	router.Get("/prendas/{nombre}/{talla}/ventas", getVentasNombreTalla)
	router.Get("/prendas/{nombre}/{talla}/ventas/{fecha}", getVentaFecha)
	router.Put("/prendas/{nombre}/{talla}/ventas/{fecha}", postVenta)
	router.Delete("/prendas/{nombre}/{talla}/ventas/{fecha}", deleteVenta)

	return router
}

func getPrendas(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	bd := models.GetBDPrueba()
	prendas, err := bd.ObtenerPrenda(nombre)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prendas); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getPrendaTalla(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	talla := chi.URLParam(r, "talla")
	bd := models.GetBDPrueba()
	prenda, err := bd.ObtenerPrendaTalla(nombre, models.Talla(talla))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prenda); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func postPrendaTalla(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	talla := chi.URLParam(r, "talla")
	cantidad := chi.URLParam(r, "cantidad")
	cantidadInt, errNum := strconv.Atoi(cantidad)
	if errNum != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bd := models.GetBDPrueba()
	err := bd.InsertarRopa(nombre, models.Talla(talla), cantidadInt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func deletePrendaTalla(w http.ResponseWriter, r *http.Request) {}

func postVenta(w http.ResponseWriter, r *http.Request) {}

func getVentaFecha(w http.ResponseWriter, r *http.Request) {}

func getVentasNombreTalla(w http.ResponseWriter, r *http.Request) {}

func deleteVenta(w http.ResponseWriter, r *http.Request) {}
