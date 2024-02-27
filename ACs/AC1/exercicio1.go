package main

import "fmt"

func calculaMedia(valores ...float64) float64 {
	if len(valores) == 0 {
		return 0
	}

	soma := 0.0
	for _, valor := range valores {
		soma += valor
	}

	media := soma / float64(len(valores))
	return media
}

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


func verificaParidade(numero int) bool {
    return numero%2 == 0
}

func converteCelciusParaFahrenheit(celcius float64) float64 {
	return (celcius * 9 / 5) + 32
}