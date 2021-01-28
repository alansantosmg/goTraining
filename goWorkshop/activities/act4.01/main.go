package main

import "fmt"

func fillArr(arr [10]int) [10]int {
	for i := range arr {
		arr[i] = i + 1
	}
	return arr
}

func main() {
	arr := [10]int{}
	fmt.Println(fillArr(arr))
}
