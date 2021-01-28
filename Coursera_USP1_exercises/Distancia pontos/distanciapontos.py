def main():
    import math
    x1 = int(input("Digite coordenada x do ponto 1: "))
    y1 = int(input("Digite coordenada y do ponto 1: "))
    x2 = int(input("Digite coordenada x do ponto 2: "))
    y2 = int(input("Digite coordenada y do ponto 2: "))
    hipotenusa = math.sqrt((x1 - x2) ** 2 + (y1 - y2) ** 2)
    if hipotenusa >= 10:
        print("longe")
    else:
        print("perto")
main()