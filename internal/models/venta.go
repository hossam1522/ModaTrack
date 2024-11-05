package models

import "time"

type Venta struct {
	fecha         time.Time
	itemsVendidos map[Ropa]int
}

func NuevaVenta(itemsVendidos map[Ropa]int) Venta {
	return Venta{
		fecha:         time.Now(),
		itemsVendidos: itemsVendidos,
	}
}
