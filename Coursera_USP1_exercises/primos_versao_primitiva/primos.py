def main():


    numero = int(input("Digite um número inteiro: "))   
    
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
        print("não primo")
    else:
        print("primo")

    
main()
