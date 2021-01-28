package main

import (
	"fmt"
)

// funcao que retorna tupla de 3 valores
func retThree(x int) (int, int, int) {
	return 2 * x, x * x, -x
}

func main() {
	fmt.Println(retThree(10))
	n1, n2, n3 := retThree(20)
	fmt.Println(n1, n2, n3)

	//fazendo swapping entre 2 variáveis
	// sem usar varíaveis temporárias
	n1, n2 = n2, n1

	// tuplas também aceitam expressões
	x1, x2, x3 := n1*2, n1*n1, -n1
	fmt.Println(x1, x2, x3)

}
