package main

import (
    "fmt"
)

func main() {
    valor1 := 7
    valor2 := 9
    x := soma(valor1, valor2)
    s := subtracao(valor1, valor2)
    y := multiplicacao(valor1, valor2)
    z := divisao(valor1, valor2)
    fmt.Println("Soma:", x)
    fmt.Println("Subtração:", s)
    fmt.Println("Multiplicação:", y)
    fmt.Println("Divisão:", z)
}

func soma(i ...int) int {
    total := 0
    for _, v := range i {
        total += v
    }
    return total
}

func subtracao(i ...int) int {
    if len(i) == 0 {
        return 0 // Caso nenhum argumento seja passado, retorna 0
    }
    total := i[0]
    for _, v := range i[1:] {
        total -= v
    }
    return total
}

func multiplicacao(i ...int) int {
    total := 1
    for _, v := range i {
        total *= v
    }
    return total
}

func divisao(i ...int) int {
    if len(i) == 0 {
        return 0 // Caso nenhum argumento seja passado, retorna 0
    }
    total := i[0]
    for _, v := range i[1:] {
        if v == 0 {
            return 0 // Evita divisão por zero
        }
        total /= v
    }
    return total
}
