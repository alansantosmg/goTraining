package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// funcao de tratamento de erro
func exitOnError(err error) {
	// se erro for diferente de nada imprime erro e encerra
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// criar tipo de struct location
	type location struct {

		// para ser exportado como json deve ser em maiuscula
		// possivel usar tags json e xml para renomear keys
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude" xml:"longitude"`
	}

	// criar var que usa struct location
	curiosity := location{-4.5895, 137.4417}

	// transforma var em bytes (string)
	bytes, err := json.Marshal(curiosity)

	// funcao criada p/ tratamento de erro
	exitOnError(err)

	// imprimir string json
	fmt.Println(string(bytes))

	// transforma var em bytes indentados
	bytes2, err2 := json.MarshalIndent(curiosity, "", "   ")
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}

	fmt.Println(string(bytes2))
}
