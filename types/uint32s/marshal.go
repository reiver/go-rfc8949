package uint32s

import (
	"github.com/reiver/go-rfc8949/initialbyte"
	"github.com/reiver/go-rfc8949/types/uint16s"
)

func Marshal(value uint32) ([]byte, error) {

	switch {
	case value < 65536:
		return uint16s.Marshal(uint16(value))
	}

	return []byte{initialbyte.Uint32,
		byte( (value & 0xff_00_00_00) >> (8*3) ),
		byte( (value & 0x00_ff_00_00) >> (8*2) ),
		byte( (value & 0x00_00_ff_00) >> (8*1) ),
		byte( (value & 0x00_00_00_ff)          ),
	}, nil
}
