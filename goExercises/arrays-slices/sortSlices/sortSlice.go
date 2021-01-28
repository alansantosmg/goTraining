package main

import (
	"fmt"
	"sort"
)

// define um tipo de estrutura (objeto
type aStructure struct {
	person string
	height int
	weight int
}

func main() {

	// Cria slice para armazenar estrutura do tipo
	mySlice := make([]aStructure, 0)

	// insere dados no array seguindo estrutura
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 40})

	// imprime array
	fmt.Println("0:", mySlice)

	// ordena o slice pela propriedade height da estrutura
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	// orderna o slice pela propriedade height da estrutura (inverso)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)

}
