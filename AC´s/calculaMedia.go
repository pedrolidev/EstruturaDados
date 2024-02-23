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



