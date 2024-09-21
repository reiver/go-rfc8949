package prefix

import (
	"github.com/reiver/go-rfc8949/initialbyte/additionalinfo"
)

func Prefix(majorType byte, length uint64) []byte {

	var p []byte

	switch {
	case length <= 23:
		var first byte = majorType | byte(length)

		p = append(p, first)

	case length < 256: // == 2^8
		var first byte = majorType | additionalinfo.Uint8

		p = append(p, first)
		p = append(p, uint8(length))

	case length < 65536 : // == 2^16
		var first byte = majorType | additionalinfo.Uint16

		p = append(p, first)
		p = append(p, uint8((0xFF_00 & length) >> 8))
		p = append(p, uint8((0x00_FF & length)     ))

	case length < 4294967296: // == 2^32
		var first byte = majorType | additionalinfo.Uint32

		p = append(p, first)
		p = append(p, uint8((0xFF_00_00_00 & length) >> (8*3)))
		p = append(p, uint8((0x00_FF_00_00 & length) >> (8*2)))
		p = append(p, uint8((0x00_00_FF_00 & length) >>  8   ))
		p = append(p, uint8((0x00_00_00_FF & length)         ))

	default:
		var first byte = majorType | additionalinfo.Uint64

		p = append(p, first)
		p = append(p, uint8((0xFF_00_00_00_00_00_00_00 & length) >> (8*7)))
		p = append(p, uint8((0x00_FF_00_00_00_00_00_00 & length) >> (8*6)))
		p = append(p, uint8((0x00_00_FF_00_00_00_00_00 & length) >> (8*5)))
		p = append(p, uint8((0x00_00_00_FF_00_00_00_00 & length) >> (8*4)))
		p = append(p, uint8((0x00_00_00_00_FF_00_00_00 & length) >> (8*3)))
		p = append(p, uint8((0x00_00_00_00_00_FF_00_00 & length) >> (8*2)))
		p = append(p, uint8((0x00_00_00_00_00_00_FF_00 & length) >>  8   ))
		p = append(p, uint8((0x00_00_00_00_00_00_00_FF & length)         ))
	}

	return p
}
