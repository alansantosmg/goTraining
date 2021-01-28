package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args

	// Imprime erro se numero de argumentos menor que 2
	if len(arguments) < 2 {
		fmt.Printf("usage: selectColumn column <file1> [<file2>[...<fileN]]\n")
		os.Exit(1)
	}

	// converte 1o argumento em int
	temp, err := strconv.Atoi(arguments[1])
	// se não converter imprime erro
	if err != nil {
		fmt.Println("Column value is not an integer:", temp)
		return
	}

	// se inteiro do 1o arg for menor que zero, imprime erro
	column := temp
	if column < 0 {
		fmt.Println("invalid Column number!")
		os.Exit(1)
	}

	// para cada arquivo dado como argumento a partir do 2o argumento
	for _, filename := range arguments[2:] {

		// imprime nome do arquio
		fmt.Println("\t\t", filename)

		// abre arquivo
		f, err := os.Open(filename)
		// se não abrir imprime erro e retoma nova iteração do for
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		// se adia fechamento durante a execução da função
		defer f.Close()

		// cria buffer de leitura
		r := bufio.NewReader(f)
		for {
			// le console ate encontrar '\n' e guarda em line
			line, err := r.ReadingString('\n')

			// se chegar ao fim do arquivo, encerra leitura (sai do loop)
			if err == io.EOF {
				// sai do loop de leitura
				break

				// se der erro de leitura
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}
			// quebra linha em com base em espaços em branco
			// armazena strings em um slice
			data := strings.Fields(line)

			// se a linha for maior ou igual a coluna
			if len(data) >= column {
				// imprime dado da coluna
				fmt.Println((data[column-1]))
			}
		}
	}
}
