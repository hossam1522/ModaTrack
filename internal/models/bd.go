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
