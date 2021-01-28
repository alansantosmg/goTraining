package main

import "fmt"

/* cria tipo de estrutura relatorio
 	Contem estrutura temperatura
	Contem estrutura location
*/

// embed temperature type into report struct
type report struct {
	sol
	temperature
	location
}

// cria tipo de estrutura temperatura
type temperature struct {
	high, low celsius
}

// cria tipo de estrutura location
type location struct {
	lat, long float64
}

// cria tipos
type celsius float64
type sol int

// cria metodo average
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (r report) average() celsius {
	return r.temperature.average()
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func main() {

	// cria location bradbury do tipo location
	bradbury := location{-4.5895, 137.4417}

	// cria temperatura t do tipo temperatura
	t := temperature{-1.0, -78}

	// cria relatorio report do tipo report
	report := report{sol: 15, temperature: t, location: bradbury}

	// imprime dados
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v C\n", report.temperature.high)

	// imprime a media da temperatura a partir do metodo report
	// Best solution
	fmt.Println(report.average())

	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))
}
