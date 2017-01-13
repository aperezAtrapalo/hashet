package main

import (
	"fmt"

	"github.com/marc-gr/hashet"
)

func main() {
	b1 := []byte{0, 2, 4, 8}
	b2 := []byte{2, 4, 8, 16}
	b3 := []byte{1, 1, 2, 3}
	l := len(b1)

	// we generate a new hash for the set
	h1, _ := hashet.NewFromSet(l, b1, b2, b3)

	// now we "corrupt" the data and generate another hash
	b2[3] = 15
	h2, _ := hashet.NewFromSet(l, b1, b2, b3)

	// they are different
	fmt.Printf("different:\n%s\n%s\n", h1, h2)

	// but if the set is the same
	b2[3] = 16
	h3 := hashet.New(l)
	h3.Rehash(b1) // we can also add values to the set in an additive way
	h3.Rehash(b2, b3)

	// they are equal
	fmt.Printf("equal:\n%s\n%s\n", h1, h3)
}
