package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func init() {
	router.Get("/prenda/{id}", GetPrenda)
	router.Get("/prenda/{id}/{talla}", GetPrendaTalla)
}

func GetPrenda(w http.ResponseWriter, r *http.Request) {}

func GetPrendaTalla(w http.ResponseWriter, r *http.Request) {}
