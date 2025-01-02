package models

import (
	"ModaTrack/internal/log"
	"errors"
)

type Stock struct {
	inventario map[Ropa]int `json:"inventario"`
}

var ErrPrendaNoEncontrada = errors.New("no existe ninguna prenda en el inventario con esas características")

func NewStock() *Stock {
	return &Stock{inventario: make(map[Ropa]int)}
}

func (s *Stock) GetRopa(nombre string) ([]Ropa, error) {
	log.GetLogger().Info().Msg("Buscando ropa por nombre en el inventario")
	var prendas []Ropa

	for ropa := range s.inventario {
		if ropa.nombre == nombre && s.inventario[ropa] > 0 {
			prendas = append(prendas, ropa)
		}
	}

	if len(prendas) == 0 {
		return nil, ErrPrendaNoEncontrada
	}

	return prendas, nil
}

func (s *Stock) GetRopaTalla(nombre string, talla Talla) (Ropa, error) {
	log.GetLogger().Info().Msg("Buscando ropa por nombre y talla en el inventario")
	for ropa := range s.inventario {
		if ropa.nombre == nombre && ropa.talla == talla {
			return ropa, nil
		}
	}
	return Ropa{}, ErrPrendaNoEncontrada
}

func (s *Stock) GetStock(nombre string, tallas ...Talla) int {
	log.GetLogger().Info().Msg("Calculando stock de ropa en el inventario")
	var stock int = 0
	for ropa, cantidad := range s.inventario {
		if ropa.nombre == nombre {
			if len(tallas) == 0 {
				stock += cantidad
			} else {
				for _, talla := range tallas {
					if ropa.talla == talla {
						stock += cantidad
						break
					}
				}
			}
		}
	}
	return stock
}
