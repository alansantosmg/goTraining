import time



while True: 
    current_time = time.localtime()
    hora = str(current_time.tm_hour)
    minuto = str(current_time.tm_min)
    segundo = str(current_time.tm_sec)

    print('\033[K',hora, ':',minuto,':',segundo ,end='\r')
    time.sleep(1)


   

        
