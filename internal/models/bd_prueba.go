package models

func BDPrueba() *BD {
	bd_prueba := NewBD()
	bd_prueba.InsertarRopa("camisa", M, 1)
	bd_prueba.InsertarRopa("camisa", L, 1)
	bd_prueba.InsertarRopa("pantalon", M, 1)
	bd_prueba.InsertarRopa("pantalon", L, 1)

	return bd_prueba
}

var bd_prueba = BDPrueba()

func GetBDPrueba() *BD {
	return bd_prueba
}
