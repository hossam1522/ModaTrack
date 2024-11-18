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

// ¿Qué pasa si hay más de una prenda con el mismo nombre?
// Debería devolver un vector con todas las prendas
func TestStockConPrendaDuplicada(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: M}] = 1
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: L}] = 1

	ropas, _ := stock.GetRopa("camisa")

	if _, ok := interface{}(ropas).([]Ropa); !ok {
		t.Errorf("Se esperaba un slice de Ropa, pero se obtuvo %T", ropas)
	}

	if len(ropas) != 2 {
		t.Error("Se esperaban dos prendas")
	}

	if ropas[0].talla != M || ropas[1].talla != L {
		t.Error("Las tallas no son las esperadas")
	}

	if ropas[0].nombre != "camisa" || ropas[1].nombre != "camisa" {
		t.Error("Los nombres no son los esperados")
	}
}
