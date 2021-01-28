from valida_entradas import *
import pickle

class Contact:              # Define  a classe Contact
    
    # define parametros de entrada de dados da classe
    def __init__(self,name, address, phone):

        # define atributos da classe e entra com dados dos parametros
        self.name = name 
        self.address = address
        self.phone = phone
    
    # obs: self quer dizer "a propria classe ou seja:
    # será construída uma classe com: 
    # seu próprio nome self = Contact
    # parametros para entrada de dados:  name, address, phone
    # atributos da classe: Contact.name, Contact.address, Contact.phone
    

    pass


def new_contact():
    ''' Cadastra novo contato
        Recebe nome, telefone, email
    '''
    print('\nCriar um novo contato')    
    
    # obter dados do usuário
    name = read_text('Entre com nome do contato: ')
    address = read_text('Entre com endereço: ')
    phone = read_text('Entre com telefone: ')

    # cria uma instancia da classe Contact
    # preenche os atributos da classe com as entradas do usuário
    new_contact = Contact(name = name, address = address, phone = phone)  

    
    # armazena objeto new_contact em lista
    contacts.append(new_contact)

    # informa entrada concluída
    print('\nContato armazenado.')
    input('Pressione qualquer tecla para continuar. ')
  



def find_contact(search_name): 
    search_name = search_name.strip()
    search_name = search_name.lower()
    
    for contact in contacts:
        name = contact.name
        name = name.strip()
        name = name.lower()     

        if name.startswith(search_name):
           return  contact
    return None




def edit_contact():
    
    print('\nEditar contato')    
    search_name = read_text('Entre com nome do contato: ')  
    
    contact = find_contact(search_name)
    if contact != None:
        name = read_text('Entre com novo nome ou . para cancelar alteração: ')
        if name != '.': 
            contact.name = name
        address = read_text('Entre com novo endereço ou . para cancelar alteração: ')
        if address != '.':
            contact.address = address
        phone = read_text('Entre com novo endereço ou . para cancelar alteração: ')
        if phone != ',': 
            contact.phone = phone
    else:
        input('Contato não encontrado.')
    input('\nEdição Concluída. Pressione qualquer tecla.')
   
       


def display_contact():
    search_name = read_text('Entre com nome do contato: ')  
    contact = find_contact(search_name) 
    if contact != None:          
        print('\nName: ',contact.name)
        print('Address: ',contact.address)
        print('Phone: ',contact.phone,'\n')
        input('Pressione qualquer tecla para continuar. ')
    else: 
        input('\nContato não encontrado. Pressione qualquer tecla')
  



def save_contacts(file_path):   
    with open(file_path, 'wb') as out_file:
        pickle.dump(contacts, out_file)
    print('\nContatos Salvos. Pressione qualquer tecla.')




def load_contacts(file_path): 
    global contacts     
    with open(file_path, 'rb') as input_file:
        contacts = pickle.load(input_file)

   


contacts =[]
myfile = 'contacts.pickle'
try:
    load_contacts(myfile)
except:
    print('\nArquivo não encontrado.')
    contacts = []



# menu principal
menu ='''
Tiny Contacts

1. Novo contato
2. Encontrar contato
3. Editar contato
4. Salvar contatos

5. Sair do programa

Escolha sua opção: '''


while True: 
    menu_command = read_range_int(menu, min_value = 0, max_value = 6 )
                                
    if menu_command == 1:
        new_contact()
    elif menu_command == 2: 
        display_contact()
    elif menu_command == 3:
        edit_contact()
    elif menu_command == 4: 
        save_contacts('contacts.pickle')
    elif menu_command == 6:     
        load_contacts(myfile)

    elif menu_command == 5:
        break # Vai para saida.
    continue # Se naõ for 1,2 ou 3 continua no menu (em loop)

# saida
print('\nObrigado por usar o Tiny contacts.\n')
            
    
    
    
