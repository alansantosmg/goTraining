package main

import "fmt"

func compArrays() (bool, bool, [10]int) {

	var arr1 [10]int
	arr2 := [...]int{9: 0}
	arr3 := [10]int{1, 9: 10, 4: 5}

	return arr1 == arr2, arr1 == arr3, arr3
}

func main() {

	comp1, comp2, comp3 := compArrays()

	fmt.Printf("%#v %#v %#v \n ", comp1, comp2, comp3)

}
