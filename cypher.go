// Package cypherDecipher contains functions for adding and removing salt from the text.
package main

import (
	"unicode/utf8"

	"github.com/tjarratt/babble"
)

// CypherPassword adds salt to the password. It takes in a normal string and returns a salted string and the length of the salt.
func CypherPassword(password string) (saltedPassword string, pCount int, spCount int) {
	word := randomWordGenerator()
	pCount = utf8.RuneCountInString(password)
	saltedPassword = password[:utf8.RuneCountInString(password)/2] + word + password[utf8.RuneCountInString(password)/2:]
	spCount = utf8.RuneCountInString(saltedPassword)
	return
}

// DecipherPassword removes the salt from the password. It takes in a salted string, the length of the original password and length of the salt and returns the original password.
func DecipherPassword(saltedPassword string, pCount, spCount int) string {
	return saltedPassword[:pCount/2] + saltedPassword[pCount/2+spCount-pCount:]
}

// randomWordGenerator returns a random word from the default unix dictionary(Works for Windows users too.)
func randomWordGenerator() string {
	b := babble.NewBabbler()
	b.Count = 1
	word := b.Babble()
	return word
}
