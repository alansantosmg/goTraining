package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	fmt.Println("Script started at:", start)
	fmt.Println("Saving de Worldl")
	time.Sleep(2 * time.Second)
	end := time.Now()
	fmt.Println("The script has completed at", end)

}
