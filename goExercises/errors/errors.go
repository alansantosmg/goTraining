package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Se usuário não tiver digitado nenhum argumento
	// 1 significa que digitou somente o proprio comando
	// pois ele é considerado o primeiro argumento
	if len(os.Args) == 1 {

		fmt.Println("Por favor me dê um numero de ponto flutuante")
		os.Exit(1)
	}

	// definição de variaveis
	arguments := os.Args
	var err error = errors.New("An error")
	k := 1
	var n float64

	// Enquanto erro não for nada
	for err != nil {

		// e se k for maior que numero de argumentos
		if k > len(arguments) {
			fmt.Println("Nenhum argumento é numero de ponto flutuante")
			return

		}
		//senao
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	min, max := n, n

	// define um loop
	// enquanto tiver menos que 2 argumentos

	for i := 2; i < len(arguments); i++ {

		// converta n para float
		n, err := strconv.ParseFloat(arguments[i], 64)

		// se erro for vazio
		if err == nil {

			// e se n for menor que minimo
			if n < min {

				// então minimo passa a ser n
				min = n
			}

			//  mas se n for maior que maximo
			if n > max {

				// n passa a ser maximo
				max = n
			}
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
