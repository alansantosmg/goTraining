package main

import (
	"fmt"
)

func main() {

	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km

	fmt.Println()
	fmt.Println("---/ Tempo de viagem da Terra à Marte /---")

	fmt.Printf("Distancia mínima de %4v km entre órbitas, à velocidade da luz: ", distance)
	fmt.Println(distance/lightSpeed, "segundos.")

	distance = 401000000

	fmt.Printf("Distancia máxima de %4v km entre órbitas, à velocidade da luz: ", distance)
	fmt.Println(distance/lightSpeed, "segundos.")

}
