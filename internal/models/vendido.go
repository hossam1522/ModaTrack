package models

// ItemVendido representa un producto y la cantidad vendida
type ProductoVendido struct {
	Producto    *Ropa
	Cantidad    int
	PrecioVenta int
}
