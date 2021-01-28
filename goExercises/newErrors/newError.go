package main

import (
	"errors"
	"fmt"
)

// função exemplo que tem geração de erro
func returnError(a, b int) error {
	if a == b {
		// cria variavel err do tipo error
		// atribui novo erro String a variavel
		err := errors.New("error in returnError() function")
		// retorna err
		return err
	}
	return nil

}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}
	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}
	// verifica estado da variavel err com método Error()
	if err.Error() == "error in returnError() function" {
		fmt.Println("!!")
	}
}
