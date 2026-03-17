package fulldates_test

import (
	"testing"

	"bytes"
	"time"

	"github.com/reiver/go-rfc8949/types/tags/fulldates"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    time.Time
		Expected []byte
	}{
		// RFC 8943 examples.
		{
			// 1004("1940-10-09")
			Value:    time.Date(1940, 10, 9, 0, 0, 0, 0, time.UTC),
			Expected: []byte{
				0xd9, 0x03, 0xec, // tag 1004
				0x6a,             // text string, length 10
				0x31, 0x39, 0x34, 0x30, 0x2d, 0x31, 0x30, 0x2d, 0x30, 0x39, // "1940-10-09"
			},
		},
		{
			// 1004("1980-12-08")
			Value:    time.Date(1980, 12, 8, 0, 0, 0, 0, time.UTC),
			Expected: []byte{
				0xd9, 0x03, 0xec, // tag 1004
				0x6a,             // text string, length 10
				0x31, 0x39, 0x38, 0x30, 0x2d, 0x31, 0x32, 0x2d, 0x30, 0x38, // "1980-12-08"
			},
		},



		// Unix epoch.
		{
			// 1004("1970-01-01")
			Value:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			Expected: []byte{
				0xd9, 0x03, 0xec, // tag 1004
				0x6a,             // text string, length 10
				0x31, 0x39, 0x37, 0x30, 0x2d, 0x30, 0x31, 0x2d, 0x30, 0x31, // "1970-01-01"
			},
		},



		// Time-of-day is discarded.
		{
			// 1004("2013-03-21") — even though input has 20:04:00 UTC
			Value:    time.Date(2013, 3, 21, 20, 4, 0, 0, time.UTC),
			Expected: []byte{
				0xd9, 0x03, 0xec, // tag 1004
				0x6a,             // text string, length 10
				0x32, 0x30, 0x31, 0x33, 0x2d, 0x30, 0x33, 0x2d, 0x32, 0x31, // "2013-03-21"
			},
		},



		// Timezone is discarded (date taken from the time.Time's location).
		{
			// 1004("2013-03-21") — input is in +05:30 but we just get the date in that zone
			Value:    time.Date(2013, 3, 21, 20, 4, 0, 0, time.FixedZone("IST", 5*3600+30*60)),
			Expected: []byte{
				0xd9, 0x03, 0xec, // tag 1004
				0x6a,             // text string, length 10
				0x32, 0x30, 0x31, 0x33, 0x2d, 0x30, 0x33, 0x2d, 0x32, 0x31, // "2013-03-21"
			},
		},



		// Go zero value.
		{
			// 1004("0001-01-01")
			Value:    time.Time{},
			Expected: []byte{
				0xd9, 0x03, 0xec, // tag 1004
				0x6a,             // text string, length 10
				0x30, 0x30, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x2d, 0x30, 0x31, // "0001-01-01"
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := fulldates.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %s", test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				t.Logf("VALUE: %s", test.Value)
				continue
			}
		}
	}
}
