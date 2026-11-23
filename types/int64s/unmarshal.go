package int64s

import (
	"math"

	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/majortype"
)

// Unmarshal decodes a CBOR integer — unsigned (major type 0) or negative (major type 1) — from the start of 'p',
// into an int64.
// It returns the value, and the rest of 'p' after the data item.
//
// If the integer does not fit in an int64, Unmarshal returns [ErrOverflow].
//
// Unmarshal accepts any well-formed encoding of the argument, including a non-minimal one.
// (Use [head.Decode] and [head.Head.IsMinimal] to require minimal encoding.)
//
// Example usage:
//
//	value, rest, err := int64s.Unmarshal(p)
func Unmarshal(p []byte) (value int64, rest []byte, err error) {
	h, err := head.Decode(p)
	if nil != err {
		return 0, p, err
	}
	if h.IsIndefinite() {
		return 0, p, ErrNotInteger
	}

	switch h.MajorType {
	case majortype.UnsignedInteger:
		if math.MaxInt64 < h.Argument {
			return 0, p, ErrOverflow
		}
		return int64(h.Argument), p[h.Length:], nil
	case majortype.NegativeInteger:
		// The value is -1 minus the argument.
		if math.MaxInt64 < h.Argument {
			return 0, p, ErrOverflow
		}
		return -1 - int64(h.Argument), p[h.Length:], nil
	default:
		return 0, p, ErrNotInteger
	}
}
