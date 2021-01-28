package main

// test
import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dwarfs = append(dwarfs, "Orcus")

	fmt.Println(dwarfs)

	dwarfs = append(dwarfs, "Salacia", "Quaorar", "Sedna")

	fmt.Println(dwarfs)
	fmt.Println(len(dwarfs))
}
