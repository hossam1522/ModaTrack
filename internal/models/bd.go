package models

import (
	"ModaTrack/internal/log"
	"errors"
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
