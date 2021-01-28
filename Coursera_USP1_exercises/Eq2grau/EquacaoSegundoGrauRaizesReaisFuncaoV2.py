def main():
    import math
    # entra valores da equacao
    a1 = float(input("Entre valor de a: "))
    b1 = float(input("Entre valor de b: "))
    c1 = float(input("Entre valor de c: "))
    
    #calcula delta function
    def calculaDelta(a,b,c):
        return (b ** 2) - 4 * a * c
          
    #calcula quadrado delta function
    def calculaQuadradoDelta(d):
        return math.sqrt(d)
    
    #calcula raiz positiva
    def raizPositiva (a,b):
        return (-b + quadradoDelta) / (2 *a)

    #calcula raiz negativa
    def raizNegativa (a,b):
        return (-b - quadradoDelta) / (2 *a)

    delta = calculaDelta(a1,b1,c1)
    quadradoDelta = calculaQuadradoDelta(delta)
  
    # verifica se delta é > que zero (duas raizes reais)
    if delta > 0:
        x1 = raizPositiva(a1,b1)
        x2 = raizNegativa(a1,b1)
        if x1 <= x2:
            print("as raízes da equação são",x1,"e",x2)
        else:
            print("as raízes da equação são",x2,"e",x1)            
    else: 
    # verificar se delta é - 0 (uma raiz real positiva)
   
        if delta == 0:
            x1 = raizPositiva(a1,b1)
            print("a raiz desta equação é",x1)
        else:
    # informa quando delta for menor que zero (não tem raizes reais)
            print("esta equação não possui raízes reais")

main()

