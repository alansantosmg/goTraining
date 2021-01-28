package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var WaitGroup sync.WaitGroup

	WaitGroup.Add(2)

	go func() {
		escrever("Olá Mundo.")
		WaitGroup.Done()
	}()

	go func() {
		escrever2("Estou programando.")
		WaitGroup.Done()
	}()

	WaitGroup.Wait()
	fmt.Println("Fim do waitgroup")
}

func escrever(text string) {

	for i := 0; i < 5; i++ {

		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func escrever2(text string) {

	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
