package models

type Ropa struct {
	nombre string
	precio float64
	talla  Talla
}

func NewRopa(nombre string, precio float64, talla Talla) Ropa {
	return Ropa{nombre: nombre, precio: precio, talla: talla}
}
