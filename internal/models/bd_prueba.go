package models

import "time"

func BDPrueba() *BD {
	bd_prueba := NewBD()
	bd_prueba.InsertarRopa("camisa", M, 1)
	bd_prueba.InsertarRopa("camisa", L, 5)
	bd_prueba.InsertarRopa("pantalon", M, 5)
	bd_prueba.InsertarRopa("pantalon", L, 1)
	bd_prueba.InsertarVenta("camisa", L)
	bd_prueba.InsertarVenta("camisa", L)

	// "2024-06-12 15:30:45"
	fecha, _ := time.Parse(time.RFC3339, "2024-06-12T15:30:45Z")
	bd_prueba.InsertarVenta("pantalon", M, fecha)
	bd_prueba.InsertarVenta("pantalon", M, fecha.AddDate(-1, 0, 0))

	return bd_prueba
}

var bd_prueba = BDPrueba()

func GetBDPrueba() *BD {
	return bd_prueba
}
