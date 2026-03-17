package embeddedjson

import (
	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes raw JSON bytes as a CBOR tagged embedded JSON object (tag 262, major type 6).
//
// The result is tag 262 wrapping a CBOR byte string containing the JSON data.
// No JSON validation is performed — the caller is responsible for providing
// valid JSON.
//
// See: https://github.com/toravir/CBOR-Tag-Specs/blob/master/embeddedJSON.md
func Marshal(value []byte) ([]byte, error) {
	return tags.Marshal(tagnumber.EmbeddedJSON, value)
}
