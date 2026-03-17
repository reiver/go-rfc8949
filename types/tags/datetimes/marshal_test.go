package datetimes_test

import (
	"testing"

	"bytes"
	"time"

	"github.com/reiver/go-rfc8949/types/tags/datetimes"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    time.Time
		Expected []byte
	}{
		// RFC 8949 Appendix A example.
		{
			// 0("2013-03-21T20:04:00Z")
			Value:    time.Date(2013, 3, 21, 20, 4, 0, 0, time.UTC),
			Expected: []byte{0xc0, 0x74, 0x32, 0x30, 0x31, 0x33, 0x2d, 0x30, 0x33, 0x2d, 0x32, 0x31, 0x54, 0x32, 0x30, 0x3a, 0x30, 0x34, 0x3a, 0x30, 0x30, 0x5a},
		},



		// Sub-second precision preserved.
		{
			// 0("2013-03-21T20:04:00.5Z")
			Value: time.Date(2013, 3, 21, 20, 4, 0, 500000000, time.UTC),
			Expected: append(
				[]byte{0xc0},
				append(
					[]byte{0x76},
					[]byte("2013-03-21T20:04:00.5Z")...,
				)...,
			),
		},
		{
			// 0("2013-03-21T20:04:00.123456789Z")
			Value: time.Date(2013, 3, 21, 20, 4, 0, 123456789, time.UTC),
			Expected: append(
				[]byte{0xc0},
				append(
					[]byte{0x78, 0x1e},
					[]byte("2013-03-21T20:04:00.123456789Z")...,
				)...,
			),
		},



		// Timezone offset.
		{
			// 0("2013-03-21T20:04:00+05:30")
			Value: time.Date(2013, 3, 21, 20, 4, 0, 0, time.FixedZone("IST", 5*3600+30*60)),
			Expected: append(
				[]byte{0xc0},
				append(
					[]byte{0x78, 0x19},
					[]byte("2013-03-21T20:04:00+05:30")...,
				)...,
			),
		},
		{
			// 0("2013-03-21T20:04:00-08:00")
			Value: time.Date(2013, 3, 21, 20, 4, 0, 0, time.FixedZone("PST", -8*3600)),
			Expected: append(
				[]byte{0xc0},
				append(
					[]byte{0x78, 0x19},
					[]byte("2013-03-21T20:04:00-08:00")...,
				)...,
			),
		},



		// Unix epoch.
		{
			// 0("1970-01-01T00:00:00Z")
			Value:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			Expected: append(
				[]byte{0xc0},
				append(
					[]byte{0x74},
					[]byte("1970-01-01T00:00:00Z")...,
				)...,
			),
		},



		// Go zero value.
		{
			// 0("0001-01-01T00:00:00Z")
			Value:    time.Time{},
			Expected: append(
				[]byte{0xc0},
				append(
					[]byte{0x74},
					[]byte("0001-01-01T00:00:00Z")...,
				)...,
			),
		},
	}

	for testNumber, test := range tests {

		actual, err := datetimes.Marshal(test.Value)

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
