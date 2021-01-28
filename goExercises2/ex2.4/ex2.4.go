// Assignement operators - shortcuts
package main

import "fmt"

func main() {

	var weight = 149.0
	weight = weight * 0.383

	// shortcut
	var weight2 = 149.0
	weight2 *= 0.383

	fmt.Println(weight)
	fmt.Println(weight2)

}
