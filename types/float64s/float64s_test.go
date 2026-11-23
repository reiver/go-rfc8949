package float64s_test

import (
	"bytes"
	"math"
	"testing"

	"github.com/reiver/go-rfc8949/types/float64s"
)

func TestUnmarshal(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected float64
		ExpectedBits int
	}{
		// RFC 8949 Appendix A examples.
		{
			// 0.0
			Data:         []byte{0xf9, 0x00, 0x00},
			Expected:     0.0,
			ExpectedBits: 16,
		},
		{
			// 1.0
			Data:         []byte{0xf9, 0x3c, 0x00},
			Expected:     1.0,
			ExpectedBits: 16,
		},
		{
			// 1.5
			Data:         []byte{0xf9, 0x3e, 0x00},
			Expected:     1.5,
			ExpectedBits: 16,
		},
		{
			// 65504.0
			Data:         []byte{0xf9, 0x7b, 0xff},
			Expected:     65504.0,
			ExpectedBits: 16,
		},
		{
			// 5.960464477539063e-8
			Data:         []byte{0xf9, 0x00, 0x01},
			Expected:     5.960464477539063e-8,
			ExpectedBits: 16,
		},
		{
			// -4.0
			Data:         []byte{0xf9, 0xc4, 0x00},
			Expected:     -4.0,
			ExpectedBits: 16,
		},
		{
			// 100000.0
			Data:         []byte{0xfa, 0x47, 0xc3, 0x50, 0x00},
			Expected:     100000.0,
			ExpectedBits: 32,
		},
		{
			// 3.4028234663852886e+38
			Data:         []byte{0xfa, 0x7f, 0x7f, 0xff, 0xff},
			Expected:     3.4028234663852886e+38,
			ExpectedBits: 32,
		},
		{
			// 1.1
			Data:         []byte{0xfb, 0x3f, 0xf1, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a},
			Expected:     1.1,
			ExpectedBits: 64,
		},
		{
			// 1.0e+300
			Data:         []byte{0xfb, 0x7e, 0x37, 0xe4, 0x3c, 0x88, 0x00, 0x75, 0x9c},
			Expected:     1.0e+300,
			ExpectedBits: 64,
		},
		{
			// -4.1
			Data:         []byte{0xfb, 0xc0, 0x10, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66},
			Expected:     -4.1,
			ExpectedBits: 64,
		},
		{
			// Infinity
			Data:         []byte{0xf9, 0x7c, 0x00},
			Expected:     math.Inf(1),
			ExpectedBits: 16,
		},
		{
			// -Infinity
			Data:         []byte{0xfb, 0xff, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			Expected:     math.Inf(-1),
			ExpectedBits: 64,
		},
	}

	for testNumber, test := range tests {

		actual, bits, rest, err := float64s.Unmarshal(test.Data)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		if test.Expected != actual || test.ExpectedBits != bits || 0 != len(rest) {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v (%d bits)", test.Expected, test.ExpectedBits)
			t.Logf("ACTUAL:   %v (%d bits)", actual, bits)
			t.Logf("REST: %#v", rest)
			continue
		}
	}
}

func TestUnmarshal_nan(t *testing.T) {
	// NaN (RFC 8949 Appendix A)
	actual, _, _, err := float64s.Unmarshal([]byte{0xf9, 0x7e, 0x00})
	if nil != err {
		t.Fatalf("Did not expect an error but actually got one: (%T) %s", err, err)
	}
	if !math.IsNaN(actual) {
		t.Errorf("Expected NaN but actually got %v.", actual)
	}
}

func TestUnmarshal_errors(t *testing.T) {

	tests := [][]byte{
		{0xf4},       // false
		{0xf6},       // null
		{0x01},       // 1
		{0xf9, 0x00}, // truncated
	}

	for testNumber, test := range tests {

		_, _, _, err := float64s.Unmarshal(test)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("DATA: %#v", test)
			continue
		}
	}
}

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value float64
		Expected []byte
	}{
		{
			Value:    1.1,
			Expected: []byte{0xfb, 0x3f, 0xf1, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a},
		},
		{
			Value:    -4.1,
			Expected: []byte{0xfb, 0xc0, 0x10, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66},
		},
		{
			// Always 64 bits, even when 16 would do.
			Value:    1.0,
			Expected: []byte{0xfb, 0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
	}

	for testNumber, test := range tests {

		actual, err := float64s.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		if !bytes.Equal(test.Expected, actual) {
			t.Errorf("For test #%d, the actual encoding is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", test.Expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}
	}
}
