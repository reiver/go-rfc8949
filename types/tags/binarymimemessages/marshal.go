package binarymimemessages

import (
	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a binary MIME message as a CBOR tagged byte string (tag 257, major type 6).
//
// The result is tag 257 wrapping a CBOR byte string containing a complete
// MIME message (all headers and body) per RFC 2045. Unlike tag 36, which
// uses a text string and requires valid UTF-8, tag 257 uses a byte string
// and can represent MIME messages with binary or 8-bit transfer encodings.
//
// No MIME validation is performed — the caller is responsible for providing
// a valid MIME message.
//
// See: http://peteroupc.github.io/CBOR/binarymime.html
func Marshal(value []byte) ([]byte, error) {
	return tags.Marshal(tagnumber.BinaryMIMEMessage, value)
}
