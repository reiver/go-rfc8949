package textstrings

import (
	"fmt"
	"unicode/utf8"

	"github.com/reiver/go-rfc8949/initialbyte"
)

func Marshal(value string) ([]byte, error) {

	if !utf8.ValidString(value) {
		return nil, fmt.Errorf("rfc8949: text string is not valid UTF-8")
	}

	var content []byte = []byte(value)
	length := len(content)

	switch {
	case length <= 23:
		switch length {
		case                                                  0:
			return append([]byte{initialbyte.TextStringLen0},  content...), nil
		case                                                  1:
			return append([]byte{initialbyte.TextStringLen1},  content...), nil
		case                                                  2:
			return append([]byte{initialbyte.TextStringLen2},  content...), nil
		case                                                  3:
			return append([]byte{initialbyte.TextStringLen3},  content...), nil
		case                                                  4:
			return append([]byte{initialbyte.TextStringLen4},  content...), nil
		case                                                  5:
			return append([]byte{initialbyte.TextStringLen5},  content...), nil
		case                                                  6:
			return append([]byte{initialbyte.TextStringLen6},  content...), nil
		case                                                  7:
			return append([]byte{initialbyte.TextStringLen7},  content...), nil
		case                                                  8:
			return append([]byte{initialbyte.TextStringLen8},  content...), nil
		case                                                  9:
			return append([]byte{initialbyte.TextStringLen9},  content...), nil
		case                                                  10:
			return append([]byte{initialbyte.TextStringLen10}, content...), nil
		case                                                  11:
			return append([]byte{initialbyte.TextStringLen11}, content...), nil
		case                                                  12:
			return append([]byte{initialbyte.TextStringLen12}, content...), nil
		case                                                  13:
			return append([]byte{initialbyte.TextStringLen13}, content...), nil
		case                                                  14:
			return append([]byte{initialbyte.TextStringLen14}, content...), nil
		case                                                  15:
			return append([]byte{initialbyte.TextStringLen15}, content...), nil
		case                                                  16:
			return append([]byte{initialbyte.TextStringLen16}, content...), nil
		case                                                  17:
			return append([]byte{initialbyte.TextStringLen17}, content...), nil
		case                                                  18:
			return append([]byte{initialbyte.TextStringLen18}, content...), nil
		case                                                  19:
			return append([]byte{initialbyte.TextStringLen19}, content...), nil
		case                                                  20:
			return append([]byte{initialbyte.TextStringLen20}, content...), nil
		case                                                  21:
			return append([]byte{initialbyte.TextStringLen21}, content...), nil
		case                                                  22:
			return append([]byte{initialbyte.TextStringLen22}, content...), nil
		case                                                  23:
			return append([]byte{initialbyte.TextStringLen23}, content...), nil
		}

	case length < 256: // 2^8
		var result []byte
		result = append(result, initialbyte.TextStringLenUint8, byte(length))
		result = append(result, content...)
		return result, nil

	case length < 65536: // 2^16
		var result []byte
		result = append(result, initialbyte.TextStringLenUint16,
			byte((length & 0xFF_00) >> 8),
			byte((length & 0x00_FF)     ),
		)
		result = append(result, content...)
		return result, nil

	case length < 4294967296: // 2^32
		var result []byte
		result = append(result, initialbyte.TextStringLenUint32,
			byte((length & 0xFF_00_00_00) >> (8*3)),
			byte((length & 0x00_FF_00_00) >> (8*2)),
			byte((length & 0x00_00_FF_00) >> (8*1)),
			byte((length & 0x00_00_00_FF)         ),
		)
		result = append(result, content...)
		return result, nil

	default:
		var x uint64 = uint64(length)
		var result []byte
		result = append(result, initialbyte.TextStringLenUint64,
			byte((x & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
			byte((x & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
			byte((x & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
			byte((x & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
			byte((x & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
			byte((x & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
			byte((x & 0x00_00_00_00_00_00_FF_00) >> (8*1)),
			byte((x & 0x00_00_00_00_00_00_00_FF)         ),
		)
		result = append(result, content...)
		return result, nil
	}

	// This should never be reached, but the compiler requires it.
	return nil, nil
}
