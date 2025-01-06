package api

import (
	"ModaTrack/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()
var bd = models.GetBDPrueba()

func capturarParametrosRuta(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parametrosRuta := chi.RouteContext(r.Context()).URLParams
		parametrosMap := make(map[string]string, len(parametrosRuta.Keys))
		for i, key := range parametrosRuta.Keys {
			parametrosMap[key] = parametrosRuta.Values[i]
		}
		ctx := context.WithValue(r.Context(), "parametrosMap", parametrosMap)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getRouter() *chi.Mux {
	router.With(capturarParametrosRuta).Get("/prendas/{nombre}", getPrendas)
	router.With(capturarParametrosRuta).Get("/prendas/{nombre}/{talla}", getPrendaTalla)
	router.With(capturarParametrosRuta).Post("/prendas/{nombre}/{talla}/{cantidad}", postPrendaTalla)
	router.With(capturarParametrosRuta).Get("/prendas/{nombre}/{talla}/ventas", getVentasPrenda)
	router.With(capturarParametrosRuta).Get("/prendas/{nombre}/{talla}/ventas/{fecha}", getVentaFecha)
	router.With(capturarParametrosRuta).Put("/prendas/{nombre}/{talla}/ventas/{fecha}", putVenta)
	router.With(capturarParametrosRuta).Delete("/prendas/{nombre}/{talla}/ventas/{fecha}", deleteVenta)

	return router
}

func getPrendas(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("parametrosMap").(map[string]string)
	prendas, err := models.ObtenerPrenda(bd, params["nombre"])
	if err != nil {
		http.Error(w, fmt.Sprintf("No se ha podido obtener la prenda '%s' de la base de datos", params["nombre"]), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prendas); err != nil {
		http.Error(w, fmt.Sprintf("No se ha podido codificar la prenda '%s' a formato JSON", params["nombre"]), http.StatusBadRequest)
	}
}

func getPrendaTalla(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("parametrosMap").(map[string]string)
	prenda, err := models.ObtenerPrendaTalla(bd, params["nombre"], models.Talla(params["talla"]))
	if err != nil {
		http.Error(w, fmt.Sprintf("No se ha podido obtener la prenda '%s' con talla '%s' de la base de datos", params["nombre"], params["talla"]), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(prenda); err != nil {
		http.Error(w, fmt.Sprintf("No se ha podido codificar la prenda '%s' con talla '%s' a formato JSON", params["nombre"], params["talla"]), http.StatusBadRequest)
	}
}

func postPrendaTalla(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("parametrosMap").(map[string]string)
	cantidadInt, errNum := strconv.Atoi(params["cantidad"])
	if errNum != nil {
		http.Error(w, fmt.Sprintf("La cantidad '%s' no es un número válido", params["cantidad"]), http.StatusBadRequest)
		return
	}
	err := models.InsertarRopa(bd, params["nombre"], models.Talla(params["talla"]), cantidadInt)
	if err != nil {
		http.Error(w, fmt.Sprintf("No se ha podido insertar la prenda '%s' con talla '%s' en la base de datos", params["nombre"], params["talla"]), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getVentasPrenda(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("parametrosMap").(map[string]string)
	ventas, err := models.ObtenerVentas(bd, params["nombre"], models.Talla(params["talla"]))
	if err != nil {
		http.Error(w, fmt.Sprintf("No se han podido obtener las ventas de la prenda '%s' con talla '%s' de la base de datos", params["nombre"], params["talla"]), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ventas); err != nil {
		http.Error(w, fmt.Sprintf("No se han podido codificar las ventas de la prenda '%s' con talla '%s' a formato JSON", params["nombre"], params["talla"]), http.StatusBadRequest)
	}
}

func getVentaFecha(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("parametrosMap").(map[string]string)
	fechaVenta, err := time.Parse(time.RFC3339, params["fecha"])
	if err != nil {
		http.Error(w, fmt.Sprintf("La fecha '%s' no tiene el formato correcto, debe ser RFC3339", params["fecha"]), http.StatusUnprocessableEntity)
		return
	}
	venta, err := models.ObtenerVentas(bd, params["nombre"], models.Talla(params["talla"]), fechaVenta)
	if err != nil {
		http.Error(w, fmt.Sprintf("No se han podido obtener las ventas de la prenda '%s' con talla '%s' y fecha '%s' de la base de datos", params["nombre"], params["talla"], params["fecha"]), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(venta); err != nil {
		http.Error(w, fmt.Sprintf("No se han podido codificar las ventas de la prenda '%s' con talla '%s' y fecha '%s' a formato JSON", params["nombre"], params["talla"], params["fecha"]), http.StatusBadRequest)
	}
}

func putVenta(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("parametrosMap").(map[string]string)
	fechaVenta, err := time.Parse(time.RFC3339, params["fecha"])
	if err != nil {
		http.Error(w, fmt.Sprintf("La fecha '%s' no tiene el formato correcto, debe ser RFC3339", params["fecha"]), http.StatusUnprocessableEntity)
		return
	}
	err = models.InsertarVenta(bd, params["nombre"], models.Talla(params["talla"]), fechaVenta)
	if err != nil {
		http.Error(w, fmt.Sprintf("No se ha podido insertar la venta de la prenda '%s' con talla '%s' y fecha '%s' en la base de datos", params["nombre"], params["talla"], params["fecha"]), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func deleteVenta(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("parametrosMap").(map[string]string)
	fechaVenta, err := time.Parse(time.RFC3339, params["fecha"])
	if err != nil {
		http.Error(w, fmt.Sprintf("La fecha '%s' no tiene el formato correcto, debe ser RFC3339", params["fecha"]), http.StatusUnprocessableEntity)
		return
	}
	err = models.EliminarVenta(bd, params["nombre"], models.Talla(params["talla"]), fechaVenta)
	if err != nil {
		http.Error(w, fmt.Sprintf("No se ha podido eliminar la venta de la prenda '%s' con talla '%s' y fecha '%s' en la base de datos", params["nombre"], params["talla"], params["fecha"]), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
