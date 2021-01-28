package main

import (
	"fmt"
	"time"
)

func main() {

	// obter tempo "agora" em segundos desde 1/01/1970 00:00:00
	fmt.Println("Epoch time:", time.Now().Unix())

	// obter tempo "agora"
	t := time.Now()

	fmt.Println(t, t.Format(time.RFC3339))

	//imprime frações do tempo
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	// gera delay de 1 segundo
	time.Sleep(time.Second)

	// gera delai de 5 segundos
	time.Sleep(time.Second * 5)

	// diferença entre t1 e t (um segundo antes)
	t1 := time.Now()
	fmt.Println("Time difference:", t1.Sub(t))

	// formata data  dia, mes ano
	formatT := t.Format("01 January 2008")
	fmt.Println(formatT)

	// carrega fuso horário de paris em loc
	loc, _ := time.LoadLocation("America/Sao_Paulo")

	// horário de londres convertido para fuso de paris
	londonTime := t.In(loc)
	fmt.Println("Paris:", londonTime)

}
