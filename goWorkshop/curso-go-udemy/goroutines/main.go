package main

import (
	"fmt"
	"time"
)

func main() {

	go escrever()
	escrever2()
}

func escrever() {

	for {

		fmt.Println("Ola mundo")

		time.Sleep(2 * time.Second)
	}
}

func escrever2() {

	for {
		fmt.Println("I am the best")

		time.Sleep(2 * time.Second)
	}
}
