package tags_test

import (
	"bytes"
	"testing"

	"github.com/reiver/go-rfc8949/types/tags"
)

func TestUnmarshalHead(t *testing.T) {

	tests := []struct{
		Data []byte
		ExpectedTagNumber uint64
		ExpectedRest []byte
		ExpectedError error
	}{
		// RFC 8949 Appendix A examples.
		{
			// 1(1363896240)
			Data:              []byte{0xc1, 0x1a, 0x51, 0x4b, 0x67, 0xb0},
			ExpectedTagNumber: 1,
			ExpectedRest:      []byte{0x1a, 0x51, 0x4b, 0x67, 0xb0},
		},
		{
			// 32("http://www.example.com") (only checking the head)
			Data:              []byte{0xd8, 0x20, 0x76},
			ExpectedTagNumber: 32,
			ExpectedRest:      []byte{0x76},
		},



		{
			// 42(h'00...') — the head of a CID tag
			Data:              []byte{0xd8, 0x2a, 0x58, 0x25, 0x00},
			ExpectedTagNumber: 42,
			ExpectedRest:      []byte{0x58, 0x25, 0x00},
		},



		{
			Data:          []byte{0x80},
			ExpectedRest:  []byte{0x80},
			ExpectedError: tags.ErrNotTagged,
		},
	}

	for testNumber, test := range tests {

		tagNumber, rest, err := tags.UnmarshalHead(test.Data)

		if test.ExpectedError != err {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", test.ExpectedError)
			t.Logf("ACTUAL:   %v", err)
			continue
		}

		if test.ExpectedTagNumber != tagNumber {
			t.Errorf("For test #%d, the actual tag number is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", test.ExpectedTagNumber)
			t.Logf("ACTUAL:   %d", tagNumber)
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
