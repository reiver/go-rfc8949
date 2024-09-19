package undefined_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/xtype/undefined"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Expected []byte
	}{
		{
			Expected: []byte{0xf7},
		},
	}

	for testNumber, test := range tests {

		actual, err := undefined.Marshal()

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
