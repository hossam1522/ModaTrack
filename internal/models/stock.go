package models

import "errors"

type Stock struct {
	inventario map[Ropa]int
}

var ErrPrendaNoEncontrada = errors.New("no existe ninguna prenda en el inventario con esas características")

func NewStock() *Stock {
	return &Stock{inventario: make(map[Ropa]int)}
}

func (s *Stock) GetRopa(nombre string) (interface{}, error) {
	var prendas []Ropa

	for ropa := range s.inventario {
		if ropa.nombre == nombre {
			prendas = append(prendas, ropa)
		}
	}

	switch len(prendas) {
	case 0:
		return nil, ErrPrendaNoEncontrada
	case 1:
		return prendas[0], nil
	default:
		return prendas, nil
	}
}

func (s *Stock) GetRopaTalla(nombre string, talla Talla) (Ropa, error) {
	for ropa := range s.inventario {
		if ropa.nombre == nombre && ropa.talla == talla {
			return ropa, nil
		}
	}
	return Ropa{}, ErrPrendaNoEncontrada
}

func (s *Stock) GetStock(nombre string, talla ...Talla) int {
	var stock int = 0
	for ropa, cantidad := range s.inventario {
		if ropa.nombre == nombre {
			if len(talla) == 0 {
				stock += cantidad
			} else if ropa.talla == talla[0] {
				stock += cantidad
			}
		}
	}
	return stock
}
