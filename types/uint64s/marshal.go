package uint64s

import (
	"github.com/reiver/go-rfc8949/initialbyte"
	"github.com/reiver/go-rfc8949/types/uint32s"
)

func Marshal(value uint64) ([]byte, error) {

	switch {
	case value < 4294967296:
		return uint32s.Marshal(uint32(value))
	}

	return []byte{initialbyte.Uint64,
		byte( (value & 0xff_00_00_00_00_00_00_00) >> (8*7) ),
		byte( (value & 0x00_ff_00_00_00_00_00_00) >> (8*6) ),
		byte( (value & 0x00_00_ff_00_00_00_00_00) >> (8*5) ),
		byte( (value & 0x00_00_00_ff_00_00_00_00) >> (8*4) ),
		byte( (value & 0x00_00_00_00_ff_00_00_00) >> (8*3) ),
		byte( (value & 0x00_00_00_00_00_ff_00_00) >> (8*2) ),
		byte( (value & 0x00_00_00_00_00_00_ff_00) >> (8*1) ),
		byte( (value & 0x00_00_00_00_00_00_00_ff)          ),
	}, nil
}
