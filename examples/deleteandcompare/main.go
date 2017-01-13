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

	// to delete an element of the set also from the hash value we apply it again
	h1.Rehash(b1)

	// then we generate a new hash for a set with the same values except the removed one
	h2, _ := hashet.NewFromSet(l, b3, b2)

	// they are equal
	fmt.Printf("equal:\n%s\n%s\n", h1, h2)
}
