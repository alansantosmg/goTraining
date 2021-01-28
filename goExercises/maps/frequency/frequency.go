package main

import "fmt"

func main() {

	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	// cria um map vazio
	frequency := make(map[float64]int)

	// interage sobre o slice pegando cada temperatura
	// repete para cada chave do indice
	for _, t := range temperatures {

		// adiciona a temperatura como indice do map
		// adiciona 1 como valor para indice
		// se na iteracao repetir valor chave
		// adiciona 1 para aquela chave
		// lembrando que map não repete chave de indice
		frequency[t]++
		fmt.Println(frequency)
	}

	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)

	}

}
