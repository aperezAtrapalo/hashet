# hashet

[![Build Status](https://travis-ci.org/marc-gr/hashet.svg?branch=master)](https://travis-ci.org/marc-gr/hashet)[![Go Report Card](https://goreportcard.com/badge/github.com/marc-gr/hashet?style=flat-square)](https://goreportcard.com/report/marc-gr/hashet)

## getting the package

```sh
go get -u -t -i github.com/marc-gr/hashet
```

## using it

```go
import "github.com/marc-gr/hashet"

...
// creates a hash of length 20
h := hashet.New(20)
...

// we add some new value and recalculate the hash value
h.Rehash(someBytes)
...
```

## use case

This package is intended to provide an easy way to assert the integrity of an unordered set by generating a hash for a set of data, you can add data gradually to the set and the hash will update accordingly.
Two sets with the same data, no matter the order of it, will generate the same hash.
If you want to remove the data from the set, just rehash the value again with the removed register and it will recalculate it.

## not use when

If you are looking for a cryptographic safe hash, this is not for you. This package only provide a fast and cheap integrity check.