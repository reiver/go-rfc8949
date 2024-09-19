package int32s

import (
	"github.com/reiver/go-rfc8949/initialbyte"
	"github.com/reiver/go-rfc8949/types/int16s"
	"github.com/reiver/go-rfc8949/types/uint32s"
)

func Marshal(value int32) ([]byte, error) {

	if 0 <= value {
		return uint32s.Marshal(uint32(value))
	}

	if -32768 <= value {
		return int16s.Marshal(int16(value))
	}

	{
		var x uint32 = uint32((value + 1)*-1)
		return []byte{initialbyte.Int32,
			byte((x & 0xFF_00_00_00) >> (8*3)),
			byte((x & 0x00_FF_00_00) >> (8*2)),
			byte((x & 0x00_00_FF_00) >> (8*1)),
			byte((x & 0x00_00_00_FF)     ),
		}, nil
	}
}
