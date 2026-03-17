package embeddedjson_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/tags/embeddedjson"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    []byte
		Expected []byte
	}{
		// Empty JSON object.
		{
			// 262(h'7B7D') — "{}"
			Value:    []byte("{}"),
			Expected: []byte{
				0xd9, 0x01, 0x06, // tag 262
				0x42,             // byte string, length 2
				0x7B, 0x7D,       // "{}"
			},
		},



		// Empty JSON array.
		{
			// 262(h'5B5D') — "[]"
			Value:    []byte("[]"),
			Expected: []byte{
				0xd9, 0x01, 0x06, // tag 262
				0x42,             // byte string, length 2
				0x5B, 0x5D,       // "[]"
			},
		},



		// Simple JSON object.
		{
			// 262({"name":"Alice"})
			Value:    []byte(`{"name":"Alice"}`),
			Expected: append(
				[]byte{
					0xd9, 0x01, 0x06, // tag 262
					0x50,             // byte string, length 16
				},
				[]byte(`{"name":"Alice"}`)...,
			),
		},



		// JSON with nested structure.
		{
			// 262({"a":1,"b":[2,3]})
			Value:    []byte(`{"a":1,"b":[2,3]}`),
			Expected: append(
				[]byte{
					0xd9, 0x01, 0x06, // tag 262
					0x51,             // byte string, length 17
				},
				[]byte(`{"a":1,"b":[2,3]}`)...,
			),
		},



		// JSON with Unicode (😀 is 4 UTF-8 bytes, so total is 16 bytes).
		{
			// 262({"emoji":"😀"})
			Value:    []byte(`{"emoji":"😀"}`),
			Expected: append(
				[]byte{
					0xd9, 0x01, 0x06, // tag 262
					0x50,             // byte string, length 16
				},
				[]byte(`{"emoji":"😀"}`)...,
			),
		},



		// JSON null literal.
		{
			// 262("null")
			Value:    []byte("null"),
			Expected: []byte{
				0xd9, 0x01, 0x06, // tag 262
				0x44,             // byte string, length 4
				0x6E, 0x75, 0x6C, 0x6C, // "null"
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := embeddedjson.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %s", test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				t.Logf("VALUE: %s", test.Value)
				continue
			}
		}
	}
}
