package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Name     string `json:"name"`
	Pedigree string `json:"pedigree"`
	Age      uint   `json:"age"`
}

func main() {

	c := cachorro{"Rex", "Dalmata", 3}

	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

}
