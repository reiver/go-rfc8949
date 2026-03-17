package bytestrings

import (
	"github.com/reiver/go-rfc8949/initialbyte"
)

func Marshal(value []byte) ([]byte, error) {

	length := len(value)

	switch {
	case length <= 23:
		switch length {
		case                                                  0:
			return append([]byte{initialbyte.ByteStringLen0},  value...), nil
		case                                                  1:
			return append([]byte{initialbyte.ByteStringLen1},  value...), nil
		case                                                  2:
			return append([]byte{initialbyte.ByteStringLen2},  value...), nil
		case                                                  3:
			return append([]byte{initialbyte.ByteStringLen3},  value...), nil
		case                                                  4:
			return append([]byte{initialbyte.ByteStringLen4},  value...), nil
		case                                                  5:
			return append([]byte{initialbyte.ByteStringLen5},  value...), nil
		case                                                  6:
			return append([]byte{initialbyte.ByteStringLen6},  value...), nil
		case                                                  7:
			return append([]byte{initialbyte.ByteStringLen7},  value...), nil
		case                                                  8:
			return append([]byte{initialbyte.ByteStringLen8},  value...), nil
		case                                                  9:
			return append([]byte{initialbyte.ByteStringLen9},  value...), nil
		case                                                  10:
			return append([]byte{initialbyte.ByteStringLen10}, value...), nil
		case                                                  11:
			return append([]byte{initialbyte.ByteStringLen11}, value...), nil
		case                                                  12:
			return append([]byte{initialbyte.ByteStringLen12}, value...), nil
		case                                                  13:
			return append([]byte{initialbyte.ByteStringLen13}, value...), nil
		case                                                  14:
			return append([]byte{initialbyte.ByteStringLen14}, value...), nil
		case                                                  15:
			return append([]byte{initialbyte.ByteStringLen15}, value...), nil
		case                                                  16:
			return append([]byte{initialbyte.ByteStringLen16}, value...), nil
		case                                                  17:
			return append([]byte{initialbyte.ByteStringLen17}, value...), nil
		case                                                  18:
			return append([]byte{initialbyte.ByteStringLen18}, value...), nil
		case                                                  19:
			return append([]byte{initialbyte.ByteStringLen19}, value...), nil
		case                                                  20:
			return append([]byte{initialbyte.ByteStringLen20}, value...), nil
		case                                                  21:
			return append([]byte{initialbyte.ByteStringLen21}, value...), nil
		case                                                  22:
			return append([]byte{initialbyte.ByteStringLen22}, value...), nil
		case                                                  23:
			return append([]byte{initialbyte.ByteStringLen23}, value...), nil
		}

	case length < 256: // 2^8
		var result []byte
		result = append(result, initialbyte.ByteStringLenUint8, byte(length))
		result = append(result, value...)
		return result, nil

	case length < 65536: // 2^16
		var result []byte
		result = append(result, initialbyte.ByteStringLenUint16,
			byte((length & 0xFF_00) >> 8),
			byte((length & 0x00_FF)     ),
		)
		result = append(result, value...)
		return result, nil

	case length < 4294967296: // 2^32
		var result []byte
		result = append(result, initialbyte.ByteStringLenUint32,
			byte((length & 0xFF_00_00_00) >> (8*3)),
			byte((length & 0x00_FF_00_00) >> (8*2)),
			byte((length & 0x00_00_FF_00) >> (8*1)),
			byte((length & 0x00_00_00_FF)         ),
		)
		result = append(result, value...)
		return result, nil

	default:
		var x uint64 = uint64(length)
		var result []byte
		result = append(result, initialbyte.ByteStringLenUint64,
			byte((x & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
			byte((x & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
			byte((x & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
			byte((x & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
			byte((x & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
			byte((x & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
			byte((x & 0x00_00_00_00_00_00_FF_00) >> (8*1)),
			byte((x & 0x00_00_00_00_00_00_00_FF)         ),
		)
		result = append(result, value...)
		return result, nil
	}

	// This should never be reached, but the compiler requires it.
	return nil, nil
}
