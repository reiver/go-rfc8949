/*
Package head decodes the "head" of a CBOR data item.

From IETF RFC-8949:

	"""
	The initial byte and any additional bytes consumed to construct the argument are collectively referred to as the head of the data item.
	"""

The head is: the initial byte (3-bit major type + 5-bit additional information), followed by 0, 1, 2, 4, or 8 bytes of argument.

	┌───────────────────────────┬─────────────────────────────────────────────────┐
	│ Additional Information    │ Argument                                        │
	├───────────────────────────┼─────────────────────────────────────────────────┤
	│ 0..23                     │ the additional information itself               │
	│ 24                        │ the following 1 byte                            │
	│ 25                        │ the following 2 bytes (big-endian)              │
	│ 26                        │ the following 4 bytes (big-endian)              │
	│ 27                        │ the following 8 bytes (big-endian)              │
	│ 28, 29, 30                │ reserved — not well-formed                      │
	│ 31                        │ indefinite length (or "break", for major type 7)│
	└───────────────────────────┴─────────────────────────────────────────────────┘
*/
package head

import (
	"github.com/reiver/go-rfc8949/initialbyte/additionalinfo"
	"github.com/reiver/go-rfc8949/majortype"
)

// Head is the decoded head of a CBOR data item.
type Head struct {
	// MajorType is the major type, in the same form as the constants in package majortype
	// (i.e., the high-order 3 bits of the initial byte, still in place — ex: majortype.TextString).
	MajorType byte

	// AdditionalInformation is the low-order 5 bits of the initial byte (0..31).
	AdditionalInformation byte

	// Argument is the argument.
	//
	// For major types 0..6 it is the value, the length, the count, or the tag number.
	// For major type 7 with additional information 25, 26, or 27, it is the raw bits of the floating-point number.
	// For major type 7 with additional information 0..24, it is the simple value.
	// It is zero when additional information is 31 (indefinite length / "break").
	Argument uint64

	// Length is the number of bytes the head takes up (1, 2, 3, 5, or 9).
	Length int
}

// IsIndefinite returns whether the head has additional information 31 —
// i.e., an indefinite-length byte string, text string, array, or map; or (for major type 7) the "break" stop code.
func (receiver Head) IsIndefinite() bool {
	return 31 == receiver.AdditionalInformation
}

// IsMinimal returns whether the argument was encoded in the fewest bytes possible
// (the "preferred serialization" of RFC 8949 section 4.1, which deterministic encodings require).
//
// For major type 7, floating-point numbers (additional information 25, 26, 27) are not judged by this — it returns true for them.
// A simple value encoded with additional information 24 is minimal only if it is 32 or greater (RFC 8949 section 3.3).
func (receiver Head) IsMinimal() bool {
	if majortype.MajorType7 == receiver.MajorType {
		switch receiver.AdditionalInformation {
		case additionalinfo.Uint8:
			return 32 <= receiver.Argument
		default:
			return true
		}
	}

	switch receiver.AdditionalInformation {
	case additionalinfo.Uint8:
		return 24 <= receiver.Argument
	case additionalinfo.Uint16:
		return 0xFF < receiver.Argument
	case additionalinfo.Uint32:
		return 0xFF_FF < receiver.Argument
	case additionalinfo.Uint64:
		return 0xFF_FF_FF_FF < receiver.Argument
	default:
		return true
	}
}

// Decode decodes the head at the start of 'p'.
//
// Decode does not look past the head — for example, it does not check that a byte string's content is all there.
//
// Example usage:
//
//	h, err := head.Decode(p)
//	if nil != err {
//		return err
//	}
//
//	var rest []byte = p[h.Length:]
func Decode(p []byte) (Head, error) {
	var nada Head

	if len(p) <= 0 {
		return nada, ErrEmpty
	}

	var initialByte byte = p[0]

	var h = Head{
		MajorType:             majortype.MajorTypeFromInitialByte(initialByte),
		AdditionalInformation: initialByte & 0b000_11111,
	}

	var size int
	switch h.AdditionalInformation {
	case additionalinfo.Uint8:
		size = 1
	case additionalinfo.Uint16:
		size = 2
	case additionalinfo.Uint32:
		size = 4
	case additionalinfo.Uint64:
		size = 8
	case 28, 29, 30:
		return nada, ErrReservedAdditionalInformation
	case 31:
		h.Length = 1
		return h, nil
	default:
		h.Argument = uint64(h.AdditionalInformation)
		h.Length = 1
		return h, nil
	}

	if len(p) < 1+size {
		return nada, ErrUnexpectedEnd
	}

	var argument uint64
	for i:=1; i<=size; i++ {
		argument = (argument << 8) | uint64(p[i])
	}

	h.Argument = argument
	h.Length = 1 + size

	return h, nil
}
