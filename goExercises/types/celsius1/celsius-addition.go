package main

import "fmt"

type celsius float64

var f = func(temperature celsius) celsius {
	return temperature + 10
}

func main() {

	const degrees = 20

	//	var temperature celsius = degrees

	//	temperature += 10

	fmt.Println(f(degrees))

	//	var warmUp float64 = 10

	// necessaŕio converter warmUp para celsius
	// pois celsius é um tipo definido pelo programador
	// ele não é um tipo float64.
	// Somente tem propriedades de float64, mas é celsius!
	//	temperature += celsius(warmUp)

}
