def main():
	#mostra digito das dezenas de um número inteiro
	#entrar número
	numeroEntrada = input("Digite um número inteiro: ")
	# desprezar casas além da dezena
	desprezaAlemDezenas = int(numeroEntrada) % 100
	# desprezar casa das unidades
	digitoDezena = int(desprezaAlemDezenas) // 10
	#Mostrar digito da dezena
	print("O dígito das dezenas é",digitoDezena)
main()
