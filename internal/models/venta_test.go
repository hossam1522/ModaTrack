package models

import (
	"testing"
	"time"
)

func TestVentaStock(t *testing.T) {
	stock := NewStock()
	ropa := Ropa{"camisa", 10, M}
	stock.inventario[ropa] = 1
	_, err := NuevaVenta(map[Ropa]int{ropa: 1}, stock)
	if stock.inventario[ropa] != 0 {
		t.Error("El stock no se ha reducido")
	}
	if err != nil {
		t.Error("No se esperaba un error")
	}
}

func TestVentaStockVacio(t *testing.T) {
	stock := NewStock()
	_, err := NuevaVenta(map[Ropa]int{{"camisa", 10, M}: 1}, stock)
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

func TestVentaStockInsuficiente(t *testing.T) {
	stock := NewStock()
	ropa := Ropa{"camisa", 10, M}
	stock.inventario[ropa] = 0
	_, err := NuevaVenta(map[Ropa]int{ropa: 1}, stock)
	if err == nil {
		t.Error("Se esperaba un error")
	}
}

func TestVentaFechaActual(t *testing.T) {
	stock := NewStock()
	ropa := Ropa{"camisa", 10, M}
	stock.inventario[ropa] = 1
	venta, _ := NuevaVenta(map[Ropa]int{ropa: 1}, stock)
	if time.Until(venta.fecha) > time.Second {
		t.Error("La fecha no es la actual")
	}
}

func TestVentaFechaPasada(t *testing.T) {
	stock := NewStock()
	ropa := Ropa{"camisa", 10, M}
	stock.inventario[ropa] = 1
	fecha := time.Now().AddDate(-1, 0, 0)
	venta, _ := NuevaVenta(map[Ropa]int{ropa: 1}, stock, fecha)
	if venta.fecha != fecha {
		t.Errorf("La fecha no es la esperada: %v", venta.fecha)
	}
}

func TestVentaFechaAnteriorDosAños(t *testing.T) {
	stock := NewStock()
	ropa := Ropa{"camisa", 10, M}
	stock.inventario[ropa] = 1
	fecha := time.Now().AddDate(-2, 0, 0)
	_, err := NuevaVenta(map[Ropa]int{ropa: 1}, stock, fecha)
	if err == nil {
		t.Errorf("No se esperaba que la venta fuera exitosa")
	}
}

// Test para comprobar cuál ha sido el producto más vendido
func TestMasVendido(t *testing.T) {
	vectorVentas := []Venta{
		{fecha: time.Now(), itemsVendidos: map[Ropa]int{{"camisa", 10, M}: 1}},
		{fecha: time.Now(), itemsVendidos: map[Ropa]int{{"camisa", 10, M}: 2}},
		{fecha: time.Now(), itemsVendidos: map[Ropa]int{{"pantalón", 20, M}: 1}},
	}
	productoMasVendido, _ := MasVendido(vectorVentas)
	if productoMasVendido.nombre != "camisa" {
		t.Errorf("El producto más vendido no es el esperado: %s", productoMasVendido.nombre)
	}
}
