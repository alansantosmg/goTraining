package main

import "fmt"

func main() {
	km := 2

	//erro de semantica
	// Era para andar a pe somente se km fosse inferior a 2
	if km > 2 {
		fmt.Println("Take the car")
	} else {
		fmt.Println("Go to the destiny")
	}
}
