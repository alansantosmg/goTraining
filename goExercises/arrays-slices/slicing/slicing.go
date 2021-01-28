package main

import "fmt"

func main() {

	planets := [...]string{
		"Mercury",
		"Venus",
		"Earh",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	ice := iceGiants

	fmt.Println(terrestrial, gasGiants, iceGiants)

	// slices are views of arrays

	// if your modify a slice item
	// you're modifying the origin array
	ice[0] = "test"

	fmt.Println(iceGiants)
	fmt.Println(ice)
	fmt.Println(ice[0])
	fmt.Println(planets)
}
