package models

import "errors"

type Stock struct {
	inventario map[Ropa]int
}

func NewStock() *Stock {
	return &Stock{inventario: make(map[Ropa]int)}
}

// Método para buscar una prenda por nombre y talla
func (s *Stock) GetRopa(nombre string) (Ropa, error) {
	for ropa := range s.inventario {
		if ropa.nombre == nombre {
			return ropa, nil
		}
	}
	return Ropa{}, errors.New("no se ha encontrado la prenda")
}

// Método para buscar una prenda por nombre y talla
func (s *Stock) GetRopaTalla(nombre string, talla Talla) (Ropa, error) {
	for ropa := range s.inventario {
		if ropa.nombre == nombre && ropa.talla == talla {
			return ropa, nil
		}
	}
	return Ropa{}, errors.New("no se ha encontrado la prenda")
}
