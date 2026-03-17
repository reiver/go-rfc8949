package arrays

import (
	"fmt"

	"github.com/reiver/go-rfc8949/initialbyte"
	"github.com/reiver/go-rfc8949/types/bools"
	"github.com/reiver/go-rfc8949/types/bytestrings"
	"github.com/reiver/go-rfc8949/types/int8s"
	"github.com/reiver/go-rfc8949/types/int16s"
	"github.com/reiver/go-rfc8949/types/int32s"
	"github.com/reiver/go-rfc8949/types/int64s"
	"github.com/reiver/go-rfc8949/types/nils"
	"github.com/reiver/go-rfc8949/types/textstrings"
	"github.com/reiver/go-rfc8949/types/uint8s"
	"github.com/reiver/go-rfc8949/types/uint16s"
	"github.com/reiver/go-rfc8949/types/uint32s"
	"github.com/reiver/go-rfc8949/types/uint64s"
)

func Marshal(value []any) ([]byte, error) {

	length := len(value)

	var result []byte

	switch {
	case length <= 23:
		switch length {
		case  0:
			result = append(result, initialbyte.ArrayLen0)
		case  1:
			result = append(result, initialbyte.ArrayLen1)
		case  2:
			result = append(result, initialbyte.ArrayLen2)
		case  3:
			result = append(result, initialbyte.ArrayLen3)
		case  4:
			result = append(result, initialbyte.ArrayLen4)
		case  5:
			result = append(result, initialbyte.ArrayLen5)
		case  6:
			result = append(result, initialbyte.ArrayLen6)
		case  7:
			result = append(result, initialbyte.ArrayLen7)
		case  8:
			result = append(result, initialbyte.ArrayLen8)
		case  9:
			result = append(result, initialbyte.ArrayLen9)
		case 10:
			result = append(result, initialbyte.ArrayLen10)
		case 11:
			result = append(result, initialbyte.ArrayLen11)
		case 12:
			result = append(result, initialbyte.ArrayLen12)
		case 13:
			result = append(result, initialbyte.ArrayLen13)
		case 14:
			result = append(result, initialbyte.ArrayLen14)
		case 15:
			result = append(result, initialbyte.ArrayLen15)
		case 16:
			result = append(result, initialbyte.ArrayLen16)
		case 17:
			result = append(result, initialbyte.ArrayLen17)
		case 18:
			result = append(result, initialbyte.ArrayLen18)
		case 19:
			result = append(result, initialbyte.ArrayLen19)
		case 20:
			result = append(result, initialbyte.ArrayLen20)
		case 21:
			result = append(result, initialbyte.ArrayLen21)
		case 22:
			result = append(result, initialbyte.ArrayLen22)
		case 23:
			result = append(result, initialbyte.ArrayLen23)
		}

	case length < 256: // 2^8
		result = append(result, initialbyte.ArrayLenUint8, byte(length))

	case length < 65536: // 2^16
		result = append(result, initialbyte.ArrayLenUint16,
			byte((length & 0xFF_00) >> 8),
			byte((length & 0x00_FF)     ),
		)

	case length < 4294967296: // 2^32
		result = append(result, initialbyte.ArrayLenUint32,
			byte((length & 0xFF_00_00_00) >> (8*3)),
			byte((length & 0x00_FF_00_00) >> (8*2)),
			byte((length & 0x00_00_FF_00) >> (8*1)),
			byte((length & 0x00_00_00_FF)         ),
		)

	default:
		var x uint64 = uint64(length)
		result = append(result, initialbyte.ArrayLenUint64,
			byte((x & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
			byte((x & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
			byte((x & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
			byte((x & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
			byte((x & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
			byte((x & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
			byte((x & 0x00_00_00_00_00_00_FF_00) >> (8*1)),
			byte((x & 0x00_00_00_00_00_00_00_FF)         ),
		)
	}

	for i, item := range value {
		encoded, err := marshalItem(item)
		if nil != err {
			return nil, fmt.Errorf("rfc8949: array item [%d]: %w", i, err)
		}
		result = append(result, encoded...)
	}

	return result, nil
}

func marshalItem(item any) ([]byte, error) {

	switch v := item.(type) {

	case nil:
		return nils.Marshal()

	case bool:
		return bools.Marshal(v)

	case uint8:
		return uint8s.Marshal(v)
	case uint16:
		return uint16s.Marshal(v)
	case uint32:
		return uint32s.Marshal(v)
	case uint64:
		return uint64s.Marshal(v)
	case uint:
		return uint64s.Marshal(uint64(v))

	case int8:
		return int8s.Marshal(v)
	case int16:
		return int16s.Marshal(v)
	case int32:
		return int32s.Marshal(v)
	case int64:
		return int64s.Marshal(v)
	case int:
		return int64s.Marshal(int64(v))

	case string:
		return textstrings.Marshal(v)

	case []byte:
		return bytestrings.Marshal(v)

	case []any:
		return Marshal(v)

	default:
		return nil, fmt.Errorf("rfc8949: unsupported type: %T", item)
	}
}
