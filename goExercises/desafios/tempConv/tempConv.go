package main

import (
	"fmt"
	"os"
	"strconv"
)

// farToCelsius converts temperature from farenheit to celsius
func farToCelsius(temp float64) float64 {
	return (temp - 32.0) * 5.0 / 9.0
}

// celsiusToFar converts temperature from celsius to farenheit
func celsiusToFar(temp float64) float64 {
	return (temp*9)/5 + 32
}

// case incorret arguments are given
var helpMessage = `
tempConv - Converts temperatures to/from celsius x farenheint.

Usage: tempConv <temp>  <tempType>

Examples:

tempConv 100 c  - Converts 100c to the respective farenheint temperature: 212f
tempConv 212 f  - Converts 212f to the respective celsius temperature: 100c

`

// Converts a given temperature from/to celsius or from/to farenheit
func main() {

	if len(os.Args) == 3 { // check number of command line arguments

		// get command line args
		var tempInput = os.Args[1]
		var tempType = os.Args[2]

		// convert temperature string to float
		t, _ := strconv.ParseFloat(tempInput, 64)

		// convert temperature
		var c = farToCelsius(t)
		var f = celsiusToFar(t)

		// print converted temperature or helpMessage
		switch tempType {

		case "f":
			fmt.Printf("Farenheint to Celsius: \n%g \u2109  = %g \u2103\n", t, c)

		case "c":

			fmt.Printf("Celsius to Farenheint: \n%g \u2103  = %g \u2109\n", t, f)
		default:

			fmt.Printf("%s", helpMessage)
		}

	} else {
		fmt.Printf("%s", helpMessage)
	}
}
