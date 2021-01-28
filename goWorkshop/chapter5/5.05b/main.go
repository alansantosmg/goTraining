package main

import "fmt"

func main() {
	i := []int{5, 10, 15}
	fmt.Println(sum(5, 4))
	fmt.Println(sum(i...))
}

func sum(number ...int) (result int) {

	for i := range number {
		result += number[i]
	}
	return result
}
