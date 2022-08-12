package main

import (
	"github.com/tjarratt/babble"
)

func CypherPassword(password string) (saltedPassword string, spCount int) {
	word := randomWordGenerator()
	return password[:len(password)/2] + word + password[len(password)/2:], len(word)
}

func DecypherPassword(saltedPassword string, pCount, spCount int) string {
	return saltedPassword[:spCount/2] + saltedPassword[pCount/2+spCount:]
}

func randomWordGenerator() string {
	b := babble.NewBabbler()
	b.Count = 1
	word := b.Babble()
	return word
}
