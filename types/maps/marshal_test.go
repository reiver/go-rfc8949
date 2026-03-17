package maps_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/maps"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value map[any]any
		Expected []byte
	}{
		// RFC 8949 Appendix A examples.
		{
			// {}
			Value:    map[any]any{},
			Expected: []byte{0xa0},
		},
		{
			// {1: 2, 3: 4}
			// Keys sorted: 1 (0x01) < 3 (0x03)
			Value:    map[any]any{uint64(1): uint64(2), uint64(3): uint64(4)},
			Expected: []byte{0xa2, 0x01, 0x02, 0x03, 0x04},
		},
		{
			// {"a": 1, "b": [2, 3]}
			// Keys sorted: "a" (0x6161) < "b" (0x6162)
			Value:    map[any]any{"a": uint64(1), "b": []any{uint64(2), uint64(3)}},
			Expected: []byte{0xa2, 0x61, 0x61, 0x01, 0x61, 0x62, 0x82, 0x02, 0x03},
		},
		{
			// {"a": "A", "b": "B", "c": "C", "d": "D", "e": "E"}
			Value: map[any]any{"a": "A", "b": "B", "c": "C", "d": "D", "e": "E"},
			Expected: []byte{
				0xa5,
				0x61, 0x61, 0x61, 0x41,
				0x61, 0x62, 0x61, 0x42,
				0x61, 0x63, 0x61, 0x43,
				0x61, 0x64, 0x61, 0x44,
				0x61, 0x65, 0x61, 0x45,
			},
		},



		// nil value (treated as empty map).
		{
			Value:    nil,
			Expected: []byte{0xa0},
		},



		// Deterministic key sorting (Section 4.2.1).
		// Keys of different types/lengths sort by their encoded bytes.
		{
			// {10: 0, 100: 0, -1: 0}
			// Encoded keys: 10 -> 0x0a, 100 -> 0x1864, -1 -> 0x20
			// Sorted: 0x0a < 0x1864 < 0x20
			Value: map[any]any{uint64(10): uint64(0), uint64(100): uint64(0), int64(-1): uint64(0)},
			Expected: []byte{
				0xa3,
				0x0a, 0x00,
				0x18, 0x64, 0x00,
				0x20, 0x00,
			},
		},
		{
			// {"z": 0, "aa": 0}
			// Encoded keys: "z" -> 0x617a (2 bytes), "aa" -> 0x626161 (3 bytes)
			// Sorted: 0x617a < 0x626161
			Value: map[any]any{"z": uint64(0), "aa": uint64(0)},
			Expected: []byte{
				0xa2,
				0x61, 0x7a, 0x00,
				0x62, 0x61, 0x61, 0x00,
			},
		},



		// Heterogeneous keys and values.
		{
			// {1: true, "key": nil}
			// Encoded keys: 1 -> 0x01, "key" -> 0x636b6579
			// Sorted: 0x01 < 0x636b6579
			Value: map[any]any{uint64(1): true, "key": nil},
			Expected: []byte{
				0xa2,
				0x01, 0xf5,
				0x63, 0x6b, 0x65, 0x79, 0xf6,
			},
		},



		// Nested map.
		{
			// {"a": {"b": 1}}
			Value: map[any]any{"a": map[any]any{"b": uint64(1)}},
			Expected: []byte{
				0xa1,
				0x61, 0x61,
				0xa1,
				0x61, 0x62, 0x01,
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := maps.Marshal(test.Value)

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

func TestMarshalStringKeys(t *testing.T) {

	tests := []struct{
		Value map[string]any
		Expected []byte
	}{
		// RFC 8949 Appendix A examples.
		{
			// {}
			Value:    map[string]any{},
			Expected: []byte{0xa0},
		},
		{
			// {"a": 1, "b": [2, 3]}
			Value:    map[string]any{"a": uint64(1), "b": []any{uint64(2), uint64(3)}},
			Expected: []byte{0xa2, 0x61, 0x61, 0x01, 0x61, 0x62, 0x82, 0x02, 0x03},
		},
		{
			// {"a": "A", "b": "B", "c": "C", "d": "D", "e": "E"}
			Value: map[string]any{"a": "A", "b": "B", "c": "C", "d": "D", "e": "E"},
			Expected: []byte{
				0xa5,
				0x61, 0x61, 0x61, 0x41,
				0x61, 0x62, 0x61, 0x42,
				0x61, 0x63, 0x61, 0x43,
				0x61, 0x64, 0x61, 0x44,
				0x61, 0x65, 0x61, 0x45,
			},
		},



		// nil value (treated as empty map).
		{
			Value:    nil,
			Expected: []byte{0xa0},
		},



		// Deterministic key sorting by encoded bytes.
		{
			// {"z": 0, "aa": 0}
			// Encoded keys: "z" -> 0x617a (2 bytes), "aa" -> 0x626161 (3 bytes)
			// Sorted: 0x617a < 0x626161
			Value: map[string]any{"z": uint64(0), "aa": uint64(0)},
			Expected: []byte{
				0xa2,
				0x61, 0x7a, 0x00,
				0x62, 0x61, 0x61, 0x00,
			},
		},
		{
			// {"b": 1, "a": 2}
			// Sorted by encoded bytes: "a" < "b"
			Value: map[string]any{"b": uint64(1), "a": uint64(2)},
			Expected: []byte{
				0xa2,
				0x61, 0x61, 0x02,
				0x61, 0x62, 0x01,
			},
		},



		// Nested map[string]any.
		{
			// {"outer": {"inner": 42}}
			Value: map[string]any{"outer": map[string]any{"inner": uint64(42)}},
			Expected: []byte{
				0xa1,
				0x65, 0x6f, 0x75, 0x74, 0x65, 0x72,
				0xa1,
				0x65, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x2a,
			},
		},



		// Mixed value types.
		{
			// {"bool": true, "int": 1, "nil": nil, "str": "hello"}
			Value: map[string]any{"bool": true, "int": uint64(1), "nil": nil, "str": "hello"},
			Expected: []byte{
				0xa4,
				0x63, 0x69, 0x6e, 0x74, 0x01,
				0x63, 0x6e, 0x69, 0x6c, 0xf6,
				0x63, 0x73, 0x74, 0x72, 0x65, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
				0x64, 0x62, 0x6f, 0x6f, 0x6c, 0xf5,
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := maps.MarshalStringKeys(test.Value)

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

func TestMarshalUnsupportedType(t *testing.T) {

	tests := []struct{
		Value map[any]any
	}{
		{
			Value: map[any]any{"key": float64(1.0)},
		},
		{
			Value: map[any]any{float64(1.0): "value"},
		},
	}

	for testNumber, test := range tests {

		_, err := maps.Marshal(test.Value)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			continue
		}
	}
}
