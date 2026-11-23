package textstrings_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/types/textstrings"
)

func TestUnmarshal(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected string
		ExpectedRest []byte
	}{
		// RFC 8949 Appendix A examples.
		{
			Data:     []byte{0x60},
			Expected: "",
		},
		{
			Data:     []byte{0x61, 0x61},
			Expected: "a",
		},
		{
			Data:     []byte{0x64, 0x49, 0x45, 0x54, 0x46},
			Expected: "IETF",
		},
		{
			Data:     []byte{0x62, 0x22, 0x5c},
			Expected: "\"\\",
		},
		{
			Data:     []byte{0x62, 0xc3, 0xbc},
			Expected: "ü",
		},
		{
			Data:     []byte{0x63, 0xe6, 0xb0, 0xb4},
			Expected: "水",
		},
		{
			Data:     []byte{0x64, 0xf0, 0x90, 0x85, 0x91},
			Expected: "\U00010151",
		},



		{
			Data:         []byte{0x61, 0x61, 0x61, 0x62},
			Expected:     "a",
			ExpectedRest: []byte{0x61, 0x62},
		},
	}

	for testNumber, test := range tests {

		actual, rest, err := textstrings.Unmarshal(test.Data)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA: %#v", test.Data)
			continue
		}

		if test.Expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", test.Expected)
			t.Logf("ACTUAL:   %q", actual)
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

func TestUnmarshal_errors(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected error
	}{
		{
			Data:     []byte{},
			Expected: head.ErrEmpty,
		},
		{
			// says 4 bytes, but has only 2
			Data:     []byte{0x64, 0x49, 0x45},
			Expected: head.ErrUnexpectedEnd,
		},
		{
			// h''
			Data:     []byte{0x40},
			Expected: textstrings.ErrNotTextString,
		},
		{
			// (_ "strea", "ming")
			Data:     []byte{0x7f, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x67, 0xff},
			Expected: textstrings.ErrIndefiniteLength,
		},
		{
			// not valid UTF-8
			Data:     []byte{0x62, 0xc3, 0x28},
			Expected: textstrings.ErrInvalidUTF8,
		},
	}

	for testNumber, test := range tests {

		_, _, err := textstrings.Unmarshal(test.Data)

		if test.Expected != err {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", test.Expected)
			t.Logf("ACTUAL:   %v", err)
			t.Logf("DATA: %#v", test.Data)
			continue
		}
	}
}

func TestUnmarshal_roundTrip(t *testing.T) {

	values := []string{
		"",
		"a",
		"hello world",
		strings.Repeat("x", 23),
		strings.Repeat("x", 24),
		strings.Repeat("x", 255),
		strings.Repeat("x", 256),
		strings.Repeat("x", 65536),
		"ü水\U00010151",
	}

	for testNumber, value := range values {

		p, err := textstrings.Marshal(value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error from Marshal but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		actual, rest, err := textstrings.Unmarshal(p)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error from Unmarshal but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		if value != actual || 0 != len(rest) {
			t.Errorf("For test #%d, the round-trip did not work.", testNumber)
			t.Logf("LEN(EXPECTED): %d", len(value))
			t.Logf("LEN(ACTUAL):   %d", len(actual))
			t.Logf("REST: %#v", rest)
			continue
		}
	}
}
