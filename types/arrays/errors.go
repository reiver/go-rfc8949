package arrays

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrIndefiniteLength = erorr.Error("rfc8949: indefinite-length CBOR array not supported")
	ErrNotArray         = erorr.Error("rfc8949: not a CBOR array (major type 4)")
)
