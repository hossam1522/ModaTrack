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

// Comprobar que no puedo obtener una prenda con una talla que no existe
func TestStockVacioTalla(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", precio: 10, talla: L}] = 1
	_, err := stock.GetRopaTalla("camisa", M)
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
	// Verificar el tipo de retorno
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
