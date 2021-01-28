package main

import "fmt"

func message() string {

	m := ""
	arr := [4]int{1, 2, 3, 4}

	for i := range arr {

		arr[i] = arr[i] * arr[i]

		m += fmt.Sprintf("%v %v \n", i, arr[i])
	}
	return m
}

func main() {
	fmt.Println(message())
}
