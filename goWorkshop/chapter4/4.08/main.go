package main

import (
	"fmt"
	"os"
)

func getPassedArgs(minArgs int) []string {

	// verifica se quantidade de argumentos passados
	// é 3 (proprio comando mais 2 argumentos)
	// mínimo tem que ser 3 pois assim teremos no mínimo
	// dois argumentos para comparações
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed", minArgs)
		os.Exit(1)
	}

	// carrega argumentos para um slice
	// a partir do segundo argumento
	//retorna o slice
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

// recebe slice de argumentos
// inicia variavel longest como zero
// compara argumento com longest
// se argumento for maior, guarda em longest
// compara proximo argumento com longest
// no final retorna maior argumento guardado em longest
func findLongest(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}
	return longest
}

func main() {
	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest)
	} else {
		fmt.Println("There was an error")
		os.Exit(1)
	}
}
