package initialbyte_test

import (
	"testing"

	"github.com/reiver/go-rfc8949/initialbyte"
)

func TestConstants(t *testing.T) {

	tests := []struct{
		Name string
		Actual byte
		Expected byte
	}{
		// Major type 0: Unsigned Integer (high 3 bits = 000)
		// Inline values 0–23
		{
			Name:     "UnsignedInteger0",
			Actual:   initialbyte.UnsignedInteger0,
			Expected: 0x00,
		},
		{
			Name:     "UnsignedInteger1",
			Actual:   initialbyte.UnsignedInteger1,
			Expected: 0x01,
		},
		{
			Name:     "UnsignedInteger2",
			Actual:   initialbyte.UnsignedInteger2,
			Expected: 0x02,
		},
		{
			Name:     "UnsignedInteger3",
			Actual:   initialbyte.UnsignedInteger3,
			Expected: 0x03,
		},
		{
			Name:     "UnsignedInteger4",
			Actual:   initialbyte.UnsignedInteger4,
			Expected: 0x04,
		},
		{
			Name:     "UnsignedInteger5",
			Actual:   initialbyte.UnsignedInteger5,
			Expected: 0x05,
		},
		{
			Name:     "UnsignedInteger6",
			Actual:   initialbyte.UnsignedInteger6,
			Expected: 0x06,
		},
		{
			Name:     "UnsignedInteger7",
			Actual:   initialbyte.UnsignedInteger7,
			Expected: 0x07,
		},
		{
			Name:     "UnsignedInteger8",
			Actual:   initialbyte.UnsignedInteger8,
			Expected: 0x08,
		},
		{
			Name:     "UnsignedInteger9",
			Actual:   initialbyte.UnsignedInteger9,
			Expected: 0x09,
		},
		{
			Name:     "UnsignedInteger10",
			Actual:   initialbyte.UnsignedInteger10,
			Expected: 0x0a,
		},
		{
			Name:     "UnsignedInteger11",
			Actual:   initialbyte.UnsignedInteger11,
			Expected: 0x0b,
		},
		{
			Name:     "UnsignedInteger12",
			Actual:   initialbyte.UnsignedInteger12,
			Expected: 0x0c,
		},
		{
			Name:     "UnsignedInteger13",
			Actual:   initialbyte.UnsignedInteger13,
			Expected: 0x0d,
		},
		{
			Name:     "UnsignedInteger14",
			Actual:   initialbyte.UnsignedInteger14,
			Expected: 0x0e,
		},
		{
			Name:     "UnsignedInteger15",
			Actual:   initialbyte.UnsignedInteger15,
			Expected: 0x0f,
		},
		{
			Name:     "UnsignedInteger16",
			Actual:   initialbyte.UnsignedInteger16,
			Expected: 0x10,
		},
		{
			Name:     "UnsignedInteger17",
			Actual:   initialbyte.UnsignedInteger17,
			Expected: 0x11,
		},
		{
			Name:     "UnsignedInteger18",
			Actual:   initialbyte.UnsignedInteger18,
			Expected: 0x12,
		},
		{
			Name:     "UnsignedInteger19",
			Actual:   initialbyte.UnsignedInteger19,
			Expected: 0x13,
		},
		{
			Name:     "UnsignedInteger20",
			Actual:   initialbyte.UnsignedInteger20,
			Expected: 0x14,
		},
		{
			Name:     "UnsignedInteger21",
			Actual:   initialbyte.UnsignedInteger21,
			Expected: 0x15,
		},
		{
			Name:     "UnsignedInteger22",
			Actual:   initialbyte.UnsignedInteger22,
			Expected: 0x16,
		},
		{
			Name:     "UnsignedInteger23",
			Actual:   initialbyte.UnsignedInteger23,
			Expected: 0x17,
		},

		// Unsigned integer additional-info markers
		{
			Name:     "Uint8",
			Actual:   initialbyte.Uint8,
			Expected: 0x18,
		},
		{
			Name:     "Uint16",
			Actual:   initialbyte.Uint16,
			Expected: 0x19,
		},
		{
			Name:     "Uint32",
			Actual:   initialbyte.Uint32,
			Expected: 0x1a,
		},
		{
			Name:     "Uint64",
			Actual:   initialbyte.Uint64,
			Expected: 0x1b,
		},

		// Major type 1: Negative Integer (high 3 bits = 001, offset = 0x20)
		// CBOR encodes -N as major type 1 with argument N-1
		{
			Name:     "NegativeIntegerNeg1",
			Actual:   initialbyte.NegativeIntegerNeg1,
			Expected: 0x20,
		},
		{
			Name:     "NegativeIntegerNeg2",
			Actual:   initialbyte.NegativeIntegerNeg2,
			Expected: 0x21,
		},
		{
			Name:     "NegativeIntegerNeg3",
			Actual:   initialbyte.NegativeIntegerNeg3,
			Expected: 0x22,
		},
		{
			Name:     "NegativeIntegerNeg4",
			Actual:   initialbyte.NegativeIntegerNeg4,
			Expected: 0x23,
		},
		{
			Name:     "NegativeIntegerNeg5",
			Actual:   initialbyte.NegativeIntegerNeg5,
			Expected: 0x24,
		},
		{
			Name:     "NegativeIntegerNeg6",
			Actual:   initialbyte.NegativeIntegerNeg6,
			Expected: 0x25,
		},
		{
			Name:     "NegativeIntegerNeg7",
			Actual:   initialbyte.NegativeIntegerNeg7,
			Expected: 0x26,
		},
		{
			Name:     "NegativeIntegerNeg8",
			Actual:   initialbyte.NegativeIntegerNeg8,
			Expected: 0x27,
		},
		{
			Name:     "NegativeIntegerNeg9",
			Actual:   initialbyte.NegativeIntegerNeg9,
			Expected: 0x28,
		},
		{
			Name:     "NegativeIntegerNeg10",
			Actual:   initialbyte.NegativeIntegerNeg10,
			Expected: 0x29,
		},
		{
			Name:     "NegativeIntegerNeg11",
			Actual:   initialbyte.NegativeIntegerNeg11,
			Expected: 0x2a,
		},
		{
			Name:     "NegativeIntegerNeg12",
			Actual:   initialbyte.NegativeIntegerNeg12,
			Expected: 0x2b,
		},
		{
			Name:     "NegativeIntegerNeg13",
			Actual:   initialbyte.NegativeIntegerNeg13,
			Expected: 0x2c,
		},
		{
			Name:     "NegativeIntegerNeg14",
			Actual:   initialbyte.NegativeIntegerNeg14,
			Expected: 0x2d,
		},
		{
			Name:     "NegativeIntegerNeg15",
			Actual:   initialbyte.NegativeIntegerNeg15,
			Expected: 0x2e,
		},
		{
			Name:     "NegativeIntegerNeg16",
			Actual:   initialbyte.NegativeIntegerNeg16,
			Expected: 0x2f,
		},
		{
			Name:     "NegativeIntegerNeg17",
			Actual:   initialbyte.NegativeIntegerNeg17,
			Expected: 0x30,
		},
		{
			Name:     "NegativeIntegerNeg18",
			Actual:   initialbyte.NegativeIntegerNeg18,
			Expected: 0x31,
		},
		{
			Name:     "NegativeIntegerNeg19",
			Actual:   initialbyte.NegativeIntegerNeg19,
			Expected: 0x32,
		},
		{
			Name:     "NegativeIntegerNeg20",
			Actual:   initialbyte.NegativeIntegerNeg20,
			Expected: 0x33,
		},
		{
			Name:     "NegativeIntegerNeg21",
			Actual:   initialbyte.NegativeIntegerNeg21,
			Expected: 0x34,
		},
		{
			Name:     "NegativeIntegerNeg22",
			Actual:   initialbyte.NegativeIntegerNeg22,
			Expected: 0x35,
		},
		{
			Name:     "NegativeIntegerNeg23",
			Actual:   initialbyte.NegativeIntegerNeg23,
			Expected: 0x36,
		},
		{
			Name:     "NegativeIntegerNeg24",
			Actual:   initialbyte.NegativeIntegerNeg24,
			Expected: 0x37,
		},

		// Negative integer additional-info markers
		{
			Name:     "Int8",
			Actual:   initialbyte.Int8,
			Expected: 0x38,
		},
		{
			Name:     "Int16",
			Actual:   initialbyte.Int16,
			Expected: 0x39,
		},
		{
			Name:     "Int32",
			Actual:   initialbyte.Int32,
			Expected: 0x3a,
		},
		{
			Name:     "Int64",
			Actual:   initialbyte.Int64,
			Expected: 0x3b,
		},

		// Major type 2: Byte String (high 3 bits = 010, offset = 0x40)
		{
			Name:     "ByteStringLen0",
			Actual:   initialbyte.ByteStringLen0,
			Expected: 0x40,
		},
		{
			Name:     "ByteStringLen1",
			Actual:   initialbyte.ByteStringLen1,
			Expected: 0x41,
		},
		{
			Name:     "ByteStringLen2",
			Actual:   initialbyte.ByteStringLen2,
			Expected: 0x42,
		},
		{
			Name:     "ByteStringLen3",
			Actual:   initialbyte.ByteStringLen3,
			Expected: 0x43,
		},
		{
			Name:     "ByteStringLen4",
			Actual:   initialbyte.ByteStringLen4,
			Expected: 0x44,
		},
		{
			Name:     "ByteStringLen5",
			Actual:   initialbyte.ByteStringLen5,
			Expected: 0x45,
		},
		{
			Name:     "ByteStringLen6",
			Actual:   initialbyte.ByteStringLen6,
			Expected: 0x46,
		},
		{
			Name:     "ByteStringLen7",
			Actual:   initialbyte.ByteStringLen7,
			Expected: 0x47,
		},
		{
			Name:     "ByteStringLen8",
			Actual:   initialbyte.ByteStringLen8,
			Expected: 0x48,
		},
		{
			Name:     "ByteStringLen9",
			Actual:   initialbyte.ByteStringLen9,
			Expected: 0x49,
		},
		{
			Name:     "ByteStringLen10",
			Actual:   initialbyte.ByteStringLen10,
			Expected: 0x4a,
		},
		{
			Name:     "ByteStringLen11",
			Actual:   initialbyte.ByteStringLen11,
			Expected: 0x4b,
		},
		{
			Name:     "ByteStringLen12",
			Actual:   initialbyte.ByteStringLen12,
			Expected: 0x4c,
		},
		{
			Name:     "ByteStringLen13",
			Actual:   initialbyte.ByteStringLen13,
			Expected: 0x4d,
		},
		{
			Name:     "ByteStringLen14",
			Actual:   initialbyte.ByteStringLen14,
			Expected: 0x4e,
		},
		{
			Name:     "ByteStringLen15",
			Actual:   initialbyte.ByteStringLen15,
			Expected: 0x4f,
		},
		{
			Name:     "ByteStringLen16",
			Actual:   initialbyte.ByteStringLen16,
			Expected: 0x50,
		},
		{
			Name:     "ByteStringLen17",
			Actual:   initialbyte.ByteStringLen17,
			Expected: 0x51,
		},
		{
			Name:     "ByteStringLen18",
			Actual:   initialbyte.ByteStringLen18,
			Expected: 0x52,
		},
		{
			Name:     "ByteStringLen19",
			Actual:   initialbyte.ByteStringLen19,
			Expected: 0x53,
		},
		{
			Name:     "ByteStringLen20",
			Actual:   initialbyte.ByteStringLen20,
			Expected: 0x54,
		},
		{
			Name:     "ByteStringLen21",
			Actual:   initialbyte.ByteStringLen21,
			Expected: 0x55,
		},
		{
			Name:     "ByteStringLen22",
			Actual:   initialbyte.ByteStringLen22,
			Expected: 0x56,
		},
		{
			Name:     "ByteStringLen23",
			Actual:   initialbyte.ByteStringLen23,
			Expected: 0x57,
		},

		// Byte string additional-info markers
		{
			Name:     "ByteStringLenUint8",
			Actual:   initialbyte.ByteStringLenUint8,
			Expected: 0x58,
		},
		{
			Name:     "ByteStringLenUint16",
			Actual:   initialbyte.ByteStringLenUint16,
			Expected: 0x59,
		},
		{
			Name:     "ByteStringLenUint32",
			Actual:   initialbyte.ByteStringLenUint32,
			Expected: 0x5a,
		},
		{
			Name:     "ByteStringLenUint64",
			Actual:   initialbyte.ByteStringLenUint64,
			Expected: 0x5b,
		},

		// Major type 7: Simple values (high 3 bits = 111, offset = 0xe0)
		{
			Name:     "False",
			Actual:   initialbyte.False,
			Expected: 0xf4,
		},
		{
			Name:     "True",
			Actual:   initialbyte.True,
			Expected: 0xf5,
		},
		{
			Name:     "Null",
			Actual:   initialbyte.Null,
			Expected: 0xf6,
		},
		{
			Name:     "Undefined",
			Actual:   initialbyte.Undefined,
			Expected: 0xf7,
		},
	}

	for testNumber, test := range tests {

		actual   := test.Actual
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value of %s is not what was expected.", testNumber, test.Name)
			t.Logf("EXPECTED: 0x%02x", expected)
			t.Logf("ACTUAL:   0x%02x", actual)
			continue
		}
	}
}
