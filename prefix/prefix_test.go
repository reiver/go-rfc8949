package prefix_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/prefix"
)

func TestPrefix(t *testing.T) {

	tests := []struct{
		MajorType byte
		Length uint64
		Expected []byte
	}{
		{
			MajorType:       0b101_00000,
			Length:              0b00000, // 0
			Expected: []byte{0b101_00000},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b00001,// 1
			Expected: []byte{0b101_00001},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b00010, // 2
			Expected: []byte{0b101_00010},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b00011, // 3
			Expected: []byte{0b101_00011},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b00100, // 4
			Expected: []byte{0b101_00100},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b00101, // 5
			Expected: []byte{0b101_00101},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b00110, // 6
			Expected: []byte{0b101_00110},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b00111, // 7
			Expected: []byte{0b101_00111},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b01000, // 8
			Expected: []byte{0b101_01000},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b01001, // 9
			Expected: []byte{0b101_01001},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b01010, // 10
			Expected: []byte{0b101_01010},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b01011, // 11
			Expected: []byte{0b101_01011},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b01100, // 12
			Expected: []byte{0b101_01100},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b01101, // 13
			Expected: []byte{0b101_01101},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b01110, // 14
			Expected: []byte{0b101_01110},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b01111, // 15
			Expected: []byte{0b101_01111},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b10000, // 16
			Expected: []byte{0b101_10000},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b10001, // 17
			Expected: []byte{0b101_10001},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b10010, // 18
			Expected: []byte{0b101_10010},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b10011, // 19
			Expected: []byte{0b101_10011},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b10100, // 20
			Expected: []byte{0b101_10100},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b10101, // 21
			Expected: []byte{0b101_10101},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b10110, // 22
			Expected: []byte{0b101_10110},
		},
		{
			MajorType:       0b101_00000,
			Length:              0b10111, // 23
			Expected: []byte{0b101_10111},
		},









		{
			MajorType:       0b101_00000,
			Length:                       24,
			Expected: []byte{0b101_11000, 24},
		},
		{
			MajorType:       0b101_00000,
			Length:                       25,
			Expected: []byte{0b101_11000, 25},
		},
		{
			MajorType:       0b101_00000,
			Length:                       26,
			Expected: []byte{0b101_11000, 26},
		},
		{
			MajorType:       0b101_00000,
			Length:                       27,
			Expected: []byte{0b101_11000, 27},
		},
		{
			MajorType:       0b101_00000,
			Length:                       28,
			Expected: []byte{0b101_11000, 28},
		},
		{
			MajorType:       0b101_00000,
			Length:                       29,
			Expected: []byte{0b101_11000, 29},
		},
		{
			MajorType:       0b101_00000,
			Length:                       30,
			Expected: []byte{0b101_11000, 30},
		},
		{
			MajorType:       0b101_00000,
			Length:                       31,
			Expected: []byte{0b101_11000, 31},
		},
		{
			MajorType:       0b101_00000,
			Length:                       32,
			Expected: []byte{0b101_11000, 32},
		},



		{
			MajorType:       0b101_00000,
			Length:                       250,
			Expected: []byte{0b101_11000, 250},
		},
		{
			MajorType:       0b101_00000,
			Length:                       251,
			Expected: []byte{0b101_11000, 251},
		},
		{
			MajorType:       0b101_00000,
			Length:                       252,
			Expected: []byte{0b101_11000, 252},
		},
		{
			MajorType:       0b101_00000,
			Length:                       253,
			Expected: []byte{0b101_11000, 253},
		},
		{
			MajorType:       0b101_00000,
			Length:                       254,
			Expected: []byte{0b101_11000, 254},
		},
		{
			MajorType:       0b101_00000,
			Length:                       255,
			Expected: []byte{0b101_11000, 255},
		},









		{
			MajorType:       0b101_00000,
			Length:       256,
			Expected: []byte{0b101_11001,
				byte((256 & 0xFF_00) >> 8),
				byte((256 & 0x00_FF)     ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       257,
			Expected: []byte{0b101_11001,
				byte((257 & 0xFF_00) >> 8),
				byte((257 & 0x00_FF)     ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       258,
			Expected: []byte{0b101_11001,
				byte((258 & 0xFF_00) >> 8),
				byte((258 & 0x00_FF)     ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       259,
			Expected: []byte{0b101_11001,
				byte((259 & 0xFF_00) >> 8),
				byte((259 & 0x00_FF)     ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       260,
			Expected: []byte{0b101_11001,
				byte((260 & 0xFF_00) >> 8),
				byte((260 & 0x00_FF)     ),
			},
		},



		{
			MajorType:       0b101_00000,
			Length:       65531,
			Expected: []byte{0b101_11001,
				byte((65531 & 0xFF_00) >> 8),
				byte((65531 & 0x00_FF)     ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       65532,
			Expected: []byte{0b101_11001,
				byte((65532 & 0xFF_00) >> 8),
				byte((65532 & 0x00_FF)     ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       65533,
			Expected: []byte{0b101_11001,
				byte((65533 & 0xFF_00) >> 8),
				byte((65533 & 0x00_FF)     ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       65534,
			Expected: []byte{0b101_11001,
				byte((65534 & 0xFF_00) >> 8),
				byte((65534 & 0x00_FF)     ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       65535,
			Expected: []byte{0b101_11001,
				byte((65535 & 0xFF_00) >> 8),
				byte((65535 & 0x00_FF)     ),
			},
		},









		{
			MajorType:       0b101_00000,
			Length:       65536,
			Expected: []byte{0b101_11010,
				byte((65536 & 0xFF_00_00_00) >> (8*3)),
				byte((65536 & 0x00_FF_00_00) >> (8*2)),
				byte((65536 & 0x00_00_FF_00) >>  8   ),
				byte((65536 & 0x00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       65537,
			Expected: []byte{0b101_11010,
				byte((65537 & 0xFF_00_00_00) >> (8*3)),
				byte((65537 & 0x00_FF_00_00) >> (8*2)),
				byte((65537 & 0x00_00_FF_00) >>  8   ),
				byte((65537 & 0x00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       65536,
			Expected: []byte{0b101_11010,
				byte((65536 & 0xFF_00_00_00) >> (8*3)),
				byte((65536 & 0x00_FF_00_00) >> (8*2)),
				byte((65536 & 0x00_00_FF_00) >>  8   ),
				byte((65536 & 0x00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       65538,
			Expected: []byte{0b101_11010,
				byte((65538 & 0xFF_00_00_00) >> (8*3)),
				byte((65538 & 0x00_FF_00_00) >> (8*2)),
				byte((65538 & 0x00_00_FF_00) >>  8   ),
				byte((65538 & 0x00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       65539,
			Expected: []byte{0b101_11010,
				byte((65539 & 0xFF_00_00_00) >> (8*3)),
				byte((65539 & 0x00_FF_00_00) >> (8*2)),
				byte((65539 & 0x00_00_FF_00) >>  8   ),
				byte((65539 & 0x00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       65540,
			Expected: []byte{0b101_11010,
				byte((65540 & 0xFF_00_00_00) >> (8*3)),
				byte((65540 & 0x00_FF_00_00) >> (8*2)),
				byte((65540 & 0x00_00_FF_00) >>  8   ),
				byte((65540 & 0x00_00_00_FF)         ),
			},
		},



		{
			MajorType:       0b101_00000,
			Length:       4294967290,
			Expected: []byte{0b101_11010,
				byte((4294967290 & 0xFF_00_00_00) >> (8*3)),
				byte((4294967290 & 0x00_FF_00_00) >> (8*2)),
				byte((4294967290 & 0x00_00_FF_00) >>  8   ),
				byte((4294967290 & 0x00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       4294967291,
			Expected: []byte{0b101_11010,
				byte((4294967291 & 0xFF_00_00_00) >> (8*3)),
				byte((4294967291 & 0x00_FF_00_00) >> (8*2)),
				byte((4294967291 & 0x00_00_FF_00) >>  8   ),
				byte((4294967291 & 0x00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       4294967292,
			Expected: []byte{0b101_11010,
				byte((4294967292 & 0xFF_00_00_00) >> (8*3)),
				byte((4294967292 & 0x00_FF_00_00) >> (8*2)),
				byte((4294967292 & 0x00_00_FF_00) >>  8   ),
				byte((4294967292 & 0x00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       4294967293,
			Expected: []byte{0b101_11010,
				byte((4294967293 & 0xFF_00_00_00) >> (8*3)),
				byte((4294967293 & 0x00_FF_00_00) >> (8*2)),
				byte((4294967293 & 0x00_00_FF_00) >>  8   ),
				byte((4294967293 & 0x00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       4294967294,
			Expected: []byte{0b101_11010,
				byte((4294967294 & 0xFF_00_00_00) >> (8*3)),
				byte((4294967294 & 0x00_FF_00_00) >> (8*2)),
				byte((4294967294 & 0x00_00_FF_00) >>  8   ),
				byte((4294967294 & 0x00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       4294967295,
			Expected: []byte{0b101_11010,
				byte((4294967295 & 0xFF_00_00_00) >> (8*3)),
				byte((4294967295 & 0x00_FF_00_00) >> (8*2)),
				byte((4294967295 & 0x00_00_FF_00) >>  8   ),
				byte((4294967295 & 0x00_00_00_FF)         ),
			},
		},









		{
			MajorType:       0b101_00000,
			Length:       4294967296,
			Expected: []byte{0b101_11011,
				byte((4294967296 & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
				byte((4294967296 & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
				byte((4294967296 & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
				byte((4294967296 & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
				byte((4294967296 & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
				byte((4294967296 & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
				byte((4294967296 & 0x00_00_00_00_00_00_FF_00) >>  8   ),
				byte((4294967296 & 0x00_00_00_00_00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       4294967297,
			Expected: []byte{0b101_11011,
				byte((4294967297 & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
				byte((4294967297 & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
				byte((4294967297 & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
				byte((4294967297 & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
				byte((4294967297 & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
				byte((4294967297 & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
				byte((4294967297 & 0x00_00_00_00_00_00_FF_00) >>  8   ),
				byte((4294967297 & 0x00_00_00_00_00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       4294967298,
			Expected: []byte{0b101_11011,
				byte((4294967298 & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
				byte((4294967298 & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
				byte((4294967298 & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
				byte((4294967298 & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
				byte((4294967298 & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
				byte((4294967298 & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
				byte((4294967298 & 0x00_00_00_00_00_00_FF_00) >>  8   ),
				byte((4294967298 & 0x00_00_00_00_00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       4294967299,
			Expected: []byte{0b101_11011,
				byte((4294967299 & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
				byte((4294967299 & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
				byte((4294967299 & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
				byte((4294967299 & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
				byte((4294967299 & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
				byte((4294967299 & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
				byte((4294967299 & 0x00_00_00_00_00_00_FF_00) >>  8   ),
				byte((4294967299 & 0x00_00_00_00_00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       4294967300,
			Expected: []byte{0b101_11011,
				byte((4294967300 & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
				byte((4294967300 & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
				byte((4294967300 & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
				byte((4294967300 & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
				byte((4294967300 & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
				byte((4294967300 & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
				byte((4294967300 & 0x00_00_00_00_00_00_FF_00) >>  8   ),
				byte((4294967300 & 0x00_00_00_00_00_00_00_FF)         ),
			},
		},



		{
			MajorType:       0b101_00000,
			Length:       18446744073709551613,
			Expected: []byte{0b101_11011,
				byte((18446744073709551613 & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
				byte((18446744073709551613 & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
				byte((18446744073709551613 & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
				byte((18446744073709551613 & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
				byte((18446744073709551613 & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
				byte((18446744073709551613 & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
				byte((18446744073709551613 & 0x00_00_00_00_00_00_FF_00) >>  8   ),
				byte((18446744073709551613 & 0x00_00_00_00_00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       18446744073709551614,
			Expected: []byte{0b101_11011,
				byte((18446744073709551614 & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
				byte((18446744073709551614 & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
				byte((18446744073709551614 & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
				byte((18446744073709551614 & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
				byte((18446744073709551614 & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
				byte((18446744073709551614 & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
				byte((18446744073709551614 & 0x00_00_00_00_00_00_FF_00) >>  8   ),
				byte((18446744073709551614 & 0x00_00_00_00_00_00_00_FF)         ),
			},
		},
		{
			MajorType:       0b101_00000,
			Length:       18446744073709551615, // 2^64 - 1
			Expected: []byte{0b101_11011,
				byte((18446744073709551615 & 0xFF_00_00_00_00_00_00_00) >> (8*7)),
				byte((18446744073709551615 & 0x00_FF_00_00_00_00_00_00) >> (8*6)),
				byte((18446744073709551615 & 0x00_00_FF_00_00_00_00_00) >> (8*5)),
				byte((18446744073709551615 & 0x00_00_00_FF_00_00_00_00) >> (8*4)),
				byte((18446744073709551615 & 0x00_00_00_00_FF_00_00_00) >> (8*3)),
				byte((18446744073709551615 & 0x00_00_00_00_00_FF_00_00) >> (8*2)),
				byte((18446744073709551615 & 0x00_00_00_00_00_00_FF_00) >>  8   ),
				byte((18446744073709551615 & 0x00_00_00_00_00_00_00_FF)         ),
			},
		},
	}

	for testNumber, test := range tests {

		actual := prefix.Prefix(test.MajorType, test.Length)
		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual prefix it not what was expected." , testNumber)
			t.Logf("EXPECTED: (%d) %#v", len(expected), expected)
			t.Logf("ACTUAL  : (%d) %#v", len(actual), actual)
			t.Logf("EXPECTED MAJOR-TYPE: %#03b", (expected[0] & 0b111_00000) >> 5)
			t.Logf("ACTUAL   MAJOR-TYPE: %#03b", (actual[0]   & 0b111_00000) >> 5)
			t.Logf("EXPECTED EXTRA-INFO: %#05b (%d)", (expected[0] & 0b000_11111), (expected[0] & 0b000_11111))
			t.Logf("ACTUAL   EXTRA-INFO: %#05b (%d)", (actual[0]   & 0b000_11111), (actual[0]   & 0b000_11111))
			t.Logf("LENGTH: %d", test.Length)
			continue
		}
	}
}
