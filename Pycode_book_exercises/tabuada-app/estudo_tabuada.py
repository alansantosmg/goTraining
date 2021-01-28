import random
print('''
Treino de tabuada - vrs 1.0
---------------------------
''')

print('''Instruções
Digite sua resposta ou tecle x para sair:
-----------------------------------------
        ''')
acertos = 0
erros = 0
while True:
    try:
        
        fator_primeiro = random.randint(0,10)
        fator_segundo = random.randint(0,10)
        resultado = fator_primeiro * fator_segundo
        
        

        print('Pergunta:',fator_primeiro, ' x ', fator_segundo, ' = ? ' )
        entrada_resposta = input('Resposta: ')
    
        # verifica se usuário quer sair
        if entrada_resposta.upper() == 'X':
            break
        # se continuar manda perguntas
        resposta = int(entrada_resposta)
        if resposta == resultado: 
            print('Resposta correta!\n ')
            acertos = acertos +1
            continue
        print('Resposta incorreta. A resposta certa é:', resultado, '\n' )
        erros = erros +1
        continue
    except ValueError:
        print('\nInválido. Digite um número inteiro positivo ou x para sair.\n')


print('\n\nEstatísticas:')
print('-------------')
tentativas = acertos + erros
try:
    percentual_acertos = 0
    percentual_acertos = (acertos / tentativas) *100
except ZeroDivisionError:
    print('')
print('Tentativas: ', tentativas)
print('Respostas corretas: ', acertos)
print('Respostas incorretas: ', erros)
print('Progresso: ', int(percentual_acertos), '%')
if percentual_acertos == 100:
    print('\nParabéns, a tabuada está na ponta da língua!\n\n')
elif percentual_acertos > 89:
    print('\nVocê está indo muito bem! Quase perfeito!\n\n')
elif percentual_acertos > 69:
    print('\nQuase lá! Você está indo bem. Estude só mais um pouquinho...\n\n')
elif ((percentual_acertos > 49) and (percentual_acertos < 70)):
    print('\nVocê está no caminho. Precisa estudar mais!\n\n')
elif (percentual_acertos < 49 and tentativas !=0):
    print('\nVocê precisa melhorar. Não desanime. Continue estudando! Você chega lá!\n\n')
else: 
    print('\nVolte depois para testar seu conhecimento! Hasta la vista, baby!\n\n')

