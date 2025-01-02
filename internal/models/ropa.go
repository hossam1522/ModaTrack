package models

type Ropa struct {
	nombre string
	talla  Talla
}

func (r Ropa) GetTalla() Talla {
	return r.talla
}

func (r Ropa) GetNombre() string {
	return r.nombre
}
