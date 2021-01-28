def main():
    cartaoCliente = int(input("Digite o número do seu cartão de crédito: "))
    cartaoLido = 1
    encontreiMeuCartaoNaLista = False
    while cartaoLido != 0 and not encontreiMeuCartaoNaLista:
        cartaoLido = int(input("Digite o número do próximo cartão de crédito: "))
        if cartaoLido == cartaoCliente:
            encontreiMeuCartaoNaLista = True
    if encontreiMeuCartaoNaLista:
        print("Cartão encontrado",cartaoCliente)
    else:
        print("Cartão não encontrado.")
            
main()