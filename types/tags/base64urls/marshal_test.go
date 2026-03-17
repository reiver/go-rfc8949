package base64urls_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/tags/base64urls"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    []byte
		Expected []byte
	}{
		// Empty input produces empty base64url string.
		{
			// 33("")
			Value:    []byte{},
			Expected: []byte{
				0xd8, 0x21, // tag 33
				0x60,       // text string, length 0
			},
		},



		// Single byte.
		{
			// 33("AQ")
			Value:    []byte{0x01},
			Expected: append(
				[]byte{0xd8, 0x21}, // tag 33
				append(
					[]byte{0x62}, // text string, length 2
					[]byte("AQ")...,
				)...,
			),
		},



		// Four bytes — no padding in output.
		{
			// 33("AQIDBA")
			Value:    []byte{0x01, 0x02, 0x03, 0x04},
			Expected: append(
				[]byte{0xd8, 0x21}, // tag 33
				append(
					[]byte{0x66}, // text string, length 6
					[]byte("AQIDBA")...,
				)...,
			),
		},



		// URL-safe alphabet: bytes that produce + and / in standard base64
		// should produce - and _ in base64url.
		{
			// {0xFB, 0xEF, 0xFF} → standard base64 "++//" → base64url "--__"
			Value:    []byte{0xFB, 0xEF, 0xFF},
			Expected: append(
				[]byte{0xd8, 0x21}, // tag 33
				append(
					[]byte{0x64}, // text string, length 4
					[]byte("--__")...,
				)...,
			),
		},



		// Three bytes — exact multiple of 3, no padding needed even in standard base64.
		{
			// 33("AQID")
			Value:    []byte{0x01, 0x02, 0x03},
			Expected: append(
				[]byte{0xd8, 0x21}, // tag 33
				append(
					[]byte{0x64}, // text string, length 4
					[]byte("AQID")...,
				)...,
			),
		},



		// Longer input (16 bytes, like a UUID).
		{
			// 33("AQIDBAUGBwgJCgsMDQ4P")
			Value:    []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
			Expected: append(
				[]byte{0xd8, 0x21}, // tag 33
				append(
					[]byte{0x74}, // text string, length 20
					[]byte("AQIDBAUGBwgJCgsMDQ4P")...,
				)...,
			),
		},
	}

	for testNumber, test := range tests {

		actual, err := base64urls.Marshal(test.Value)

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
