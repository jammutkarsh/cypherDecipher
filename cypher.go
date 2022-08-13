/*The main objective of this package is to learn
1. How to build conflict free packages
2. How to publish packages
3. How to use `git tag` and use them to have a versioned release of your package.
*/

// Package cypherDecipher contains functions for adding and removing salt from the text.
package cypherDecipher

import (
	"github.com/tjarratt/babble"
)

// CypherPassword adds salt to the password. It takes in a normal string and returns a salted string and the length of the salt.
func CypherPassword(password string) (saltedPassword string, spCount int) {
	word := randomWordGenerator()
	return password[:len(password)/2] + word + password[len(password)/2:], len(word)
}

// DecipherPassword removes the salt from the password. It takes in a salted string, the length of the original password and length of the salt and returns the original password.
func DecipherPassword(saltedPassword string, pCount, spCount int) string {
	return saltedPassword[:spCount/2] + saltedPassword[pCount/2+spCount:]
}

// randomWordGenerator returns a random word from the default unix dictionary(Works for Windows users too.)
func randomWordGenerator() string {
	b := babble.NewBabbler()
	b.Count = 1
	word := b.Babble()
	return word
}
