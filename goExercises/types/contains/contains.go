package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Encontre você mesmo em uma caverna pouco iluminada.")

	var command = "Caminhar para fora"
	var exit = strings.Contains(command, "fora")
	fmt.Println("you leave the cave:", exit)

}
