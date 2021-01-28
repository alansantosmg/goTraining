package main

import "fmt"

func main() {

	// Channel Buffer demo
	// If we don't create a buffer, we will get a deadlock after
	// we send the message

	// The program will wait for the message return, but
	// it will not execute the code with the message return.

	// Using buffers, the program keeps the message in the buffer and
	// goes on and can execute the code with the message return.

	// Best practice: Avoid create channels and return channels in the main function.

	// Demonstração de buffer de canal
	// Se não for criado buffer a mensagem enviada pelo canal bloqueia o programa (deadlock)
	// porque ele ficará esperando uma resposta e não avançará para o retorno do canal

	//Se o local da resposta está depois do ponto de envio, o programa não chegará la.

	// melhor pratica é usar canais em funções fora de main
	// retornando para outras funcoes ou para main

	//************** EXEMPLO ****************************************************


	
	// define um canal com buffer 2
	canal := make(chan string, 2)



	// envia duas strings dentro canal
	canal <- "ola Mundo"
	canal <- "ola Mundo"
	
	
	// avisa fechamento do canal
	close(canal)

	
	
	// recebe saida do canal quantas vezes necessário
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}
