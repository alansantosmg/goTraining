package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {

	// como valor passado é nada,
	// se for descomentada a linha de adicionar
	// ira gerar panic, pois não é possível
	// mais de nada ao nada.
	//p.age++

}

func main() {

	// var do tipo person que é uma struct aponta para nada
	var nobody *person

	// imprime nada
	fmt.Println(nobody)

	// chama metodo e passa parametro nada.

	nobody.birthday()
}
