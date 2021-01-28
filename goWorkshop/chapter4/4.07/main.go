package main

import "fmt"

func fillArray(arr [10]int) [10]int {
	for i := range arr {
		arr[i] = i + 1
	}
	return arr
}

func multiplyArray(arr [10]int) [10]int {
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}

func main() {
	var arr [10]int
	arr = fillArray(arr)
	fmt.Println(arr)
	arr = multiplyArray(arr)
	fmt.Println(arr)
}
