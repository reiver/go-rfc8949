package uint64s

import (
	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/majortype"
)

// Unmarshal decodes a CBOR unsigned integer (major type 0) from the start of 'p'.
// It returns the value, and the rest of 'p' after the data item.
//
// Unmarshal accepts any well-formed encoding of the argument, including a non-minimal one.
// (Use [head.Decode] and [head.Head.IsMinimal] to require minimal encoding.)
//
// Example usage:
//
//	value, rest, err := uint64s.Unmarshal(p)
func Unmarshal(p []byte) (value uint64, rest []byte, err error) {
	h, err := head.Decode(p)
	if nil != err {
		return 0, p, err
	}
	if majortype.UnsignedInteger != h.MajorType {
		return 0, p, ErrNotUnsignedInteger
	}
	if h.IsIndefinite() {
		return 0, p, ErrNotUnsignedInteger
	}

	return h.Argument, p[h.Length:], nil
}
