package main

import (
	"math/big"
	"fmt"
)

func main() {

	// Ao usar pacote big todas as variávels envolvidas
	// na operação também devem ser do tipo big.

	// cria um variavel do tipo big.Int
	lightSpeed := big.NewInt(299702)

	// cria variavel do tipo big.Int
	secondsPerDay := big.NewInt(86400)

	// cria variável do tipo bigInt sem inicializar
	distance := new(big.Int)

	// inicializa variavel usando string
	// pois não é possível representá-lo como um float64
	// cuidado pois string tem aspas.
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is ", distance, "km away")

	// cria variavel do tipo big int
	seconds := new(big.Int)

	// divide distancia por tempo
	// metodo do pacote big
	seconds.Div(distance, lightSpeed)

	// cria variavel big int
	days := new(big.Int)

	// divide total sm segundos por segundos por dia
	// para obter quantidade de dias de viagem
	days.Div(seconds, secondsPerDay)
	fmt.Println("That is ", days, "days to travel at light speed.")

}
