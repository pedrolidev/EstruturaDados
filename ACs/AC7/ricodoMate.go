package main

import (
    "fmt"
)

func main() {
    var valor int
    var quantidadeGarrafa, quantidadeCuia float64
    var listaParticipantes []string

    fmt.Println("Digite o número de participantes, quantidade de garrafas e quantidade de cuia:")
    fmt.Scan(&valor, &quantidadeGarrafa, &quantidadeCuia)

    for i := 0; i < valor; i++ {
        var nome string
        fmt.Printf("Digite o nome do participante %d: ", i+1)
        fmt.Scan(&nome)
        listaParticipantes = append(listaParticipantes, nome)
    }

    mostraValor := 0
    for quantidadeGarrafa >= quantidadeCuia {
        quantidadeGarrafa -= quantidadeCuia
        mostraValor++
    }

    mostraValor = mostraValor % valor
    fmt.Printf("Participante: %s, Quantidade restante na garrafa: %.1f\n", listaParticipantes[mostraValor], quantidadeGarrafa)
}
