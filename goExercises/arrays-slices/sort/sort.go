package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	// sorting long way
	// package short, type StringSlice method sort
	// sort.StringSlice(planets).sort()

	// sort short way
	sort.Strings(planets)

	fmt.Println(planets)

}
