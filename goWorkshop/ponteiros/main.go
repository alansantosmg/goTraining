package main

import "fmt"

func main() {
	temperatura := 100
	fatorEbulicao := 2

	// Ponteiro nao aceita valor, logo deve ser passado endereço 
	// do valor (da variavel que armazena valor)
	pontoEbulicao(&temperatura, fatorEbulicao)
	fmt.Println(temperatura)
}
]//declaracao de parametro do tipo ponteiro
// esse parametro vai receber um endereço de memoria de alguma variavel
func pontoEbulicao(temp *int, fe int) {

	// Neste caso a funçao nao precisa de retorno
	// pois já está alterando valor fora dela
	// no caso, a variavel temperatura
	*temp = *temp * fe
}
