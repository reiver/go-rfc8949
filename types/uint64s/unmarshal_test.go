package uint64s_test

import (
	"bytes"
	"testing"

	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/types/uint64s"
)

func TestUnmarshal(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected uint64
		ExpectedRest []byte
	}{
		// RFC 8949 Appendix A examples.
		{
			Data:     []byte{0x00},
			Expected: 0,
		},
		{
			Data:     []byte{0x01},
			Expected: 1,
		},
		{
			Data:     []byte{0x0a},
			Expected: 10,
		},
		{
			Data:     []byte{0x17},
			Expected: 23,
		},
		{
			Data:     []byte{0x18, 0x18},
			Expected: 24,
		},
		{
			Data:     []byte{0x18, 0x19},
			Expected: 25,
		},
		{
			Data:     []byte{0x18, 0x64},
			Expected: 100,
		},
		{
			Data:     []byte{0x19, 0x03, 0xe8},
			Expected: 1000,
		},
		{
			Data:     []byte{0x1a, 0x00, 0x0f, 0x42, 0x40},
			Expected: 1000000,
		},
		{
			Data:     []byte{0x1b, 0x00, 0x00, 0x00, 0xe8, 0xd4, 0xa5, 0x10, 0x00},
			Expected: 1000000000000,
		},
		{
			Data:     []byte{0x1b, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			Expected: 18446744073709551615,
		},



		// The rest is returned.
		{
			Data:         []byte{0x18, 0x64, 0xf6, 0x01},
			Expected:     100,
			ExpectedRest: []byte{0xf6, 0x01},
		},



		// Non-minimal encodings are accepted.
		{
			Data:     []byte{0x19, 0x00, 0x01},
			Expected: 1,
		},
	}

	for testNumber, test := range tests {

		actual, rest, err := uint64s.Unmarshal(test.Data)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA: %#v", test.Data)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("DATA: %#v", test.Data)
				continue
			}
		}

		{
			expected := test.ExpectedRest

			if !bytes.Equal(expected, rest) {
				t.Errorf("For test #%d, the actual rest is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", rest)
				t.Logf("DATA: %#v", test.Data)
				continue
			}
		}
	}
}

func TestUnmarshal_errors(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected error
	}{
		{
			Data:     nil,
			Expected: head.ErrEmpty,
		},
		{
			Data:     []byte{0x19, 0x03},
			Expected: head.ErrUnexpectedEnd,
		},
		{
			// -1
			Data:     []byte{0x20},
			Expected: uint64s.ErrNotUnsignedInteger,
		},
		{
			// ""
			Data:     []byte{0x60},
			Expected: uint64s.ErrNotUnsignedInteger,
		},
		{
			// indefinite (not well-formed for major type 0)
			Data:     []byte{0x1f},
			Expected: uint64s.ErrNotUnsignedInteger,
		},
	}

	for testNumber, test := range tests {

		_, rest, err := uint64s.Unmarshal(test.Data)

		{
			expected := test.Expected
			actual   := err

			if expected != actual {
				t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED: %v", expected)
				t.Logf("ACTUAL:   %v", actual)
				t.Logf("DATA: %#v", test.Data)
				continue
			}
		}

		if !bytes.Equal(test.Data, rest) {
			t.Errorf("For test #%d, on error, the rest should be the original data.", testNumber)
			t.Logf("EXPECTED: %#v", test.Data)
			t.Logf("ACTUAL:   %#v", rest)
			continue
		}
	}
}

func TestUnmarshal_roundTrip(t *testing.T) {

	values := []uint64{
		0, 1, 23, 24, 255, 256, 65535, 65536, 4294967295, 4294967296, 18446744073709551615,
	}

	for testNumber, value := range values {

		p, err := uint64s.Marshal(value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error from Marshal but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		actual, rest, err := uint64s.Unmarshal(p)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error from Unmarshal but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		if value != actual || 0 != len(rest) {
			t.Errorf("For test #%d, the round-trip did not work.", testNumber)
			t.Logf("EXPECTED: %d", value)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("REST: %#v", rest)
			continue
		}
	}
}
