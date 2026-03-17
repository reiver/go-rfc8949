package jsonnumbers

import (
	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a JSON numeric value as a CBOR tagged text string (tag 284, major type 6).
//
// The result is tag 284 wrapping a CBOR text string containing the textual
// representation of a JSON number. This preserves exact numeric precision
// beyond I-JSON (RFC 7493) limits.
//
// The value should conform to RFC 8259's JSON number grammar. No validation
// is performed — the caller is responsible for providing a valid JSON number.
//
// See: https://gist.github.com/theory/ef667af1c725240e6e30d525786d58e6
func Marshal(value string) ([]byte, error) {
	return tags.Marshal(tagnumber.JSONNumber, value)
}
