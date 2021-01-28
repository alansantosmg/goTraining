// Unix Echo app simulator - vrs 3.
// os package has the var os.Args
// os.Args is a slice
// The first position is the command
// The next positions are the arguments
// Go print the arguments we need to start at position one
package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string

	// In this form, there is not the for counting var
	// Instead we use a range of itens
	// the range has a index andvalue
	// The index is used to run the n times according to any itens
	// the range's value can be manipulated via arg var
	// The arg's value changes every time the loop runs
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "

	}

	fmt.Println(s)

}
