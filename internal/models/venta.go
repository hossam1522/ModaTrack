package models

import "time"

type Venta struct {
	fecha         time.Time
	itemsVendidos []ProductoVendido
}

func NuevaVenta(itemsVendidos []ProductoVendido) Venta {
	return Venta{
		fecha:         time.Now(),
		itemsVendidos: itemsVendidos,
	}
}
