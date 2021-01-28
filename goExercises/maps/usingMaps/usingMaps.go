package main

import (
	"fmt"
)

func main() {

	// criando map Imap com make
	iMap := make(map[string]int)
	iMap["k1"] = 12
	iMap["k2"] = 13
	fmt.Println("iMap:", iMap)

	// criando map Imap literal adicionando valores
	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	}
	fmt.Println("anotherMap:", anotherMap)

	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	fmt.Println("anotherMap:", anotherMap)

	// verificando se iMap possui item
	_, ok := iMap["doestItExist"]
	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does NOT exist")
	}

	// iterando com map
	for key, value := range iMap {
		fmt.Println(key, value)
	}
}
