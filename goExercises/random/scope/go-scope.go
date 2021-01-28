package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {

	// define ano, mes aleatorio e dia

	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31

	// define numero de dias conforme mês
	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	// define dia aleatorio
	day := rand.Intn(daysInMonth) + 1

	// imprime data
	fmt.Println()
	fmt.Printf("A data é: %v/%v/%v, %v.\n", day, month, year, era)

}
