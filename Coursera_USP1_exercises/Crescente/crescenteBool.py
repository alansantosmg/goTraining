def main():
    decrescente = True
    anterior = int(input("Digite o primeiro numero da sequencia: "))
    valor = 1
    while valor != 0 and decrescente == True:
        valor = int(input("Digite o proximo número da sequência: "))
        if valor > anterior:
            decrescente = False
        anterior = valor
    anterior = valor
    if decrescente == True: 
        print("A sequencia está em ordem decrescente")
    else:
        print("A sequencia está em ordem crescente")
main()

