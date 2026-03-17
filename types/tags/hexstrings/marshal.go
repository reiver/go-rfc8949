package hexstrings

import (
	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a byte slice as a CBOR tagged hexadecimal string (tag 263, major type 6).
//
// The result is tag 263 wrapping a CBOR byte string. Tag 263 signals that
// the byte string should be displayed or converted as hexadecimal when
// rendered in a text-based format.
//
// The bytes are stored raw — no hex encoding is performed. The tag is a
// display hint, not a data transformation.
//
// See: https://github.com/toravir/CBOR-Tag-Specs/blob/master/hexString.md
func Marshal(value []byte) ([]byte, error) {
	return tags.Marshal(tagnumber.HexString, value)
}
