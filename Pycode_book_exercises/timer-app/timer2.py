# importa funções
import time

# Msg boas vindas
print("\n****** Timer Simples *******\n")

# Entradas do usuario
lembrete = input('Quero me lembrar de: ')
user_timer = input('Quero ser lembrado daqui (segundos): ')

# Iniciar timer
print('\nSeu timer foi setado! Pressione enter para iniciar...')

timer_interno = int(user_timer) # converte entrada do usuário em int

time.sleep(timer_interno) # seta timer

# Conclui timer
print('''
Fim do tempo! Atenção! Você deve ser lembrado de:

''',lembrete )
