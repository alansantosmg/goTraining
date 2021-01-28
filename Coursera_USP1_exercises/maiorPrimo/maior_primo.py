def main(): 
    firstNumber = int(input("Entre com primeiro número: "))

    #funçao testa primo
    def maior_primo(numero):   
        # Inicializa passagem de laço
        primo = True
        # Se numero for zero ou 1 não é primo por definicao matematica
        if (numero <= 1) == True:
                primo = False
        else:
            # Se numero for maior que 1 testa divisores até metade do numero
            # por definição nenhum numero real é divisivel por mais que sua metade
            divisor = 2
        while divisor <= (numero /2): 
            if numero % divisor != 0: 
                divisor = divisor + 1
            else:
                # se encontrar um divisor além de 1 até metade do numero
                # numero será considerado não primo 
                # laco será encerrado
                primo = False
                divisor = (numero /2 + 1)         
        if primo == False:
            return 0
        else:
            return 1
    #Guarda primo ou subtrai -1 do numero e testa primo novamente
    while (firstNumber > 1):
        
        if (maior_primo(firstNumber) != 0): 
            print(firstNumber)
            firstNumber = 0
        else: 
            firstNumber = firstNumber -1 

main()