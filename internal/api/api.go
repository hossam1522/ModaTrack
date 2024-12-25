package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func init() {
	router.Get("/prendas/{nombre}", getPrendas)
	router.Get("/prendas/{nombre}/{talla}", getPrendaTalla)
	router.Post("/prendas/{nombre}/{talla}/{cantidad}", putPrendaTalla)
	router.Delete("/prendas/{nombre}/{talla}", deletePrendaTalla)
	router.Put("/ventas/{nombre}/{talla}", postVenta)
	router.Get("/ventas/{nombre}/{talla}/{fecha}", getVentaFecha)
	router.Get("/ventas/{nombre}", getVentasNombre)
	router.Get("/ventas/{nombre}/{talla}", getVentasNombreTalla)
	router.Delete("/ventas/{nombre}/{talla}/{fecha}", deleteVenta)
}

func getPrendas(w http.ResponseWriter, r *http.Request) {}

func getPrendaTalla(w http.ResponseWriter, r *http.Request) {}

func putPrendaTalla(w http.ResponseWriter, r *http.Request) {}

func deletePrendaTalla(w http.ResponseWriter, r *http.Request) {}

func postVenta(w http.ResponseWriter, r *http.Request) {}

func getVentaFecha(w http.ResponseWriter, r *http.Request) {}

func getVentasNombre(w http.ResponseWriter, r *http.Request) {}

func getVentasNombreTalla(w http.ResponseWriter, r *http.Request) {}

func deleteVenta(w http.ResponseWriter, r *http.Request) {}
