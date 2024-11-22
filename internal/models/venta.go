package models

import (
	"errors"
	"time"
)

var ErrFechaAnteriorDosAños = errors.New("la fecha no puede ser anterior a dos años")
var ErrNoStock = errors.New("no hay stock disponible")
var ErrStockVacio = errors.New("el stock está vacío, no puedes vender")
var ErrInsuficientesPrendas = errors.New("no hay suficientes prendas en el inventario")

type Venta struct {
	fecha         time.Time
	itemsVendidos map[Ropa]int
}

func NuevaVenta(itemsVendidos map[Ropa]int, inventario *Stock, fecha ...time.Time) (Venta, error) {
	fechaActual := time.Now()
	fechaLimite := fechaActual.AddDate(-2, 0, 0)

	var fechaVenta time.Time
	if len(fecha) > 0 {
		fechaVenta = fecha[0]
		if fechaVenta.Before(fechaLimite) {
			return Venta{}, ErrFechaAnteriorDosAños
		}
	} else {
		fechaVenta = fechaActual
	}

	if inventario == nil {
		return Venta{}, ErrStockVacio
	}

	for item, cantidadSolicitada := range itemsVendidos {
		stockDisponible := inventario.inventario[item]
		if stockDisponible == 0 {
			return Venta{}, ErrNoStock
		}
		if stockDisponible < cantidadSolicitada {
			return Venta{}, ErrInsuficientesPrendas
		}
		inventario.inventario[item] -= cantidadSolicitada
	}

	return Venta{
		fecha:         fechaVenta,
		itemsVendidos: itemsVendidos,
	}, nil
}

func MasVendido(ventas []Venta) (Ropa, int) {
	ventasPorPrenda := make(map[Ropa]int)

	for _, venta := range ventas {
		for prenda, cantidad := range venta.itemsVendidos {
			ventasPorPrenda[prenda] += cantidad
		}
	}

	var prendaMasVendida Ropa
	var cantidadMasVendida int
	for prenda, cantidad := range ventasPorPrenda {
		if cantidad > cantidadMasVendida {
			prendaMasVendida = prenda
			cantidadMasVendida = cantidad
		}
	}

	return prendaMasVendida, cantidadMasVendida
}
