package main

import (
	"fmt"
)

func main() {
	// Para i = zero
	// equanto ir for menor que 100
	// Rode o loop e add 1 em i
	for i := 0; i < 100; i++ {
		// se o resto da divisão de i por 20 for zero
		if i%20 == 0 {
			// volte ao inicio do loop
			continue
		}
		// se i for igual a 95 encerre o loop
		if i == 95 {
			break
		}
		// imprima i
		fmt.Print(i, " ")
	}

	// de uma linha em branco
	fmt.Println()

	// i é igual a 10
	i := 10

	// entre em loop
	for {
		// se i for maior que 10 saia
		if i < 10 {
			break
		}
		// imprima i e subtraia 1 de i.
		fmt.Print(i, " ")
		i--
	}

	// linha em branco
	fmt.Println()

	// i é igual a zero
	i = 0
	// uma expressão é verdadeira
	anExpression := true

	// ok é verdadeiro
	// enqunto ok for verdadeiro
	// ok igual a anExpression
	for ok := true; ok; ok = anExpression {
		// se i  for maior que 10 AnExpression é falso
		if i > 10 {
			anExpression = false

		}

		// imprimir i e adicionar 1 a i
		fmt.Print(i, " ")
		i++

	}
	fmt.Println()

	// criar array de 5 posições e popular

	anArray := [5]int{0, 1, -1, 2, -2}

	// criar loop
	// i indice e valores na faixa do array
	for i, value := range anArray {

		// imprimir indice e valores
		fmt.Println("index:", i, "value: ", value)
	}

}
