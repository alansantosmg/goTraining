package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Name     string `json:"name:"`
	Pedigree string `json:"pedigree"`
	Age      uint   `json:"age"`
}

func main() {

	cachorroEmJSON := `{"name":"Rex","pedigree":"Dalmata","age":3}`
	var c cachorro
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

}
