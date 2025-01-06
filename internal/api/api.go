package api

import (
	"ModaTrack/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()
var bd = models.GetBDPrueba()

func getRouter() *chi.Mux {
	router.Get("/prendas/{nombre}", getPrendas)
	router.Get("/prendas/{nombre}/{talla}", getPrendaTalla)
	router.Post("/prendas/{nombre}/{talla}/{cantidad}", postPrendaTalla)
	router.Get("/prendas/{nombre}/{talla}/ventas", getVentasPrenda)
	router.Get("/prendas/{nombre}/{talla}/ventas/{fecha}", getVentaFecha)
	router.Put("/prendas/{nombre}/{talla}/ventas/{fecha}", putVenta)
	router.Delete("/prendas/{nombre}/{talla}/ventas/{fecha}", deleteVenta)

	return router
}

func getPrendas(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	prendas, err := models.ObtenerPrenda(bd, nombre)
	if err != nil {
		http.Error(w, "No se ha podido obtener la prenda de la base de datos", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prendas); err != nil {
		http.Error(w, "No se ha podido codificar la prenda a formato JSON", http.StatusBadRequest)
	}
}

func getPrendaTalla(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	talla := chi.URLParam(r, "talla")
	prenda, err := models.ObtenerPrendaTalla(bd, nombre, models.Talla(talla))
	if err != nil {
		http.Error(w, "No se ha podido obtener la prenda de la base de datos", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prenda); err != nil {
		http.Error(w, "No se ha podido codificar la prenda a formato JSON", http.StatusBadRequest)
	}
}

func postPrendaTalla(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	talla := chi.URLParam(r, "talla")
	cantidad := chi.URLParam(r, "cantidad")
	cantidadInt, errNum := strconv.Atoi(cantidad)
	if errNum != nil {
		http.Error(w, "La cantidad debe ser un número entero", http.StatusBadRequest)
		return
	}
	err := models.InsertarRopa(bd, nombre, models.Talla(talla), cantidadInt)
	if err != nil {
		http.Error(w, "No se ha podido insertar la prenda en la base de datos", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getVentasPrenda(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	talla := chi.URLParam(r, "talla")
	ventas, err := models.ObtenerVentas(bd, nombre, models.Talla(talla))
	if err != nil {
		http.Error(w, "No se han podido obtener las ventas de la prenda", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ventas); err != nil {
		http.Error(w, "No se han podido codificar las ventas a formato JSON", http.StatusBadRequest)
	}
}

func getVentaFecha(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	talla := chi.URLParam(r, "talla")
	fecha := chi.URLParam(r, "fecha")
	fechaVenta, err := time.Parse(time.RFC3339, fecha)
	if err != nil {
		http.Error(w, "La fecha no tiene el formato correcto, debe ser RFC3339", http.StatusUnprocessableEntity)
		return
	}
	venta, err := models.ObtenerVentas(bd, nombre, models.Talla(talla), fechaVenta)
	if err != nil {
		http.Error(w, "No se ha podido obtener la venta de la prenda de la base de datos", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(venta); err != nil {
		http.Error(w, "No se ha podido codificar la venta a formato JSON", http.StatusBadRequest)
	}
}

func putVenta(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	talla := chi.URLParam(r, "talla")
	fecha := chi.URLParam(r, "fecha")
	fechaVenta, err := time.Parse(time.RFC3339, fecha)
	if err != nil {
		http.Error(w, "La fecha no tiene el formato correcto, debe ser RFC3339", http.StatusUnprocessableEntity)
		return
	}
	err = models.InsertarVenta(bd, nombre, models.Talla(talla), fechaVenta)
	if err != nil {
		http.Error(w, "No se ha podido insertar la venta en la base de datos", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func deleteVenta(w http.ResponseWriter, r *http.Request) {
	nombre := chi.URLParam(r, "nombre")
	talla := chi.URLParam(r, "talla")
	fecha := chi.URLParam(r, "fecha")
	fechaVenta, err := time.Parse(time.RFC3339, fecha)
	if err != nil {
		http.Error(w, "La fecha no tiene el formato correcto, debe ser RFC3339", http.StatusUnprocessableEntity)
		return
	}
	err = models.EliminarVenta(bd, nombre, models.Talla(talla), fechaVenta)
	if err != nil {
		http.Error(w, "No se ha podido eliminar la venta de la base de datos", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
