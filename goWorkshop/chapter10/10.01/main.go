package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func whatstheclock() string {
	return time.Now().Format(time.ANSIC)
}

func main() {

	fmt.Println(whatstheclock())

	host, err := os.Hostname()
	if err != nil {
		log.Fatalf("Hostname não encontrado: %v", err)
	}

	fmt.Println(host)

}
