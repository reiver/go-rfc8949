package nils_test

import (
	"bytes"
	"testing"

	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/types/nils"
)

func TestUnmarshal(t *testing.T) {

	tests := []struct{
		Data []byte
		ExpectedRest []byte
		ExpectedError error
	}{
		{
			Data:          nil,
			ExpectedRest:  nil,
			ExpectedError: head.ErrEmpty,
		},



		{
			// null (RFC 8949 Appendix A)
			Data:         []byte{0xf6},
			ExpectedRest: []byte{},
		},
		{
			Data:         []byte{0xf6, 0x01},
			ExpectedRest: []byte{0x01},
		},



		{
			// undefined
			Data:          []byte{0xf7},
			ExpectedRest:  []byte{0xf7},
			ExpectedError: nils.ErrNotNull,
		},
	}

	for testNumber, test := range tests {

		rest, err := nils.Unmarshal(test.Data)

		if test.ExpectedError != err {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", test.ExpectedError)
			t.Logf("ACTUAL:   %v", err)
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
