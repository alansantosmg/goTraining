package main

import "fmt"

func main() {
	canada := "Canada"

	var home *string

	home = &canada

	fmt.Println(*home)

}