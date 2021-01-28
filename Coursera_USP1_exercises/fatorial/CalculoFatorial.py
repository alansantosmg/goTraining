def main():
    # Digita número para multiplicacao fatorial
    entrada = int(input("Digite o valor de n: "))

    # inicializa contador do laço
    fator = entrada

    # Laco rodara fatores enquanto fator > 0
    # Menor fator não pode ser zero pois zero * n = 0
    # Decresce contador para proximo fator a multiplicar pela entrada
    # Multiplica fator 
    while fator > 1:
        fator = fator - 1
        entrada = entrada * fator


    # Mostra resultado
    print(entrada)
main()

