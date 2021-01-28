from valida_entradas import *
import pickle

class Contact:              # Define  a classe Contact

    # variaveis 'globais da classe
    # podem ser usados em qualquer método
    # não dependem de objeto para existir (existem na classe)
    # o uso de 2 underscore protege os atributos
    # Assim eles não podem ser invocados externamente à classe
    # Só poderiam ser invocados através de um método existente 
    # na classe que tivesse a capacidade de dar um get ou set
   
    __min_session_length = 0.5
    __max_sessoin_length = 3
    __min_text_length = 4
    __open_fee = 30
    __hourly_fee  = 50

    # decorator static altera a propriedade do método 
    # para que ele se torne estático, ou seja, 
    # seja parte da classe (não necessitando de um objeto para existir)
    # é interessante usar métodos estáticos para validações, pois elas
    # podem ser invocadas mesmo se o objeto ainda não existir. 

    @staticmethod
    def valid_text(text): 
        ''' 
        Valida do texto dos contatos

        '''
        if len(text) < Contact.__min_text_length: 
            return False
        else: 
            return True 

    @staticmethod
    def validate_session_length(session_length):
        if session_length < Contact.__min_session_length: 
            return
        if session_length > Contact.__max_sessoin_length:
            return
        return session_length 
    
    # Metodos não estáticos, só existem na instancia do objeto


    # define parametros de entrada de dados da classe
    # é o construtor da classe
    # o metodo init é estatico por natureza pois ele só pode ser invocado pela 
    # classe que deverá construir

    def __init__(self,name, address, phone, hours_worked = 0, billing_amount = 0):
              
        #atributos públicos
        # podem ser alterados diretamente pelo programa
        # sem precisar de ter um metodo na classe para isso
        self.name = name 
        self.address = address
        self.phone = phone
       
        # atributo privado (2 underscores)
        # só pode ser alterado internamente na classe
        # externamente só poderia ser alterado por um metodo
        # Esse metodo é que seria invocado externamente 
        self.__hours_worked = 0
        self.__billing_amount = 0
        self.__version = 2
    
    
    # obs: self quer dizer "a propria classe ou seja:
    # será construída uma classe com: 
    # seu próprio nome self = Contact
    # parametros para entrada de dados:  name, address, phone
    # atributos da classe: Contact.name, Contact.address, Contact.phone
    # obtem (get) o dado do atributo __hours_worked
   
    def get_hours_worked(self):
        return self.__hours_worked

    # altera (SET) a propriedade __hours_worked
    def add_session(self, session_length):
        # se a funçao de validar não validar gerar excecao para programador
        # poderia ser substituido por if. Vide versão 2
        # onde usado juntamente com try geraria msg de erro mas
        # deixaria o programa funcionando
        if not Contact.validate_session_length(session_length):
            raise Exception('Quantidade de trabalho inválida')

        self.__hours_worked = self.__hours_worked + session_length
        amount_to_bill = Contact.__open_fee + (Contact.__hourly_fee * session_length)
        self.__billing_amount = self.__billing_amount + amount_to_bill

        return True
    

    # metodo que checa a versão da classe
    def check_version(self):
        if self.__version == 1: 
            self.__billing_amount = 0
            self.__version = 2

    # metodo que imprime o objeto como string
    
    def __str__(self):       
        template = '''Nome: {0}
Endereço: {1}
Telefone: {2}
Horas executadas: {3}
Valor a faturar: {4}'''
        return template.format(self.name, self.address, self.phone, self.__hours_worked, self.billing_amount)
   
    #property é um decorator para criar facilmente métodos GET e SET para 
    # um atributo da classe. Ele pode ser usado ao invés de criar métodos GET e SET
    # a facilidade reside no fato que ele usa o mesmo nome do atributo.

    @property
    def name(self): 
        return self.__name

    @name.setter
    def name(self,name): 
        if not Contact.valid_text(name): 
            raise Exception('Invalid name')
        self.__name = name
   
    @property
    def address(self): 
       return self.__address

    @address.setter
    def address(self,address):
        if not Contact.valid_text(address):
            raise Exception('Invalid Address')
        self.__address = address


    @property
    def billing_amount(self):
        return self.__billing_amount
    
    

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
        name = name.strip() # retira espaços em branco
        name = name.lower() # passa tudo para minusculo 

        if name.startswith(search_name): # busca caracteres começando no parametro
           return  contact
    return None  # retornar nada




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
        phone = read_text('Entre com novo telefone ou . para cancelar alteração: ')
        if phone != ',': 
            contact.phone = phone
    else:
        input('Contato não encontrado.')
    input('\nEdição Concluída. Pressione qualquer tecla.')
   
       


def display_contact():
    search_name = read_text('\nEntre com nome do contato: ')  
    contact = find_contact(search_name) 
    if contact != None:        
        print()    
        print(contact)
        input('\nPressione qualquer tecla para continuar. ')
        
      
    else: 
        input('\nContato não encontrado. Pressione qualquer tecla')
  



def save_contacts(file_path):   
    with open(file_path, 'wb') as out1_file:  # wb write binary
        pickle.dump(contacts, out_file)  # picar a informação e despejar no arquivo binário
    print('\nContatos Salvos. Pressione qualquer tecla.')




def load_contacts(file_path): 
    global contacts  #variavel é global definida fora do escopo da funçao
    contacts.clear() 

    # clausula with é usada para fechar arquivo
    # ao final da operação ou caso ocorra erro. 
    # poderia colocar um try também para mostrar caso erro ocorra

    with open(file_path, 'rb') as input_file:  # r de read  b de binary
        contacts = pickle.load(input_file)  # carregar arquivo binário
    
    for contact in contacts:
        contact.check_version()




def add_session():
    print('Add hours')
    search_name = read_text('Entre com nome do contato: ')  
    contact = find_contact(search_name)
    if contact != None: 
        print(contact.name)
        
        print ('Horas trabalhadas até o momento: ', contact.get_hours_worked())
        
        prompt = 'Trabalho atual: '
        session_length = read_float(prompt)
        contact.add_session(session_length)
                  
        print('Horas trabalhadas atualizadas: ', contact.get_hours_worked())
    else: 
        print(contact.name, 'não encontrado.\n')
        input('\nPressione qualquer tecla\n')



contacts = [] 
myfile = 'contacts.pickle' # pickle é um arquivo binário

# prepara e imprime msg caso haja exceção
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
4. Adicionar sessão
5. Salvar contatos

6. Sair do programa

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
        add_session()
    elif menu_command == 5:
        save_contacts('contacts.pickle')
    elif menu_command == 6:
        break # Vai para saida.
    continue # Se naõ for 1,2 ou 3 continua no menu (em loop)

# saida
print('\nObrigado por usar o Tiny contacts.\n')
            