package main

import (
	"fmt"
	"time"
)

func main() {

	// cria variavel de canal e informa o que vai passar no canal
	canal := make(chan string)

	// passa canal como parametro pra funcao
	// diz pra funcão qual canal será usado
	go escrever("olá mundo", canal)




	/**************************************************
	// recebe mensagem enquanto canal estiver aberto
	for {
		message, open := <-canal

		// verifica se canal está aberto
		// se não estiver mais e o canal enviar um close
		// fecha o canal
		if !open {
			break
		}

		// executa codigo recebido pelo canal
		// continua o programa apos receber todas mensagens
		fmt.Println(message)
	}
	****************************************************/


	// outra forma mais compacta de verficar se o 
	//canal continua aberto para receber mensagens
		for message := range canal {
		fmt.Println(message)
	}

	fmt.Println("fim do programa.")

}

// funcao que vai ter canal precisa receber
// um canal como parametro e o que ele vai passar
func escrever(text string, canal chan string) {

	for i := 0; i < 5; i++ {

		// recebe parametro e envia pelo canal
		canal <- text
		time.Sleep(time.Second)
	}
	// depois que rodar 5x fecha o canal
	close(canal)
}



func escrever2(text string) {

	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}

}
