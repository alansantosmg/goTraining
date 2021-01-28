def verify_osfile(my_list):
    import os.path
    file_path = read_text('Salvar dados no arquivo (nomedoarquivo.txt): ')
    if os.path.isfile(file_path):
        escolha = read_text('O arquivo já existe, deseja continuar? (s/n)')
        if escolha.upper() == 'S':  
            save_data(file_path,my_list)
        else:
            print('\nOk. Arquivo não foi sobreescrito.')
            input('Pressione qualquer tecla para continuar.')
    else: 
        save_data(file_path,my_list)






def save_data(file_path,my_list):
    try:

        with open(file_path, 'w') as my_output_file:                
            for entry in my_list:
                my_output_file.write(str(entry)+'\n')        
            my_output_file.close()
            print('\nArquivo salvo.', file_path)               
    except:
        print('\nAlgo errado ao salvar o arquivo\n')
        