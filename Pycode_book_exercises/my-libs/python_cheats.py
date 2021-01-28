#
# convert avoid case errors on string inputs
# use upper() or lower method
# example:
#
name = 'Rob'
if name.upper() == Rob:
        return name




# 
# using randon function
#
import random
random.randint(1,6) # show random number between 1-6




#
# time funcion 
#
import time
mytime = time.localtime() # get time
hour = mytime.tm_hour   # extract hours
minute = mytime.tm_minute # extract minutes
time.sleep(30) # wait 30 seconds




# 
# print multiline and preserve blank lines
#
print('''This 

is the test.''')




#
# metodo e funcões funcionam da mesma forma
# função não depende da existencia de objeto.
# exemplo: print input
# metodo depende da existencia de objeto.
# exemplo: upper()
# depende de (ou é aplicado a) um objeto string
#




#
# forma mais verborragica
# é mais recomendado porque: 
# computador é mais rapido executando 2 comandos curtos do que um longo
# comandos curtos usa menos memória na execucao
# é possível reaproveitar a string entrada
#
ride_number_text = input('please enter the ride number you want: ')
ride_number = int(ride_number_text)




#
# forma mais longa
# menos linhas de código
# uma variável a menos
# execucao mais lenta pois é instrução mais longa
# valor entrado em string é perdido, só guarda como int
# ocupa mais memória na execução
# 
ride_number_text = int(input('Please enter the ride number you want:'))






#
# melhor forma de validar  entradas do usuario e tratar execoes
# 
#
ride_number_flag = False # cria flag como falso
while ride_number_flag == False: # enquanto a flag for falsa solicita entrada
    try:
        ride_number_text = input('please enter the ride number you want: ')
        ride_number = int(ride_number_text)
        ride_number_flag = True # muda flag pra True quando entrada for inteiro
        print('You have entered', ride_number)
    except ValueError: # excecao se não entrar com valor compativel com tipo 
        print('invalid number') # msg exceçao volta no loop, pois nao executou flag true
    except KeyboardInterrupt: # excecao se pressionar control C para interromper 
        print('Please, do not try stop the program') 





# 
# no python 2 precisa especificar no inicio do programa que é unicode 
# se tiver carcteres além do ASCII padrão
# no python 3 não precisa da declaração de unicode
# 
#





#
# validando entrada de usuário - 2a forma
# controlando o loop com break
#
while True: # roda loop até ter um break
    try:
        ride_number_text = input('please enter the ride number you want: ')
        ride_number = int(ride_number_text)
        # se não entrou int gera excecao e volta no loop eterno
        print('You have entered', ride_number)
        break # quando entra int execucao prossegue e break quebra o loop
    except ValueError: # excecao se não entrar com valor compativel com tipo 
        print('invalid number') # msg exceçao volta no loop, pois nao executou flag true
    except KeyboardInterrupt: # excecao se pressionar control C para interromper 
        print('Please, do not try stop the program') 




#
# uso do "break"
# Terminando loop mais cedo usando if e break
#
count=0
while count<5:
    print('Inside loop')
    count = count+1
    if count == 3:
        break
print('Outside loop')




#
#
# # Uso do "continue"
# Limitando entradas do usuario a algumas opcoes
# e ao mesmo tempo controlando excecoes com While, if, continue e break
#
#
while True:
    try:
        escolhaopcao = input('entre com opcao 1 ou 2:  ')
        opcaoescolhida = int(escolhaopcao)
        if (opcaoescolhida > 2 and opcaoescolhida < 1 ):
            print('Opção não disponível.')
            continue
        break
    except ValueError:
        print('Opção inválida.')
print('Voce escolheu opcao: ', opcaoescolhida)  



#
# caso especial laço for
# quando count for 5 volta ao inicio do loop
# resultado 1,2,3,4,6,7,8,9,10,11,12
# o for (vai guardando vezes que loop rodou até o final)
# o continue volta o loop antes que chegue ao final
# daí ele pula a contagem e continua
# 
for count in range(1,13):
    if count == 5:
        continue
    print(count)
print('finished')





#
# caso estranho do FOR
# resultado vai ser hello na vertical
# cada letra da string será um item 
# como uma lista separada por virgulas
#
for letter in 'hello':
    print(letter)
    



#
# Erro de recursão. Função chamando a si mesma
# estoura a pilha do buffer de retorno
#
def m1():
    print('hello')
    m1()

m1()




#
# notação de nome para função
# Verbo (acao)_ nome
# 
def display_menu():




#
# limpando a linha, imprimindo e retornando carro
#
    print('\033[K',hora, ':', minuto,':', segundo , end='\r')
time.sleep(1)




#
# funcao com parametro
#
def teste(x, y=5):

#
# Cuidado com shadowing
# no caso abaixo o resultado será
# 100
# 99 
# porque o computador cria var local e global com mesmo nome
# 
    chesse = 99
    def func():
        chesse = 100
        print('local Chesse is',chesse)
    func()
    print('Global chesse is:', chesse)




#
# para usar variavel global (com mesmo nome)
# dentro de uma funcao, usar parametro:  global
# No caso abaixo imprimirá o mesmo valor 2x
chesse = 99
def func():
    global chesse
    print(chesse)
func()
print(chesse)

# sempre que usar global deixe bem comentado para 
# para não confundir-se no codigo.




#
# a primeira statement de uma funcao pode ser 
# documentada. Essa documentação depois pode 
# ser lida por utilitarios. Exemplo abaixo: 
#
def read_text(prompt):
    '''
    Mostra um prompt que lê uma string de 
    texto e valida se o usuário não está
    usando sequencias de escape para abortar
    '''
