package float64s

import (
	"math"

	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/initialbyte/additionalinfo"
	"github.com/reiver/go-rfc8949/majortype"
)

// Unmarshal decodes a CBOR floating-point number — half-precision (16-bit), single-precision (32-bit), or double-precision (64-bit) —
// from the start of 'p', into a float64.
// It returns the value, and how many bits it was encoded in (16, 32, or 64), and the rest of 'p' after the data item.
//
// Every half- and single-precision value can be represented exactly as a float64.
//
// Example usage:
//
//	value, bits, rest, err := float64s.Unmarshal(p)
func Unmarshal(p []byte) (value float64, bits int, rest []byte, err error) {
	h, err := head.Decode(p)
	if nil != err {
		return 0, 0, p, err
	}
	if majortype.FloatingPoint != h.MajorType {
		return 0, 0, p, ErrNotFloat
	}

	switch h.AdditionalInformation {
	case additionalinfo.Uint16:
		return float16bitsToFloat64(uint16(h.Argument)), 16, p[h.Length:], nil
	case additionalinfo.Uint32:
		return float64(math.Float32frombits(uint32(h.Argument))), 32, p[h.Length:], nil
	case additionalinfo.Uint64:
		return math.Float64frombits(h.Argument), 64, p[h.Length:], nil
	default:
		return 0, 0, p, ErrNotFloat
	}
}

// float16bitsToFloat64 converts the bits of an IEEE 754 half-precision number into a float64.
// (This follows the decode_half example in RFC 8949 Appendix D.)
func float16bitsToFloat64(half uint16) float64 {
	var exponent int = int((half >> 10) & 0x1f)
	var mantissa int = int(half & 0x3ff)

	var value float64
	switch exponent {
	case 0:
		value = math.Ldexp(float64(mantissa), -24)
	case 31:
		if 0 == mantissa {
			value = math.Inf(1)
		} else {
			value = math.NaN()
		}
	default:
		value = math.Ldexp(float64(mantissa+1024), exponent-25)
	}

	if 0 != (half & 0x8000) {
		return -value
	}
	return value
}
