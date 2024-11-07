package models

import "time"

type Venta struct {
	fecha         time.Time
	itemsVendidos map[Ropa]int
}

func NuevaVenta(itemsVendidos map[Ropa]int, fecha ...time.Time) Venta {
	fechaActual := time.Now()
	fechaLimite := fechaActual.AddDate(-2, 0, 0)

	var fechaVenta time.Time
	if len(fecha) > 0 {
		fechaVenta = fecha[0]
		if fechaVenta.Before(fechaLimite) {
			panic("la fecha de la venta no puede ser anterior a la fecha límite de dos años")
		}
	} else {
		fechaVenta = fechaActual
	}

	return Venta{
		fecha:         fechaVenta,
		itemsVendidos: itemsVendidos,
	}
}
