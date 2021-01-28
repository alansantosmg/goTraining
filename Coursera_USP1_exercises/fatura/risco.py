def main():
	# Cálculo de risco
	#
	#
	# Entrada dados para cálculo
	print("Calculadora de risco")
	print("====================")
	print()
	print("ENTRADA DE DADOS")
	patrimonioTotal = input("Digite Patrimonio total: ")
	percentualRiscoTotal = input("Digite % risco: ")
	precoCompra = input("Digite preço de compra: ")
	precoStopLoss = input("Digite preço de Stop Loss: ")
     	#
	# Calcula risco total do patrimônio
	valorRiscoTotal = ((float(patrimonioTotal) * 1000) * float(percentualRiscoTotal)) / 100
	round(valorRiscoTotal, 2)
	#
	#
	# calcula valor de risco por ação
	valorRiscoAcao = float(precoCompra) - float(precoStopLoss)
	round(valorRiscoAcao, 2)
	#
	#
	# Calcula percentual de risco por ação
	percentualRiscoAcao = (valorRiscoAcao / float(precoCompra)) * 100
	round(percentualRiscoAcao, 1)
	#
	#
	# Calcula quantidade de ações permitida por risco de operação
	quantidadePorOperacao = (float(valorRiscoTotal) / float(valorRiscoAcao))
	round(quantidadePorOperacao, 2)
	#
	#
	# Mostra risco patrimonial
	print()
	print()
	print("CÁLCULO DE RISCO PATRIMONIAL ")
	print("Patrimônio: R$", patrimonioTotal, "K")
	print("Risco:", percentualRiscoTotal, "%")
	print("Valor de risco: R$ ", valorRiscoTotal)
	print()
	print()
       	#
	#
	# Mostra risco por ação
	print("CÁLCULO DE RISCO POR AÇÃO")
	print("Preço compra: R$", precoCompra)
	print("Preço stop loss: R$", precoStopLoss)
	print("Risco por ação: R$", valorRiscoAcao)
	print("Risco da operação:", percentualRiscoAcao, "%")
	print()
	#
	# Mostra quantidade permitida para compra na operação dentro do risco calculado
	print("GESTÃO DE RISCO")
	print("Quantidade permitida: ", quantidadePorOperacao)
main()


