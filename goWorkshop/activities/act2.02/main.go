package main

import "fmt"

func main() {

	words := map[string]int{
		"gonna": 3,
		"you":   3,
		"give":  8,
		"never": 1,
		"up":    4,
	}

	myvalue := 0
	mykey := ""
	for key, value := range words {

		if value >= myvalue {
			myvalue = value
			mykey = key
		}
	}
	fmt.Println(words[mykey])
}
