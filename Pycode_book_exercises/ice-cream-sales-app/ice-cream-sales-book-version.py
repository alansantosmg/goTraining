
# Ice cream Sales

from valida_entradas import read_int, read_range,read_text
import math
from decimal import *
import os.path

#getcontext().prec = 2

# variaveis globais
sales = [] # lista de vendas




def show_sales():
    '''
    imprime o titulo do relatório de vendas por stand
    inicializa contador 
    mostra numero do respectivo stand e sua venda
    adiciona contador para próximo numero de stand
    '''
    print('\nsales figures\n')
    count = 1
    for sales_value in sales: 
        print('Sales for stand', count, 'are', sales_value)
        count = count + 1
    input('\nPressione qualquer tecla para continuar')



def sort_low_to_high():
    '''
    Imprimir lista de vendas ordenadas da menor para maior
    '''
    for _passagem in range(0,len(sales)):
        done_swap = False
        for count in range(0,len(sales)-1-_passagem):
            if sales[count] > sales[count+1]:
                bucket = sales[count]
                sales[count] = sales[count+1]
                sales[count+1] = bucket
                done_swap = True
                
        if done_swap == False:
            break
    print('Vendas ordenadas da menor para maior')
    print('')
    show_sales()



def sort_high_to_low():
    for _passagem in range(0,len(sales)):
        done_swap = False
        for count in range(0,len(sales)-1-_passagem):
            if sales[count] < sales[count+1]:
                bucket = sales[count]
                sales[count] = sales[count+1]
                sales[count+1] = bucket
                done_swap = True
                
        if done_swap == False:
            break
    print('Vendas ordenadas da maior para menor')
    print('')
    show_sales()



def highest(): 
    '''
    Imprimir maior venda
    '''
    highest = 0
    for count in range(0, len(sales)): 
            if highest < sales[count]:
                highest = sales[count]
                count = count +1
    print('\nMaior venda: ', highest)



def lowest(): 
    '''
    Imprimir menor venda
    '''
    lowest = sales[0]
    for count in range(0, len(sales)): 
            if lowest > sales[count]:
                lowest = sales[count]
                count = count +1
    print('\nMenor venda: ', lowest)
   


def total_sales():
    '''
    Imprimir total de vendas
    '''
    # print('\nTotal das vendas: '+ str(math.sum(sales)))
    value = 0
    for count in range(0, len(sales)):         
        value = value + sales[count]
        count = count +1
    print('\nTotal de vendas é: ', value)
    input('\nPressione qualquer tecla\n')



def avegare_sales():
    '''
    Imprimir média das vendas
    '''
    # print('\nMedia das vendas: ' + str(math.fsum (sales)))
    value = 0
    for count in range(0, len(sales)): 
                value = sales[count] + value
    value = value / len(sales)
    media = Decimal(value).quantize(Decimal('0.01'))
    print ('\nMedia das vendas: ', media)
    input('\nPressione qualquer tecla\n')



def get_sales(nr_sales):
    '''
    recebe como parametro quantidade de stands com vendas
    limpa a lista de vendas
    inicia o contador para cada stand e suas vendas
    Solicitar entrar venda para cada stand até entrar venda de todos
    '''
    sales.clear()
    for count in range(1,nr_sales+1):
        prompt='\nEnter the sales for stand ' + str(count) + ': '  
        sales.append(read_int(prompt))
    print('\nEntradas finalizadas.')
    input('Pressione qualquer tecla para continuar')
    


def save_sales(file_path):
    print('Salvar dados no arquivo (nomedoarquivo.txt): ', file_path)
    if os.path.isfile(file_path):
        escolha = read_text('O arquivo já existe, deseja continuar? (s/n)')
        if escolha.upper() == 'S':  
            try:

                with open(file_path, 'w') as my_output_file:
            
                    

                    for sale in sales:
                        my_output_file.write(str(sale)+'\n')        
                    my_output_file.close()
                    print('\nArquivo salvo.', file_path)
                    

            except:
                print('\nAlgo errado ao salvar o arquivo\n')
        else:
            print('\nOk. Arquivo não foi sobreescrito.')
        input('Pressione qualquer tecla para continuar.')



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
        print('\nArquivo ' + file_path + 'carregado.')
    except:
        print('\nAlgo errado ao carregar o arquivo.\n')
    finally:
        my_input_file.close()



''' Another form of use file with exception
    def save_sales(file_path):
        print('Salvar vendas em: ', file_path)
        try:
        my_output_file = open(file_path,'w')
        for sale in sales:
            my_output_file.write(str(sale)+'\n')        
        my_output_file.close()
        print('\n Vendas salvas no arquivo ', file_path)

   except:
        print('\nAlgo errado ao salvar o arquivo\n')
    finally:
        my_output_file.close()
'''



# Menu principal
while True:  
    print('''
    Ice-Cream Sales
    
    1: Print the Sales
    2: Sort Low to high
    3: Sort high to low
    4: Hightest and Lowest
    5: Total Sales
    6: Average Sales
    7: Enter Figures
    8: Save Sales
    9: Load Sales
    0: Good bye
    ''')

    # prompt do menu
    choose_menu = read_range('Enter your command: ', min_value = 0, max_value = 9)
    
    if choose_menu == 1:
        show_sales()
    elif choose_menu == 2:
        sort_high_to_low()
    elif choose_menu == 3:
        sort_low_to_high()
    elif choose_menu == 4: 
        highest()
        lowest ()
        input('\nPressione qualquer tecla\n')
    elif choose_menu == 5: 
        total_sales()
    elif choose_menu == 6:
        avegare_sales() 
    elif choose_menu == 7: # item de menu Enter Figures
        # prompt quantidade de stands que tiveram venda
        prompt = 'Entre quantidade de stands com vendas registradas: '
        # registrar e imprimir vendas até terminar stands com vendas
        get_sales(read_int(prompt))        
        continue
    elif choose_menu == 8: 
        prompt = 'Entre com nome do arquivo: '
        save_sales(read_text(prompt))
    elif choose_menu == 9: 
        prompt = 'Entre com nome do arquivo: '
        load_sales(read_text(prompt))
    elif choose_menu == 0: # item de menu Good bye
        break
    continue

# exit
print('Obrigado por usar o sistema') 
  
