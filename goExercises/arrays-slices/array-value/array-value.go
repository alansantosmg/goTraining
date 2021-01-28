package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	// copy array
	planetsMarkII := planets
	planets[2] = "Whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

}
