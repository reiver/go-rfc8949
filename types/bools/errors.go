package bools

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrNotBool = erorr.Error("rfc8949: not a CBOR boolean (simple value 20 or 21)")
)
