package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println()
	fmt.Println("Computador adivinhando número!")

	// variáveis iniciais
	var computerNumber int

	// escolhe um numero aleatório
	var myNumber = rand.Intn(10000000)

	// guarda data/ hora inicial
	var initTime = time.Now()

	// enquanto computador não adivinhar o número
	for computerNumber != myNumber {

		// escolha um número para comparar com numero inicial
		computerNumber = rand.Intn(10000000)
		fmt.Println("Escolhi:", computerNumber)

		// Compara números
		switch {
		case computerNumber > myNumber:
			fmt.Print("Errei para mais!  Vou escolher outro número. ")
		case computerNumber < myNumber:
			fmt.Print("Errei para menos! Vou escolher outro número. ")
		default:
			fmt.Printf("%v - Agora acertei!\n", computerNumber)
		}
	}

	// guarda data/hora final
	var finalTime = time.Now()

	// contabiliza difença entre data/hora inicial e final
	difTime := finalTime.Sub(initTime)
	fmt.Printf("Levei: %v para acertar a resposta", difTime)
}
