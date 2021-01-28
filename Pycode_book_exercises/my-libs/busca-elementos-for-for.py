lista1 = [1,2,3,4]
lista2 = [1,5,6]

def verifica_lista(lista1,lista2):
    for elemento1 in(lista1):
        for elemento2 in(lista2):
            if elemento1 == elemento2:
                resultado = True
                return print(resultado)
    resultado = False
    return print(resultado)
verifica_lista(lista1,lista2)

