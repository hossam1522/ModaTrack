package models

import "testing"

func TestInsertarRopaBD(t *testing.T) {
	bd := NewBD()
	err := bd.InsertarRopa("camisa", M, 1)
	if err != nil {
		t.Error("Se esperaba un error")
	}
}
