package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday()
	hour := time.Now().Hour()

	fmt.Println("Dia:", day, "Hora:", hour)

	if day.String() == "Monday" && hour > 1 {
		fmt.Println("Script completo")
	}
	fmt.Println("Script resumido")
}
