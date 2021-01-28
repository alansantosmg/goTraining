package main

import (
	"fmt"
	"time"
)

//cria interface stardater
type stardater interface {
	YearDay() int
	Hour() int
}

func stardate(t stardater) float64 {
	// converte time para float64
	doy := float64(t.YearDay())
	h := float64(t.Hour())
	// calcula data estelar
	return 1000 + +doy + h
}

type sol int

func (s sol) Hour() int {
	return 0
}

func (s sol) YearDay() int {
	return int(s % 668) // 668 sols in a martian year
}

func main() {

	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	s := sol(1422)
	fmt.Printf("%.1f happy birthday\n", stardate(s))

}
