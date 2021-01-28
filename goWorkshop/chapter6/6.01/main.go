package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8}
	total := 0

	//ira gerar erro pois i só pode ser adicionado até 3 posicao
	for i := 0; i <= 10; i++ {
		total += nums[i]
	}
	fmt.Println("Total:", total)
}
