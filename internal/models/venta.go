package models

import "time"

type Venta struct {
	Fecha         time.Time
	ItemsVendidos []ProductoVendido
}
