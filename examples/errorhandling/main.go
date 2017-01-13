package main

import (
	"fmt"
	"log"

	"github.com/marc-gr/hashet"
)

func main() {
	b1 := []byte{0, 2, 4, 8}
	b2 := []byte{2, 4, 8}
	l := len(b1)

	// we generate a new hash for the set
	_, err := hashet.NewFromSet(l, b1, b2)
	if err != nil {
		// the package returns errors as an implementation of an interface Mismatch
		// which gives you a safe way to handle them and more information if desired
		// about the nature of the errors
		errMismatch, ok := err.(hashet.Mismatch)
		if !ok {
			log.Fatalf("the error is from an unexpected type %T", err)
		}
		switch errMismatch.Mismatch() {
		case hashet.HashLengthLesserThanValue:
			fmt.Println("the length of the hash is less than one of the values length")
		case hashet.HashLengthGreaterThanValue:
			fmt.Println("the length of the hash is greater than one of the values length")
		}
	}
}
