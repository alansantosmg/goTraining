package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Compilador p/ arquitetura:", runtime.GOARCH)
	fmt.Println("Versão do Go:", runtime.Version())
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de Go routines", runtime.NumGoroutine())
}
