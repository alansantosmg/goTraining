package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("Invalid hourly rate")
	ErrHoursWorked = errors.New("Invalid hours worked per week")
)

func main() {
	pay := payDay(81, 50)
	fmt.Println(pay)


	if hourlyRate
}

func payDay(hoursWorked, hourlyRate int) int{
	report := func() {
	fmt.Printf("HoursWorked: %d\nHourlyRate: %d\n", hoursWorked, hourlyRate)
}
	defer report()
}


