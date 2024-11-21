package models

import (
	"testing"
)

func TestStockVacio(t *testing.T) {
	stock := NewStock()
	_, err := stock.GetRopa("camisa")
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

func TestStockVacioTalla(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: L}] = 1
	_, err := stock.GetRopaTalla("camisa", M)
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

func TestStockConPrenda(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: M}] = 1
	ropa, err := stock.GetRopa("camisa")
	if err != nil {
		t.Error("Se esperaba una prenda")
	}

	switch prenda := ropa.(type) {
	case Ropa:
		if prenda.nombre != "camisa" {
			t.Error("La prenda no es la esperada")
		}
	case []Ropa:
		t.Error("Se esperaba una sola prenda, pero se devolvió un slice")
	default:
		t.Error("El resultado tiene un tipo inesperado")
	}
}

func TestStockConPrendaTalla(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: M}] = 1
	ropa, err := stock.GetRopaTalla("camisa", M)
	if err != nil {
		t.Error("Se esperaba una prenda")
	}
	if ropa.nombre != "camisa" {
		t.Error("La prenda no es la esperada")
	}
	if ropa.talla != M {
		t.Error("La talla no es la esperada")
	}
}

func TestStockConPrendaDuplicada(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: M}] = 1
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: L}] = 1

	ropas, _ := stock.GetRopa("camisa")

	// Asegurar que es un slice
	prendas, ok := ropas.([]Ropa)
	if !ok {
		t.Errorf("Se esperaba un slice de Ropa, pero se obtuvo %T", ropas)
	}

	// Validaciones
	if len(prendas) != 2 {
		t.Errorf("Se esperaban 2 prendas, pero se obtuvieron %d", len(prendas))
	}

	if prendas[0].talla != M || prendas[1].talla != L {
		t.Error("Las tallas no son las esperadas")
	}

	if prendas[0].nombre != "camisa" || prendas[1].nombre != "camisa" {
		t.Error("Los nombres no son los esperados")
	}
}

func TestCantidadStockPrenda(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: M}] = 5
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: L}] = 10
	stock.inventario[Ropa{nombre: "pantalones", precio: 20, talla: M}] = 3
	stock.inventario[Ropa{nombre: "pantalones", precio: 20, talla: L}] = 7

	camisas := stock.GetStock("camisa")
	if camisas != 15 {
		t.Errorf("Se esperaban 15 camisas, pero se obtuvieron %d", camisas)
	}
}
