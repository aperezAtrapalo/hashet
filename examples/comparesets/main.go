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

	// then we generate a new hash for a set with values in a different order
	h2, _ := hashet.NewFromSet(l, b3, b1, b2)

	// they are equal
	fmt.Printf("equal:\n%s\n%s\n", h1, h2)
}
