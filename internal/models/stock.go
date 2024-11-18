package models

type Stock struct {
	inventario map[Ropa]int
}

func NewStock() *Stock {
	return &Stock{inventario: make(map[Ropa]int)}
}
