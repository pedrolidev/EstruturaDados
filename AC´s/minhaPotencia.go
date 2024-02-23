package main

import "fmt"

func minhaPotencia(base, expoente int) int {
	if expoente == 0 {
		return 1
	}
	if base == 0 {
		return 0
	}

	resultado := 1
	absExpoente := expoente
	if expoente < 0 {
		absExpoente = -expoente
	}

	for i := 0; i < absExpoente; i++ {
		resultado *= base
	}

	if expoente < 0 {
		return 1 / resultado
	}
	return resultado
}
