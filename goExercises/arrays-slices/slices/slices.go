package main

import (
	"fmt"
)

func main() {
	// cria slice
	aSlice := []int{1, 2, 3, 4, 5}
	// imprime o slice
	fmt.Println(aSlice)

	// cria um slice  integer vazio tamanho 2
	integer := make([]int, 2)

	// imprime slice integer
	fmt.Println("integer:", integer)

	// atribui nada a integer
	integer = nil
	// imprime integer
	fmt.Println("integer:", integer)

	// criar anArray 5 inteiros, 1- a -5
	anArray := [5]int{-1, -2, -3, -4, -5}

	// cria refAnArray e atribui anArray
	refAnArray := anArray[:]

	// imprime anArray
	fmt.Println("anArray:", anArray)

	// imprime refAnArray
	fmt.Println("refAnArray:", refAnArray)

	// anArray p4 = 100
	anArray[4] = -100

	// imprime refAnArray
	fmt.Println("refAnArray:", refAnArray)

	// cria slice s  byte de 5 p
	s := make([]byte, 5)

	// imprime s
	fmt.Println("s:", s)

	// cria slice twoD bidimensional inteiro 3 p
	twoD := make([][]int, 2)
	// imprime twoD
	fmt.Println("twoD:", twoD)

	// imprime linha em branco
	fmt.Println()

	// cria um for
	// enquanto i for menor que tamando twoD soma 1 em i

	for i := 0; i < len(twoD); i++ {

		// cria for
		// enquanto j for menor que 2 soma 1 em j
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j)

			for _, x := range twoD {
				for i, y := range x {
					fmt.Println("index i:", i, "value y", y)
				}
				fmt.Println()
			}
		}
	}
}
