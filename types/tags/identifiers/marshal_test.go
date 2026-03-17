package identifiers_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/tags/identifiers"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    any
		Expected []byte
	}{
		// Small unsigned integer identifier.
		{
			// 39(0)
			Value:    uint64(0),
			Expected: []byte{
				0xd8, 0x27, // tag 39
				0x00,       // unsigned integer 0
			},
		},



		// Inline unsigned integer identifier.
		{
			// 39(23)
			Value:    uint64(23),
			Expected: []byte{
				0xd8, 0x27, // tag 39
				0x17,       // unsigned integer 23
			},
		},



		// 1-byte unsigned integer identifier.
		{
			// 39(255)
			Value:    uint64(255),
			Expected: []byte{
				0xd8, 0x27, // tag 39
				0x18, 0xff, // unsigned integer 255
			},
		},



		// Large unsigned integer identifier (JavaScript-unsafe range).
		{
			// 39(9007199254740993) — 2^53 + 1, beyond JS Number.MAX_SAFE_INTEGER
			Value:    uint64(9007199254740993),
			Expected: []byte{
				0xd8, 0x27, // tag 39
				0x1b, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // uint64
			},
		},



		// String identifier.
		{
			// 39("user-42")
			Value:    "user-42",
			Expected: append(
				[]byte{0xd8, 0x27}, // tag 39
				append(
					[]byte{0x67}, // text string, length 7
					[]byte("user-42")...,
				)...,
			),
		},



		// Byte string identifier.
		{
			// 39(h'DEADBEEF')
			Value:    []byte{0xDE, 0xAD, 0xBE, 0xEF},
			Expected: []byte{
				0xd8, 0x27, // tag 39
				0x44,       // byte string, length 4
				0xDE, 0xAD, 0xBE, 0xEF,
			},
		},



		// Boolean identifier (edge case — tag 39 accepts any type).
		{
			// 39(true)
			Value:    true,
			Expected: []byte{
				0xd8, 0x27, // tag 39
				0xf5,       // true
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := identifiers.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
				continue
			}
		}
	}
}
