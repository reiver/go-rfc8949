package bytestrings

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrIndefiniteLength = erorr.Error("rfc8949: indefinite-length CBOR byte string not supported")
	ErrNotByteString    = erorr.Error("rfc8949: not a CBOR byte string (major type 2)")
)
