package main

import (
	"io"
	"os"
)

func main() {

	myString := ""

	//armazena argumentos
	arguments := os.Args

	// se argumentos for 1
	// o proprio comando é um argumento
	if len(arguments) == 1 {

		myString = "Please, give me one argument!"

	} else {

		myString = arguments[1]

	}

	// se for mais de um argumento
	// escreve argumento no stdOut (tela_
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")

}
