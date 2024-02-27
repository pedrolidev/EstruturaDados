package main

import (
	"fmt"
	"math"
)

// #1

func number() {
	list := [10]int{3, 7, 9, 12, 17, 22, 23, 25, 32, 36}

	for i :=0; i < len(list); i++{
		fmt.Println(list[i])
	}
}

// #2

// Inversão de string
func upsideDown(input string) string {
	runes := []rune(input)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// Função para compilar a interação do usuário com o código
func main() {
	var input string
	fmt.Println("Insira uma string: ")
	fmt.Scanln(&input)

	// Chaamadq de função
	reverse := upsideDown(input)

	fmt.Println(("String reversa: ", reverse))
}

// #3

type Ponto struct {
	X, Y float64
}

func (p Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	ponto := Ponto{X: 3, Y: 4}

	distancia := ponto.DistanciaOrigem()

	fmt.Printf("Os pontos se distanciam (%.2f, %.2f) até a sua origem é: %.2f\n", ponto.X, ponto.Y, distancia)
}

// #4 (DÚVIDAS) -> Importação de pacotes e main

func AreaRetangulo(base, altura float64) float64 {
	return base * altura
}

func PerimetroRetangulo(base, altura float64) float64 {
	return 2 * (base + altura)
}

