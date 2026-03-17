package uuids

import (
	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a [16]byte as a CBOR tagged binary UUID (tag 37, major type 6).
//
// The result is a tag 37 wrapping a 16-byte CBOR byte string.
//
// This function does not validate the UUID variant or version bits.
// Any 16 bytes are accepted as a valid tag 37 payload.
//
// The [16]byte parameter type is compatible with google/uuid.UUID,
// which is defined as [16]byte.
func Marshal(value [16]byte) ([]byte, error) {
	return tags.Marshal(tagnumber.BinaryUUID, value[:])
}
