def main():
    
    # Define valor de entrada para ir direto ao laco
    # entrada = 1

    # Solicita entrada até usuário digitar numero com duas casas
    # O numero entrado pode ser positivo ou negativo
    # while entrada < 10 and entrada > -10:
    entrada = int(input("Digite um número inteiro: "))

    
    # Se numero entrado possuir 2 ou mais casas e for negativo converte para positivo
    # Instruções só funcionam corretamente com número positivo
    if entrada <= -10:
        entrada = entrada * -1

      
    # Define valores iniciais para isolamento de digitos
    moduloDigitoDireita = 10
    moduloDigitoEsquerda = 100
    isolamentoDigitoDireita = 1
    isolamentoDigitoEsquerda = 10
    
    # Define saida de fluxo 
    naoAchouAdjacente = True
      
    # Conta quantidade de digitos (casas) do valor digitado
    # necessario para saber quantas vezes teste rodará e variaveis de isolamento mudarão de valor    
    casas = 0
    numeroDivisaocasas = entrada
    while numeroDivisaocasas > 0:
        numeroDivisaocasas = numeroDivisaocasas // 10
        casas = casas + 1

    # laco roda enquanto numero de casas for maior que zero e nenhum par adjacente tenha sido encontrado
    while casas > 0 and naoAchouAdjacente == True:  
        # Separa digito a direita do numero
        digitoDireita = ((entrada % moduloDigitoDireita) // isolamentoDigitoDireita)
        # Separa digito adjacente ao digito direito do numero
        digitoEsquerda = ((entrada % moduloDigitoEsquerda) // isolamentoDigitoEsquerda)
        # Compara se os dois digitos separados do numero são iguais
        if digitoDireita == digitoEsquerda:
            # Se digitos forem iguais imprime mensagem
            print("sim")
            # encerra o laço relatando que não achou adjacente é falso
            # neste caso vai para teste de numero de casas = 0
            # como casas não foram zeradas não imprimirá mensagem de digitos nao encontrados
            naoAchouAdjacente = False
        else:
            # Se os 2 digitos separados não forem iguais continua teste
            # Contador de casas decresce para rodar somente com casas existentes
            # Se contador casas chegar a zero o laço é abortado
            # e vai para teste de casas.
            # como este é zero será impresso que nenhum numero adjacente igual foi encontrado.
            casas = casas -1
            # multiplica variaveis de isolamento de digito por 10
            # objetivo é testar próximas 2 casas além das 2 testadas
            # Exemplo: na primeira vez roda para comparar digitos da unidade e dezena
            # Ao multiplicar variaveis de isolamento por 10 passara a verifar dezena e centena e assim por diante
            moduloDigitoDireita = moduloDigitoDireita * 10
            moduloDigitoEsquerda = moduloDigitoEsquerda * 10                
            isolamentoDigitoDireita = isolamentoDigitoDireita * 10
            isolamentoDigitoEsquerda = isolamentoDigitoEsquerda * 10
          
    if  casas == 0:
        # se contador casas for zero significa que valor digitado não tem algarismos adjacentes iguais
        print("não")                
               
    
main()