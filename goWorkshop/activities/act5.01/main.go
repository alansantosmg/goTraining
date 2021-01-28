package main

import "fmt"

type weekday int

const (
	monday weekday = iota
	tuesday
	wednesday
	thursday
	friday
	saturday
)

// tipo empregado strutct
type employee struct {
	id        int
	firstName string
	lastName  string
}

type developer struct {
	individual employee
	hourlyRate int
	workWeek   [7]int
}

func (d *developer) LogHours(day weekday, hours int) {
	d.workWeek[day] = hours
}

func (d *developer) HoursWorked() int {
	total := 0
	for _, v := range d.workWeek {
		total += v
	}
	return total
}

func main() {

	d1 := developer{
		individual: employee{
			1212,
			"Alan",
			"Santos"},
		hourlyRate: 12,
	}

	d1.LogHours(monday, 8)
	d1.LogHours(tuesday, 10)

	fmt.Println("Hours worked on Monday: ", d1.workWeek[monday])
	fmt.Println("Hours worked on Tuesday: ", d1.workWeek[tuesday])
	fmt.Printf("Hours worked this week: %d", d1.HoursWorked())
}
