package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(3,2)
	resultado := 5

	if teste != resultado {
        t.Errorf("A soma de 3 + 2 deve ser 5, mas retornou %d", teste)
    }
}

func TestSub(t *testing.T) {
	teste := subtracao(3,2)
	resultado := 1

	if teste != resultado {
        t.Errorf("A subtracao de 3 - 2 deve ser 1, mas retornou %d", teste)
    }
}

func TestMulti(t *testing.T) {
	teste := multiplicacao(3,2)
	resultado := 6

	if teste != resultado {
        t.Errorf("A subtracao de 3 * 2 deve ser 6, mas retornou %d", teste)
    }
}

func TestDivi(t *testing.T) {
	teste := divisao(6,2)
	resultado := 3

	if teste != resultado {
        t.Errorf("A subtracao de 6 * 2 deve ser 3, mas retornou %d", teste)
    }
}
