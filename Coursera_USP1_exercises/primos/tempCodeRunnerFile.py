def main(): 
    
    from primos import verificaPrimo

    firstNumber = int(input("Entre com primeiro número: "))
    contador = firstNumber
    while (contador > 1):
        
        if (verificaPrimo(firstnumber) != 0): 
            print(firstNumber)
            firstNumber = firstnumber -1
        else: 
            firstNumber = firstnumber -1 

main()