package mimemessages

import (
	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a MIME message as a CBOR tagged text string (tag 36, major type 6).
//
// The result is tag 36 wrapping a CBOR text string containing a complete
// MIME message (all headers and body) per RFC 2045. No MIME validation
// is performed — the caller is responsible for providing a valid MIME message.
//
// See: RFC 8949 Section 3.4.5.3.
func Marshal(value string) ([]byte, error) {
	return tags.Marshal(tagnumber.MIMEMessage, value)
}
