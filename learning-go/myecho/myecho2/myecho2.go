// Unix Echo app simulator vrs. 2.0
// os package has the var os.Args
// os.Args is a slice
// The first position is the command
// The next positions are the arguments
// Go print the arguments we need to start at position one
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// using method Join to separate os.Args pararms
	// beggining at index 1 cause index zero is the command
	fmt.Println(strings.Join(os.Args[1:], ", "))

}
