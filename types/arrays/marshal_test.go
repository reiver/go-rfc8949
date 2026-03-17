package arrays_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/arrays"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value []any
		Expected []byte
	}{
		// RFC 8949 Appendix A examples.
		{
			// []
			Value:    []any{},
			Expected: []byte{0x80},
		},
		{
			// [1, 2, 3]
			Value:    []any{uint64(1), uint64(2), uint64(3)},
			Expected: []byte{0x83, 0x01, 0x02, 0x03},
		},
		{
			// [1, [2, 3], [4, 5]]
			Value:    []any{uint64(1), []any{uint64(2), uint64(3)}, []any{uint64(4), uint64(5)}},
			Expected: []byte{0x83, 0x01, 0x82, 0x02, 0x03, 0x82, 0x04, 0x05},
		},
		{
			// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25]
			Value:    []any{
				uint64(1), uint64(2), uint64(3), uint64(4), uint64(5),
				uint64(6), uint64(7), uint64(8), uint64(9), uint64(10),
				uint64(11), uint64(12), uint64(13), uint64(14), uint64(15),
				uint64(16), uint64(17), uint64(18), uint64(19), uint64(20),
				uint64(21), uint64(22), uint64(23), uint64(24), uint64(25),
			},
			Expected: []byte{
				0x98, 0x19,
				0x01, 0x02, 0x03, 0x04, 0x05,
				0x06, 0x07, 0x08, 0x09, 0x0a,
				0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
				0x10, 0x11, 0x12, 0x13, 0x14,
				0x15, 0x16, 0x17, 0x18, 0x18, 0x18, 0x19,
			},
		},



		// nil value (treated as empty array).
		{
			Value:    nil,
			Expected: []byte{0x80},
		},



		// Heterogeneous arrays.
		{
			// [true, false, nil]
			Value:    []any{true, false, nil},
			Expected: []byte{0x83, 0xf5, 0xf4, 0xf6},
		},
		{
			// [1, "hello"]
			Value:    []any{uint64(1), "hello"},
			Expected: []byte{0x82, 0x01, 0x65, 0x68, 0x65, 0x6c, 0x6c, 0x6f},
		},
		{
			// [h'AABB', "text"]
			Value:    []any{[]byte{0xAA, 0xBB}, "text"},
			Expected: []byte{0x82, 0x42, 0xAA, 0xBB, 0x64, 0x74, 0x65, 0x78, 0x74},
		},



		// Negative integers.
		{
			// [-1]
			Value:    []any{int64(-1)},
			Expected: []byte{0x81, 0x20},
		},
		{
			// [-10, 10]
			Value:    []any{int64(-10), uint64(10)},
			Expected: []byte{0x82, 0x29, 0x0a},
		},



		// Go int and uint types.
		{
			// [int(42), uint(100)]
			Value:    []any{int(42), uint(100)},
			Expected: []byte{0x82, 0x18, 0x2a, 0x18, 0x64},
		},



		// Inline array lengths (0-23 items).
		{
			Value:             []any{},
			Expected:          []byte{0x80},
		},
		{
			Value:             []any{uint8(0)},
			Expected:          []byte{0x81, 0x00},
		},
		{
			Value:             []any{uint8(0), uint8(0)},
			Expected:          []byte{0x82, 0x00, 0x00},
		},
		{
			Value:             []any{uint8(0), uint8(0), uint8(0)},
			Expected:          []byte{0x83, 0x00, 0x00, 0x00},
		},



		// Deeply nested array.
		{
			// [[[]]]
			Value:    []any{[]any{[]any{}}},
			Expected: []byte{0x81, 0x81, 0x80},
		},



		// All signed integer widths.
		{
			Value:    []any{int8(1), int16(2), int32(3), int64(4)},
			Expected: []byte{0x84, 0x01, 0x02, 0x03, 0x04},
		},



		// All unsigned integer widths.
		{
			Value:    []any{uint8(1), uint16(2), uint32(3), uint64(4)},
			Expected: []byte{0x84, 0x01, 0x02, 0x03, 0x04},
		},
	}

	for testNumber, test := range tests {

		actual, err := arrays.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}

func TestMarshalUnsupportedType(t *testing.T) {

	tests := []struct{
		Value []any
	}{
		{
			Value: []any{float64(1.0)},
		},
		{
			Value: []any{complex128(1+2i)},
		},
		{
			Value: []any{struct{}{}},
		},
	}

	for testNumber, test := range tests {

		_, err := arrays.Marshal(test.Value)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
