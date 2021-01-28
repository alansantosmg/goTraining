package main

import (
	"fmt"
	"tdd1/enderecos"
	"time"
)

func main() {

	start := time.Now()
	fmt.Println("Início", start)
	tipoEndereco := enderecos.TipoDeEndereco("Avenida paulista")
	fmt.Println(tipoEndereco)

	end := time.Now()
	fmt.Println("fim", end)
}
