package main

import "fmt"

// Speaker define interface (comportamento do metodo)
type Speaker interface {
	Speak() string
}

// cria struct com nome da pessoa
type person struct {
	name      string
	age       int
	isMarried bool
}

func main() {

	//cria pessoa usando struct
	p := person{
		name:      "Cailyn",
		age:       33,
		isMarried: false,
	}

	// chama método falar
	fmt.Println(p.Speak())

	// chama metodo String
	fmt.Println(p)
}

// metodo pela interface Stringer do pacote padrao
func (p person) String() string {
	return fmt.Sprintf("%v (%v years old). \nMarried status: %v ", p.name, p.age, p.isMarried)
}

// metodo definido pela interface falar
func (p person) Speak() string {
	return "Hi. My name is " + p.name
}
