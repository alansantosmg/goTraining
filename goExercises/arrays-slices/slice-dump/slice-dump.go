package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: lenght %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])

	// dwarfs1 - capacity = 5
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	// dwarfs2 - copy of array - capacity = 10
	// quando atinge limite dobra capacidade
	// copiando os itens para novo array implicito
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")

	// imprime slice do array com capacidade 10
	fmt.Println(dwarfs2)

	// imprime slice do array com capacidade 5
	fmt.Println(dwarfs1)

	// altera slice do array com capacidade 10
	dwarfs2[1] = "Plutão"

	// imprime slice do array com capacidade 10 alterado
	fmt.Println(dwarfs2)

	// prova que dwarfs2 é slice de outro array
	// visto que elementos no slice do array original
	// não foi alterado
	fmt.Println(dwarfs1)

	// prova que dwarfs3 é slice do mesmo array do slice dwarfs2
	// porque elemento alterado de dwarfs2 está presente
	fmt.Println(dwarfs3)

	// Quando um array (que está por tras do slice
	// excede a capacidade, Go copia os itens do
	// array original para um novo array com dobro da
	// capacidade do array original

}
