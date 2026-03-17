package base64s_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/tags/base64s"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    []byte
		Expected []byte
	}{
		// Empty input produces empty base64 string.
		{
			// 34("")
			Value:    []byte{},
			Expected: []byte{
				0xd8, 0x22, // tag 34
				0x60,       // text string, length 0
			},
		},



		// Single byte — 1 byte produces 4-character base64 with 2 padding chars.
		{
			// 34("AQ==")
			Value:    []byte{0x01},
			Expected: append(
				[]byte{0xd8, 0x22}, // tag 34
				append(
					[]byte{0x64}, // text string, length 4
					[]byte("AQ==")...,
				)...,
			),
		},



		// Two bytes — produces 4-character base64 with 1 padding char.
		{
			// 34("AQI=")
			Value:    []byte{0x01, 0x02},
			Expected: append(
				[]byte{0xd8, 0x22}, // tag 34
				append(
					[]byte{0x64}, // text string, length 4
					[]byte("AQI=")...,
				)...,
			),
		},



		// Three bytes — exact multiple of 3, no padding needed.
		{
			// 34("AQID")
			Value:    []byte{0x01, 0x02, 0x03},
			Expected: append(
				[]byte{0xd8, 0x22}, // tag 34
				append(
					[]byte{0x64}, // text string, length 4
					[]byte("AQID")...,
				)...,
			),
		},



		// Four bytes — produces 8-character base64 with 2 padding chars.
		{
			// 34("AQIDBA==")
			Value:    []byte{0x01, 0x02, 0x03, 0x04},
			Expected: append(
				[]byte{0xd8, 0x22}, // tag 34
				append(
					[]byte{0x68}, // text string, length 8
					[]byte("AQIDBA==")...,
				)...,
			),
		},



		// Standard alphabet: bytes that produce + and / (not - and _).
		{
			// {0xFB, 0xEF, 0xFF} → standard base64 "++//"
			Value:    []byte{0xFB, 0xEF, 0xFF},
			Expected: append(
				[]byte{0xd8, 0x22}, // tag 34
				append(
					[]byte{0x64}, // text string, length 4
					[]byte("++//")...,
				)...,
			),
		},



		// Longer input (16 bytes).
		{
			// 34("AQIDBAUGBwgJCgsMDQ4P")
			Value:    []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
			Expected: append(
				[]byte{0xd8, 0x22}, // tag 34
				append(
					[]byte{0x74}, // text string, length 20
					[]byte("AQIDBAUGBwgJCgsMDQ4P")...,
				)...,
			),
		},
	}

	for testNumber, test := range tests {

		actual, err := base64s.Marshal(test.Value)

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
