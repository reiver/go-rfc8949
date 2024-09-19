package bools_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/bools"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value bool
		Expected []byte
	}{
		{
			Value: false,
			Expected: []byte{0xf4},
		},
		{
			Value: true,
			Expected: []byte{0xf5},
		},
	}

	for testNumber, test := range tests {

		actual, err := bools.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				continue
			}
		}
	}
}
