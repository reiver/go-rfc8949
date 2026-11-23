package uint64s

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrNotUnsignedInteger = erorr.Error("rfc8949: not a CBOR unsigned integer (major type 0)")
)
