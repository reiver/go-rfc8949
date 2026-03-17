package additionalinfo_test

import (
	"testing"

	"github.com/reiver/go-rfc8949/initialbyte/additionalinfo"
)

func TestConstants(t *testing.T) {

	tests := []struct{
		Name string
		Actual byte
		Expected byte
	}{
		{
			Name:     "Uint8",
			Actual:   additionalinfo.Uint8,
			Expected: 24,
		},
		{
			Name:     "Uint16",
			Actual:   additionalinfo.Uint16,
			Expected: 25,
		},
		{
			Name:     "Uint32",
			Actual:   additionalinfo.Uint32,
			Expected: 26,
		},
		{
			Name:     "Uint64",
			Actual:   additionalinfo.Uint64,
			Expected: 27,
		},
	}

	for testNumber, test := range tests {

		actual   := test.Actual
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value of %s is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED: %d (0x%02x)", expected, expected)
			t.Logf("ACTUAL:   %d (0x%02x)", actual,   actual)
			continue
		}
	}
}
