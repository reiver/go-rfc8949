package mimemessages_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/tags/mimemessages"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    string
		Expected []byte
	}{
		// Minimal MIME message: single header + body.
		{
			// 36("Content-Type: text/plain\r\n\r\nHello")
			Value:    "Content-Type: text/plain\r\n\r\nHello",
			Expected: append(
				[]byte{
					0xd8, 0x24,       // tag 36
					0x78, 0x21,       // text string, length 33 (1-byte length)
				},
				[]byte("Content-Type: text/plain\r\n\r\nHello")...,
			),
		},



		// Multiple headers.
		{
			// 36("Content-Type: text/plain\r\nSubject: Test\r\n\r\nBody here.")
			Value:    "Content-Type: text/plain\r\nSubject: Test\r\n\r\nBody here.",
			Expected: append(
				[]byte{
					0xd8, 0x24,       // tag 36
					0x78, 0x35,       // text string, length 53 (1-byte length)
				},
				[]byte("Content-Type: text/plain\r\nSubject: Test\r\n\r\nBody here.")...,
			),
		},



		// MIME message with charset parameter.
		{
			// 36("Content-Type: text/plain; charset=utf-8\r\n\r\nHello, world!")
			Value:    "Content-Type: text/plain; charset=utf-8\r\n\r\nHello, world!",
			Expected: append(
				[]byte{
					0xd8, 0x24,       // tag 36
					0x78, 0x38,       // text string, length 56 (1-byte length)
				},
				[]byte("Content-Type: text/plain; charset=utf-8\r\n\r\nHello, world!")...,
			),
		},



		// Empty body.
		{
			// 36("Content-Type: text/plain\r\n\r\n")
			Value:    "Content-Type: text/plain\r\n\r\n",
			Expected: append(
				[]byte{
					0xd8, 0x24,       // tag 36
					0x78, 0x1c,       // text string, length 28 (1-byte length)
				},
				[]byte("Content-Type: text/plain\r\n\r\n")...,
			),
		},



		// Multipart MIME message.
		{
			Value: "MIME-Version: 1.0\r\n" +
				"Content-Type: multipart/mixed; boundary=frontier\r\n" +
				"\r\n" +
				"--frontier\r\n" +
				"Content-Type: text/plain\r\n" +
				"\r\n" +
				"Hello\r\n" +
				"--frontier--\r\n",
			Expected: append(
				[]byte{
					0xd8, 0x24,       // tag 36
					0x78, 0x84,       // text string, length 132 (1-byte length)
				},
				[]byte(
					"MIME-Version: 1.0\r\n"+
						"Content-Type: multipart/mixed; boundary=frontier\r\n"+
						"\r\n"+
						"--frontier\r\n"+
						"Content-Type: text/plain\r\n"+
						"\r\n"+
						"Hello\r\n"+
						"--frontier--\r\n",
				)...,
			),
		},
	}

	for testNumber, test := range tests {

		actual, err := mimemessages.Marshal(test.Value)

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
