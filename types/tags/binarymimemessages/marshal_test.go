package binarymimemessages_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/tags/binarymimemessages"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    []byte
		Expected []byte
	}{
		// Minimal MIME message: single header + body.
		{
			// 257(h'Content-Type: text/plain\r\n\r\nHello')
			Value:    []byte("Content-Type: text/plain\r\n\r\nHello"),
			Expected: append(
				[]byte{
					0xd9, 0x01, 0x01, // tag 257
					0x58, 0x21,       // byte string, length 33 (1-byte length)
				},
				[]byte("Content-Type: text/plain\r\n\r\nHello")...,
			),
		},



		// Multiple headers.
		{
			// 257(h'Content-Type: text/plain\r\nSubject: Test\r\n\r\nBody here.')
			Value:    []byte("Content-Type: text/plain\r\nSubject: Test\r\n\r\nBody here."),
			Expected: append(
				[]byte{
					0xd9, 0x01, 0x01, // tag 257
					0x58, 0x35,       // byte string, length 53 (1-byte length)
				},
				[]byte("Content-Type: text/plain\r\nSubject: Test\r\n\r\nBody here.")...,
			),
		},



		// Binary body content — the key advantage over tag 36.
		{
			Value: append(
				[]byte("Content-Type: application/octet-stream\r\nContent-Transfer-Encoding: binary\r\n\r\n"),
				[]byte{0x00, 0x01, 0xFF, 0xFE, 0x80, 0x90}...,
			),
			Expected: append(
				[]byte{
					0xd9, 0x01, 0x01, // tag 257
					0x58, 0x53,       // byte string, length 83 (1-byte length)
				},
				append(
					[]byte("Content-Type: application/octet-stream\r\nContent-Transfer-Encoding: binary\r\n\r\n"),
					[]byte{0x00, 0x01, 0xFF, 0xFE, 0x80, 0x90}...,
				)...,
			),
		},



		// Empty body.
		{
			Value:    []byte("Content-Type: text/plain\r\n\r\n"),
			Expected: append(
				[]byte{
					0xd9, 0x01, 0x01, // tag 257
					0x58, 0x1c,       // byte string, length 28 (1-byte length)
				},
				[]byte("Content-Type: text/plain\r\n\r\n")...,
			),
		},



		// Empty input.
		{
			Value:    []byte{},
			Expected: []byte{
				0xd9, 0x01, 0x01, // tag 257
				0x40,             // byte string, length 0
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := binarymimemessages.Marshal(test.Value)

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
