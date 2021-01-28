package main

import (
	"fmt"
	"math"
)

func main() {

	temperaturas := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	// cria um map float64 e slice float64
	groups := make(map[float64][]float64)

	// para cada temperatura encontrada no range
	for _, t := range temperaturas {

		// divide a temperatura por 10
		// trunca os decimais
		// multiplica por 10 para obter faixas de 10 em 10
		g := math.Trunc(t/10) * 10

		// usa cada faixa como indice
		// e acrescenta temperatura no indice
		groups[g] = append(groups[g], t)

	}

	// imprime indices (faixas de 10 e temperaturas)
	for g, temperaturas := range groups {
		fmt.Printf("%v: %v\n", g, temperaturas)
	}
}
