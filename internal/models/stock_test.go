package models

import (
	"testing"
)

// Comprobar que no puedo obtener una prenda que no existe
func TestStockVacio(t *testing.T) {
	stock := NewStock()
	_, err := stock.GetRopa("camisa")
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

// Comprobar que puedo obtener una prenda que existe
func TestStockConPrenda(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: M}] = 1
	ropa, err := stock.GetRopa("camisa")
	if err != nil {
		t.Error("Se esperaba una prenda")
	}
	if ropa.nombre != "camisa" {
		t.Error("La prenda no es la esperada")
	}
}

// Comprobar que puedo buscar una prenda por nombre y talla
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
