package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func init() {
	router.Get("/prenda/{nombre}", GetPrenda)
	router.Put("/prenda/{nombre}", PutPrenda)
	router.Delete("/prenda/{nombre}", DeletePrenda)
	router.Get("/prenda/{nombre}/{talla}", GetPrendaTalla)
	router.Put("/prenda/{nombre}/{talla}", PutPrendaTalla)
	router.Delete("/prenda/{nombre}/{talla}", DeletePrendaTalla)
	router.Post("/venta", PostVenta)
	router.Get("/venta/{id}/fecha", GetFechaVenta)
	router.Get("/venta/{id}/prendas", GetPrendasVenta)
	router.Delete("/venta/{id}", DeleteVenta)
}

func GetPrenda(w http.ResponseWriter, r *http.Request) {}

func PutPrenda(w http.ResponseWriter, r *http.Request) {}

func DeletePrenda(w http.ResponseWriter, r *http.Request) {}

func GetPrendaTalla(w http.ResponseWriter, r *http.Request) {}

func PutPrendaTalla(w http.ResponseWriter, r *http.Request) {}

func DeletePrendaTalla(w http.ResponseWriter, r *http.Request) {}

func PostVenta(w http.ResponseWriter, r *http.Request) {}

func GetPrendasVenta(w http.ResponseWriter, r *http.Request) {}

func GetFechaVenta(w http.ResponseWriter, r *http.Request) {}

func DeleteVenta(w http.ResponseWriter, r *http.Request) {}
