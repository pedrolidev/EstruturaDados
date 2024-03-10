package main

import "fmt"

// 1) Primeiro exercício (Torre de Hanói)

func hanoiTower(num int, origem, destino, teste string) {
	if num == 1 {
		fmt.Printf("\nMova a pedra de " + origem + "para o" + destino)
		return
	}

	explainedHanoi(num-1, origem, teste, destino)


	fmt.Printf("\nMova a pedra de " + origem + "para o" + destino)

	explainedHanoi(num-1, origem, teste, destino, origem)
}

// 2) Segundo exercício

func addingFinder(array []int, alvo int) (int, int) {
	esquerda, direita := 0, len(array)-1

	for esquerda < direita {
		soma := array[esquerda] + array[direita]
		if soma == alvo {
			return array[esquerda], array[direita]
		} else if soma < alvo {
			esquerda++
		} else {
			direita--
		}
	}

	return -1, -1
}

// Função main para rodar os códigos!!

func main() {
	// Torre de Hanói
	n := 3
	fmt.Printf("A resposta encontrada para resolver o jogo foi feita com %d discos:\n", n)
	hanoiTower(n, "A", "B", "C")

	// Algoritmo
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	alvo := 10

	resultado1, resultado2 := addingFinder(array, alvo)

	if resultado1 == -1 && resultado2 == -1 {
		fmt.Println("Nenhum par foi econtrado.")
	} else {
		fmt.Printf("O par foi encontrado com sucesso: (%d, %d)\n", resultado1, resultado2)
	}
}
