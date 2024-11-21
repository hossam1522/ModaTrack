package models

import (
	"errors"
	"time"
)

type Venta struct {
	fecha         time.Time
	itemsVendidos map[Ropa]int
}

/* func NuevaVenta(itemsVendidos map[Ropa]int, fecha ...time.Time) (Venta, error) {
	fechaActual := time.Now()
	fechaLimite := fechaActual.AddDate(-2, 0, 0)

	var fechaVenta time.Time
	if len(fecha) > 0 {
		fechaVenta = fecha[0]
		if fechaVenta.Before(fechaLimite) {
			return Venta{}, errors.New("la fecha no puede ser anterior a dos años")
		}
	} else {
		fechaVenta = fechaActual
	}

	return Venta{
		fecha:         fechaVenta,
		itemsVendidos: itemsVendidos,
	}, nil
}
*/

func NuevaVenta(itemsVendidos map[Ropa]int, inventario *Stock, fecha ...time.Time) (Venta, error) {
	fechaActual := time.Now()
	fechaLimite := fechaActual.AddDate(-2, 0, 0)

	var fechaVenta time.Time
	if len(fecha) > 0 {
		fechaVenta = fecha[0]
		if fechaVenta.Before(fechaLimite) {
			return Venta{}, errors.New("la fecha no puede ser anterior a dos años")
		}
	} else {
		fechaVenta = fechaActual
	}

	// Si stock está vacío, no hacer nada. Si no lo está, actualizar
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

		return Venta{}, errors.New("no hay stock disponible")
	}

	return Venta{}, errors.New("stock vacío, no puedes vender")

}
