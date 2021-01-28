package main

import "fmt"

func main() {

	// cria um map do tipo string/int
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]

	fmt.Printf("On average the Earth is %v c.\n", temp)

	// altera item do map
	temperature["Earth"] = 16

	// insere novo item no map
	temperature["Venus"] = 464

	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon) // será zero pois não está no map

	// Comma syntax

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %v º C.\n", moon)
	} else {

		fmt.Println("Where is the moon?")
	}

	// maps cannot be copied like vars and arrays

	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars:": "Sector ZZ9",
	}

	planetsMarkII := planets

	planetsMarkII["Earth"] = "SEctor 10"
	// a alteração acima deve refletir em planets e planetsMarkII, já que eles são o mesmo map!

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

}
