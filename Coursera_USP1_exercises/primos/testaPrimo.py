def main(): 
    
    from primos import verificaPrimo

    firstNumber = int(input("Entre com primeiro número: "))
    
    while (firstNumber > 1):
        
        if (verificaPrimo(firstNumber) != 0): 
            print(firstNumber)
            firstNumber = firstNumber -1
        else: 
            firstNumber = firstNumber -1 
    
main()