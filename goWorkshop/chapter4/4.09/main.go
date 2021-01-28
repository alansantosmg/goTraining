package main

import (
	"fmt"
	"os"
)

// obtem parametros de linha de comando
func getPassedArgs() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

// atualizar lista de locais
func getLocals(extraLocals []string) []string {
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocals...)
	return locales
}

func main() {

	// obter locais de parametros de linha de comando obtidos
	locales := getLocals(getPassedArgs())
	fmt.Println("Locales to use:", locales)
}
