package maps

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrIndefiniteLength = erorr.Error("rfc8949: indefinite-length CBOR map not supported")
	ErrNotMap           = erorr.Error("rfc8949: not a CBOR map (major type 5)")
)
