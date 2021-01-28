def main():
	# conversao celsius Farhenheit
    print("Conversor de temperatura 1.0")
	# entrar temperatura Farhenheit
    temperaturaFarheineit = input("Digite uma temperatura em Farheineit: ")
    # converter temperatura
    temperaturaCelsius = (float(temperaturaFarheineit) -32) *5 // 9
    # mostrar conversao
    print(temperaturaFarheineit,"F", "= ", temperaturaCelsius,"C")
main()