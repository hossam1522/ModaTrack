package models

import (
	"ModaTrack/internal/log"
	"errors"
	"time"
)

type BD struct {
	stock  *Stock
	ventas []Venta
}

func NewBD() *BD {
	return &BD{
		stock:  NewStock(),
		ventas: []Venta{},
	}
}

func (bd *BD) InsertarRopa(nombre string, talla Talla, cantidad int) error {
	log.GetLogger().Info().Msg("Insertando ropa en la base de datos")

	cantidadActual := bd.stock.inventario[Ropa{nombre: nombre, talla: talla}]
	bd.stock.inventario[Ropa{nombre: nombre, talla: talla}] = cantidadActual + cantidad

	if bd.stock.inventario[Ropa{nombre: nombre, talla: talla}] != cantidadActual+cantidad {
		return errors.New("no se ha podido insertar la ropa")
	}

	return nil
}

func (bd *BD) EliminarRopa(nombre string, talla Talla) error {
	log.GetLogger().Info().Msg("Eliminando ropa de la base de datos")

	if bd.stock.inventario[Ropa{nombre: nombre, talla: talla}] == 0 {
		return errors.New("no se ha podido eliminar la ropa")
	}

	bd.stock.inventario[Ropa{nombre: nombre, talla: talla}]--

	return nil
}

func (bd *BD) InsertarVenta(nombre string, talla Talla, fecha ...time.Time) error {
	log.GetLogger().Info().Msg("Insertando venta en la base de datos")

	venta, err := NuevaVenta(map[Ropa]int{{nombre, talla}: 1}, bd.stock, fecha...)
	if err != nil {
		return err
	}

	bd.ventas = append(bd.ventas, venta)

	return nil
}

func (bd *BD) ObtenerVentas(nombre string, talla Talla, fecha ...time.Time) []Venta {
	log.GetLogger().Info().Msg("Obteniendo ventas de la base de datos")

	var ventas []Venta
	for _, venta := range bd.ventas {
		if _, ok := venta.itemsVendidos[Ropa{nombre, talla}]; ok {
			if len(fecha) == 0 || venta.fecha == fecha[0] {
				ventas = append(ventas, venta)
			}
		}
	}

	return ventas
}
