package main

import (
	"fmt"
)

func main() {

	var administrator *string

	scolese := "Christopher J. Scolese"
	administrator = &scolese

	fmt.Println(*administrator)

	bolden := "Charles Bolden"
	administrator = &bolden

	fmt.Println(*administrator)
}
