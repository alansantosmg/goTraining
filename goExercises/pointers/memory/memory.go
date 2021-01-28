package main

import "fmt"

func main() {

	answer := 42
	// imprime endereço da memoria de answer
	fmt.Println(&answer)

	// atribui endereço de answer a adress
	// address aponta para endereço de answer
	address := &answer

	// imprime o valor de answer
	// uma vez que o endereço de answer foi atribuido a address
	// mostra para onde address está apontando
	fmt.Println(*address)

	// imprime o valor de answer
	// answer aponta para o próprio endereço
	fmt.Println(*&answer)

}
