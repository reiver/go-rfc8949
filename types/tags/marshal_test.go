package tags_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/tags"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		TagNumber uint64
		Content   any
		Expected  []byte
	}{
		// RFC 8949 Appendix A examples.
		{
			// 0("2013-03-21T20:04:00Z")
			TagNumber: 0,
			Content:   "2013-03-21T20:04:00Z",
			Expected:  []byte{0xc0, 0x74, 0x32, 0x30, 0x31, 0x33, 0x2d, 0x30, 0x33, 0x2d, 0x32, 0x31, 0x54, 0x32, 0x30, 0x3a, 0x30, 0x34, 0x3a, 0x30, 0x30, 0x5a},
		},
		{
			// 1(1363896240)
			TagNumber: 1,
			Content:   uint64(1363896240),
			Expected:  []byte{0xc1, 0x1a, 0x51, 0x4b, 0x67, 0xb0},
		},
		{
			// 23(h'01020304')
			TagNumber: 23,
			Content:   []byte{0x01, 0x02, 0x03, 0x04},
			Expected:  []byte{0xd7, 0x44, 0x01, 0x02, 0x03, 0x04},
		},
		{
			// 24(h'6449455446')
			TagNumber: 24,
			Content:   []byte{0x64, 0x49, 0x45, 0x54, 0x46},
			Expected:  []byte{0xd8, 0x18, 0x45, 0x64, 0x49, 0x45, 0x54, 0x46},
		},
		{
			// 32("http://www.example.com")
			TagNumber: 32,
			Content:   "http://www.example.com",
			Expected:  []byte{0xd8, 0x20, 0x76, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d},
		},



		// Tag number boundary values.
		{
			// Tag 0 (inline): smallest tag number
			TagNumber: 0,
			Content:   uint64(0),
			Expected:  []byte{0xc0, 0x00},
		},
		{
			// Tag 23 (inline): largest inline tag number
			TagNumber: 23,
			Content:   uint64(0),
			Expected:  []byte{0xd7, 0x00},
		},
		{
			// Tag 24 (1-byte): first 1-byte tag number
			TagNumber: 24,
			Content:   uint64(0),
			Expected:  []byte{0xd8, 0x18, 0x00},
		},
		{
			// Tag 255 (1-byte): last 1-byte tag number
			TagNumber: 255,
			Content:   uint64(0),
			Expected:  []byte{0xd8, 0xff, 0x00},
		},
		{
			// Tag 256 (2-byte): first 2-byte tag number
			TagNumber: 256,
			Content:   uint64(0),
			Expected:  []byte{0xd9, 0x01, 0x00, 0x00},
		},
		{
			// Tag 65535 (2-byte): last 2-byte tag number
			TagNumber: 65535,
			Content:   uint64(0),
			Expected:  []byte{0xd9, 0xff, 0xff, 0x00},
		},
		{
			// Tag 65536 (4-byte): first 4-byte tag number
			TagNumber: 65536,
			Content:   uint64(0),
			Expected:  []byte{0xda, 0x00, 0x01, 0x00, 0x00, 0x00},
		},
		{
			// Tag 4294967295 (4-byte): last 4-byte tag number
			TagNumber: 4294967295,
			Content:   uint64(0),
			Expected:  []byte{0xda, 0xff, 0xff, 0xff, 0xff, 0x00},
		},
		{
			// Tag 4294967296 (8-byte): first 8-byte tag number
			TagNumber: 4294967296,
			Content:   uint64(0),
			Expected:  []byte{0xdb, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00},
		},



		// Tag 55799: self-described CBOR.
		{
			// 55799("CBOR")
			TagNumber: 55799,
			Content:   "CBOR",
			Expected:  []byte{0xd9, 0xd9, 0xf7, 0x64, 0x43, 0x42, 0x4f, 0x52},
		},



		// Nested tag.
		{
			// Tag wrapping nil.
			TagNumber: 1,
			Content:   nil,
			Expected:  []byte{0xc1, 0xf6},
		},
	}

	for testNumber, test := range tests {

		actual, err := tags.Marshal(test.TagNumber, test.Content)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("TAG-NUMBER: %d", test.TagNumber)
			t.Logf("CONTENT: %#v", test.Content)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				t.Logf("TAG-NUMBER: %d", test.TagNumber)
				t.Logf("CONTENT: %#v", test.Content)
				continue
			}
		}
	}
}

func TestMarshalUnsupportedContent(t *testing.T) {

	tests := []struct{
		TagNumber uint64
		Content   any
	}{
		{
			TagNumber: 0,
			Content:   float64(1.0),
		},
		{
			TagNumber: 0,
			Content:   struct{}{},
		},
	}

	for testNumber, test := range tests {

		_, err := tags.Marshal(test.TagNumber, test.Content)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("TAG-NUMBER: %d", test.TagNumber)
			t.Logf("CONTENT: %#v", test.Content)
			continue
		}
	}
}
