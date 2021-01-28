def main(): 
    
    coeficiente = int(input("entre com coeficiente: "))
    denominador = int(input("entre com denominador: "))

    

    def fatorial(n):
        #inicializa fatorial = 0
        f = 0
        # Fatorial de 0 é 1
        if n == 0:
            f = 1

        # Se n não for zero calcula
        else:
            # inicializa contador
            f = n 
        # Fatoriza enquanto número ainda for maior que 1
        while n > 1:            
            # Subtrai 1 do número para obter multiplicador
            n = n - 1
            # multiplica multiplicador pelo fator
            f = f * n
        return f
    
    def coeficienteBinomial(x,y):
        return  fatorial(x) / ((fatorial(x - y)) * fatorial(y))

    print(coeficienteBinomial(coeficiente,denominador))
main()