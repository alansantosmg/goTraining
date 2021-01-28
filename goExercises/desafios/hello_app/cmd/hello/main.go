// use main for app entry
package main

// You can import from local or internet
import (
	"fmt"

	"../../internal/greet"
)

// App entry point
func main() {

	// It's ok get function returns inside vars
	// when initializing var use abrev. form  :=
	greeting := greet.Hello("alan", "francês")

	// imported functions always starts with upper case
	fmt.Println(greeting)
}

//author: Alan Santos
