package models

import (
	"ModaTrack/internal/log"
	"encoding/json"
	"errors"
	"time"
)

var ErrFechaAnteriorDosAños = errors.New("la fecha no puede ser anterior a dos años")
var ErrNoStock = errors.New("no hay stock disponible")
var ErrStockVacio = errors.New("el stock está vacío, no puedes vender")
var ErrInsuficientesPrendas = errors.New("no hay suficientes prendas en el inventario")

type Venta struct {
	fecha         time.Time
	itemsVendidos map[Ropa]int
}

func (v Venta) GetFecha() time.Time {
	return v.fecha
}

func (v Venta) GetItemsVendidos() map[Ropa]int {
	return v.itemsVendidos
}

func (v Venta) MarshalJSON() ([]byte, error) {

	type ItemVendido struct {
		Ropa     Ropa `json:"ropa"`
		Cantidad int  `json:"cantidad"`
	}

	var items []ItemVendido
	for ropa, cantidad := range v.itemsVendidos {
		items = append(items, ItemVendido{
			Ropa:     ropa,
			Cantidad: cantidad,
		})
	}

	return json.Marshal(&struct {
		Fecha         time.Time     `json:"fecha"`
		ItemsVendidos []ItemVendido `json:"items_vendidos"`
	}{
		Fecha:         v.fecha,
		ItemsVendidos: items,
	})
}

func (v *Venta) UnmarshalJSON(data []byte) error {

	type ItemVendido struct {
		Ropa     Ropa `json:"ropa"`
		Cantidad int  `json:"cantidad"`
	}

	aux := struct {
		Fecha         time.Time     `json:"fecha"`
		ItemsVendidos []ItemVendido `json:"items_vendidos"`
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	v.fecha = aux.Fecha
	v.itemsVendidos = make(map[Ropa]int)
	for _, item := range aux.ItemsVendidos {
		v.itemsVendidos[item.Ropa] = item.Cantidad
	}

	return nil
}

func NuevaVenta(itemsVendidos map[Ropa]int, inventario *Stock, fecha ...time.Time) (Venta, error) {
	log.GetLogger().Info().Msg("Creando nueva venta")
	fechaActual := time.Now()
	fechaLimite := fechaActual.AddDate(-2, 0, 0)

	var fechaVenta time.Time
	if len(fecha) > 0 {
		fechaVenta = fecha[0]
		if fechaVenta.Before(fechaLimite) {
			return Venta{}, ErrFechaAnteriorDosAños
		}
	} else {
		fechaVenta = fechaActual
	}

	if inventario == nil {
		return Venta{}, ErrStockVacio
	}

	for item, cantidadSolicitada := range itemsVendidos {
		stockDisponible := inventario.inventario[item]
		if stockDisponible == 0 {
			return Venta{}, ErrNoStock
		}
		if stockDisponible < cantidadSolicitada {
			return Venta{}, ErrInsuficientesPrendas
		}
		inventario.inventario[item] -= cantidadSolicitada
	}

	return Venta{
		fecha:         fechaVenta,
		itemsVendidos: itemsVendidos,
	}, nil
}

func MasVendido(ventas []Venta) (Ropa, int) {
	log.GetLogger().Info().Msg("Calculando prenda más vendida")
	ventasPorPrenda := make(map[Ropa]int)

	for _, venta := range ventas {
		for prenda, cantidad := range venta.itemsVendidos {
			ventasPorPrenda[prenda] += cantidad
		}
	}

	var prendaMasVendida Ropa
	var cantidadMasVendida int
	for prenda, cantidad := range ventasPorPrenda {
		if cantidad > cantidadMasVendida {
			prendaMasVendida = prenda
			cantidadMasVendida = cantidad
		}
	}

	return prendaMasVendida, cantidadMasVendida
}

func TotalVendido(ventas []Venta, nombre string, talla ...Talla) int {
	log.GetLogger().Info().Msg("Calculando total vendido de una prenda")
	var cantidadVendida int
	for _, venta := range ventas {
		for prenda, cantidad := range venta.itemsVendidos {
			if prenda.nombre == nombre {
				if len(talla) == 0 || (len(talla) > 0 && prenda.talla == talla[0]) {
					cantidadVendida += cantidad
				}
			}

		}
	}
	return cantidadVendida
}

func TallaMasVendida(ventas []Venta) Talla {
	log.GetLogger().Info().Msg("Calculando talla más vendida")
	tallasVendidas := make(map[Talla]int)
	for _, venta := range ventas {
		for prenda, cantidad := range venta.itemsVendidos {
			tallasVendidas[prenda.talla] += cantidad
		}
	}
	var tallaMasVendida Talla
	var cantidadMasVendida int
	for talla, cantidad := range tallasVendidas {
		if cantidad > cantidadMasVendida {
			tallaMasVendida = talla
			cantidadMasVendida = cantidad
		}
	}
	return tallaMasVendida
}
