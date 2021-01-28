
import time

def digitar_devagar(texto, delay=0.1):
    for letra in (texto):
        print(letra, end=' ',flush=True)
        time.sleep(delay)

meu_texto = input('meu texto: ')
digitar_devagar(meu_texto)  