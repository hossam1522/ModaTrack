package models

import "time"

var fecha_prueba, _ = time.Parse(time.RFC3339, "2024-06-12T15:30:45Z")

func BDPrueba() BD {
	bd_prueba := BD{
		Stock: &Stock{
			inventario: map[Ropa]int{
				{nombre: "camisa", talla: M}:   1,
				{nombre: "camisa", talla: L}:   5,
				{nombre: "pantalon", talla: M}: 5,
				{nombre: "pantalon", talla: L}: 1,
			},
		},
		Ventas: []Venta{
			{
				itemsVendidos: map[Ropa]int{
					{nombre: "camisa", talla: L}: 1,
				},
				fecha: time.Now(),
			},
			{
				itemsVendidos: map[Ropa]int{
					{nombre: "camisa", talla: L}: 1,
				},
				fecha: time.Now(),
			},
			{
				itemsVendidos: map[Ropa]int{
					{nombre: "pantalon", talla: M}: 1,
				},
				fecha: time.Now(),
			},
			{
				itemsVendidos: map[Ropa]int{
					{nombre: "pantalon", talla: M}: 1,
				},
				fecha: fecha_prueba,
			},
		},
	}

	return bd_prueba
}

var bd_prueba = BDPrueba()

func GetBDPrueba() *BD {
	return &bd_prueba
}
