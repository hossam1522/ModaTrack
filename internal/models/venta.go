package models

import (
	"errors"
	"time"
)

var ErrFechaAnteriorDosAños = errors.New("la fecha no puede ser anterior a dos años")
var ErrNoStock = errors.New("no hay stock disponible")
var ErrStockVacio = errors.New("el stock está vacío, no puedes vender")

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

	if inventario != nil {
		for item, cantidad := range itemsVendidos {
			if inventario.Existe(item) {
				err := inventario.Restar(item, cantidad)
				if err != nil {
					return Venta{}, err
				}
				return Venta{
					fecha:         fechaVenta,
					itemsVendidos: itemsVendidos,
				}, nil
			}
		}

		return Venta{}, ErrNoStock
	}

	return Venta{}, ErrStockVacio

}
