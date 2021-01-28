package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {

	// prevents panic
	if p == nil {
		return
	}
	p.age++

}

func main() {

	// var do tipo person que é uma struct aponta para nada
	var nobody *person

	// imprime nada
	fmt.Println(nobody)

	// chama metodo e passa parametro nada.

	nobody.birthday()
}
