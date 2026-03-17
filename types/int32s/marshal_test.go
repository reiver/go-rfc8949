package int32s_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/int32s"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value int32
		Expected []byte
	}{

		// Positive values — delegated to uint32s.Marshal

		// 0–23: single byte, inline in initial byte (major type 0)
		{
			Value:    0,
			Expected: []byte{0x00},
		},
		{
			Value:    1,
			Expected: []byte{0x01},
		},
		{
			Value:    23,
			Expected: []byte{0x17},
		},

		// 24–255: 2 bytes (initial byte 0x18 + 1 value byte)
		{
			Value:    24,
			Expected: []byte{0x18, 24},
		},
		{
			Value:    255,
			Expected: []byte{0x18, 255},
		},

		// 256–65535: 3 bytes (initial byte 0x19 + 2 value bytes)
		{
			Value:    256,
			Expected: []byte{0x19, 0x01, 0x00},
		},
		{
			Value:    65535,
			Expected: []byte{0x19, 0xFF, 0xFF},
		},

		// 65536–2147483647: 5 bytes (initial byte 0x1a + 4 value bytes)
		{
			Value:    65536,
			Expected: []byte{0x1a, 0x00, 0x01, 0x00, 0x00},
		},
		{
			Value:    1000000,
			Expected: []byte{0x1a, 0x00, 0x0F, 0x42, 0x40},
		},
		{
			Value:    2147483647, // math.MaxInt32
			Expected: []byte{0x1a, 0x7F, 0xFF, 0xFF, 0xFF},
		},



		// Negative values delegated to int16s.Marshal (value >= -32768)

		// -1 to -24: single byte, inline in initial byte (major type 1)
		{
			Value:    -1,
			Expected: []byte{0x20},
		},
		{
			Value:    -10,
			Expected: []byte{0x29},
		},
		{
			Value:    -24,
			Expected: []byte{0x37},
		},

		// -25 to -256: 2 bytes (initial byte 0x38 + 1 argument byte)
		{
			Value:    -25,
			Expected: []byte{0x38, 24},
		},
		{
			Value:    -100,
			Expected: []byte{0x38, 99},
		},
		{
			Value:    -256,
			Expected: []byte{0x38, 255},
		},

		// -257 to -32768: 3 bytes (initial byte 0x39 + 2 argument bytes)
		{
			Value:    -257,
			Expected: []byte{0x39, 0x01, 0x00},
		},
		{
			Value:    -1000,
			Expected: []byte{0x39, 0x03, 0xE7},
		},

		// Boundary: -32768 (minimum int16, still delegated to int16s)
		{
			Value:    -32768,
			Expected: []byte{0x39, 0x7F, 0xFF},
		},



		// Values handled directly by int32s: -32769 to -65536
		// 3 bytes (initial byte 0x39 + 2 argument bytes)

		{
			Value:    -32769,
			Expected: []byte{0x39, 0x80, 0x00},
		},
		{
			Value:    -50000,
			Expected: []byte{0x39, 0xC3, 0x4F},
		},
		{
			Value:    -65535,
			Expected: []byte{0x39, 0xFF, 0xFE},
		},
		{
			Value:    -65536,
			Expected: []byte{0x39, 0xFF, 0xFF},
		},



		// Values handled directly by int32s: below -65536
		// 5 bytes (initial byte 0x3a + 4 argument bytes)

		{
			Value:    -65537,
			Expected: []byte{0x3a, 0x00, 0x01, 0x00, 0x00},
		},
		{
			Value:    -1000000,
			Expected: []byte{0x3a, 0x00, 0x0F, 0x42, 0x3F},
		},
		{
			Value:    -100000000,
			Expected: []byte{0x3a, 0x05, 0xF5, 0xE0, 0xFF},
		},
		{
			Value:    -2147483647,
			Expected: []byte{0x3a, 0x7F, 0xFF, 0xFF, 0xFE},
		},
		{
			Value:    -2147483648, // math.MinInt32
			Expected: []byte{0x3a, 0x7F, 0xFF, 0xFF, 0xFF},
		},
	}

	for testNumber, test := range tests {

		actual, err := int32s.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("VALUE: %d", test.Value)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("VALUE:    %d", test.Value)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				continue
			}
		}
	}
}
