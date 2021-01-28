package main

import "fmt"

func taxCalc(cost float64, salesTax float64) float64 {

	return cost * salesTax
}

func main() {

	cake := taxCalc(0.99, 0.075)
	milk := taxCalc(2.75, 0.015)
	butter := taxCalc(0.87, 0.02)

	total := cake + butter + milk

	fmt.Println("Sales Tax Total: ", total)
}
