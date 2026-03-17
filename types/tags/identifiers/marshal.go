package identifiers

import (
	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a value as a CBOR tagged identifier (tag 39, major type 6).
//
// The result is tag 39 wrapping the CBOR encoding of value. Tag 39 signals
// that the wrapped value has identifier semantics — it is an ID, not a
// quantity intended for arithmetic.
//
// The value can be any type supported by the internal marshaling dispatcher
// (integers, strings, byte strings, etc.).
//
// See: https://github.com/lucas-clemente/cbor-specs/blob/master/id.md
func Marshal(value any) ([]byte, error) {
	return tags.Marshal(tagnumber.Identifier, value)
}
