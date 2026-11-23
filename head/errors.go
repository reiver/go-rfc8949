package head

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrEmpty                         = erorr.Error("rfc8949: empty data — no CBOR data item")
	ErrReservedAdditionalInformation = erorr.Error("rfc8949: reserved additional information (28, 29, or 30) — not well-formed")
	ErrUnexpectedEnd                 = erorr.Error("rfc8949: unexpected end of data")
)
