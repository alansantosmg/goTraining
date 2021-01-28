
print('''
Gerador de tabuada - vrs 1.0
----------------------------''')
while True:
    
    try:
        # menu de opcoes
        print('''

Instruções:
- Digite um número inteiro de 1 a 10
- Para sair digite 0
                ''')
        fator = input('Tabuada de: ')
        fator_escolhido = int(fator)
        
        # verifica se tabuada é de 1 a 10
        if fator_escolhido < 0 or fator_escolhido > 10:
            print('Número fora do intervalo') 
            continue
        
        #gerador de tabuada
        elif fator_escolhido != 0:                         
                print('-----------------')
                multiplicador = 1
                while multiplicador < 10:
                    resultado = fator_escolhido * multiplicador
                    print(fator_escolhido, ' x ', multiplicador, '  =  ', resultado)
                    multiplicador = multiplicador + 1 
                    continue
                
        #saida do programa
        else:
            break 
        
    except ValueError:
        print('Entrada incorreta.')
print('obrigado por usar o tabuada!')

