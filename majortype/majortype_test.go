package majortype_test

import (
	"testing"

	"github.com/reiver/go-rfc8949/majortype"
)

func TestConstants(t *testing.T) {

	tests := []struct{
		Name string
		Actual byte
		Expected byte
	}{
		{
			Name:     "MajorType0",
			Actual:   majortype.MajorType0,
			Expected: 0b000_00000,
		},
		{
			Name:     "MajorType1",
			Actual:   majortype.MajorType1,
			Expected: 0b001_00000,
		},
		{
			Name:     "MajorType2",
			Actual:   majortype.MajorType2,
			Expected: 0b010_00000,
		},
		{
			Name:     "MajorType3",
			Actual:   majortype.MajorType3,
			Expected: 0b011_00000,
		},
		{
			Name:     "MajorType4",
			Actual:   majortype.MajorType4,
			Expected: 0b100_00000,
		},
		{
			Name:     "MajorType5",
			Actual:   majortype.MajorType5,
			Expected: 0b101_00000,
		},
		{
			Name:     "MajorType6",
			Actual:   majortype.MajorType6,
			Expected: 0b110_00000,
		},
		{
			Name:     "MajorType7",
			Actual:   majortype.MajorType7,
			Expected: 0b111_00000,
		},

		// Named aliases
		{
			Name:     "UnsignedInteger",
			Actual:   majortype.UnsignedInteger,
			Expected: 0b000_00000,
		},
		{
			Name:     "NegativeInteger",
			Actual:   majortype.NegativeInteger,
			Expected: 0b001_00000,
		},
		{
			Name:     "ByteString",
			Actual:   majortype.ByteString,
			Expected: 0b010_00000,
		},
		{
			Name:     "TextString",
			Actual:   majortype.TextString,
			Expected: 0b011_00000,
		},
		{
			Name:     "Array",
			Actual:   majortype.Array,
			Expected: 0b100_00000,
		},
		{
			Name:     "Map",
			Actual:   majortype.Map,
			Expected: 0b101_00000,
		},
		{
			Name:     "Tagged",
			Actual:   majortype.Tagged,
			Expected: 0b110_00000,
		},
		{
			Name:     "FloatingPoint",
			Actual:   majortype.FloatingPoint,
			Expected: 0b111_00000,
		},
		{
			Name:     "SimpleValue",
			Actual:   majortype.SimpleValue,
			Expected: 0b111_00000,
		},
		{
			Name:     "BreakStopCode",
			Actual:   majortype.BreakStopCode,
			Expected: 0b111_00000,
		},
	}

	for testNumber, test := range tests {

		actual   := test.Actual
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value of %s is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED: 0x%02x (0b%08b)", expected, expected)
			t.Logf("ACTUAL:   0x%02x (0b%08b)", actual,   actual)
			continue
		}
	}
}

