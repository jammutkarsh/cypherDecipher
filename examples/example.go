package main

import (
	"fmt"

	"github.com/JammUtkarsh/cypherDecipher"
)

func main() {
	var username, password string
	fmt.Print("Enter username: ")
	fmt.Scan(&username)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)
	salt, pCount, spCount := cypherDecipher.CypherPassword(password)
	fmt.Println(username)
	fmt.Println(salt)
	fmt.Println(pCount)
	fmt.Println(spCount)
}