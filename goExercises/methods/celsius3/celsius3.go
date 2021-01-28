package main

import "fmt"

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}
func main() {
	var k kelvin = 294.0
	c := k.celsius()
	fmt.Print(k, "o K is ", c, "o C")
}
