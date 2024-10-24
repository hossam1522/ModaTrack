package models

type Ropa struct {
	id              int
	nombre          string
	cantidadEnStock int
	precio          float64
	talla           Talla
}
