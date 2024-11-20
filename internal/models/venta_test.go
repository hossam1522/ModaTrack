package models

import (
	"testing"
	"time"
)

// Test para añadir una venta sin fecha
func TestVenta(t *testing.T) {
	ropa := map[Ropa]int{
		{nombre: "camisa", precio: 10, talla: M}: 1,
	}
	venta := NuevaVenta(ropa)
	if len(venta.itemsVendidos) != 1 {
		t.Error("Se esperaba un item vendido")
	}
}

// Test para comprobar que la fecha es correcta si no se especifica
func TestVentaFechaActual(t *testing.T) {
	ropa := map[Ropa]int{
		{nombre: "camisa", precio: 10, talla: M}: 1,
	}
	venta := NuevaVenta(ropa)
	// Si la diferencia es de más de un segundo, se considera incorrecto
	if time.Until(venta.fecha) > time.Second {
		t.Error("La fecha no es la actual")
	}
}

// Test para comprobar que se pueden introducir ventas
// anterores con fechas pasadas
func TestVentaFechaPasada(t *testing.T) {
	ropa := map[Ropa]int{
		{nombre: "camisa", precio: 10, talla: M}: 1,
	}
	fecha := time.Now().AddDate(-1, 0, 0)
	venta := NuevaVenta(ropa, fecha)
	if venta.fecha != fecha {
		t.Errorf("La fecha no es la esperada: %v", venta.fecha)
	}
}
