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
	stock.inventario[Ropa{nombre: "camisa", talla: L}] = 1
	_, err := stock.GetRopaTalla("camisa", M)
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

func TestStockConPrenda(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", talla: M}] = 1
	ropa, err := stock.GetRopa("camisa")
	if err != nil {
		t.Error("Se esperaba un elemento")
	}
	if len(ropa) != 1 {
		t.Errorf("Se esperaba un elemento, se encontraron %d", len(ropa))
	}
}

func TestStockConPrendaTalla(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", talla: M}] = 1
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
	stock.inventario[Ropa{nombre: "camisa", talla: M}] = 1
	stock.inventario[Ropa{nombre: "camisa", talla: L}] = 1

	ropas, _ := stock.GetRopa("camisa")

	if len(ropas) != 2 {
		t.Errorf("Se esperaban 2 prendas, pero se obtuvieron %d", len(ropas))
	}

	if !(ropas[0].talla == M && ropas[1].talla == L) && !(ropas[0].talla == L && ropas[1].talla == M) {
		t.Error("Las tallas no son las esperadas")
	}

	if ropas[0].nombre != "camisa" || ropas[1].nombre != "camisa" {
		t.Error("Los nombres no son los esperados")
	}
}

func TestCantidadStockPrenda(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", talla: M}] = 5
	stock.inventario[Ropa{nombre: "camisa", talla: L}] = 10
	stock.inventario[Ropa{nombre: "pantalones", talla: M}] = 3
	stock.inventario[Ropa{nombre: "pantalones", talla: L}] = 7

	camisas := stock.GetStock("camisa")
	if camisas != 15 {
		t.Errorf("Se esperaban 15 camisas, pero se obtuvieron %d", camisas)
	}
}

func TestCantidadStockPrendaTalla(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", talla: M}] = 5
	stock.inventario[Ropa{nombre: "camisa", talla: L}] = 10
	stock.inventario[Ropa{nombre: "pantalones", talla: M}] = 3
	stock.inventario[Ropa{nombre: "pantalones", talla: L}] = 7

	camisaM := stock.GetStock("camisa", M)
	if camisaM != 5 {
		t.Errorf("Se esperaban 5 camisas talla M, pero se obtuvieron %d", camisaM)
	}

	camisaL := stock.GetStock("camisa", L)
	if camisaL != 10 {
		t.Errorf("Se esperaban 10 camisas talla L, pero se obtuvieron %d", camisaL)
	}
}

func TestCantidadStockPrendaVariasTallas(t *testing.T) {
	stock := NewStock()
	stock.inventario[Ropa{nombre: "camisa", talla: M}] = 5
	stock.inventario[Ropa{nombre: "camisa", talla: L}] = 10
	stock.inventario[Ropa{nombre: "pantalones", talla: M}] = 3
	stock.inventario[Ropa{nombre: "pantalones", talla: L}] = 7
	stock.inventario[Ropa{nombre: "pantalones", talla: XL}] = 2

	pantalones := stock.GetStock("pantalones", L, XL)
	if pantalones != 9 {
		t.Errorf("Se esperaban 9 pantalones talla L o XL, pero se obtuvieron %d", pantalones)
	}
}
