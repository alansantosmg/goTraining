package main

import "fmt"

// tipos de objetos do  tipo float64
type fahrenheit float64
type celsius float64
type kelvin float64

// método  celsius para objetos do tipo fahrenheit
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

// metodo celsius para objetos do tipo kelvin
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// metodo fahrenheit para objetos do tipo celsius
func (c celsius) fahrenheit() farenheit {
	return fahrenheit(c*9/5) + 32
}


func main() {

	// cria objeto (Variável) do tipo kelvin
	// e atribui valor 300
	var k kelvin = 300.0

	// cria variáveis e deixa Go inferir o
	// tipo de objeto das durante a atribuição
	c := k.celsius()
	f := fahrenheit(c)

	// imprime resultado com 2 casas decimais
	fmt.Printf("%.2f kelvin = %.2f celsius = %.2f fahrenheit", k, c, f)

}
