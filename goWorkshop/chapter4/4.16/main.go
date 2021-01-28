package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
	"305": "sue",
	"204": "Bob",
	"631": "Jake",
	"073": "tracy",
}

func deleteUser(id string) {
	delete(users, id)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User id not passed")
		os.Exit(1)
	}

	userID := os.Args[1]
	deleteUser(userID)
	fmt.Println("Users: ", users)
	fmt.Printf("user deleted:", userID)
}
