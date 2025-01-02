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

func (bd *BD) ObtenerPrenda(nombre string) ([]Ropa, error) {
	log.GetLogger().Info().Msg("Obteniendo prenda de la base de datos")

	var prendas []Ropa
	for ropa := range bd.stock.inventario {
		if ropa.nombre == nombre && bd.stock.inventario[ropa] > 0 {
			prendas = append(prendas, ropa)
		}
	}

	if len(prendas) == 0 {
		return nil, errors.New("no se ha podido obtener la prenda")
	}

	return prendas, nil
}

func (bd *BD) ObtenerPrendaTalla(nombre string, talla Talla) (Ropa, error) {
	log.GetLogger().Info().Msg("Obteniendo prenda de la base de datos")

	for ropa := range bd.stock.inventario {
		if ropa.nombre == nombre && ropa.talla == talla && bd.stock.inventario[ropa] > 0 {
			return ropa, nil
		}
	}

	return Ropa{}, errors.New("no se ha podido obtener la prenda")
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

func (bd *BD) EliminarVenta(nombre string, talla Talla, fecha time.Time) error {
	log.GetLogger().Info().Msg("Eliminando venta de la base de datos")

	var index int = -1
	for i, venta := range bd.ventas {
		if _, ok := venta.itemsVendidos[Ropa{nombre, talla}]; ok && venta.fecha == fecha {
			index = i
			break
		}
	}

	if index < 0 {
		return errors.New("no se ha podido eliminar la venta")
	}

	bd.ventas = append(bd.ventas[:index], bd.ventas[index+1:]...)

	return nil
}
