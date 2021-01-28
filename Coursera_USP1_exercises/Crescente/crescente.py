def main():
    firstNumber = int(input("Entre com primeiro número: "))
    secondNumber = int(input("Entre com segundo número: "))
    thirdNumber = int(input("Entre com terceiro número: "))
    if ((firstNumber < secondNumber and secondNumber < thirdNumber) or 
        (firstNumber == secondNumber and secondNumber == thirdNumber)):
        print("crescente")
    else:
        print("não está em ordem crescente")
main()
