package main

import (
	"fmt"
)

func main() {
	// define um TIPO de structure XYZ
	type XYZ struct {
		X int
		Y int
		Z int
	}

	// inicializa variável s1 com tipo XYZ
	var s1 XYZ

	// atualiza variavel s1
	s1.X = 10
	s1.Y = 20
	s1.Z = 30
	fmt.Printf("x: %d, y: %d, z: %d\n", s1.X, s1.Y, s1.Z)

	fmt.Println()

	// cria variaveis p1 e p2 usando tipy XYZ literal
	p1 := XYZ{23, 12, -2}
	p2 := XYZ{Z: 12, Y: 13}
	fmt.Println(p1)
	fmt.Println(p2)

	// cria um array de Structure XYZ
	pSlice := [4]XYZ{}

	pSlice[2] = p1
	pSlice[0] = p2

	// Se uma structure for alterada isso não alterará o array
	// pois ao criar um array, c/ structure estamos criando
	// uma cópia da structure inicial dentro do array.
	p2 = XYZ{1, 2, 3}

	fmt.Println(pSlice)
}
