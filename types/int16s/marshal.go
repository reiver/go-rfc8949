package int16s

import (
	"github.com/reiver/go-rfc8949/initialbyte"
	"github.com/reiver/go-rfc8949/types/int8s"
	"github.com/reiver/go-rfc8949/types/uint16s"
)

func Marshal(value int16) ([]byte, error) {

	if 0 <= value {
		return uint16s.Marshal(uint16(value))
	}

	if -128 <= value {
		return int8s.Marshal(int8(value))
	}

	if -256 <= value {
		return []byte{initialbyte.Int8, byte((value + 1)*-1)}, nil
	}

	{
		var x uint16 = uint16((value + 1)*-1)
		return []byte{initialbyte.Int16,
			byte((x & 0xFF_00) >> 8),
			byte((x & 0x00_FF)     ),
		}, nil
	}
}
