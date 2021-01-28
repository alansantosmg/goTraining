package main

import (
	"fmt"
)

func main() {

	type location struct {
		lat, long float64
	}

	// composite literal struct 1a forma
	oportunity := location{lat: -1.9462, long: 354.4734}

	fmt.Println(oportunity)

	insight := location{lat: 4.5, long: 135.9}

	fmt.Println(insight)

	// composite literal struct forma 2
	curiosity := location{-14.5684, 175.472636}

	fmt.Println(curiosity)

	// imprime value
	fmt.Printf("%v\n", curiosity)

	// imprime key and value
	fmt.Printf("%+v\n", curiosity)

}
