package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for {

		number := rand.Intn(99)

		if number%3 == 0 {
			fmt.Println("divisible by 3")
			continue
		} else if number%8 == 0 {
			fmt.Println("divisible by 2")
			break
		}
		fmt.Println(number)

	}
}
