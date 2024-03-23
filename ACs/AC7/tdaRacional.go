package main

import (
    "fmt"
)

type Fracao struct {
    numerador   int
    denominador int
}

func somar(frac1 Fracao, frac2 Fracao) Fracao {
    volta := Fracao{}
    volta.numerador = frac1.numerador*frac2.denominador + frac2.numerador*frac1.denominador
    volta.denominador = frac1.denominador * frac2.denominador
    return volta
}

func diminuir(frac1 Fracao, frac2 Fracao) Fracao {
    volta := Fracao{}
    volta.numerador = frac1.numerador*frac2.denominador - frac2.numerador*frac1.denominador
    volta.denominador = frac1.denominador * frac2.denominador
    return volta
}

func multiplicar(frac1 Fracao, frac2 Fracao) Fracao {
    volta := Fracao{}
    volta.numerador = frac1.numerador * frac2.numerador
    volta.denominador = frac1.denominador * frac2.denominador
    return volta
}

func dividir(frac1 Fracao, frac2 Fracao) Fracao {
    volta := Fracao{}
    volta.numerador = frac1.numerador * frac2.denominador
    volta.denominador = frac1.denominador * frac2.numerador
    return volta
}

func main() {
    var frac1, frac2 Fracao
    var operacao rune
    var N int

    fmt.Println("Quantidade de operações:")
    fmt.Scanln(&N)

    for i := 0; i < N; i++ {
        fmt.Printf("Operação %d: ", i+1)
        fmt.Scanf("%d / %d %c %d / %d\n", &frac1.numerador, &frac1.denominador, &operacao, &frac2.numerador, &frac2.denominador)

        var resultado Fracao
        switch operacao {
        case '+':
            resultado = somar(frac1, frac2)
        case '-':
            resultado = diminuir(frac1, frac2)
        case '*':
            resultado = multiplicar(frac1, frac2)
        case '/':
            resultado = dividir(frac1, frac2)
        }
        fmt.Printf("Resultado da operação %d: %d/%d\n", i+1, resultado.numerador, resultado.denominador)
    }
}
