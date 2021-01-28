package main

import (
	"fmt"
)

func itensSold() {
	itens := make(map[string]int)
	itens["john"] = 412
	itens["Celina"] = 109
	itens["mica"] = 24

	for k, v := range itens {
		fmt.Printf("%s sold %d itens and ", k, v)

		if v < 40 {
			fmt.Println("is below expectations")

		} else if v > 40 && v <= 100 {
			fmt.Println("meets expectations")
		} else if v > 100 {

			fmt.Println("exceeded expectations")
		}

	}
}

func main() {

	itensSold()
}
