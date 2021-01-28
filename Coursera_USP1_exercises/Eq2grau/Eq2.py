def main():
    # importa funcoes de calculo da biblioteca de funcoes
    from calculos import calculaDelta, calculaQuadradoDelta, raizPositiva, raizNegativa

    # entra valores da equacao
    a1 = float(input("Entre valor de a: "))
    b1 = float(input("Entre valor de b: "))
    c1 = float(input("Entre valor de c: "))
    
    # econtrando valor de delta com funcao calcula delta  
    delta = calculaDelta(a1,b1,c1)

    # encontrando raiz de delta com função calculaQuadradoDelta
    quadradoDelta = calculaQuadradoDelta(delta)
  
    # verifica se delta é > que zero (duas raizes reais)
    if delta > 0:
        x1 = raizPositiva(a1,b1,quadradoDelta)
        x2 = raizNegativa(a1,b1,quadradoDelta)

        # Ordena raizes reais da maior para menor
        if x1 <= x2:
            print("as raízes da equação são",x1,"e",x2)
        else:
            print("as raízes da equação são",x2,"e",x1)            
    else: 
    # verificar se valor de delta é zero (uma raiz real positiva)
   
        if delta == 0:
            x1 = raizPositiva(a1,b1,quadradoDelta)
            print("a raiz desta equação é",x1)
        else:
    # informa quando delta for menor que zero (não tem raizes reais)
            print("esta equação não possui raízes reais")

main()

