package models

import (
	"testing"
	"time"
)

// Test para comprobar que el stock se reduce al realizar una venta
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

// Test para comprobar que no se puede realizar una venta con stock vacío
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

// Test para comprobar que la fecha es correcta si no se especifica
func TestVentaFechaActual(t *testing.T) {
	stock := NewStock()
	ropa := Ropa{"camisa", 10, M}
	stock.inventario[ropa] = 1
	venta, _ := NuevaVenta(map[Ropa]int{ropa: 1}, stock)
	// Si la diferencia es de más de un segundo, se considera incorrecto
	if time.Until(venta.fecha) > time.Second {
		t.Error("La fecha no es la actual")
	}
}

// Test para comprobar que se pueden introducir ventas
// anteriores con fechas pasadas
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

// Test para comprobar que no se pueden introducir ventas
// anterores a dos años
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
