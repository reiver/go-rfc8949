package epochs_test

import (
	"testing"

	"bytes"
	"time"

	"github.com/reiver/go-rfc8949/types/tags/epochs"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    time.Time
		Expected []byte
	}{
		// RFC 8949 Appendix A example.
		{
			// 1(1363896240)
			Value:    time.Date(2013, 3, 21, 20, 4, 0, 0, time.UTC),
			Expected: []byte{0xc1, 0x1a, 0x51, 0x4b, 0x67, 0xb0},
		},



		// Unix epoch.
		{
			// 1(0)
			Value:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			Expected: []byte{0xc1, 0x00},
		},



		// Small positive values.
		{
			// 1(1) — one second after epoch
			Value:    time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC),
			Expected: []byte{0xc1, 0x01},
		},
		{
			// 1(10) — ten seconds after epoch
			Value:    time.Date(1970, 1, 1, 0, 0, 10, 0, time.UTC),
			Expected: []byte{0xc1, 0x0a},
		},



		// Negative value (before epoch).
		{
			// 1(-1) — one second before epoch
			Value:    time.Date(1969, 12, 31, 23, 59, 59, 0, time.UTC),
			Expected: []byte{0xc1, 0x20},
		},
		{
			// 1969-01-01T00:00:00Z
			Value:    time.Date(1969, 1, 1, 0, 0, 0, 0, time.UTC),
			Expected: []byte{0xc1, 0x3a, 0x01, 0xe1, 0x33, 0x7f},
		},



		// Sub-second precision is truncated.
		{
			// 1(1363896240) — nanoseconds ignored
			Value:    time.Date(2013, 3, 21, 20, 4, 0, 500000000, time.UTC),
			Expected: []byte{0xc1, 0x1a, 0x51, 0x4b, 0x67, 0xb0},
		},



		// Large value.
		{
			// 2100-01-01T00:00:00Z = 4102444800
			Value:    time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
			Expected: []byte{0xc1, 0x1a, 0xf4, 0x86, 0x57, 0x00},
		},
	}

	for testNumber, test := range tests {

		actual, err := epochs.Marshal(test.Value)

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
				t.Logf("UNIX: %d", test.Value.Unix())
				continue
			}
		}
	}
}
