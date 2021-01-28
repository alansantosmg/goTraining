from valida_entradas import *





class Contact:
    pass




contacts = []






def new_contact():
    ''' Cadastra novo contato
        Recebe nome, telefone, email
    '''
    print('\nCriar um novo contato')

    
     #instancia classe contato para new_contact
    new_contact = Contact()  

    # entra com atributos de new_contact
    new_contact.name = read_text('Entre com nome do contato: ')
    new_contact.address = read_text('Entre com endereço: ')
    new_contact.phone = read_text('Entre com telefone: ')

    # armazena objeto new_contact em lista
    contacts.append(new_contact)

    # informa entrada concluída
    print('\nContato armazenado.')
    input('Pressione qualquer tecla para continuar. ')
  






def find_contact():
    ''' Busca novo contato
        Recebe string e retorna contato
    '''

    # entra com contato a buscar
    print('\nEncontrar contato')    
    search_name = read_text('Entre com nome do contato: ')
    search_name = search_name.strip()
    search_name = search_name.lower()

    result = None

    for contact in contacts:
        name = contact.name
        name = name.strip()
        name = name.lower()
   
        if name.startswith(search_name):
            # se contato exite exibe contato
            result = contact
            break
    if result != None:        
        print('\nName: ',result.name)
        print('Address: ',result.address)
        print('Phone: ',result.phone,'\n')
    else: 
        # se contato nao existe informa usuario        
        print('\nContato não encontrado.')

    input('Pressione qualquer tecla para continuar. ')
  

def save_contacts():
    input('Função não implementada. Pressione qualquer tecla.')
    pass




# menu principal
menu ='''
Tiny Contacts

1. Novo contato
2. Encontrar contato
3. Salvar contatos
4. Sair do programa

Escolha sua opção: '''


while True: 
    menu_command = read_range_int(menu, min_value = 0, max_value = 4 )
                                    
    if menu_command == 1:
        new_contact()
    elif menu_command == 2: 
        find_contact()
    elif menu_command == 3:
        save_contacts()
    elif menu_command == 4:
        break # Vai para saida.
    continue # Se naõ for 1,2 ou 3 continua no menu (em loop)

# saida
print('\nObrigado por usar o Tiny contacts.\n')
        
        
    
    