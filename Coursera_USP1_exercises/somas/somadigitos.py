def main():
    
   
    # Entrada de numero cujos digitos serão somados
    entrada = int(input("Digite um número inteiro: "))

          
    # Define valores iniciais para isolamento de digitos
    moduloDigitoDireita = 10
    isolamentoDigitoDireita = 1
 
      
    # Conta quantidade de digitos (casas) do valor digitado
    # necessario para saber quantas vezes teste rodará e variaveis de isolamento mudarão de valor    
    casas = 0
    numeroDivisaocasas = entrada
    while numeroDivisaocasas > 0:
        numeroDivisaocasas = numeroDivisaocasas // 10
        casas = casas + 1


    # Inicializa variavel que receberá soma de digitos
    SomaDigitos = 0
    
    # laco roda enquanto numero de casas for maior que zero 
    while casas > 0:  
        # Separa digito a direita do numero
        digitoDireita = ((entrada % moduloDigitoDireita) // isolamentoDigitoDireita)
        # Soma digito separado
        SomaDigitos = SomaDigitos + digitoDireita
        # muda variaveis para isolamento da próxima casa a esquerda
        moduloDigitoDireita = moduloDigitoDireita * 10                     
        isolamentoDigitoDireita = isolamentoDigitoDireita * 10
        # decresce contador de casas que necessitam separar digito
        casas = casas -1
    print(SomaDigitos)

                   
               
    
main()