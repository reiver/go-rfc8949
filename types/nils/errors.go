package nils

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrNotNull = erorr.Error("rfc8949: not a CBOR null (simple value 22)")
)
