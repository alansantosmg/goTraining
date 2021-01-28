package main

import (
	"fmt"
)

func main() {

	const LIGHTSPEED = 2992792 // km/s
	const SECONDSPERDAY = 86400

	var distance int64 = 41.3e12
	fmt.Println("Distância de Alpha Centauri:", distance)

	days := distance / LIGHTSPEED / SECONDSPERDAY

	fmt.Println("Tempo de viagem (dias) para Alpha Centauri:", days)

	years := days / 365
	fmt.Println("tempo de viagem (anos) para Alpha Centauri:", years)

}
