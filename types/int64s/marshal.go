package int64s

import (
	"github.com/reiver/go-rfc8949/initialbyte"
	"github.com/reiver/go-rfc8949/types/int32s"
	"github.com/reiver/go-rfc8949/types/uint64s"
)

func Marshal(value int64) ([]byte, error) {

	if 0 <= value {
		return uint64s.Marshal(uint64(value))
	}

	if -4294967296 <= value {
		return int32s.Marshal(int32(value))
	}

	{
		var x uint64 = uint64((value + 1)*-1)
		return []byte{initialbyte.Int64,
			byte((x & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
			byte((x & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
			byte((x & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
			byte((x & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
			byte((x & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
			byte((x & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
			byte((x & 0x00_00_00_00_00_00_FF_00) >> (8*1)),
			byte((x & 0x00_00_00_00_00_00_00_FF)         ),
		}, nil
	}
}
