package base64urls

import (
	"encoding/base64"

	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes raw bytes as a CBOR tagged base64url text string (tag 33, major type 6).
//
// The result is tag 33 wrapping a CBOR text string containing the base64url
// encoding of value, using the URL-safe alphabet (RFC 4648 Section 5) with
// no padding characters, per RFC 8949 Section 3.4.5.3.
func Marshal(value []byte) ([]byte, error) {
	encoded := base64.RawURLEncoding.EncodeToString(value)
	return tags.Marshal(tagnumber.Base64URL, encoded)
}
