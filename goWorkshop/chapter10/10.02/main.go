package main

import (
	"fmt"
	"time"

	t "10.02/timelapse"
)

func main() {

	start := time.Now()
	time.Sleep(2 * time.Second)

	end := time.Now()

	fmt.Println(t.ElapsedTime(start, end))
}