# segundo exemplo comentário de linha única.
def read_text(prompt):
    'Mostra um prompt que lê uma string de...'




#
# pydoc pode ler documentação de funções
#
import pydoc
pydoc.help(read_text)  # vide exemplo acima

#
# pydoc pode ajudar a conhecer funçoes internas 
#
import pydoc
pydoc.help(print)


# 
# importar módulos I
#
#
import ARQUIVO
name = ARQUIVO.funcao1()


#
# importar módulos II
#
from ARQUIVO import funcao1
name = funcao1()




## 
# criar e popular listas
#
sales[]
sales.append(99)





# Bubble sort compara dois itens e dois 
#
# remaneja do menor para maior ou do maior para menor
# basta trocar o sinal do if
# para rodar para qualquer lista
# introduzir um laço for com um contador
# que rodara no range 0 (posicao inicial)
# até o range da lista  lenght - 1
# menos 1 pq compara de 2 em dois
# senão na ultima comparação pega item fora da lista
# e dá erro
#
# if lista[0] < lista[1]:
#    temp = lista[0]
#    lista[0] = lista[1]
#    lista[1] = temp
#
# importante: como os ultimos itens 
# ficam ordenados a cada passada, 
# ir diminuindo o range de pesquisa da função melhora o desempenho
#Exemplo abaixo (veja que é só subistuir por count):
#
def sort_high_to_low():
    for _passagem in range(0,len(sales)):
        done_swap = False
        for count in range(0,len(sales)-1-_passagem):
            if sales[count] < sales[count+1]:
                bucket = sales[count]
                sales[count] = sales[count+1]
                sales[count+1] = bucket
                done_swap = True
                print(sales)
        if done_swap = False:
            break

#
# Trabalhando com moedas em python
# primeira solução
#
#
import decimal
valor = 1.5675654
convertido = Decimal(valor).quantize(Decimal('0.01'))

#
# uma funcao é algo que um programa pode fazer
# um metodo é algo que um objeto pode fazer
#


#
# gravar lista em arquivo
#

def save_sales(file_path):
    print('Salvar vendas em: ', file_path)
    my_output_file = open(file_path,'w')
    for sale in sales:
        my_output_file.write(str(sale)+'\n')
    my_output_file.close()

#
#
# carregar arquivo em lista
#
def load_sales(file_path):
  
    print('Carregar vendas de: ', file_path)
    try:
        my_input_file = open(file_path, 'r')
        sales.clear()
        for line in my_input_file:
            line = line.strip() # retira o \n de cada linha do arquivo. Tb vale: lstrip rstrip
            print(line)
            sales.append(int(line)) 
        my_input_file.close()
    except:
        print('Algo errado ao carregar o arquivo.')
    finally:
        my_input_file.close()

    #
    # carregar um aquivo inteiro (com \n e tudo)
    #
    #
    input_file = open(file_path)
    total_file = input_file.read()
    print(total_file)
    input_file.close()


    #
    # Salvar arquivo. Mostra tambem como usar
    # a proteção de exeções com Try
    # Uso do with para tratar abertura de arquivo como objeto
    # E garantir que em caso de exceção
    # O arquivo seja fechado
    # Neste caso não e preciso usar finally para tentar fechar
    #
    def save_sales(file_path):
    print('Salvar vendas em: ', file_path)
    try:

         with open(file_path, 'w') as my_output_file:
            for sale in sales:
                my_output_file.write(str(sale)+'\n')        
            my_output_file.close()
            print('\n Vendas salvas no arquivo ', file_path)

    except:
        print('\nAlgo errado ao salvar o arquivo\n')


# 
#  usar import os.path para garantir avisar arquivo já 
#  existe. util para gerar um prompt para usuário 
#  perguntando se ele quer sobreescrever o arquivo.
#  
import os.path
if os.path.isfile('vendas.txt'):
    print('O arquivo já existe')

#
# None para definir uma variável vazia
#
contact = !None

#
# Tirar espaços de uma cadeia de strings
# usado para limpar uma cadeia de strings para busca
#
name = name.strip() # tira espacos a direita e esquerda
name = name.rstrip() # tira espaços a direita
name = name.lstrip() # tira espaços a esquerda


#
#  usando sets
#
#
set1.add(1) #cria set e adiciona 1 elemento
set1={1,2,3} #define set com vários elementos
set1.difference(set2) # lista elementos não comuns entre sets
set1.union(set2) #união de elementos do set (não lista elemento repetido)
set1.intersection(set2 # elementos em comum entre sets
set1.isdisjoint(set2) #set1 não está contido em set2
set1.issuperset(set2) # set2 está contido em set1
set1.issubset(set2) #set1 está contido em set2

if apple_potion.issubset(pocket):
	print('you have the ingredientes for the apple potion')

# exemplo de uso de subsets
# jogador de video game tem na bolsa alguns itens
# para fazer uma pocao ele precisa de 2 itens
# usando subsets é possível saber se entre os items que tem da pra fazer a poção
# é um exemplo de filtragem de dados usando SETs. 
pocket = {'axe', 'apple', 'herbs', 'flashlight'}
apple_potion = {'apple', 'herbs'}
if apple_potion.issubset(pocket):
	print('you have the ingredientes for the apple potion')

#Para encontrar determinados elementos dentro do set usando exemplo acima
apple_potion.intersection(pocket)
# irá retornar maçãs e ervas




