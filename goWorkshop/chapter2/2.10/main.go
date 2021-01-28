package main

import "fmt"

func main() {

	config := map[string]string{
		"debug":    "1",
		"loglevel": "warn",
		"version":  "1.21",
	}
	for key, value := range config {
		fmt.Println(key, value)

		names := []string{"Alan", "casa", "teste"}

		for index, _ := range names {
			fmt.Println(index)
		}
	}

}
