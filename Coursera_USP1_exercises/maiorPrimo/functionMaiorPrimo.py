
def eprimo(k):
        # Inicializa passagem de laço
        primo = True
        # Se k for zero ou 1 não é primo por definicao matematica
        if (k <= 1) == True:
            primo = False
        else:
            # Se k for maior que 1 testa divisores até metade do k
            # por definição nenhum k real é divisivel por mais que sua metade
            divisor = 2
        while divisor <= (k / 2):
            if k % divisor != 0:
                divisor = divisor + 1
            else:
                # se encontrar um divisor além de 1 até metade do k
                # k será considerado não primo
                # laco será encerrado
                primo = False
                divisor = (k / 2 + 1)
        if primo == False:
            return 0
        else:
            return 1


def maior_primo(m):
    # Guarda primo ou subtrai -1 do k e testa primo novamente
    while (m > 1):

        if (eprimo(m) != 0):
            return m                       
        else:
            m= m- 1
