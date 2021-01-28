package main

import (
	"fmt"
)

// getPointer recebe um ponteiro n inteiro
func getPointer(n *int) {

	//Não precisa retornar nada, pois a função atualiza
	//o ponteiro *n fora dela.
	*n = *n * *n
}

// recebe um inteiro n como parametro de entrada,
func returnPointer(n int) *int {
	v := n * n
	// retorna um ponteiro para um inteiro
	return &v
}

func main() {
	i := -10
	j := 25

	// ponteiro pI aponta para i
	pI := &i
	// ponteiro pJ aponta para j
	pJ := &j

	fmt.Println("pI memory:", pI)
	fmt.Println("pJ memory:", pJ)
	fmt.Println("pI value:", *pI)
	fmt.Println("pJ value:", *pJ)

	// atualiza valor de i
	*pI = 123456
	// subtrai 1 do valor de i
	*pI--
	fmt.Println("i:", i)

	// passa o ponteiro pJ que endereça j  para função
	getPointer(pJ)
	fmt.Println("j:", j)

	// passa 12 para funcao returnPointer
	k := returnPointer(12)
	// imprime ponteiro para k
	fmt.Println(*k)
	// imprime k
	fmt.Println(k)
}
