package bytestrings_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/bytestrings"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected []byte
	}{
		// RFC 8949 Appendix A examples.
		{
			// h'' (empty byte string)
			Value:    []byte{},
			Expected: []byte{0x40},
		},
		{
			// nil treated as empty byte string
			Value:    nil,
			Expected: []byte{0x40},
		},
		{
			// h'01020304'
			Value:    []byte{0x01, 0x02, 0x03, 0x04},
			Expected: []byte{0x44, 0x01, 0x02, 0x03, 0x04},
		},



		// Inline lengths (0-23).
		{
			Value:             []byte{},
			Expected: append([]byte{0x40}),
		},
		{
			Value:             []byte{0xAA},
			Expected: append([]byte{0x41}, 0xAA),
		},
		{
			Value:             []byte{0xAA, 0xBB},
			Expected: append([]byte{0x42}, 0xAA, 0xBB),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 3),
			Expected: append([]byte{0x43}, bytes.Repeat([]byte{0xFF}, 3)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 4),
			Expected: append([]byte{0x44}, bytes.Repeat([]byte{0xFF}, 4)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 5),
			Expected: append([]byte{0x45}, bytes.Repeat([]byte{0xFF}, 5)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 6),
			Expected: append([]byte{0x46}, bytes.Repeat([]byte{0xFF}, 6)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 7),
			Expected: append([]byte{0x47}, bytes.Repeat([]byte{0xFF}, 7)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 8),
			Expected: append([]byte{0x48}, bytes.Repeat([]byte{0xFF}, 8)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 9),
			Expected: append([]byte{0x49}, bytes.Repeat([]byte{0xFF}, 9)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 10),
			Expected: append([]byte{0x4a}, bytes.Repeat([]byte{0xFF}, 10)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 11),
			Expected: append([]byte{0x4b}, bytes.Repeat([]byte{0xFF}, 11)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 12),
			Expected: append([]byte{0x4c}, bytes.Repeat([]byte{0xFF}, 12)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 13),
			Expected: append([]byte{0x4d}, bytes.Repeat([]byte{0xFF}, 13)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 14),
			Expected: append([]byte{0x4e}, bytes.Repeat([]byte{0xFF}, 14)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 15),
			Expected: append([]byte{0x4f}, bytes.Repeat([]byte{0xFF}, 15)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 16),
			Expected: append([]byte{0x50}, bytes.Repeat([]byte{0xFF}, 16)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 17),
			Expected: append([]byte{0x51}, bytes.Repeat([]byte{0xFF}, 17)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 18),
			Expected: append([]byte{0x52}, bytes.Repeat([]byte{0xFF}, 18)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 19),
			Expected: append([]byte{0x53}, bytes.Repeat([]byte{0xFF}, 19)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 20),
			Expected: append([]byte{0x54}, bytes.Repeat([]byte{0xFF}, 20)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 21),
			Expected: append([]byte{0x55}, bytes.Repeat([]byte{0xFF}, 21)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 22),
			Expected: append([]byte{0x56}, bytes.Repeat([]byte{0xFF}, 22)...),
		},
		{
			// length 23: last inline length
			Value:             bytes.Repeat([]byte{0xFF}, 23),
			Expected: append([]byte{0x57}, bytes.Repeat([]byte{0xFF}, 23)...),
		},



		// 1-byte length (24-255).
		{
			// length 24: first 1-byte length
			Value:             bytes.Repeat([]byte{0xFF}, 24),
			Expected: append([]byte{0x58, 0x18}, bytes.Repeat([]byte{0xFF}, 24)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 25),
			Expected: append([]byte{0x58, 0x19}, bytes.Repeat([]byte{0xFF}, 25)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 100),
			Expected: append([]byte{0x58, 0x64}, bytes.Repeat([]byte{0xFF}, 100)...),
		},
		{
			// length 255: last 1-byte length
			Value:             bytes.Repeat([]byte{0xFF}, 255),
			Expected: append([]byte{0x58, 0xFF}, bytes.Repeat([]byte{0xFF}, 255)...),
		},



		// 2-byte length (256-65535).
		{
			// length 256: first 2-byte length
			Value:             bytes.Repeat([]byte{0xFF}, 256),
			Expected: append([]byte{0x59, 0x01, 0x00}, bytes.Repeat([]byte{0xFF}, 256)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 1000),
			Expected: append([]byte{0x59, 0x03, 0xE8}, bytes.Repeat([]byte{0xFF}, 1000)...),
		},
		{
			// length 65535: last 2-byte length
			Value:             bytes.Repeat([]byte{0xFF}, 65535),
			Expected: append([]byte{0x59, 0xFF, 0xFF}, bytes.Repeat([]byte{0xFF}, 65535)...),
		},



		// 4-byte length (65536+).
		{
			// length 65536: first 4-byte length
			Value:             bytes.Repeat([]byte{0xFF}, 65536),
			Expected: append([]byte{0x5a, 0x00, 0x01, 0x00, 0x00}, bytes.Repeat([]byte{0xFF}, 65536)...),
		},
		{
			Value:             bytes.Repeat([]byte{0xFF}, 100000),
			Expected: append([]byte{0x5a, 0x00, 0x01, 0x86, 0xA0}, bytes.Repeat([]byte{0xFF}, 100000)...),
		},
	}

	for testNumber, test := range tests {

		actual, err := bytestrings.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE-LENGTH: %d", len(test.Value))
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				if len(expected) <= 32 {
					t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				} else {
					t.Logf("EXPECTED: (len=%d) [%#v ... %#v]", len(expected), expected[:8], expected[len(expected)-4:])
				}
				if len(actual) <= 32 {
					t.Logf("ACTUAL:   (len=%d) %#v", len(actual), actual)
				} else {
					t.Logf("ACTUAL:   (len=%d) [%#v ... %#v]", len(actual), actual[:8], actual[len(actual)-4:])
				}
				t.Logf("VALUE-LENGTH: %d", len(test.Value))
				continue
			}
		}
	}
}
