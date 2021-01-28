package main

import (
	"fmt"
	"math/rand"
)

// capstone 1
func main() {

	const secondsPerDay = 86400
	const marsDistance = 62100000 // km
	const litoff = "October, 13, 2020"
	const minSpeed = 16
	const minCost = 36
	var company string
	var trip string

	fmt.Printf("Company            Type         days   cost\n")
	fmt.Printf("===========================================\n")

	for i := 1; i <= 10; i++ {

		companyNumber := rand.Intn(3) + 1
		speed := rand.Intn(15) + 16

		duration := marsDistance / speed / secondsPerDay

		cost := (speed - minSpeed) + minCost

		switch companyNumber {
		case 1:
			company = "Space Adventures"

		case 2:
			company = "SpaceX"
		case 3:
			company = "Virgin Galactic"
		}

		tripType := rand.Intn(2) + 1
		if tripType == 1 {
			trip = "One-way"
		} else {
			trip = "Round Trip"
			cost *= 2

		}

		fmt.Printf("%-18v %-12v %5v $%4v\n", company, trip, duration, cost)

	}

	//speed
	// rand 16 to 30 km/s

	//price
	//36 to 50 k

	// time
	// time = D / speed

}
