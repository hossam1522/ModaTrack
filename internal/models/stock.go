package models

import (
	logger "ModaTrack/internal/logger"
	"errors"
)

type Stock struct {
	inventario map[Ropa]int
}

var ErrPrendaNoEncontrada = errors.New("no existe ninguna prenda en el inventario con esas características")

func NewStock() *Stock {
	return &Stock{inventario: make(map[Ropa]int)}
}

func (s *Stock) GetRopa(nombre string) ([]Ropa, error) {
	logger.GetLogger().Info().Msg("Buscando ropa por nombre en el inventario")
	var prendas []Ropa

	for ropa := range s.inventario {
		if ropa.nombre == nombre {
			prendas = append(prendas, ropa)
		}
	}

	if len(prendas) == 0 {
		return nil, ErrPrendaNoEncontrada
	}

	return prendas, nil
}

func (s *Stock) GetRopaTalla(nombre string, talla Talla) (Ropa, error) {
	logger.GetLogger().Info().Msg("Buscando ropa por nombre y talla en el inventario")
	for ropa := range s.inventario {
		if ropa.nombre == nombre && ropa.talla == talla {
			return ropa, nil
		}
	}
	return Ropa{}, ErrPrendaNoEncontrada
}

func (s *Stock) GetStock(nombre string, tallas ...Talla) int {
	logger.GetLogger().Info().Msg("Calculando stock de ropa en el inventario")
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
