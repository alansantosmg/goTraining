package main

import (
	"fmt"
)

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
		fmt.Println("o que vem depois do defer")
	}
	fmt.Println("o que vem depois do for")
}

func main() {
	d3()
	fmt.Println("O que vem depois do main")

}
