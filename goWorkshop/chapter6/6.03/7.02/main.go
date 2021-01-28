package main

import "fmt"

// interface com metodos comuns
// às structs que satisfacam criterios
type Shape interface {
	Area() float64
	Name() string
}

// structs para receber dados
type triangle struct {
	base  float64
	width float64
}

type retangle struct {
	lenght float64
	width  float64
}

type square struct {
	side float64
}

// metodos do triangulo
func (t triangle) Name() string {
	return "triangle"
}
func (t triangle) Area() float64 {
	return (t.base * t.width) / 2
}

//metodos do quadrado
func (s square) Name() string {
	return "square"
}
func (s square) Area() float64 {
	return s.side * s.side
}

//metodos dos retangulo
func (r retangle) Name() string {
	return "retangle"
}
func (r retangle) Area() float64 {
	return r.width * r.lenght
}

// funcão que itera com dados do metodo satisfeito pela interface
// imprime área de cada objeto

func printShapeDetails(shapes ...Shape) {
	for _, v := range shapes {
		fmt.Println(v.Name(), v.Area())
	}
}

func main() {

	// entra dados
	t := triangle{base: 15.5, width: 20.1}
	r := retangle{width: 10, lenght: 15.5}
	s := square{side: 22}

	// chama funcao de impressao
	printShapeDetails(t, r, s)

}
