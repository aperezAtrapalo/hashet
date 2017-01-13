// Package hashet exposes a Hash interface suited for incremental hashing of sets.
package hashet

import (
	"errors"
	"fmt"
)

const (
	// HashLengthLesserThanValue is used for error handling
	HashLengthLesserThanValue = -1
	// HashLengthGreaterThanValue is used for error handling
	HashLengthGreaterThanValue = 1
)

// Hash represents a hash object.
type Hash interface {
	// Rehash recalculates the hash value with the new set provided.
	// Calling it twice with the same values will return the hash to the previous state.
	// It will return an error if there is any length mismatch.
	Rehash(set ...[]byte) error

	fmt.Stringer
}

type xor struct {
	value []byte
}

// New returns an empty hash object of length l.
func New(l int) Hash {
	return &xor{value: make([]byte, l)}
}

// NewFromSet returns a hash object of length l with the value calculated from the given set.
// It will return an error if there is any length mismatch.
func NewFromSet(l int, set ...[]byte) (Hash, error) {
	h := New(l)
	if err := h.Rehash(set...); err != nil {
		return nil, err
	}
	return h, nil
}

func (h *xor) Rehash(set ...[]byte) error {
	hlen := len(h.value)
	for _, v := range set {
		vlen := len(v)
		if hlen != vlen {
			return newErrMismatch(hlen, vlen, "length mismatch")
		}

		newValue := make([]byte, hlen)
		for i := 0; i < hlen; i++ {
			newValue[i] = h.value[i] ^ v[i]
		}

		h.value = newValue
	}
	return nil
}

func (h *xor) String() string {
	return fmt.Sprintf("%x", string(h.value))
}

// Mismatch offers an interface for checking mismatching errors.
type Mismatch interface {
	// Mismatch returns a negative number if the expected value is lower than the received, and a positive number
	// otherwise.
	Mismatch() int
}

type errMismatch struct {
	expected int
	got      int
	error
}

func newErrMismatch(exp, got int, msg string) error {
	return &errMismatch{expected: exp, got: got, error: errors.New(msg)}
}

func (err *errMismatch) Mismatch() int {
	if err.expected < err.got {
		return HashLengthLesserThanValue
	}
	return HashLengthGreaterThanValue
}
