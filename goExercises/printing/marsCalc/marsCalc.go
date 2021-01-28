package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println("---/ Calculadora de vida marciana /---")
	//fmt.Println("Minha idade em Marte seria:", (41 * 365 / 687), "anos.")
	//fmt.Println("Meu peso em marte seria: ", (75 * 0.62), ".")
	fmt.Println()
	fmt.Println("Informações:")
	fmt.Println("- Um ano marciano tem 687 dias.")
	fmt.Println("- A gravidade marciana é 38% menor que a terráquea.")
	fmt.Println()

	// impressão com formatação da mais controle.
	// %v = value   %8v = 8 spaces+value

	fmt.Printf("Em Marte, uma pessoa de 41 anos teria %v anos de idade.\n", (41 * 365 / 687))
	fmt.Printf("Em Marte, uma pessoa de 76 quilos pesaria  %4v kg.\n", (76 * 0.62))

}
