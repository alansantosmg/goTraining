package main

import "fmt"

// criar type struct coordinate
// degree
// minutes
// seconds
// NSEW (North, South, East, West)
type coordinate struct {
	d float64
	m float64
	s float64
	h rune
}

// method coordinate to decimal
// para coodenada do tipo c, executar metodo decimal retornar float64
func (c coordinate) decimal() float64 {

	// determina sinal positivo
	sign := 1.0

	// caso coordenada seja sul ou Oeste inverter sinal
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	// retorna conversão de graus para decimal
	return sign * (c.d + (c.m / 60) + (c.s / 3600))

}

// tipo location struct
type location struct {
	lat, long float64
}

// funcao construtora
// recebe latitude e longitude do tipo coordinate
// devolve localização do tipo location
// Se começar com maiuscula funçao é exportada
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	// entrada de dados
	lat := coordinate{35, 30, 25, 's'}
	long := coordinate{4, 35, 22.2, 'E'}

	//saida
	// O metodo atrelado ao tipo funciona como se fosse uma
	// propriedade do tipo.  Logo, basta chama-lo com a notação de ponto.
	fmt.Println(lat.decimal(), long.decimal())
	fmt.Println()

	// constroi localização sem usar construtor
	curiosity := location{lat.decimal(), long.decimal()}

	// constroi localização usando funcao construtora
	bradbury := newLocation(lat, long)

	// saida de dados
	fmt.Printf("Curiosity: \n%+v", curiosity)
	fmt.Printf("Bradbury: \n%+v", bradbury)

}
