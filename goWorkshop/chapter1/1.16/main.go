package main

import (
	"fmt"
)

const globalLimit = 100
const maxCacheSize = 10 * globalLimit
const (
	cacheKeyBook = "_book"
	cacheKeyCd   = "_cd"
)

// cria o tipo cache derivado de mapa
var cache map[string]string

// obtem itens do cache
func cacheGet(key string) string {
	return cache[key]
}

// adiciona itens ao cache
func cacheSet(key, val string) {
	if len(cache)+1 >= maxCacheSize {
		return
	}
	cache[key] = val
}

// obert livro do cache
func getBook(isbn string) string {
	return cacheGet(cacheKeyBook + isbn)
}

// colocar livro no cache
func setBook(isbn string, name string) {
	cacheSet(cacheKeyBook+isbn, name)
}

// obter cd
func getCd(sku string) string {
	return cacheGet(cacheKeyCd + sku)
}

//colocar cd no cache
func setCd(sku string, title string) {
	cacheSet(cacheKeyCd+sku, title)
}

func main() {

	// inicializa o mapa
	cache = make(map[string]string)

	setBook("1234-5678", "Get ready to go")
	setCd("1234-5678", "Get ready audio to go")

	fmt.Println("Book: ", getBook("1234-5678"))
	fmt.Println("CD:", getCd("1234-5678"))

}
