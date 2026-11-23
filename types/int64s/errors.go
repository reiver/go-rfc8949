package int64s

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrNotInteger = erorr.Error("rfc8949: not a CBOR integer (major type 0 or 1)")
	ErrOverflow   = erorr.Error("rfc8949: CBOR integer does not fit in an int64")
)
