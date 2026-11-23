package maps

import (
	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/majortype"
)

// UnmarshalHead decodes the head of a definite-length CBOR map (major type 5) from the start of 'p'.
// It returns the number of key/value pairs in the map, and the rest of 'p' after the head — which starts with the first key.
//
// The caller decodes the keys and values themselves (key, value, key, value, …).
//
// Indefinite-length maps are not (yet) supported; UnmarshalHead returns [ErrIndefiniteLength] for them.
//
// Example usage:
//
//	count, rest, err := maps.UnmarshalHead(p)
func UnmarshalHead(p []byte) (count uint64, rest []byte, err error) {
	h, err := head.Decode(p)
	if nil != err {
		return 0, p, err
	}
	if majortype.Map != h.MajorType {
		return 0, p, ErrNotMap
	}
	if h.IsIndefinite() {
		return 0, p, ErrIndefiniteLength
	}

	return h.Argument, p[h.Length:], nil
}
