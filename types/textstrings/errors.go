package textstrings

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrIndefiniteLength = erorr.Error("rfc8949: indefinite-length CBOR text string not supported")
	ErrInvalidUTF8      = erorr.Error("rfc8949: CBOR text string is not valid UTF-8")
	ErrNotTextString    = erorr.Error("rfc8949: not a CBOR text string (major type 3)")
)
