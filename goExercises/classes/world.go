// Package main provides ...
package main

import (
	"fmt"
	"math"
)

// cria um tipo de struct chamado mundo que tem um raio
type world struct {
	radius float64
}

// cria um tipo de localização lat, long
type location struct {
	lat  float64
	long float64
}

// funcao que transforma graus em radianos
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// metodo que calcula distancia entre 2 pontos num dado mundo
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func main() {

	// entra com dados do mundo
	var mars = world{radius: 3389.5}

	// entra com dados dos pontos de localização
	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}

	// calcula distancia os pontos do mundo informado.
	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2f km\n", dist)
}
