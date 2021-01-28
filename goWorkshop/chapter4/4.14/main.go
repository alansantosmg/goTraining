package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	return map[string]string{
		"305": "sue",
		"204": "bob",
		"631": "jake",
		"073": "tracy",
	}
}

func getUser(id string) (string, bool) {
	users := getUsers()
	user, exists := users[id]
	return user, exists
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userID := os.Args[1]
	name, exists := getUser(userID)

	if !exists {
		fmt.Printf("Passed user id (%v) not found.\n Users:\n", userID)
		for key, value := range getUsers() {
			fmt.Println("ID:", key, "Name:", value)
		}
		os.Exit(1)
	}
	fmt.Println("Name:", name)
}
