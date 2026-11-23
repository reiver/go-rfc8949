package nils

import (
	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/initialbyte"
)

// Unmarshal decodes a CBOR null (major type 7, simple value 22) from the start of 'p'.
// It returns the rest of 'p' after the data item.
//
// Example usage:
//
//	rest, err := nils.Unmarshal(p)
func Unmarshal(p []byte) (rest []byte, err error) {
	if len(p) <= 0 {
		return p, head.ErrEmpty
	}

	if initialbyte.Null != p[0] {
		return p, ErrNotNull
	}

	return p[1:], nil
}
