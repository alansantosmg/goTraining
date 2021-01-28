package main

import "fmt"

func message() string {

	// define um slice
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// extrai partes de um slice e guarda em m.
	m := fmt.Sprintln("first: ", s[0], s[0:1], s[:1])

	m += fmt.Sprintln("Last: ", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
	m += fmt.Sprintln("First 5: ", s[:5])
	m += fmt.Sprintln("last 4: ", s[5:])
	m += fmt.Sprintln("middle: ", s[2:7])

	return m
}

func main() {

	fmt.Printf(message())
}