func TestMajorTypeFromInitialByte(t *testing.T) {

	tests := []struct{
		InitialByte byte
		Expected byte
	}{
		// Major type 0: initial bytes 0x00–0x1f
		{
			InitialByte: 0x00,
			Expected:    majortype.MajorType0,
		},
		{
			InitialByte: 0x17,
			Expected:    majortype.MajorType0,
		},
		{
			InitialByte: 0x1b,
			Expected:    majortype.MajorType0,
		},

		// Major type 1: initial bytes 0x20–0x3f
		{
			InitialByte: 0x20,
			Expected:    majortype.MajorType1,
		},
		{
			InitialByte: 0x37,
			Expected:    majortype.MajorType1,
		},
		{
			InitialByte: 0x3b,
			Expected:    majortype.MajorType1,
		},

		// Major type 2: initial bytes 0x40–0x5f
		{
			InitialByte: 0x40,
			Expected:    majortype.MajorType2,
		},
		{
			InitialByte: 0x57,
			Expected:    majortype.MajorType2,
		},
		{
			InitialByte: 0x5b,
			Expected:    majortype.MajorType2,
		},

		// Major type 3: initial bytes 0x60–0x7f
		{
			InitialByte: 0x60,
			Expected:    majortype.MajorType3,
		},
		{
			InitialByte: 0x77,
			Expected:    majortype.MajorType3,
		},
		{
			InitialByte: 0x7b,
			Expected:    majortype.MajorType3,
		},

		// Major type 4: initial bytes 0x80–0x9f
		{
			InitialByte: 0x80,
			Expected:    majortype.MajorType4,
		},
		{
			InitialByte: 0x97,
			Expected:    majortype.MajorType4,
		},
		{
			InitialByte: 0x9b,
			Expected:    majortype.MajorType4,
		},

		// Major type 5: initial bytes 0xa0–0xbf
		{
			InitialByte: 0xa0,
			Expected:    majortype.MajorType5,
		},
		{
			InitialByte: 0xb7,
			Expected:    majortype.MajorType5,
		},
		{
			InitialByte: 0xbb,
			Expected:    majortype.MajorType5,
		},

		// Major type 6: initial bytes 0xc0–0xdf
		{
			InitialByte: 0xc0,
			Expected:    majortype.MajorType6,
		},
		{
			InitialByte: 0xd7,
			Expected:    majortype.MajorType6,
		},
		{
			InitialByte: 0xdb,
			Expected:    majortype.MajorType6,
		},

		// Major type 7: initial bytes 0xe0–0xff
		{
			InitialByte: 0xe0,
			Expected:    majortype.MajorType7,
		},
		{
			InitialByte: 0xf4,
			Expected:    majortype.MajorType7,
		},
		{
			InitialByte: 0xf7,
			Expected:    majortype.MajorType7,
		},
		{
			InitialByte: 0xff,
			Expected:    majortype.MajorType7,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.MajorTypeFromInitialByte(test.InitialByte)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual major type from initial byte 0x%02x is not what was expected.", testNumber, test.InitialByte)
			t.Logf("EXPECTED: 0x%02x", expected)
			t.Logf("ACTUAL:   0x%02x", actual)
			continue
		}
	}
}

