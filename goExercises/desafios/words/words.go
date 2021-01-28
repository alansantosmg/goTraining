package main

import (
	"fmt"
	"strings"
)

// func formatString converte string s para minusculo
func formatString(s string) string {

	s = strings.ToLower(s)
	return s
}

// func textToSlice converte string para slice
func textToSlice(t string) []string {

	// cada palavra é convertida para um item de slice
	tSlice := strings.Fields(t)

	// retira pontuação de cada item de slice
	for i := range tSlice {
		tSlice[i] = strings.Trim(tSlice[i], ".,_:;?!-")
	}

	return tSlice
}

func main() {

	// entrada de dados para analise
	inputText := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man. eu.,!;"

	// chama funcao formatString para entrada de dados
	formatedString := formatString(inputText)

	// chama funcao textoToSlice para string formatada em minusculas
	formatedSlice := textToSlice(formatedString)

	// mostra resultado da conversao para slice
	fmt.Printf("%q\n", formatedSlice)

	// cria um map
	var textMap = make(map[string]int)

	// para cada item do slice formatado
	for _, f := range formatedSlice {

		// adiciona 1 ao indice f
		textMap[f]++

	}

	// define variaveis de estatisticas
	notRepeatedWords := 0
	totalWords := 0

	for i, t := range textMap {

		//conta palagras repetidas e não repetida
		notRepeatedWords++
		totalWords += t

		// mostra ocorrencias de cada palavra
		fmt.Printf("word \"%v\": %d\n ", i, t)
	}

	// imprime estatisticas de palavras repetidas e nao repetidas
	fmt.Printf("\nWord count:\nNot repeated words: %d \nTotal words: %d", notRepeatedWords, totalWords)
}
