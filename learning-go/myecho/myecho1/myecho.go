// Unix Echo app simulator
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

	for 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "

	}

	fmt.Println(s)

}
