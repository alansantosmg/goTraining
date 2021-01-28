package main

import "fmt"

func main() {

	message := "shalom"

	for i := 0; i <= 5; i++ {

		c := message[i]

		fmt.Printf("%c\n", c)

	}

}
