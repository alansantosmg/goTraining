 
entrada = int(input("Entre com seu número: "))
print("entrada inicial",entrada)
casas = 0
numeroDivisaocasas = entrada
while numeroDivisaocasas > 0:
    numeroDivisaocasas = numeroDivisaocasas // 10
    casas = casas + 1
print(casas)
10