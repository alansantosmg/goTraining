package main

import (
	"fmt"
	"math"
)

func main() {
	square := func(a float64) (result float64) {
		result = math.Sqrt(a)
		return result
	}

	fmt.Println(square(9))
}
