package arrays

import (
	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/majortype"
)

// UnmarshalHead decodes the head of a definite-length CBOR array (major type 4) from the start of 'p'.
// It returns the number of data items in the array, and the rest of 'p' after the head — which starts with the first data item.
//
// The caller decodes the data items themselves.
//
// Indefinite-length arrays are not (yet) supported; UnmarshalHead returns [ErrIndefiniteLength] for them.
//
// Example usage:
//
//	count, rest, err := arrays.UnmarshalHead(p)
//	if nil != err {
//		return err
//	}
//	
//	for i:=uint64(0); i<count; i++ {
//		// decode a data item from 'rest', and update 'rest'.
//	}
func UnmarshalHead(p []byte) (count uint64, rest []byte, err error) {
	h, err := head.Decode(p)
	if nil != err {
		return 0, p, err
	}
	if majortype.Array != h.MajorType {
		return 0, p, ErrNotArray
	}
	if h.IsIndefinite() {
		return 0, p, ErrIndefiniteLength
	}

	return h.Argument, p[h.Length:], nil
}
