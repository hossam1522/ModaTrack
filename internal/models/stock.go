package models

import "errors"

type Stock struct {
	inventario map[Ropa]int
}

func NewStock() *Stock {
	return &Stock{inventario: make(map[Ropa]int)}
}

// Método para buscar una prenda por nombre
func (s *Stock) GetRopa(nombre string) (interface{}, error) {
	var prendas []Ropa

	for ropa := range s.inventario {
		if ropa.nombre == nombre {
			prendas = append(prendas, ropa)
		}
	}

	switch len(prendas) {
	case 0:
		return Ropa{}, errors.New("no se ha encontrado la prenda")
	case 1:
		return prendas[0], nil
	default:
		return prendas, nil
	}
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
