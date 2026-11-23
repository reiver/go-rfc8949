package float64s

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrNotFloat = erorr.Error("rfc8949: not a CBOR floating-point number (major type 7, additional information 25, 26, or 27)")
)
