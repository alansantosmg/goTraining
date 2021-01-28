package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func birthday2(p person) person {
	p.age++
	return p
}

func main() {

	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}

	// uso como parametro endereço de memória
	// onde estão os valores de rebecca
	// Se não tivesse passado ponteiro, a funcão
	// criaria uma cópia da scruct e seria necessario
	// fazer um return modificando essa variavel.

	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)

	// neste caso como não tem ponteiro eh preciso
	//alterar manualmente a struct

	rebecca = birthday2(rebecca)
	fmt.Printsf("%+v\n", rebecca)

}
