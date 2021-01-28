package main

import "fmt"

func createSlice() []string {

	mySlice := []string{"0", "1", "2", "3", "4"}

	return mySlice
}

func moveSlice(sli []string) []string {

	fmt.Println(sli[:2])
	fmt.Println(sli[3:])
	sli = append(sli[:2], sli[3:]...)

	return sli

}

func main() {

	x := createSlice()

	fmt.Println(moveSlice(x))

}
