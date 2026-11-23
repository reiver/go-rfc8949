package bytestrings_test

import (
	"bytes"
	"testing"

	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/types/bytestrings"
)

func TestUnmarshal(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected []byte
		ExpectedRest []byte
	}{
		// RFC 8949 Appendix A examples.
		{
			Data:     []byte{0x40},
			Expected: []byte{},
		},
		{
			Data:     []byte{0x44, 0x01, 0x02, 0x03, 0x04},
			Expected: []byte{0x01, 0x02, 0x03, 0x04},
		},



		{
			Data:         []byte{0x41, 0xaa, 0xf5},
			Expected:     []byte{0xaa},
			ExpectedRest: []byte{0xf5},
		},
	}

	for testNumber, test := range tests {

		actual, rest, err := bytestrings.Unmarshal(test.Data)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		if !bytes.Equal(test.Expected, actual) {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}

		if !bytes.Equal(test.ExpectedRest, rest) {
			t.Errorf("For test #%d, the actual rest is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", test.ExpectedRest)
			t.Logf("ACTUAL:   %#v", rest)
			continue
		}
	}
}

func TestUnmarshal_copies(t *testing.T) {
	var data []byte = []byte{0x42, 0x01, 0x02}

	value, _, err := bytestrings.Unmarshal(data)
	if nil != err {
		t.Fatalf("Did not expect an error but actually got one: (%T) %s", err, err)
	}

	data[1] = 0xff

	if 0x01 != value[0] {
		t.Errorf("The returned value should not share memory with the input.")
		t.Logf("VALUE: %#v", value)
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
			Data:     []byte{0x44, 0x01},
			Expected: head.ErrUnexpectedEnd,
		},
		{
			Data:     []byte{0x60},
			Expected: bytestrings.ErrNotByteString,
		},
		{
			// (_ h'0102', h'030405')
			Data:     []byte{0x5f, 0x42, 0x01, 0x02, 0x43, 0x03, 0x04, 0x05, 0xff},
			Expected: bytestrings.ErrIndefiniteLength,
		},
	}

	for testNumber, test := range tests {

		_, _, err := bytestrings.Unmarshal(test.Data)

		if test.Expected != err {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", test.Expected)
			t.Logf("ACTUAL:   %v", err)
			t.Logf("DATA: %#v", test.Data)
			continue
		}
	}
}
