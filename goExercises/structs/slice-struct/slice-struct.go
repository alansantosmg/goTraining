package main

import "fmt"

func main() {

	// cria tipo de estrutura
	type location struct {
		name string
		lat  float64
		long float64
	}

	// cria slice do tipo estrutura
	locations := []location{

		// guarda dados no slice
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}
	// imprime slice
	fmt.Println(locations)
}
