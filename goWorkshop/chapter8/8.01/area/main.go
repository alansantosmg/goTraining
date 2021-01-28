package main

import (
	"fmt"

	"github.com/alansantosmg/goWorkshop/8.01/shape"
)

func main() {

	t := shape.Triangle{base: 15.5, height: 20.1}
	r := shape.Retangle{lenght: 20, width: 10}
	s := shape.Square{side: 10}
	shape.PrintShapeDetails(t, r, s)
	fmt.Println("vim-go")
}
