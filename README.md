# cypherDecipher

Package cypherDecipher contains functions for adding and removing salt from the text.

The main objective of this package is to learn

- How to build conflict free packages
- How to publish packages
- How to use `git tag` and use them to have a versioned release of your package.

## Usage

Download dependecies

```bash
go get -u github.com/JammUtkarsh/cypherDecipher
```

Importing in functions

```go
package main 

import "github.com/JammUtkarsh/cypherDecipher"

func main() {
password := "somerandompassword"
pCount := len(password)
 saltedPassword, spCount := cypherDecipher.CypherPassword(password)
 println(saltedPassword)
println(cypherDecipher.DecypherPassword(saltedPassword, pCount, spCount))
}
```
