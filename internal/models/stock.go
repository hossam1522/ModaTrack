package models

import "errors"

type Stock struct {
	inventario map[Ropa]int
}

var ErrPrendaNoEncontrada = errors.New("no existe ninguna prenda en el inventario con esas características")
var ErrInsuficientesPrendas = errors.New("no hay suficientes prendas en el inventario")

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

func (s *Stock) Existe(prenda Ropa) bool {
	return s.inventario[prenda] > 0
}

func (s *Stock) Restar(prenda Ropa, cantidad int) error {
	if s.inventario[prenda] >= cantidad {
		s.inventario[prenda] -= cantidad
		return nil
	}
	return ErrInsuficientesPrendas
}
