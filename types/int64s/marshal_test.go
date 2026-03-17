package int64s_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/int64s"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value int64
		Expected []byte
	}{

		// Positive values — delegated to uint64s.Marshal

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

		// 65536–4294967295: 5 bytes (initial byte 0x1a + 4 value bytes)
		{
			Value:    65536,
			Expected: []byte{0x1a, 0x00, 0x01, 0x00, 0x00},
		},
		{
			Value:    4294967295,
			Expected: []byte{0x1a, 0xFF, 0xFF, 0xFF, 0xFF},
		},

		// >= 4294967296: 9 bytes (initial byte 0x1b + 8 value bytes)
		{
			Value:    4294967296,
			Expected: []byte{0x1b, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
		},
		{
			Value:    1099511627775,
			Expected: []byte{0x1b, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		},
		{
			Value:    9223372036854775807, // math.MaxInt64
			Expected: []byte{0x1b, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		},



		// Negative values delegated to int32s.Marshal (value >= -2147483648)

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

		// -257 to -65536: 3 bytes (initial byte 0x39 + 2 argument bytes)
		{
			Value:    -257,
			Expected: []byte{0x39, 0x01, 0x00},
		},
		{
			Value:    -1000,
			Expected: []byte{0x39, 0x03, 0xE7},
		},
		{
			Value:    -65536,
			Expected: []byte{0x39, 0xFF, 0xFF},
		},

		// -65537 to -2147483648: 5 bytes (initial byte 0x3a + 4 argument bytes)
		{
			Value:    -65537,
			Expected: []byte{0x3a, 0x00, 0x01, 0x00, 0x00},
		},
		{
			Value:    -1000000,
			Expected: []byte{0x3a, 0x00, 0x0F, 0x42, 0x3F},
		},

		// Boundary: -2147483648 (minimum int32, still delegated to int32s)
		{
			Value:    -2147483648,
			Expected: []byte{0x3a, 0x7F, 0xFF, 0xFF, 0xFF},
		},



		// Values handled directly by int64s: -2147483649 to -4294967296
		// 5 bytes (initial byte 0x3a + 4 argument bytes)

		// Just past the int32s delegation boundary
		{
			Value:    -2147483649,
			Expected: []byte{0x3a, 0x80, 0x00, 0x00, 0x00},
		},
		{
			Value:    -3000000000,
			Expected: []byte{0x3a, 0xB2, 0xD0, 0x5D, 0xFF},
		},
		{
			Value:    -4294967295,
			Expected: []byte{0x3a, 0xFF, 0xFF, 0xFF, 0xFE},
		},
		{
			Value:    -4294967296,
			Expected: []byte{0x3a, 0xFF, 0xFF, 0xFF, 0xFF},
		},



		// Values handled directly by int64s: below -4294967296
		// 9 bytes (initial byte 0x3b + 8 argument bytes)

		{
			Value:    -4294967297,
			Expected: []byte{0x3b, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
		},
		{
			Value:    -1099511627776,
			Expected: []byte{0x3b, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		},
		{
			Value:    -1000000000000,
			Expected: []byte{0x3b, 0x00, 0x00, 0x00, 0xE8, 0xD4, 0xA5, 0x0F, 0xFF},
		},
		{
			Value:    -9223372036854775808, // math.MinInt64
			Expected: []byte{0x3b, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		},
	}

	for testNumber, test := range tests {

		actual, err := int64s.Marshal(test.Value)

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
