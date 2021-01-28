package main

import "fmt"

func main() {
	i := -10

	if i < 0 {
		fmt.Println("Number cannot be negative")
	} else if i%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}
