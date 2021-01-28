package main

import "fmt"

func main() {
	celsius := 21.0

	// erro de arrendondamento
	fmt.Print((celsius/5.0*9.0)+32, "F\n")
	fmt.Print((9.0/5*celsius)+32, "F\n")

	// Colocar a multiplicação primeiro
	// reduz possibilidade de erros de arredondamento
	fmt.Print((celsius*9.0/5.0)+32, "F\n")
}