func TestIsUnsignedInteger(t *testing.T) {

	tests := []struct{
		Value byte
		Expected bool
	}{
		{
			Value:    majortype.UnsignedInteger,
			Expected: true,
		},
		{
			Value:    majortype.NegativeInteger,
			Expected: false,
		},
		{
			Value:    majortype.ByteString,
			Expected: false,
		},
		{
			Value:    majortype.TextString,
			Expected: false,
		},
		{
			Value:    majortype.Array,
			Expected: false,
		},
		{
			Value:    majortype.Map,
			Expected: false,
		},
		{
			Value:    majortype.Tagged,
			Expected: false,
		},
		{
			Value:    majortype.MajorType7,
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.IsUnsignedInteger(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, IsUnsignedInteger(0x%02x) returned unexpected result.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}

func TestIsNegativeInteger(t *testing.T) {

	tests := []struct{
		Value byte
		Expected bool
	}{
		{
			Value:    majortype.UnsignedInteger,
			Expected: false,
		},
		{
			Value:    majortype.NegativeInteger,
			Expected: true,
		},
		{
			Value:    majortype.ByteString,
			Expected: false,
		},
		{
			Value:    majortype.TextString,
			Expected: false,
		},
		{
			Value:    majortype.Array,
			Expected: false,
		},
		{
			Value:    majortype.Map,
			Expected: false,
		},
		{
			Value:    majortype.Tagged,
			Expected: false,
		},
		{
			Value:    majortype.MajorType7,
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.IsNegativeInteger(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, IsNegativeInteger(0x%02x) returned unexpected result.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}

func TestIsByteString(t *testing.T) {

	tests := []struct{
		Value byte
		Expected bool
	}{
		{
			Value:    majortype.UnsignedInteger,
			Expected: false,
		},
		{
			Value:    majortype.NegativeInteger,
			Expected: false,
		},
		{
			Value:    majortype.ByteString,
			Expected: true,
		},
		{
			Value:    majortype.TextString,
			Expected: false,
		},
		{
			Value:    majortype.Array,
			Expected: false,
		},
		{
			Value:    majortype.Map,
			Expected: false,
		},
		{
			Value:    majortype.Tagged,
			Expected: false,
		},
		{
			Value:    majortype.MajorType7,
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.IsByteString(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, IsByteString(0x%02x) returned unexpected result.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}

func TestIsTextString(t *testing.T) {

	tests := []struct{
		Value byte
		Expected bool
	}{
		{
			Value:    majortype.UnsignedInteger,
			Expected: false,
		},
		{
			Value:    majortype.NegativeInteger,
			Expected: false,
		},
		{
			Value:    majortype.ByteString,
			Expected: false,
		},
		{
			Value:    majortype.TextString,
			Expected: true,
		},
		{
			Value:    majortype.Array,
			Expected: false,
		},
		{
			Value:    majortype.Map,
			Expected: false,
		},
		{
			Value:    majortype.Tagged,
			Expected: false,
		},
		{
			Value:    majortype.MajorType7,
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.IsTextString(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, IsTextString(0x%02x) returned unexpected result.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}

func TestIsArray(t *testing.T) {

	tests := []struct{
		Value byte
		Expected bool
	}{
		{
			Value:    majortype.UnsignedInteger,
			Expected: false,
		},
		{
			Value:    majortype.NegativeInteger,
			Expected: false,
		},
		{
			Value:    majortype.ByteString,
			Expected: false,
		},
		{
			Value:    majortype.TextString,
			Expected: false,
		},
		{
			Value:    majortype.Array,
			Expected: true,
		},
		{
			Value:    majortype.Map,
			Expected: false,
		},
		{
			Value:    majortype.Tagged,
			Expected: false,
		},
		{
			Value:    majortype.MajorType7,
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.IsArray(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, IsArray(0x%02x) returned unexpected result.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}

func TestIsMap(t *testing.T) {

	tests := []struct{
		Value byte
		Expected bool
	}{
		{
			Value:    majortype.UnsignedInteger,
			Expected: false,
		},
		{
			Value:    majortype.NegativeInteger,
			Expected: false,
		},
		{
			Value:    majortype.ByteString,
			Expected: false,
		},
		{
			Value:    majortype.TextString,
			Expected: false,
		},
		{
			Value:    majortype.Array,
			Expected: false,
		},
		{
			Value:    majortype.Map,
			Expected: true,
		},
		{
			Value:    majortype.Tagged,
			Expected: false,
		},
		{
			Value:    majortype.MajorType7,
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.IsMap(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, IsMap(0x%02x) returned unexpected result.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}

func TestIsTagged(t *testing.T) {

	tests := []struct{
		Value byte
		Expected bool
	}{
		{
			Value:    majortype.UnsignedInteger,
			Expected: false,
		},
		{
			Value:    majortype.NegativeInteger,
			Expected: false,
		},
		{
			Value:    majortype.ByteString,
			Expected: false,
		},
		{
			Value:    majortype.TextString,
			Expected: false,
		},
		{
			Value:    majortype.Array,
			Expected: false,
		},
		{
			Value:    majortype.Map,
			Expected: false,
		},
		{
			Value:    majortype.Tagged,
			Expected: true,
		},
		{
			Value:    majortype.MajorType7,
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.IsTagged(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, IsTagged(0x%02x) returned unexpected result.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}

func TestCouldBeFloatingPoint(t *testing.T) {

	tests := []struct{
		Value byte
		Expected bool
	}{
		{
			Value:    majortype.UnsignedInteger,
			Expected: false,
		},
		{
			Value:    majortype.NegativeInteger,
			Expected: false,
		},
		{
			Value:    majortype.ByteString,
			Expected: false,
		},
		{
			Value:    majortype.TextString,
			Expected: false,
		},
		{
			Value:    majortype.Array,
			Expected: false,
		},
		{
			Value:    majortype.Map,
			Expected: false,
		},
		{
			Value:    majortype.Tagged,
			Expected: false,
		},
		{
			Value:    majortype.MajorType7,
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.CouldBeFloatingPoint(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, CouldBeFloatingPoint(0x%02x) returned unexpected result.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}

func TestCouldBeSimpleValue(t *testing.T) {

	tests := []struct{
		Value byte
		Expected bool
	}{
		{
			Value:    majortype.UnsignedInteger,
			Expected: false,
		},
		{
			Value:    majortype.NegativeInteger,
			Expected: false,
		},
		{
			Value:    majortype.ByteString,
			Expected: false,
		},
		{
			Value:    majortype.TextString,
			Expected: false,
		},
		{
			Value:    majortype.Array,
			Expected: false,
		},
		{
			Value:    majortype.Map,
			Expected: false,
		},
		{
			Value:    majortype.Tagged,
			Expected: false,
		},
		{
			Value:    majortype.MajorType7,
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.CouldBeSimpleValue(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, CouldBeSimpleValue(0x%02x) returned unexpected result.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}

func TestCouldBeBreakStopCode(t *testing.T) {

	tests := []struct{
		Value byte
		Expected bool
	}{
		{
			Value:    majortype.UnsignedInteger,
			Expected: false,
		},
		{
			Value:    majortype.NegativeInteger,
			Expected: false,
		},
		{
			Value:    majortype.ByteString,
			Expected: false,
		},
		{
			Value:    majortype.TextString,
			Expected: false,
		},
		{
			Value:    majortype.Array,
			Expected: false,
		},
		{
			Value:    majortype.Map,
			Expected: false,
		},
		{
			Value:    majortype.Tagged,
			Expected: false,
		},
		{
			Value:    majortype.MajorType7,
			Expected: true,
		},
	}

	for testNumber, test := range tests {

		actual   := majortype.CouldBeBreakStopCode(test.Value)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, CouldBeBreakStopCode(0x%02x) returned unexpected result.", testNumber, test.Value)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			continue
		}
	}
}

func TestMajorTypeFromInitialByte_AllBytes(t *testing.T) {

	// Every possible byte value should map to one of the 8 major types.
	// Each major type owns a contiguous range of 32 initial bytes.
	for i := 0; i < 256; i++ {
		b := byte(i)
		actual := majortype.MajorTypeFromInitialByte(b)
		expected := byte((i >> 5) << 5)

		if expected != actual {
			t.Errorf("For initial byte 0x%02x, the actual major type is not what was expected.", b)
			t.Logf("EXPECTED: 0x%02x", expected)
			t.Logf("ACTUAL:   0x%02x", actual)
			continue
		}
	}
}

func TestMajorTypeFromInitialByte_StripsAdditionalInfo(t *testing.T) {

	// Verify that the additional information bits (low 5) are always stripped,
	// regardless of their value.
	for mt := 0; mt < 8; mt++ {
		expected := byte(mt << 5)

		for ai := 0; ai < 32; ai++ {
			b := byte(mt << 5) | byte(ai)
			actual := majortype.MajorTypeFromInitialByte(b)

			if expected != actual {
				t.Errorf("For major type %d with additional info %d (byte 0x%02x), got wrong major type.", mt, ai, b)
				t.Logf("EXPECTED: 0x%02x", expected)
				t.Logf("ACTUAL:   0x%02x", actual)
			}
		}
	}
}

func TestIsCheckers_WithRawMajorTypeValues(t *testing.T) {

	// Each checker function should return true for exactly one major type value.
	type checkerFunc struct {
		Name string
		Fn   func(byte) bool
		TrueFor byte
	}

	checkers := []checkerFunc{
		{"IsUnsignedInteger",  majortype.IsUnsignedInteger,  majortype.MajorType0},
		{"IsNegativeInteger",  majortype.IsNegativeInteger,  majortype.MajorType1},
		{"IsByteString",       majortype.IsByteString,        majortype.MajorType2},
		{"IsTextString",       majortype.IsTextString,        majortype.MajorType3},
		{"IsArray",            majortype.IsArray,             majortype.MajorType4},
		{"IsMap",              majortype.IsMap,               majortype.MajorType5},
		{"IsTagged",           majortype.IsTagged,            majortype.MajorType6},
	}

	majorTypes := []byte{
		majortype.MajorType0,
		majortype.MajorType1,
		majortype.MajorType2,
		majortype.MajorType3,
		majortype.MajorType4,
		majortype.MajorType5,
		majortype.MajorType6,
		majortype.MajorType7,
	}

	for _, checker := range checkers {
		for _, mt := range majorTypes {
			expected := (mt == checker.TrueFor)
			actual := checker.Fn(mt)

			if expected != actual {
				t.Errorf("%s(0x%02x): expected %t, got %t", checker.Name, mt, expected, actual)
			}
		}
	}

	// The three MajorType7 checkers should all return true for MajorType7 and false for everything else.
	mt7Checkers := []struct{
		Name string
		Fn   func(byte) bool
	}{
		{"CouldBeFloatingPoint", majortype.CouldBeFloatingPoint},
		{"CouldBeSimpleValue",   majortype.CouldBeSimpleValue},
		{"CouldBeBreakStopCode", majortype.CouldBeBreakStopCode},
	}

	for _, checker := range mt7Checkers {
		for _, mt := range majorTypes {
			expected := (mt == majortype.MajorType7)
			actual := checker.Fn(mt)

			if expected != actual {
				t.Errorf("%s(0x%02x): expected %t, got %t", checker.Name, mt, expected, actual)
			}
		}
	}
}

func TestMajorTypeFromInitialByte_RoundTrip(t *testing.T) {

	// Extracting the major type from an initial byte built from a known major type
	// should return that same major type.
	majorTypes := []struct{
		Name string
		Value byte
	}{
		{"UnsignedInteger", majortype.UnsignedInteger},
		{"NegativeInteger", majortype.NegativeInteger},
		{"ByteString",      majortype.ByteString},
		{"TextString",      majortype.TextString},
		{"Array",           majortype.Array},
		{"Map",             majortype.Map},
		{"Tagged",          majortype.Tagged},
		{"MajorType7",      majortype.MajorType7},
	}

	for _, mt := range majorTypes {
		// Try combining the major type with a few different additional-info values.
		for _, ai := range []byte{0, 1, 10, 23, 24, 27, 31} {
			initialByte := mt.Value | ai

			actual := majortype.MajorTypeFromInitialByte(initialByte)

			if mt.Value != actual {
				t.Errorf("Round-trip failed for %s with additional info %d (initial byte 0x%02x).", mt.Name, ai, initialByte)
				t.Logf("EXPECTED: 0x%02x", mt.Value)
				t.Logf("ACTUAL:   0x%02x", actual)
			}
		}
	}
}

func TestCheckers_RejectNonMajorTypeValues(t *testing.T) {

	// The Is* functions compare against exact major type byte values.
	// Values that are NOT valid major types (i.e., have low bits set) should return false.
	nonMajorTypeValues := []byte{0x01, 0x1f, 0x21, 0x3f, 0x41, 0x5f, 0x61, 0x7f, 0x81, 0x9f, 0xa1, 0xbf, 0xc1, 0xdf, 0xe1, 0xff}

	checkers := []struct{
		Name string
		Fn   func(byte) bool
	}{
		{"IsUnsignedInteger",    majortype.IsUnsignedInteger},
		{"IsNegativeInteger",    majortype.IsNegativeInteger},
		{"IsByteString",         majortype.IsByteString},
		{"IsTextString",         majortype.IsTextString},
		{"IsArray",              majortype.IsArray},
		{"IsMap",                majortype.IsMap},
		{"IsTagged",             majortype.IsTagged},
		{"CouldBeFloatingPoint", majortype.CouldBeFloatingPoint},
		{"CouldBeSimpleValue",   majortype.CouldBeSimpleValue},
		{"CouldBeBreakStopCode", majortype.CouldBeBreakStopCode},
	}

	for _, checker := range checkers {
		for _, v := range nonMajorTypeValues {
			if checker.Fn(v) {
				t.Errorf("%s(0x%02x) returned true, but value has low bits set and is not a pure major type.", checker.Name, v)
			}
		}
	}

}
