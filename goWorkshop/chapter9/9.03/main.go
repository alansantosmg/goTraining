package main

import "fmt"

type person struct {
	lname  string
	age    int
	salary float64
}

func main() {

	fname := "joe"

	grades := []int{100, 87, 67}
	states := map[string]string{"KY": "kentucky", "WV": "west Virginia", "VA": "Virginia"}
	p := person{lname: "Lincoln", age: 23, salary: 210000.00}

	fmt.Printf("fname value %#v\n", fname)
	fmt.Printf("grades value %#v\n", grades)
	fmt.Printf("states value %#v\n", states)
	fmt.Printf("p value %#v\n", p)

}
