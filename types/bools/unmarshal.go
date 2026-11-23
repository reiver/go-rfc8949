package bools

import (
	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/initialbyte"
)

// Unmarshal decodes a CBOR boolean (major type 7, simple value 20 or 21) from the start of 'p'.
// It returns the value, and the rest of 'p' after the data item.
//
// Example usage:
//
//	value, rest, err := bools.Unmarshal(p)
func Unmarshal(p []byte) (value bool, rest []byte, err error) {
	if len(p) <= 0 {
		return false, p, head.ErrEmpty
	}

	switch p[0] {
	case initialbyte.False:
		return false, p[1:], nil
	case initialbyte.True:
		return true, p[1:], nil
	default:
		return false, p, ErrNotBool
	}
}
