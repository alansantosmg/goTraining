package main

import "fmt"

func main() {

	names := []string{"Jim", "jane", "Joe", "june"}

	for i := 0; i < len(names); i++ {

		fmt.Println(names[i])
	}
}
