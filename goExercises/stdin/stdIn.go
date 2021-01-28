package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//define varivel F do tipo *arquivos do SO
	// define que f será o arquivo os.Stdin
	//

	var f *os.File
	f = os.Stdin
	defer f.Close()

	// cria uma variável scanner
	// e atribui a ela um novo metodo que scaneia o que entrar na variavel F,
	// ou seja o que for input de usuario

	scanner := bufio.NewScanner(f)

	// cria um loop infinito que vai rodar os comandos dentro dele cada vez que o usuario
	// digitar algo
	for scanner.Scan() {

		// imprime > seguindo do texto obtido pelo scanner que fica na propriedade Text.

		fmt.Println(">", scanner.Text())

		// após impressão o loop volta e repete o comando

	}

}
