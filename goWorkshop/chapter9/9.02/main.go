package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	for i := 1; i <= 255; i++ {

		fmt.Printf("Decimal:%3.d  Binary: %-8.b Hex: %-2.x\n", i, i, i)

	}
}
