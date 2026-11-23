package float64s

import (
	"math"

	"github.com/reiver/go-rfc8949/initialbyte"
)

// Marshal encodes a float64 as a CBOR double-precision (64-bit) floating-point number (major type 7, additional information 27).
//
// Marshal always uses 64 bits. It does not shorten values that would fit in 16 or 32 bits
// (which RFC 8949's "preferred serialization" would do).
//
// Example usage:
//
//	p, err := float64s.Marshal(1.1)
//	
//	// p == []byte{0xfb, 0x3f, 0xf1, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a}
func Marshal(value float64) ([]byte, error) {
	var x uint64 = math.Float64bits(value)

	return []byte{initialbyte.Float64,
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
