package main

import "fmt"

func main() {

	type location struct {
		folat float64
		long  float64
	}

	bradbury := location{-4.5895, 137.4471}
	curiosity := bradbury

	curiosity.long += 0.0106
	fmt.Println(bradbury, curiosity)
}
