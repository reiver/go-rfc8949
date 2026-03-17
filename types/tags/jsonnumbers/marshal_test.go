package jsonnumbers_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/tags/jsonnumbers"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    string
		Expected []byte
	}{
		// Zero.
		{
			// 284("0")
			Value:    "0",
			Expected: []byte{
				0xd9, 0x01, 0x1c, // tag 284
				0x61,             // text string, length 1
				0x30,             // "0"
			},
		},



		// Small positive integer.
		{
			// 284("123")
			Value:    "123",
			Expected: append(
				[]byte{0xd9, 0x01, 0x1c}, // tag 284
				append(
					[]byte{0x63}, // text string, length 3
					[]byte("123")...,
				)...,
			),
		},



		// Negative integer.
		{
			// 284("-45")
			Value:    "-45",
			Expected: append(
				[]byte{0xd9, 0x01, 0x1c}, // tag 284
				append(
					[]byte{0x63}, // text string, length 3
					[]byte("-45")...,
				)...,
			),
		},



		// Fractional number.
		{
			// 284("-45.67")
			Value:    "-45.67",
			Expected: append(
				[]byte{0xd9, 0x01, 0x1c}, // tag 284
				append(
					[]byte{0x66}, // text string, length 6
					[]byte("-45.67")...,
				)...,
			),
		},



		// Scientific notation.
		{
			// 284("1.23E+10")
			Value:    "1.23E+10",
			Expected: append(
				[]byte{0xd9, 0x01, 0x1c}, // tag 284
				append(
					[]byte{0x68}, // text string, length 8
					[]byte("1.23E+10")...,
				)...,
			),
		},



		// Large integer beyond JavaScript Number.MAX_SAFE_INTEGER.
		{
			// 284("99999999999999999999")
			Value:    "99999999999999999999",
			Expected: append(
				[]byte{0xd9, 0x01, 0x1c}, // tag 284
				append(
					[]byte{0x74}, // text string, length 20
					[]byte("99999999999999999999")...,
				)...,
			),
		},



		// Arbitrary-precision decimal.
		{
			// 284("99999999999999999999.123456789")
			Value:    "99999999999999999999.123456789",
			Expected: append(
				[]byte{
					0xd9, 0x01, 0x1c, // tag 284
					0x78, 0x1e,       // text string, length 30 (1-byte length)
				},
				[]byte("99999999999999999999.123456789")...,
			),
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonnumbers.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
