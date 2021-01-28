package main

import (
	"fmt"
)

func main() {

	fmt.Println()
	fmt.Println("The year 2100, should you leap? ")

	const year = 2100

	var leapYear = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	// fmt.Println(leapYear)

	if leapYear {
		fmt.Println("Yes it is!")
	}
	fmt.Println("No! it's not!")

}
