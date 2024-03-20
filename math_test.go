package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma e invalido: Resultado %d. O valor esperado e %d", total, 30)
	}
}

func TestSub(t *testing.T) {
	total := Sub(15, 15)
	if total != 0 {
		t.Errorf("Resultado da soma e invalido: Resultado %d. O valor esperado e %d", total, 30)
	}
}

func TestMul(t *testing.T) {
	total := Mult(5, 5)
	if total != 25 {
		t.Errorf("Resultado da soma e invalido: Resultado %d. O valor esperado e %d", total, 25)
	}

}
