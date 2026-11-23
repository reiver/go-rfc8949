package textstrings

import (
	"unicode/utf8"

	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/majortype"
)

// Unmarshal decodes a definite-length CBOR text string (major type 3) from the start of 'p'.
// It returns the string, and the rest of 'p' after the data item.
//
// A text string that is not valid UTF-8 is well-formed but invalid CBOR (RFC 8949 section 3.1);
// Unmarshal returns [ErrInvalidUTF8] for it.
//
// Indefinite-length text strings are not (yet) supported; Unmarshal returns [ErrIndefiniteLength] for them.
//
// Example usage:
//
//	value, rest, err := textstrings.Unmarshal(p)
func Unmarshal(p []byte) (value string, rest []byte, err error) {
	h, err := head.Decode(p)
	if nil != err {
		return "", p, err
	}
	if majortype.TextString != h.MajorType {
		return "", p, ErrNotTextString
	}
	if h.IsIndefinite() {
		return "", p, ErrIndefiniteLength
	}

	var remaining []byte = p[h.Length:]
	if uint64(len(remaining)) < h.Argument {
		return "", p, head.ErrUnexpectedEnd
	}

	var length int = int(h.Argument)
	var content []byte = remaining[:length]

	if !utf8.Valid(content) {
		return "", p, ErrInvalidUTF8
	}

	return string(content), remaining[length:], nil
}
