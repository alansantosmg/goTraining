package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var myDate string

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s string\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	myDate = os.Args[1]
	d, err := time.Parse("02/01/2006", myDate)
	//	d, err := time.Parse("02 January 2006", myDate)
	if err == nil {
		fmt.Println("full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	} else {
		fmt.Println(err)
	}

}
