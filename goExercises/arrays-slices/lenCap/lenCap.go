package main

// Este programa mostra que quando um sclice estoura a capacidade
// ao se adicionar 1 item a mais que seu tamanho atual
// ele dobra de capacidade.

import (
	"fmt"
)

//funcao printSlice com parametro x do tipo slice-int
func printSlice(x []int) {
	//para cada numero no range do slice x
	for _, number := range x {
		// imprima o numero
		fmt.Print(number, " ")
	}
	// imprimir linha em branco
	fmt.Println()
}

func main() {
	// declara slice com 3 elementos
	aSlice := []int{-1, 0, 4}

	// passa o aSlice para função printSlice
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	fmt.Printf("aSlice: ")
	aSlice = append(aSlice, -100)

	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	aSlice = append(aSlice, -2)
	aSlice = append(aSlice, -3)
	aSlice = append(aSlice, -4)
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

}
