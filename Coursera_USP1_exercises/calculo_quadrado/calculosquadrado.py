def main():
	# Calculo de perímetro e área de quadrados
	# Entrar lado do quadrado
	lado = input("Digite o valor correspondente ao lado de um quadrado:")
	# Calcular perímetro
	perimetro = int(lado) * 4
	# Calcular area
	area = int(lado) ** 2
	# Mostrar perimetro e area
	print("perímetro:",perimetro,"-", "área:",area)
main()
