package models

import "time"

type Venta struct {
	fecha         time.Time
	itemsVendidos []ProductoVendido
}
