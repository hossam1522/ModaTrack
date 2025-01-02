package models

import (
	"ModaTrack/internal/log"
	"errors"
	"time"
)

type BD struct {
	Stock  *Stock
	Ventas []Venta
}

func NewBD() *BD {
	return &BD{
		Stock:  NewStock(),
		Ventas: []Venta{},
	}
}

func (bd *BD) InsertarRopa(nombre string, talla Talla, cantidad int) error {
	log.GetLogger().Info().Msg("Insertando ropa en la base de datos")

	cantidadActual := bd.Stock.inventario[Ropa{nombre: nombre, talla: talla}]
	bd.Stock.inventario[Ropa{nombre: nombre, talla: talla}] = cantidadActual + cantidad

	if bd.Stock.inventario[Ropa{nombre: nombre, talla: talla}] != cantidadActual+cantidad {
		return errors.New("no se ha podido insertar la ropa")
	}

	return nil
}

func (bd *BD) ObtenerPrenda(nombre string) ([]Ropa, error) {
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

func (bd *BD) ObtenerPrendaTalla(nombre string, talla Talla) (Ropa, error) {
	log.GetLogger().Info().Msg("Obteniendo prenda de la base de datos")

	for ropa := range bd.Stock.inventario {
		if ropa.nombre == nombre && ropa.talla == talla && bd.Stock.inventario[ropa] > 0 {
			return ropa, nil
		}
	}

	return Ropa{}, errors.New("no se ha podido obtener la prenda")
}

func (bd *BD) EliminarRopa(nombre string, talla Talla) error {
	log.GetLogger().Info().Msg("Eliminando ropa de la base de datos")

	if bd.Stock.inventario[Ropa{nombre: nombre, talla: talla}] == 0 {
		return errors.New("no se ha podido eliminar la ropa")
	}

	bd.Stock.inventario[Ropa{nombre: nombre, talla: talla}]--

	return nil
}

func (bd *BD) InsertarVenta(nombre string, talla Talla, fecha ...time.Time) error {
	log.GetLogger().Info().Msg("Insertando venta en la base de datos")

	venta, err := NuevaVenta(map[Ropa]int{{nombre, talla}: 1}, bd.Stock, fecha...)
	if err != nil {
		return err
	}

	bd.Ventas = append(bd.Ventas, venta)

	return nil
}

func (bd *BD) ObtenerVentas(nombre string, talla Talla, fecha ...time.Time) ([]Venta, error) {
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

func (bd *BD) EliminarVenta(nombre string, talla Talla, fecha time.Time) error {
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
