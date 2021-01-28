package main

import "log"

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {

	name := "Thanos"

	func(name string) {
		log.Println("Demo app")
		log.Printf("%s is here", name)

		log.Print("Run!")

	}(name)

}
