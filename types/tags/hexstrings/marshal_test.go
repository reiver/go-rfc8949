package hexstrings_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/tags/hexstrings"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    []byte
		Expected []byte
	}{
		// Empty byte string.
		{
			// 263(h'')
			Value:    []byte{},
			Expected: []byte{
				0xd9, 0x01, 0x07, // tag 263
				0x40,             // byte string, length 0
			},
		},



		// Single byte.
		{
			// 263(h'FF')
			Value:    []byte{0xFF},
			Expected: []byte{
				0xd9, 0x01, 0x07, // tag 263
				0x41,             // byte string, length 1
				0xFF,
			},
		},



		// Classic hex example.
		{
			// 263(h'DEADBEEF')
			Value:    []byte{0xDE, 0xAD, 0xBE, 0xEF},
			Expected: []byte{
				0xd9, 0x01, 0x07, // tag 263
				0x44,             // byte string, length 4
				0xDE, 0xAD, 0xBE, 0xEF,
			},
		},



		// SHA-256 hash (32 bytes).
		{
			// 263(h'E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855')
			// SHA-256 of the empty string.
			Value: []byte{
				0xE3, 0xB0, 0xC4, 0x42, 0x98, 0xFC, 0x1C, 0x14,
				0x9A, 0xFB, 0xF4, 0xC8, 0x99, 0x6F, 0xB9, 0x24,
				0x27, 0xAE, 0x41, 0xE4, 0x64, 0x9B, 0x93, 0x4C,
				0xA4, 0x95, 0x99, 0x1B, 0x78, 0x52, 0xB8, 0x55,
			},
			Expected: append(
				[]byte{
					0xd9, 0x01, 0x07, // tag 263
					0x58, 0x20,       // byte string, length 32 (1-byte length)
				},
				[]byte{
					0xE3, 0xB0, 0xC4, 0x42, 0x98, 0xFC, 0x1C, 0x14,
					0x9A, 0xFB, 0xF4, 0xC8, 0x99, 0x6F, 0xB9, 0x24,
					0x27, 0xAE, 0x41, 0xE4, 0x64, 0x9B, 0x93, 0x4C,
					0xA4, 0x95, 0x99, 0x1B, 0x78, 0x52, 0xB8, 0x55,
				}...,
			),
		},



		// MAC address (6 bytes).
		{
			// 263(h'001122334455')
			Value:    []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55},
			Expected: []byte{
				0xd9, 0x01, 0x07, // tag 263
				0x46,             // byte string, length 6
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55,
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := hexstrings.Marshal(test.Value)

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
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}
