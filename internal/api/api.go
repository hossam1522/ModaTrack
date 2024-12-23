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
	router.Post("/venta", postVenta)
	router.Get("/venta/{id}/fecha", getFechaVenta)
	router.Get("/venta/{id}/prendas", getPrendasVenta)
	router.Delete("/venta/{id}", deleteVenta)
}

func getPrendas(w http.ResponseWriter, r *http.Request) {}

func deletePrendas(w http.ResponseWriter, r *http.Request) {}

func getPrendaTalla(w http.ResponseWriter, r *http.Request) {}

func putPrendaTalla(w http.ResponseWriter, r *http.Request) {}

func deletePrendaTalla(w http.ResponseWriter, r *http.Request) {}

func postVenta(w http.ResponseWriter, r *http.Request) {}

func getPrendasVenta(w http.ResponseWriter, r *http.Request) {}

func getFechaVenta(w http.ResponseWriter, r *http.Request) {}

func deleteVenta(w http.ResponseWriter, r *http.Request) {}
