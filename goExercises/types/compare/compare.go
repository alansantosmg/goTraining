package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	// comparando números
	fmt.Println("There is a sign near the entrance tha reads. 'No minors'")

	var age = 41
	var minor = age < 18
	fmt.Println()
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	// comparando strings
	var fruit1 = "apple"
	var fruit2 = fruit1 < "banana"

	fmt.Println()
	fmt.Printf("%v < banana? %v", fruit1, fruit2)
	fmt.Println()
}
