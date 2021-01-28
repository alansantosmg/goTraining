package main

import (
	"fmt"
	"time"
)

func main() {

	var count1 *int

	count2 := new(int)

	countTemp := 5

	count3 := &countTemp

	t := &time.Time{}

	fmt.Println(count1)
	fmt.Println(count2)
	fmt.Println(count3)
	fmt.Println(t)

}
