package int64s_test

import (
	"math"
	"testing"

	"github.com/reiver/go-rfc8949/types/int64s"
)

func TestUnmarshal(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected int64
	}{
		// RFC 8949 Appendix A examples.
		{
			Data:     []byte{0x00},
			Expected: 0,
		},
		{
			Data:     []byte{0x19, 0x03, 0xe8},
			Expected: 1000,
		},
		{
			Data:     []byte{0x20},
			Expected: -1,
		},
		{
			Data:     []byte{0x29},
			Expected: -10,
		},
		{
			Data:     []byte{0x38, 0x63},
			Expected: -100,
		},
		{
			Data:     []byte{0x39, 0x03, 0xe7},
			Expected: -1000,
		},



		// The int64 limits.
		{
			Data:     []byte{0x1b, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			Expected: math.MaxInt64,
		},
		{
			Data:     []byte{0x3b, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			Expected: math.MinInt64,
		},
	}

	for testNumber, test := range tests {

		actual, rest, err := int64s.Unmarshal(test.Data)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA: %#v", test.Data)
			continue
		}

		if test.Expected != actual || 0 != len(rest) {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", test.Expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("REST: %#v", rest)
			t.Logf("DATA: %#v", test.Data)
			continue
		}
	}
}

func TestUnmarshal_errors(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected error
	}{
		{
			// 18446744073709551615
			Data:     []byte{0x1b, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			Expected: int64s.ErrOverflow,
		},
		{
			// -18446744073709551616
			Data:     []byte{0x3b, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			Expected: int64s.ErrOverflow,
		},
		{
			// 2^63
			Data:     []byte{0x1b, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			Expected: int64s.ErrOverflow,
		},
		{
			// h''
			Data:     []byte{0x40},
			Expected: int64s.ErrNotInteger,
		},
	}

	for testNumber, test := range tests {

		_, _, err := int64s.Unmarshal(test.Data)

		if test.Expected != err {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", test.Expected)
			t.Logf("ACTUAL:   %v", err)
			t.Logf("DATA: %#v", test.Data)
			continue
		}
	}
}

func TestUnmarshal_roundTrip(t *testing.T) {

	values := []int64{
		0, 1, -1, 23, -24, -25, 255, -256, -257, 65535, -65536, -65537, 2147483647, -2147483648, -2147483649,
		4294967295, -4294967296, -4294967297, math.MaxInt64, math.MinInt64,
	}

	for testNumber, value := range values {

		p, err := int64s.Marshal(value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error from Marshal but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		actual, rest, err := int64s.Unmarshal(p)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error from Unmarshal but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %d", value)
			continue
		}

		if value != actual || 0 != len(rest) {
			t.Errorf("For test #%d, the round-trip did not work.", testNumber)
			t.Logf("EXPECTED: %d", value)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("REST: %#v", rest)
			continue
		}
	}
}
