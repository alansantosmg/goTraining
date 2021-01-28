

# funcao computador escolhe jogada
# retirar todas peças se limite de peças for > ou =  total de peças
# se limite de peças for menor, deixar no tabuleiro limite + 1
def computador_escolhe_jogada(n,m):                   
    if m >= n: 
        return n
    else: 
        n = n - (m +1)
        return n
 

# funcao usuario escolhe jogada
# entra com numero de peças que jogador quer retirar
# se numero for zero ou maior que limite de peças, solicita nova entrada
# retira do tabuleiro numero de peças até o limite m
def usuario_escolhe_jogada(n,m):
    try:
        moveHuman = int(input("Quantas peças você vai tirar?"))
    except ValueError:
        print("Somente numeros são aceitos. Tente novamente.")
    while moveHuman > m or moveHuman <= 0: 
        moveHuman = int(input("Oops! Jogada inválida! Tente de novo."))
    return moveHuman




# funcao que roda 1 partida     
def partida(): 
    # solicita numero de peças no tabuleiro
    # enquanto pecas igual ou menor que zero, solicita numero de pecas
    pecas = 0
    limitePecas = 0
    while pecas <=0: 
        try:
            pecas = int(input("Quantas peças?"))    
        except ValueError:
            print("Somente numeros são aceitos. Tente novamente.")

    # solicita limite de peças a retirar por jogada
    try:
        limitePecas = int(input("Limite de peças por jogada?"))
    except ValueError:
        print("Somente números são aceitos. Tente novamente.")
    # enquanto limite de peças for maior que peças no tabuleiro pede novo limite
    # ou limite de peças for igual ou menor que zero pede novo limite
    while limitePecas > pecas or limitePecas <= 0:
        try:
            limitePecas = int(input("Limite de peças por jogada deve ser menor ou igual ao numero de peças."))
        except ValueError: 
            print("Somente números são aceitos. Tente novamente.")
    # se pecas for multiplo de limite de pecas +1 usuario inicia a partida
    # se não, computador inicia a partida
    if (pecas % (limitePecas + 1) == 0): 
        usuario = "Voce"
    else:
        usuario = "Computador"
    print(usuario, "começa!") 

    # enquanto houver pecas no tabuleiro jogadores se alternam
    while pecas > 0:
        jogada = 0
        # chama funçao de jogada do humano
        if usuario == "Voce":          
            jogada = usuario_escolhe_jogada(pecas,limitePecas)
            if jogada > 1: 
                print(usuario, "retirou", jogada, "peças.")
            else:
                print(usuario, "retirou", jogada, "peça.")
            # alterna para proximo jogador
            usuario = "Computador"

        else:
            # chama funcao de jogada do computador
            jogada = computador_escolhe_jogada(pecas,limitePecas)
            if jogada > 1:
                print("Computador tirou", jogada, "peças.")
            else:
                print("Computador tirou 1 peça.")
            # alterna para humano
            usuario = "Voce"
       
        # retira peças do tabuleiro a cada jogada
        pecas = pecas - jogada            
        # mostra quantas peças restam no tabuleiro
        if pecas > 1:
            print("Agora restam", pecas, "no tabuleiro.")
        elif pecas == 1: 
            print("Agora resta apenas", pecas, "no tabuleiro.")
    
    # quando não houver mais peças:
    # como no final da jogada, jogador é alternado, 
    # ao final do jogo reverte status para último jogador (que zerou tabuleiro)
    # mostra jogador que ganhou (o que zerou tabuleiro)        
    if usuario == "Computador":
        print("Fim de jogo! Você ganhou!")
        return 0
    else:
        print("Fim de jogo! O computador ganhou!")
        return 1




# funcao campeonato
# chama funcao partida em 3 rodadas consecutivas
# Ao final das rodadas mostra quem ganhou o campeonato
def campeonato():
    partidasCampeonato = 1
    usuario = 0
    computador = 0
    while partidasCampeonato <= 3:
        if partidasCampeonato == 1:
            print()
            print(" **** Primeira rodada ****")
            print()
            partidasCampeonato +=1
            if (partida() == 0):
                usuario +=1
            else:
                computador +=1
        
        elif partidasCampeonato == 2:
            print()
            print(" **** Segunda rodada ****")
            print()
            partidasCampeonato +=1
            if (partida() == 0):
                usuario +=1
            else: 
                computador +=1
        
        else:
            print()
            print(" **** Terceira rodada ****")
            print()
            partidasCampeonato +=1
            if (partida() == 0):
                usuario += 1
            else:
                computador +=1
    print()
    print("**** Final do campeonato! ****")
    print()
    print("Placar: Você", usuario, "X", "Computador", computador)






def main(): 
    # menu inicial para escolha de partida ou campeonato.
    escolha = 0
    while escolha != 1 and escolha != 2:
        print()
        print()
        print ("Bem-vindo ao jogo do NIM! Escolha:")
        print()
        print("1 - para jogar uma partida isolada")
        try:
            escolha = int(input("2 - para jogar um campeonato"))
        except ValueError:
            print("somente numeros são aceitos. Tente novamente.")
    if (escolha == 1): 
        print()
        print("Voce escolheu partida")
        print()
        partida()
    elif (escolha == 2): 
        print()
        print()
        print("Voce escolheu um campeonato!")
        print()
        campeonato()
    
        
main()


