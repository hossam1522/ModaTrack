package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func init() {
	router.Get("/prendas/{nombre}", getPrendas)
	router.Delete("/prendas/{nombre}", deletePrendas)
	router.Get("/prenda/{nombre}/{talla}", getPrendaTalla)
	router.Put("/prenda/{nombre}/{talla}/{cantidad}", putPrendaTalla)
	router.Delete("/prenda/{nombre}/{talla}", deletePrendaTalla)
	router.Put("/venta/{nombre}/{talla}", postVenta)
	router.Get("/venta/{nombre}/{talla}/{fecha}", getVentaFecha)
	router.Get("/ventas/{nombre}", getVentasNombre)
	router.Get("/ventas/{nombre}/{talla}", getVentasNombreTalla)
	router.Delete("/venta/{nombre}/{talla}/{fecha}", deleteVenta)
}

func getPrendas(w http.ResponseWriter, r *http.Request) {}

func deletePrendas(w http.ResponseWriter, r *http.Request) {}

func getPrendaTalla(w http.ResponseWriter, r *http.Request) {}

func putPrendaTalla(w http.ResponseWriter, r *http.Request) {}

func deletePrendaTalla(w http.ResponseWriter, r *http.Request) {}

func postVenta(w http.ResponseWriter, r *http.Request) {}

func getVentaFecha(w http.ResponseWriter, r *http.Request) {}

func getVentasNombre(w http.ResponseWriter, r *http.Request) {}

func getVentasNombreTalla(w http.ResponseWriter, r *http.Request) {}

func deleteVenta(w http.ResponseWriter, r *http.Request) {}
