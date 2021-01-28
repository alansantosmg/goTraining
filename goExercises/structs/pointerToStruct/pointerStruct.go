package main

import (
	"fmt"
)

// cria structure
type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

// recebe dados para criar uma structure e verifica
func createStruct(n, s string, h int32) *myStructure {

	if h > 300 {
		h = 0
	}
	// retorna estrutura para um endereço da memória
	return &myStructure{n, s, h}
}

func retStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	// retorna valor
	return myStructure{n, s, h}
}

func main() {

	s1 := createStruct("Alan", "Santos", 123)
	s2 := retStructure("Alan", "Santos", 123)

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)
}
