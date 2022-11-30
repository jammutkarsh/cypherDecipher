package cypherDecipher

import (
	"testing"
)

var password = [5]string{"password", "1234abcd", "@#$!1234abcd", "!@#$%", "abc!21!ADA"}

func TestCypherPassword(t *testing.T) {
	for i := 0; i < len(password); i++ {
		saltPassword, pCount, spCount := CypherPassword(password[i])
		if DecipherPassword(saltPassword, pCount, spCount) != password[i] {
			t.Errorf("TestCypherPassword failed for %s", password[i])
		}
	}
}
