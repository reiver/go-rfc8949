package bytestrings

import (
	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/majortype"
)

// Unmarshal decodes a definite-length CBOR byte string (major type 2) from the start of 'p'.
// It returns the content, and the rest of 'p' after the data item.
//
// The returned content is a copy — it does not share memory with 'p'.
//
// Indefinite-length byte strings are not (yet) supported; Unmarshal returns [ErrIndefiniteLength] for them.
//
// Example usage:
//
//	value, rest, err := bytestrings.Unmarshal(p)
func Unmarshal(p []byte) (value []byte, rest []byte, err error) {
	h, err := head.Decode(p)
	if nil != err {
		return nil, p, err
	}
	if majortype.ByteString != h.MajorType {
		return nil, p, ErrNotByteString
	}
	if h.IsIndefinite() {
		return nil, p, ErrIndefiniteLength
	}

	var remaining []byte = p[h.Length:]
	if uint64(len(remaining)) < h.Argument {
		return nil, p, head.ErrUnexpectedEnd
	}

	var length int = int(h.Argument)

	value = make([]byte, length)
	copy(value, remaining[:length])

	return value, remaining[length:], nil
}
