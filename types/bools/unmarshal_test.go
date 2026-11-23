package bools_test

import (
	"testing"

	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/types/bools"
)

func TestUnmarshal(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected bool
		ExpectedError error
	}{
		{
			Data:          nil,
			ExpectedError: head.ErrEmpty,
		},



		// RFC 8949 Appendix A examples.
		{
			Data:     []byte{0xf4},
			Expected: false,
		},
		{
			Data:     []byte{0xf5},
			Expected: true,
		},



		{
			// null
			Data:          []byte{0xf6},
			ExpectedError: bools.ErrNotBool,
		},
		{
			// simple(20), but not in preferred serialization
			Data:          []byte{0xf8, 0x14},
			ExpectedError: bools.ErrNotBool,
		},
	}

	for testNumber, test := range tests {

		actual, _, err := bools.Unmarshal(test.Data)

		if test.ExpectedError != err {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", test.ExpectedError)
			t.Logf("ACTUAL:   %v", err)
			continue
		}

		if test.Expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", test.Expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}
