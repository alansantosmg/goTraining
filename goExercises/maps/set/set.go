package main

import (
	"fmt"
	"sort"
)

func main() {

	var temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	// cria um map
	set := make(map[float64]bool)

	// para cada temperatura existente no map
	for _, t := range temperatures {
		// marca temperatura com valor verdadeiro
		set[t] = true
	}

	if set[-28.0] {

		fmt.Println("set member")
	}

	fmt.Println(set)

	// não é possivel ordenar um map
	// por isso deve ser convertido para array
	unique := make([]float64, 0, len(set))

	for t := range set {
		unique = append(unique, t)

	}
	sort.Float64s(unique)
	fmt.Println(unique)
}
