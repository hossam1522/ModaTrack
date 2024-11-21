package models

import (
	"testing"
	"time"
)

// Test para comprobar que el stock se reduce al realizar una venta
func TestVentaStock(t *testing.T) {
	stock := NewStock()
	ropa := NewRopa("camisa", 10, M)
	stock.inventario[ropa] = 1
	_, err := NuevaVenta(map[Ropa]int{ropa: 1}, stock)
	if stock.inventario[ropa] != 0 {
		t.Error("El stock no se ha reducido")
	}
	if err != nil {
		t.Error("No se esperaba un error")
	}
}

// Test para añadir una venta sin fecha
func TestVenta(t *testing.T) {
	ropa := map[Ropa]int{
		{nombre: "camisa", precio: 10, talla: M}: 1,
	}
	venta, _ := NuevaVenta(ropa, &Stock{})
	if len(venta.itemsVendidos) != 1 {
		t.Error("Se esperaba un item vendido")
	}
}

// Test para comprobar que la fecha es correcta si no se especifica
func TestVentaFechaActual(t *testing.T) {
	ropa := map[Ropa]int{
		{nombre: "camisa", precio: 10, talla: M}: 1,
	}
	venta, _ := NuevaVenta(ropa, &Stock{})
	// Si la diferencia es de más de un segundo, se considera incorrecto
	if time.Until(venta.fecha) > time.Second {
		t.Error("La fecha no es la actual")
	}
}

// Test para comprobar que se pueden introducir ventas
// anteriores con fechas pasadas
func TestVentaFechaPasada(t *testing.T) {
	ropa := map[Ropa]int{
		{nombre: "camisa", precio: 10, talla: M}: 1,
	}
	fecha := time.Now().AddDate(-1, 0, 0)
	venta, _ := NuevaVenta(ropa, &Stock{}, fecha)
	if venta.fecha != fecha {
		t.Errorf("La fecha no es la esperada: %v", venta.fecha)
	}
}

// Test para comprobar que no se pueden introducir ventas
// anterores a dos años
func TestVentaFechaAnteriorDosAños(t *testing.T) {
	ropa := map[Ropa]int{
		{nombre: "camisa", precio: 10, talla: M}: 1,
	}
	fecha := time.Now().AddDate(-2, 0, 0)
	_, err := NuevaVenta(ropa, &Stock{}, fecha)
	if err == nil {
		t.Errorf("No se esperaba que la venta fuera exitosa")
	}
}
