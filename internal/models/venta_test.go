package models

import "testing"

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
