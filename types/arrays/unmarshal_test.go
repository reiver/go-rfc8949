package arrays_test

import (
	"bytes"
	"testing"

	"github.com/reiver/go-rfc8949/types/arrays"
)

func TestUnmarshalHead(t *testing.T) {

	tests := []struct{
		Data []byte
		ExpectedCount uint64
		ExpectedRest []byte
		ExpectedError error
	}{
		// RFC 8949 Appendix A examples.
		{
			// []
			Data:          []byte{0x80},
			ExpectedCount: 0,
			ExpectedRest:  []byte{},
		},
		{
			// [1, 2, 3]
			Data:          []byte{0x83, 0x01, 0x02, 0x03},
			ExpectedCount: 3,
			ExpectedRest:  []byte{0x01, 0x02, 0x03},
		},
		{
			// [1, 2, 3, ..., 25] (only checking the head)
			Data:          []byte{0x98, 0x19, 0x01},
			ExpectedCount: 25,
			ExpectedRest:  []byte{0x01},
		},



		{
			// {}
			Data:          []byte{0xa0},
			ExpectedRest:  []byte{0xa0},
			ExpectedError: arrays.ErrNotArray,
		},
		{
			// [_ ]
			Data:          []byte{0x9f, 0xff},
			ExpectedRest:  []byte{0x9f, 0xff},
			ExpectedError: arrays.ErrIndefiniteLength,
		},
	}

	for testNumber, test := range tests {

		count, rest, err := arrays.UnmarshalHead(test.Data)

		if test.ExpectedError != err {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", test.ExpectedError)
			t.Logf("ACTUAL:   %v", err)
			continue
		}

		if test.ExpectedCount != count {
			t.Errorf("For test #%d, the actual count is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", test.ExpectedCount)
			t.Logf("ACTUAL:   %d", count)
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
