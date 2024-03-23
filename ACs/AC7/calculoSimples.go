package main

import (
    "fmt"
)

func main() {
    var v1, v2 int
    var v3 float64

    fmt.Println("Digite três valores inteiros:")
    fmt.Scanln(&v1, &v2, &v3)

    var v4, v5 int
    var v6 float64

    fmt.Println("Digite mais três valores inteiros:")
    fmt.Scanln(&v4, &v5, &v6)

    resultado := float64(v2)*v3 + float64(v5)*v6
    fmt.Printf("O valor a ser pago é de: R$ %.2f\n", resultado)
}
