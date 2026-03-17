package textstrings_test

import (
	"testing"

	"bytes"
	"strings"

	"github.com/reiver/go-rfc8949/types/textstrings"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value string
		Expected []byte
	}{
		// RFC 8949 Appendix A examples.
		{
			// "" (empty text string)
			Value:    "",
			Expected: []byte{0x60},
		},
		{
			// "a"
			Value:    "a",
			Expected: []byte{0x61, 0x61},
		},
		{
			// "IETF"
			Value:    "IETF",
			Expected: []byte{0x64, 0x49, 0x45, 0x54, 0x46},
		},
		{
			// "\"\\"
			Value:    "\"\\",
			Expected: []byte{0x62, 0x22, 0x5c},
		},
		{
			// U+00FC (ü) — 2 UTF-8 bytes
			Value:    "\u00fc",
			Expected: []byte{0x62, 0xc3, 0xbc},
		},
		{
			// U+6C34 (水) — 3 UTF-8 bytes
			Value:    "\u6c34",
			Expected: []byte{0x63, 0xe6, 0xb0, 0xb4},
		},
		{
			// U+10151 (𐅑) — 4 UTF-8 bytes
			Value:    "\U00010151",
			Expected: []byte{0x64, 0xf0, 0x90, 0x85, 0x91},
		},



		// Inline lengths (0-23 bytes).
		{
			Value:             "",
			Expected: append([]byte{0x60}),
		},
		{
			Value:             "a",
			Expected: append([]byte{0x61}, 'a'),
		},
		{
			Value:             "ab",
			Expected: append([]byte{0x62}, "ab"...),
		},
		{
			Value:             "abc",
			Expected: append([]byte{0x63}, "abc"...),
		},
		{
			Value:             strings.Repeat("x", 4),
			Expected: append([]byte{0x64}, strings.Repeat("x", 4)...),
		},
		{
			Value:             strings.Repeat("x", 5),
			Expected: append([]byte{0x65}, strings.Repeat("x", 5)...),
		},
		{
			Value:             strings.Repeat("x", 6),
			Expected: append([]byte{0x66}, strings.Repeat("x", 6)...),
		},
		{
			Value:             strings.Repeat("x", 7),
			Expected: append([]byte{0x67}, strings.Repeat("x", 7)...),
		},
		{
			Value:             strings.Repeat("x", 8),
			Expected: append([]byte{0x68}, strings.Repeat("x", 8)...),
		},
		{
			Value:             strings.Repeat("x", 9),
			Expected: append([]byte{0x69}, strings.Repeat("x", 9)...),
		},
		{
			Value:             strings.Repeat("x", 10),
			Expected: append([]byte{0x6a}, strings.Repeat("x", 10)...),
		},
		{
			Value:             strings.Repeat("x", 11),
			Expected: append([]byte{0x6b}, strings.Repeat("x", 11)...),
		},
		{
			Value:             strings.Repeat("x", 12),
			Expected: append([]byte{0x6c}, strings.Repeat("x", 12)...),
		},
		{
			Value:             strings.Repeat("x", 13),
			Expected: append([]byte{0x6d}, strings.Repeat("x", 13)...),
		},
		{
			Value:             strings.Repeat("x", 14),
			Expected: append([]byte{0x6e}, strings.Repeat("x", 14)...),
		},
		{
			Value:             strings.Repeat("x", 15),
			Expected: append([]byte{0x6f}, strings.Repeat("x", 15)...),
		},
		{
			Value:             strings.Repeat("x", 16),
			Expected: append([]byte{0x70}, strings.Repeat("x", 16)...),
		},
		{
			Value:             strings.Repeat("x", 17),
			Expected: append([]byte{0x71}, strings.Repeat("x", 17)...),
		},
		{
			Value:             strings.Repeat("x", 18),
			Expected: append([]byte{0x72}, strings.Repeat("x", 18)...),
		},
		{
			Value:             strings.Repeat("x", 19),
			Expected: append([]byte{0x73}, strings.Repeat("x", 19)...),
		},
		{
			Value:             strings.Repeat("x", 20),
			Expected: append([]byte{0x74}, strings.Repeat("x", 20)...),
		},
		{
			Value:             strings.Repeat("x", 21),
			Expected: append([]byte{0x75}, strings.Repeat("x", 21)...),
		},
		{
			Value:             strings.Repeat("x", 22),
			Expected: append([]byte{0x76}, strings.Repeat("x", 22)...),
		},
		{
			// length 23: last inline length
			Value:             strings.Repeat("x", 23),
			Expected: append([]byte{0x77}, strings.Repeat("x", 23)...),
		},



		// 1-byte length (24-255 bytes).
		{
			// length 24: first 1-byte length
			Value:             strings.Repeat("x", 24),
			Expected: append([]byte{0x78, 0x18}, strings.Repeat("x", 24)...),
		},
		{
			Value:             strings.Repeat("x", 25),
			Expected: append([]byte{0x78, 0x19}, strings.Repeat("x", 25)...),
		},
		{
			Value:             strings.Repeat("x", 100),
			Expected: append([]byte{0x78, 0x64}, strings.Repeat("x", 100)...),
		},
		{
			// length 255: last 1-byte length
			Value:             strings.Repeat("x", 255),
			Expected: append([]byte{0x78, 0xFF}, strings.Repeat("x", 255)...),
		},



		// 2-byte length (256-65535 bytes).
		{
			// length 256: first 2-byte length
			Value:             strings.Repeat("x", 256),
			Expected: append([]byte{0x79, 0x01, 0x00}, strings.Repeat("x", 256)...),
		},
		{
			Value:             strings.Repeat("x", 1000),
			Expected: append([]byte{0x79, 0x03, 0xE8}, strings.Repeat("x", 1000)...),
		},
		{
			// length 65535: last 2-byte length
			Value:             strings.Repeat("x", 65535),
			Expected: append([]byte{0x79, 0xFF, 0xFF}, strings.Repeat("x", 65535)...),
		},



		// 4-byte length (65536+ bytes).
		{
			// length 65536: first 4-byte length
			Value:             strings.Repeat("x", 65536),
			Expected: append([]byte{0x7a, 0x00, 0x01, 0x00, 0x00}, strings.Repeat("x", 65536)...),
		},
		{
			Value:             strings.Repeat("x", 100000),
			Expected: append([]byte{0x7a, 0x00, 0x01, 0x86, 0xA0}, strings.Repeat("x", 100000)...),
		},
	}

	for testNumber, test := range tests {

		actual, err := textstrings.Marshal(test.Value)

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

func TestMarshalInvalidUTF8(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: string([]byte{0xFF, 0xFE}),
		},
		{
			Value: string([]byte{0x80}),
		},
		{
			Value: string([]byte{0xC0, 0xAF}),
		},
		{
			Value: "hello" + string([]byte{0xFF}) + "world",
		},
	}

	for testNumber, test := range tests {

		_, err := textstrings.Marshal(test.Value)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
