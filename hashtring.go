// Package hashtring exposes a Hash interface suited for incremental hashing of
// string sets.
package hashtring

import (
	"errors"
	"fmt"
)

// Hash represents a hash object.
type Hash interface {
	// Rehash recalculates the hash value with a new one.
	// Calling it twice with the same value will return
	// the hash to the previous state.
	// It will return an error if the length of the hash
	// and the new value mismatch.
	Rehash(v string) error
}

type xor struct {
	value []byte
}

// New returns an empty hash object of length l.
func New(l int) Hash {
	return &xor{value: make([]byte, l)}
}

// NewFromSet returns a hash object of length l
// with the value calculated from the given set.
// Returns an error if the set has any value
// of length different than l.
func NewFromSet(l int, set []string) (Hash, error) {
	h := New(l)
	for _, v := range set {
		if err := h.Rehash(v); err != nil {
			return nil, err
		}
	}
	return h, nil
}

func (h *xor) Rehash(v string) error {
	hlen := len(h.value)
	vlen := len(v)
	if hlen != vlen {
		return newErrMismatch(hlen, vlen, "length mismatch")
	}

	newValue := make([]byte, hlen)
	for i := 0; i < hlen; i++ {
		newValue[i] = h.value[i] ^ v[i]
	}

	h.value = newValue
	return nil
}

func (h *xor) String() string {
	return fmt.Sprintf("%x", string(h.value))
}

// Mismatch offers an interface for
// checking mismatching errors.
type Mismatch interface {
	// Mismatch returns a negative number if the expected
	// value is lower than the received, and a positive number
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
		return -1
	}
	return 1
}
