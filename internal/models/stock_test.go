package models

import (
	"testing"
)

// Comprobar que no puedo obtener una prenda que no existe
func TestStock(t *testing.T) {
	stock := NewStock()
	_, err := stock.GetRopa("camisa")
	if err == nil {
		t.Error("Se esperaba un error")
	}
}
