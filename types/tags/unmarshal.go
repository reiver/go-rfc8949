package tags

import (
	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/majortype"
)

// UnmarshalHead decodes the head of a CBOR tagged data item (major type 6) from the start of 'p'.
// It returns the tag number, and the rest of 'p' after the head — which starts with the tag content (a single data item).
//
// The caller decodes the tag content itself.
//
// Example usage:
//
//	tagNumber, rest, err := tags.UnmarshalHead(p)
func UnmarshalHead(p []byte) (tagNumber uint64, rest []byte, err error) {
	h, err := head.Decode(p)
	if nil != err {
		return 0, p, err
	}
	if majortype.Tagged != h.MajorType {
		return 0, p, ErrNotTagged
	}
	if h.IsIndefinite() {
		return 0, p, ErrNotTagged
	}

	return h.Argument, p[h.Length:], nil
}
