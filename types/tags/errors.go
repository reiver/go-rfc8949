package tags

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrNotTagged = erorr.Error("rfc8949: not a CBOR tagged data item (major type 6)")
)
