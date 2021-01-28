package main

import (
	"fmt"
)

func main() {

	third := 1.0 / 3

	// imprime variável sem formatar
	fmt.Println(third)

	// imprime "valor" sem formatar
	fmt.Printf("%v\n", third)

	// imprime número float formatado
	fmt.Printf("%f\n", third)

	// 3 casas decimais
	fmt.Printf("%.3f\n", third)

	// 4 caracteres 2 casas decimais
	fmt.Printf("%4.2f\n", third)

	// 5 caracteres 2 casas decimais
	fmt.Printf("%05.2f\n", third)

	myNumber1 := 10.5
	// Checar o tipo de uma variável
	fmt.Printf("Type of %T for %v\n", myNumber1, myNumber1)
	fmt.Printf("Type of %T for %[1]v\n", myNumber1)

	myNumber := 10
	// imprime número inteiro
	fmt.Printf("%d", myNumber)

	// imprime os bits de um inteiro
	fmt.Printf("Bits of %v: %[1]b\n", myNumber)

	var pi rune = 960
	var letter rune = 100
	fmt.Printf("%v %v\n", pi, letter)
	// imprime caracter correspondente
	// se for caracter imprime caracter
	// se for número converte unicode para caracter
	fmt.Printf("%c %c\n", pi, letter)

}
