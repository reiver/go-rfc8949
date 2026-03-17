package base64s

import (
	"encoding/base64"

	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes raw bytes as a CBOR tagged base64 text string (tag 34, major type 6).
//
// The result is tag 34 wrapping a CBOR text string containing the base64
// encoding of value, using the standard alphabet (RFC 4648 Section 4) with
// padding characters, per RFC 8949 Section 3.4.5.3.
func Marshal(value []byte) ([]byte, error) {
	encoded := base64.StdEncoding.EncodeToString(value)
	return tags.Marshal(tagnumber.Base64, encoded)
}
