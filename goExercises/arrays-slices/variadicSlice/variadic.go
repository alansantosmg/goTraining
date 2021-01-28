package main

import "fmt"

// Define funcão com:
//- parametro prefix
//- multiplos parametros mum slice do tipo string
func terraform(prefix string, worlds ...string) []string {

	// cria um slice do tamanho dos parametros recebidos
	newWorlds := make([]string, len(worlds))

	// para cada word
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}

	return newWorlds
}

func main() {

	//Passa parametros diretamente para func terraform
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)

	// cria um Array
	// passa itens do array como parametro para terraform
	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("Planet", planets...)

	fmt.Println(newPlanets)
	fmt.Println()
	for i := range newPlanets {

		fmt.Println(newPlanets[i])
	}

}
