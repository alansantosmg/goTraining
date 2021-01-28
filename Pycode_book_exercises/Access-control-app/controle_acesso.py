import valida_entradas
import pickle

#menu

class User:
    def __init__(self, username, password):
        self.username = username
        self.password = password

def cadastrar_user():
    cadastrar_novo = 's'
    while cadastrar_novo.lower() is True:
        print('Cadastro de usuário')
        login = read_text('Digite um novo login: ')
        if login in usuarios:
            print('Usuário já existe: ')
            continue        
        username = read_text('Entre com nome completo: ')
        password = read_text('Entre com nova senha: ')
        new_User = User(username = username, password = password)
        usuarios[login]=User()  
        cadastrar_novo.lower() = read_text_yn('Deseja cadastrar outro usuário? (s/n)','s','n')        
        
    with open(file_path, 'wb') as output_file:
        pickle.dump(usuarios, output_file)





def Validar_entrada():
    print('Validar Entrada')
    login = read_text('Entre com login: ')
    count = 0
    while count < 4:
        if login in usuarios: 
            senha = read_text('Digite sua senha: ')
                if senha == usuario[login].password: 
                    print ('acesso validado! ')
                break
        count = count + 1
    


    pass

def cadastrar_user():
    print('Cadastro de usuário')
    pass


usuarios = {}
file_path = 'users.pickle'

menu = print('''
1. Cadastrar usuários
2. Salvar cadastro
3. Validar entrada
4. Sair

Escolha sua opção: 
''')

while True: 
    command_menu = read_range_int(prompt,min_value=1, max_value=9)
    if command_menu == 1
        pass
    elif command_menu == 2
        pass
    elif command_menu == 3
        pass
    elif command_menu == 4
        pass

