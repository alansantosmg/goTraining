package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println()

	var count = 10

	for count > 0 {

		fmt.Printf("Countdown:  << %v >> \r", count)

		time.Sleep(time.Second)
		count--
	}

	fmt.Print("Countdown: <<<< Time is over! >>>>")

}
