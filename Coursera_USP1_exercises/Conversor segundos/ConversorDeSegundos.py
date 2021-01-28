def main():
	# Conversor de segundos
	entradaSegundos = input("Por favor, entre com o número de segundos que deseja converter: ")
	#calcula total de dias
	dias =  int(entradaSegundos) // 86400
	#calcula total de horas
	horas = int(entradaSegundos) // 3600
	#calcula total de minutos
	minutos = int(entradaSegundos) // 60
	# calcula sobra de horas que não completam 1 dia
	sobraHoras = int(horas) % 24
	# calcula sobra de minutos que não completam 1 hora
	sobraMinutos = int(minutos) % 60
	#calcula sobra de segundos que não completam 1 minuto
	sobraSegundos = int(entradaSegundos) % 60
	#mostra resultado
	print(dias, "dias,", sobraHoras, "horas,", sobraMinutos,"minutos e", sobraSegundos, "segundos.")
	
main()


