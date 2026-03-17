package tags

import (
	"github.com/reiver/go-rfc8949/initialbyte"
	"github.com/reiver/go-rfc8949/internal/marshalitem"
)

// Marshal encodes a CBOR tagged data item (major type 6).
//
// The tagNumber is an integer in the range 0..2^64-1 identifying the tag.
// The content is the single enclosed data item, which will be marshaled
// using the standard type dispatch.
//
// This function does not validate that the content type is appropriate
// for the given tag number. For example, it will not verify that tag 0
// wraps a text string or that tag 2 wraps a byte string. That is the
// caller's responsibility.
func Marshal(tagNumber uint64, content any) ([]byte, error) {

	var result []byte

	switch {
	case tagNumber <= 23:
		switch tagNumber {
		case  0:
			result = append(result, initialbyte.TagNum0)
		case  1:
			result = append(result, initialbyte.TagNum1)
		case  2:
			result = append(result, initialbyte.TagNum2)
		case  3:
			result = append(result, initialbyte.TagNum3)
		case  4:
			result = append(result, initialbyte.TagNum4)
		case  5:
			result = append(result, initialbyte.TagNum5)
		case  6:
			result = append(result, initialbyte.TagNum6)
		case  7:
			result = append(result, initialbyte.TagNum7)
		case  8:
			result = append(result, initialbyte.TagNum8)
		case  9:
			result = append(result, initialbyte.TagNum9)
		case 10:
			result = append(result, initialbyte.TagNum10)
		case 11:
			result = append(result, initialbyte.TagNum11)
		case 12:
			result = append(result, initialbyte.TagNum12)
		case 13:
			result = append(result, initialbyte.TagNum13)
		case 14:
			result = append(result, initialbyte.TagNum14)
		case 15:
			result = append(result, initialbyte.TagNum15)
		case 16:
			result = append(result, initialbyte.TagNum16)
		case 17:
			result = append(result, initialbyte.TagNum17)
		case 18:
			result = append(result, initialbyte.TagNum18)
		case 19:
			result = append(result, initialbyte.TagNum19)
		case 20:
			result = append(result, initialbyte.TagNum20)
		case 21:
			result = append(result, initialbyte.TagNum21)
		case 22:
			result = append(result, initialbyte.TagNum22)
		case 23:
			result = append(result, initialbyte.TagNum23)
		}

	case tagNumber < 256: // 2^8
		result = append(result, initialbyte.TagNumUint8, byte(tagNumber))

	case tagNumber < 65536: // 2^16
		result = append(result, initialbyte.TagNumUint16,
			byte((tagNumber & 0xFF_00) >> 8),
			byte((tagNumber & 0x00_FF)     ),
		)

	case tagNumber < 4294967296: // 2^32
		result = append(result, initialbyte.TagNumUint32,
			byte((tagNumber & 0xFF_00_00_00) >> (8*3)),
			byte((tagNumber & 0x00_FF_00_00) >> (8*2)),
			byte((tagNumber & 0x00_00_FF_00) >> (8*1)),
			byte((tagNumber & 0x00_00_00_FF)         ),
		)

	default:
		result = append(result, initialbyte.TagNumUint64,
			byte((tagNumber & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
			byte((tagNumber & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
			byte((tagNumber & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
			byte((tagNumber & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
			byte((tagNumber & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
			byte((tagNumber & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
			byte((tagNumber & 0x00_00_00_00_00_00_FF_00) >> (8*1)),
			byte((tagNumber & 0x00_00_00_00_00_00_00_FF)         ),
		)
	}

	encoded, err := marshalitem.MarshalItem(content)
	if nil != err {
		return nil, err
	}
	result = append(result, encoded...)

	return result, nil
}
