package models

import (
	"testing"
	"time"
)

func TestInsertarRopaBD(t *testing.T) {
	bd := NewBD()
	err := bd.InsertarRopa("camisa", M, 1)
	if err != nil {
		t.Error("Se esperaba un error")
	}
}

func TestObtenerPrenda(t *testing.T) {
	bd := NewBD()
	bd.InsertarRopa("camisa", M, 1)
	bd.InsertarRopa("camisa", L, 1)

	prendas, err := bd.ObtenerPrenda("camisa")
	if err != nil {
		t.Error("Se esperaban prendas")
	}
	if len(prendas) != 2 {
		t.Errorf("Se esperaban 2 prendas, pero se obtuvieron %d", len(prendas))
	}
}

func TestObtenerPrendaTalla(t *testing.T) {
	bd := NewBD()
	bd.InsertarRopa("camisa", M, 1)
	bd.InsertarRopa("camisa", L, 1)

	prenda, err := bd.ObtenerPrendaTalla("camisa", M)
	if err != nil {
		t.Error("Se esperaba una prenda")
	}
	if prenda.nombre != "camisa" {
		t.Error("La prenda no es la esperada")
	}
	if prenda.talla != M {
		t.Error("La talla no es la esperada")
	}
}

func TestEliminarPrenda(t *testing.T) {
	bd := NewBD()
	bd.InsertarRopa("camisa", M, 1)
	bd.InsertarRopa("camisa", L, 1)

	err := bd.EliminarRopa("camisa", M)
	if err != nil {
		t.Error("Se esperaba un error")
	}

	prendas, _ := bd.ObtenerPrenda("camisa")
	if len(prendas) != 1 {
		t.Errorf("Se esperaba una prenda, pero se obtuvieron %d", len(prendas))
	}
}

func TestEliminarPrendaNumerosas(t *testing.T) {
	bd := NewBD()
	bd.InsertarRopa("camisa", M, 2)
	bd.InsertarRopa("camisa", L, 1)

	err := bd.EliminarRopa("camisa", M)
	if err != nil {
		t.Error("Se esperaba un error")
	}

	prendas, _ := bd.ObtenerPrenda("camisa")
	if len(prendas) != 2 {
		t.Errorf("Se esperaban dos prendas, pero se obtuvieron %d", len(prendas))
	}
}

func TestEliminarPrendaInexistente(t *testing.T) {
	bd := NewBD()
	err := bd.EliminarRopa("camisa", M)
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

func TestVenta(t *testing.T) {
	bd := NewBD()
	bd.InsertarRopa("camisa", M, 1)
	bd.InsertarRopa("camisa", L, 1)

	err := bd.InsertarVenta("camisa", M)
	if err != nil {
		t.Error("Se esperaba un error")
	}

	prendas, _ := bd.ObtenerPrenda("camisa")
	if len(prendas) != 1 {
		t.Errorf("Se esperaba una prenda, pero se obtuvieron %d", len(prendas))
	}
}

func TestVentaInexistente(t *testing.T) {
	bd := NewBD()
	err := bd.InsertarVenta("camisa", M)
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

func TestVentaStockInsuficienteBD(t *testing.T) {
	bd := NewBD()
	bd.InsertarRopa("camisa", M, 0)

	err := bd.InsertarVenta("camisa", M)
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

func TestVentaFecha(t *testing.T) {
	bd := NewBD()
	bd.InsertarRopa("camisa", M, 1)
	fecha := time.Now().AddDate(-1, 0, 0)
	err := bd.InsertarVenta("camisa", M, fecha)
	if err != nil {
		t.Error("Se esperaba una venta")
	}
	venta := bd.Ventas[0]
	if venta.fecha != fecha {
		t.Errorf("La fecha no es la esperada: %v", venta.fecha)
	}
}

func TestObtenerVentas(t *testing.T) {
	bd := NewBD()
	bd.InsertarRopa("camisa", M, 5)
	bd.InsertarRopa("pantalon", L, 4)
	bd.InsertarVenta("camisa", M)
	bd.InsertarVenta("camisa", M)
	bd.InsertarVenta("pantalon", L)
	ventas, err := bd.ObtenerVentas("camisa", M)
	if err != nil {
		t.Error("Se esperaban ventas")
	}
	if len(ventas) != 2 {
		t.Errorf("Se esperaban dos ventas, pero se obtuvieron %d", len(ventas))
	}
}

func TestObtenerVentasFecha(t *testing.T) {
	bd := NewBD()
	bd.InsertarRopa("camisa", M, 5)
	bd.InsertarRopa("pantalon", L, 4)
	fecha := time.Now().AddDate(-1, 0, 0)
	bd.InsertarVenta("camisa", M, fecha)
	bd.InsertarVenta("camisa", M)
	bd.InsertarVenta("pantalon", L)
	ventas, err := bd.ObtenerVentas("camisa", M, fecha)
	if err != nil {
		t.Error("Se esperaban ventas")
	}
	if len(ventas) != 1 {
		t.Errorf("Se esperaba una venta, pero se obtuvieron %d", len(ventas))
	}
}

func TestEliminarVenta(t *testing.T) {
	bd := NewBD()
	bd.InsertarRopa("camisa", M, 5)
	fecha := time.Now().AddDate(-1, 0, 0)
	bd.InsertarVenta("camisa", M, fecha)
	bd.InsertarVenta("camisa", M)
	bd.InsertarVenta("camisa", M)
	bd.EliminarVenta("camisa", M, fecha)
	if len(bd.Ventas) != 2 {
		t.Errorf("Se esperaban dos ventas, pero se obtuvieron %d", len(bd.Ventas))
	}
}
