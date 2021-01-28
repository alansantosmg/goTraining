package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Print header
	fmt.Println()
	fmt.Println("Linha espacial   Dias Tipo       Preço")
	fmt.Println("=======================================")

	rand.Seed(time.Now().UnixNano())

	// Const init
	const DISTANCE = 62100000
	const MAXKM = 30
	const MINKM = 16
	const maxPrice = 50
	const minPrice = 36

	// while we don't have 10 tickets
	for i := 1; i <= 10; i++ {

		// set random speed
		var speed int = rand.Intn(MAXKM-MINKM) + MINKM + 1

		// calc price and travel days
		var price int = speed + 20
		var travelDays = (DISTANCE / speed) / 3600 / 24

		//  select random trip type
		var tripType string
		if tripOption := rand.Intn(2); tripOption == 0 {
			tripType = "ida"
		} else {
			tripType = "ida e volta"
			price *= 2
		}

		// select a random company
		var company string
		switch companyNumber := rand.Intn(3); companyNumber {
		case 0:
			company = "Virgin Galactic"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Space Adventures"
		}

		// print Content
		fmt.Printf("%-16v %4v %-11v $%4v\n", company, travelDays, tripType, price)

	}
}
