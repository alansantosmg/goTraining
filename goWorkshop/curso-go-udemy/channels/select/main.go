package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	// funcao que espera meio segundo
	// e manda mensagem para canal 1
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "canal 1"
		}
	}()

	// cria funcão que espera 2 segundos
	// e manda mensagem para canal 2
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "canal 2"
		}
	}()

	for {

		// se não usasse o select, a função 1 que é mais rapida
		// deveria esperar a funcao dois executar e haveria perda de tempo
		// com select, caso a mensagem seja para canal 1 executa
		// caso seja para canal 2 executa.
		// como o funcao 1 não tem que esperar a funcao 2 fica mais rapido.
		select {

		case mensagemcanal1 := <-canal1:
			fmt.Println(mensagemcanal1)
		case mensagemcanal2 := <-canal2:
			fmt.Println(mensagemcanal2)
		}

	}

}
