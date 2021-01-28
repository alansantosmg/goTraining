// 2.2b align text
package main

import "fmt"

func main() {
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
	fmt.Printf("%-25v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-25v $%4v\n", "Virgin Galactic", 100)
}
