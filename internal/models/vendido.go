package models

// ItemVendido representa un producto y la cantidad vendida
type ProductoVendido struct {
	producto    *Ropa
	cantidad    int
	precioVenta int
}
