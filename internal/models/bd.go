package models

import (
	"ModaTrack/internal/log"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type BD struct {
	Stock  *Stock
	Ventas []Venta
}

type Option func(*BD) error

func NewBD(options ...Option) *BD {
	bd := &BD{
		Stock:  NewStock(),
		Ventas: []Venta{},
	}

	for _, option := range options {
		if err := option(bd); err != nil {
			return nil
		}
	}

	return bd
}

func WithJSON(filename string) Option {
	return func(bd *BD) error {
		log.GetLogger().Info().Msg("Insertando stock en la base de datos")

		data, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("no se ha podido abrir el archivo %s: %w", filename, err)
		}
		defer data.Close()

		var datosInventario struct {
			Stock struct {
				Inventario map[string]int `json:"inventario"`
			} `json:"Stock"`
			Ventas []struct {
				Fecha         string         `json:"fecha"`
				ItemsVendidos map[string]int `json:"itemsVendidos"`
			} `json:"Ventas"`
		}

		if err := json.NewDecoder(data).Decode(&datosInventario); err != nil {
			return fmt.Errorf("no se ha podido decodificar el archivo %s: %w", filename, err)
		}

		for nombre, cantidad := range datosInventario.Stock.Inventario {
			ropa, err := parseRopaTalla(nombre)
			if err != nil {
				return fmt.Errorf("no se ha podido parsear la ropa: %w", err)
			}
			bd.Stock.inventario[ropa] = cantidad
		}

		for _, venta := range datosInventario.Ventas {
			fecha, err := time.Parse(time.RFC3339, venta.Fecha)
			if err != nil {
				return fmt.Errorf("no se ha podido parsear la fecha: %w", err)
			}

			itemsVendidos := make(map[Ropa]int)
			for nombre, cantidad := range venta.ItemsVendidos {
				ropa, err := parseRopaTalla(nombre)
				if err != nil {
					return fmt.Errorf("no se ha podido parsear la ropa: %w", err)
				}
				itemsVendidos[ropa] = cantidad
			}

			bd.Ventas = append(bd.Ventas, Venta{
				itemsVendidos: itemsVendidos,
				fecha:         fecha,
			})
		}

		return nil

	}
}

func parseRopaTalla(key string) (Ropa, error) {
	partes := strings.Split(key, "_")
	if len(partes) != 2 {
		return Ropa{}, errors.New("clave de ropa inválida: " + key)
	}
	nombre := partes[0]
	talla := Talla(partes[1])
	return Ropa{nombre: nombre, talla: talla}, nil
}

func InsertarRopa(bd *BD, nombre string, talla Talla, cantidad int) error {
	log.GetLogger().Info().Msg("Insertando ropa en la base de datos")

	cantidadActual := bd.Stock.inventario[Ropa{nombre: nombre, talla: talla}]
	bd.Stock.inventario[Ropa{nombre: nombre, talla: talla}] = cantidadActual + cantidad

	if bd.Stock.inventario[Ropa{nombre: nombre, talla: talla}] != cantidadActual+cantidad {
		return errors.New("no se ha podido insertar la ropa")
	}

	return nil
}

func ObtenerPrenda(bd *BD, nombre string) ([]Ropa, error) {
	log.GetLogger().Info().Msg("Obteniendo prenda de la base de datos")

	var prendas []Ropa
	for ropa := range bd.Stock.inventario {
		if ropa.nombre == nombre && bd.Stock.inventario[ropa] > 0 {
			prendas = append(prendas, ropa)
		}
	}

	if len(prendas) == 0 {
		return nil, errors.New("no se ha podido obtener la prenda")
	}

	return prendas, nil
}

func ObtenerPrendaTalla(bd *BD, nombre string, talla Talla) (Ropa, error) {
	log.GetLogger().Info().Msg("Obteniendo prenda de la base de datos")

	for ropa := range bd.Stock.inventario {
		if ropa.nombre == nombre && ropa.talla == talla && bd.Stock.inventario[ropa] > 0 {
			return ropa, nil
		}
	}

	return Ropa{}, errors.New("no se ha podido obtener la prenda")
}

func EliminarRopa(bd *BD, nombre string, talla Talla) error {
	log.GetLogger().Info().Msg("Eliminando ropa de la base de datos")

	if bd.Stock.inventario[Ropa{nombre: nombre, talla: talla}] == 0 {
		return errors.New("no se ha podido eliminar la ropa")
	}

	bd.Stock.inventario[Ropa{nombre: nombre, talla: talla}]--

	return nil
}

func InsertarVenta(bd *BD, nombre string, talla Talla, fecha ...time.Time) error {
	log.GetLogger().Info().Msg("Insertando venta en la base de datos")

	venta, err := NuevaVenta(map[Ropa]int{{nombre, talla}: 1}, bd.Stock, fecha...)
	if err != nil {
		return err
	}

	bd.Ventas = append(bd.Ventas, venta)

	return nil
}

func ObtenerVentas(bd *BD, nombre string, talla Talla, fecha ...time.Time) ([]Venta, error) {
	log.GetLogger().Info().Msg("Obteniendo Ventas de la base de datos")

	var Ventas []Venta
	for _, venta := range bd.Ventas {
		if _, ok := venta.itemsVendidos[Ropa{nombre, talla}]; ok {
			if len(fecha) == 0 || venta.fecha == fecha[0] {
				Ventas = append(Ventas, venta)
			}
		}
	}

	if len(Ventas) == 0 {
		return nil, errors.New("no se ha podido obtener la venta")
	}

	return Ventas, nil
}

func EliminarVenta(bd *BD, nombre string, talla Talla, fecha time.Time) error {
	log.GetLogger().Info().Msg("Eliminando venta de la base de datos")

	var index int = -1
	for i, venta := range bd.Ventas {
		if _, ok := venta.itemsVendidos[Ropa{nombre, talla}]; ok && venta.fecha == fecha {
			index = i
			break
		}
	}

	if index < 0 {
		return errors.New("no se ha podido eliminar la venta")
	}

	bd.Ventas = append(bd.Ventas[:index], bd.Ventas[index+1:]...)

	return nil
}
