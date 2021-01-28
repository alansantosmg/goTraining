package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var firstDate, secondDate string
	if len(os.Args) != 3 {
		fmt.Printf("Utilização: %s dd/mm/yyyy dd/mm/yyyy", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	firstDate = os.Args[1]
	secondDate = os.Args[2]

	firstDateInput, err := time.Parse("02/01/2006", firstDate)
	anoFirstDateInput := firstDateInput.Year()
	if err == nil {
		fmt.Println("Data inicial:", firstDateInput)
	} else {
		fmt.Println(err)
	}

	secondDateInput, err := time.Parse("02/01/2006", secondDate)
	anoSecondDateInput := secondDateInput.Year()
	if err == nil {
		fmt.Println("Data final:", secondDateInput)
	} else {
		fmt.Println(err)
	}

	intervaloDatas := anoSecondDateInput - anoFirstDateInput
	fmt.Printf("A diferença é de %d anos.", intervaloDatas)
}
