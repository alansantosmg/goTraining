package main

func main() {

	// variables
	s1 := "My String"
	var s2 = "My String"
	var s3 string = "My String"

	var s4, s5, s6 int

	// arrays
	myArray := [3]int{1, 2, 3}
	myArray2 := [2][2]int{{1, 2}, {3, 4}}
	myArray3 := [3]int{}

	// slices
	mySlice := []int{1, 2, 3}
	mySlice2 := []int{}
	mySlice3 := [][]int{}
	mySlice4 := make([]int, 3)
	mySlice5 := []int{1, 2, 3}
	mySlice6 := make([]byte, 50)

	myMap := make(map[string]int)
	myMap["first"] = 1

	mymap2 := map[string]int{
		"k1": 1,
		"k2": 2,
	}

	const HEIGHT = 100

	const S2 string = "My string"

	const (
		C1 = 200
		C2 = 100
	)

}
