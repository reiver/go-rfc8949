package int8s_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/int8s"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value int8
		Expected []byte
	}{
		{
			Value: 0,
			Expected: []byte{0x00},
		},
		{
			Value: 1,
			Expected: []byte{0x01},
		},
		{
			Value: 10,
			Expected: []byte{0x0a},
		},
		{
			Value: 23,
			Expected: []byte{0x17},
		},
		{
			Value: 24,
			Expected: []byte{0x18,0x18},
		},
		{
			Value: 25,
			Expected: []byte{0x18,0x19},
		},
		{
			Value: 100,
			Expected: []byte{0x18,0x64},
		},



		{
			Value: -1,
			Expected: []byte{0x20},
		},
		{
			Value: -10,
			Expected: []byte{0x29},
		},
		{
			Value: -100,
			Expected: []byte{0x38,0x63},
		},









		{
			Value:              -128,
			Expected: []byte{56, 128-1},
		},
		{
			Value:              -127,
			Expected: []byte{56, 127-1},
		},
		{
			Value:              -126,
			Expected: []byte{56, 126-1},
		},
		{
			Value:              -125,
			Expected: []byte{56, 125-1},
		},
		{
			Value:              -124,
			Expected: []byte{56, 124-1},
		},
		{
			Value:              -123,
			Expected: []byte{56, 123-1},
		},
		{
			Value:              -122,
			Expected: []byte{56, 122-1},
		},
		{
			Value:              -121,
			Expected: []byte{56, 121-1},
		},
		{
			Value:              -120,
			Expected: []byte{56, 120-1},
		},
		{
			Value:              -119,
			Expected: []byte{56, 119-1},
		},
		{
			Value:              -118,
			Expected: []byte{56, 118-1},
		},
		{
			Value:              -117,
			Expected: []byte{56, 117-1},
		},
		{
			Value:              -116,
			Expected: []byte{56, 116-1},
		},
		{
			Value:              -115,
			Expected: []byte{56, 115-1},
		},
		{
			Value:              -114,
			Expected: []byte{56, 114-1},
		},
		{
			Value:              -113,
			Expected: []byte{56, 113-1},
		},
		{
			Value:              -112,
			Expected: []byte{56, 112-1},
		},
		{
			Value:              -111,
			Expected: []byte{56, 111-1},
		},
		{
			Value:              -110,
			Expected: []byte{56, 110-1},
		},
		{
			Value:              -109,
			Expected: []byte{56, 109-1},
		},
		{
			Value:              -108,
			Expected: []byte{56, 108-1},
		},
		{
			Value:              -107,
			Expected: []byte{56, 107-1},
		},
		{
			Value:              -106,
			Expected: []byte{56, 106-1},
		},
		{
			Value:              -105,
			Expected: []byte{56, 105-1},
		},
		{
			Value:              -104,
			Expected: []byte{56, 104-1},
		},
		{
			Value:              -103,
			Expected: []byte{56, 103-1},
		},
		{
			Value:              -102,
			Expected: []byte{56, 102-1},
		},
		{
			Value:              -101,
			Expected: []byte{56, 101-1},
		},
		{
			Value:              -100,
			Expected: []byte{56, 100-1},
		},
		{
			Value:              -99,
			Expected: []byte{56, 99-1},
		},
		{
			Value:              -98,
			Expected: []byte{56, 98-1},
		},
		{
			Value:              -97,
			Expected: []byte{56, 97-1},
		},
		{
			Value:              -96,
			Expected: []byte{56, 96-1},
		},
		{
			Value:              -95,
			Expected: []byte{56, 95-1},
		},
		{
			Value:              -94,
			Expected: []byte{56, 94-1},
		},
		{
			Value:              -93,
			Expected: []byte{56, 93-1},
		},
		{
			Value:              -92,
			Expected: []byte{56, 92-1},
		},
		{
			Value:              -91,
			Expected: []byte{56, 91-1},
		},
		{
			Value:              -90,
			Expected: []byte{56, 90-1},
		},
		{
			Value:              -89,
			Expected: []byte{56, 89-1},
		},
		{
			Value:              -88,
			Expected: []byte{56, 88-1},
		},
		{
			Value:              -87,
			Expected: []byte{56, 87-1},
		},
		{
			Value:              -86,
			Expected: []byte{56, 86-1},
		},
		{
			Value:              -85,
			Expected: []byte{56, 85-1},
		},
		{
			Value:              -84,
			Expected: []byte{56, 84-1},
		},
		{
			Value:              -83,
			Expected: []byte{56, 83-1},
		},
		{
			Value:              -82,
			Expected: []byte{56, 82-1},
		},
		{
			Value:              -81,
			Expected: []byte{56, 81-1},
		},
		{
			Value:              -80,
			Expected: []byte{56, 80-1},
		},
		{
			Value:              -79,
			Expected: []byte{56, 79-1},
		},
		{
			Value:              -78,
			Expected: []byte{56, 78-1},
		},
		{
			Value:              -77,
			Expected: []byte{56, 77-1},
		},
		{
			Value:              -76,
			Expected: []byte{56, 76-1},
		},
		{
			Value:              -75,
			Expected: []byte{56, 75-1},
		},
		{
			Value:              -74,
			Expected: []byte{56, 74-1},
		},
		{
			Value:              -73,
			Expected: []byte{56, 73-1},
		},
		{
			Value:              -72,
			Expected: []byte{56, 72-1},
		},
		{
			Value:              -71,
			Expected: []byte{56, 71-1},
		},
		{
			Value:              -70,
			Expected: []byte{56, 70-1},
		},
		{
			Value:              -69,
			Expected: []byte{56, 69-1},
		},
		{
			Value:              -68,
			Expected: []byte{56, 68-1},
		},
		{
			Value:              -67,
			Expected: []byte{56, 67-1},
		},
		{
			Value:              -66,
			Expected: []byte{56, 66-1},
		},
		{
			Value:              -65,
			Expected: []byte{56, 65-1},
		},
		{
			Value:              -64,
			Expected: []byte{56, 64-1},
		},
		{
			Value:              -63,
			Expected: []byte{56, 63-1},
		},
		{
			Value:              -62,
			Expected: []byte{56, 62-1},
		},
		{
			Value:              -61,
			Expected: []byte{56, 61-1},
		},
		{
			Value:              -60,
			Expected: []byte{56, 60-1},
		},
		{
			Value:              -59,
			Expected: []byte{56, 59-1},
		},
		{
			Value:              -58,
			Expected: []byte{56, 58-1},
		},
		{
			Value:              -57,
			Expected: []byte{56, 57-1},
		},
		{
			Value:              -56,
			Expected: []byte{56, 56-1},
		},
		{
			Value:              -55,
			Expected: []byte{56, 55-1},
		},
		{
			Value:              -54,
			Expected: []byte{56, 54-1},
		},
		{
			Value:              -53,
			Expected: []byte{56, 53-1},
		},
		{
			Value:              -52,
			Expected: []byte{56, 52-1},
		},
		{
			Value:              -51,
			Expected: []byte{56, 51-1},
		},
		{
			Value:              -50,
			Expected: []byte{56, 50-1},
		},
		{
			Value:              -49,
			Expected: []byte{56, 49-1},
		},
		{
			Value:              -48,
			Expected: []byte{56, 48-1},
		},
		{
			Value:              -47,
			Expected: []byte{56, 47-1},
		},
		{
			Value:              -46,
			Expected: []byte{56, 46-1},
		},
		{
			Value:              -45,
			Expected: []byte{56, 45-1},
		},
		{
			Value:              -44,
			Expected: []byte{56, 44-1},
		},
		{
			Value:              -43,
			Expected: []byte{56, 43-1},
		},
		{
			Value:              -42,
			Expected: []byte{56, 42-1},
		},
		{
			Value:              -41,
			Expected: []byte{56, 41-1},
		},
		{
			Value:              -40,
			Expected: []byte{56, 40-1},
		},
		{
			Value:              -39,
			Expected: []byte{56, 39-1},
		},
		{
			Value:              -38,
			Expected: []byte{56, 38-1},
		},
		{
			Value:              -37,
			Expected: []byte{56, 37-1},
		},
		{
			Value:              -36,
			Expected: []byte{56, 36-1},
		},
		{
			Value:              -35,
			Expected: []byte{56, 35-1},
		},
		{
			Value:              -34,
			Expected: []byte{56, 34-1},
		},
		{
			Value:              -33,
			Expected: []byte{56, 33-1},
		},
		{
			Value:              -32,
			Expected: []byte{56, 32-1},
		},
		{
			Value:              -31,
			Expected: []byte{56, 31-1},
		},
		{
			Value:              -30,
			Expected: []byte{56, 30-1},
		},
		{
			Value:              -29,
			Expected: []byte{56, 29-1},
		},
		{
			Value:              -28,
			Expected: []byte{56, 28-1},
		},
		{
			Value:              -27,
			Expected: []byte{56, 27-1},
		},
		{
			Value:              -26,
			Expected: []byte{56, 26-1},
		},
		{
			Value:              -25,
			Expected: []byte{56, 25-1},
		},
		{
			Value:          -24,
			Expected: []byte{55},
		},
		{
			Value:          -23,
			Expected: []byte{54},
		},
		{
			Value:          -22,
			Expected: []byte{53},
		},
		{
			Value:          -21,
			Expected: []byte{52},
		},
		{
			Value:          -20,
			Expected: []byte{51},
		},
		{
			Value:          -19,
			Expected: []byte{50},
		},
		{
			Value:          -18,
			Expected: []byte{49},
		},
		{
			Value:          -17,
			Expected: []byte{48},
		},
		{
			Value:          -16,
			Expected: []byte{47},
		},
		{
			Value:          -15,
			Expected: []byte{46},
		},
		{
			Value:          -14,
			Expected: []byte{45},
		},
		{
			Value:          -13,
			Expected: []byte{44},
		},
		{
			Value:          -12,
			Expected: []byte{43},
		},
		{
			Value:          -11,
			Expected: []byte{42},
		},
		{
			Value:          -10,
			Expected: []byte{41},
		},
		{
			Value:           -9,
			Expected: []byte{40},
		},
		{
			Value:           -8,
			Expected: []byte{39},
		},
		{
			Value:           -7,
			Expected: []byte{38},
		},
		{
			Value:           -6,
			Expected: []byte{37},
		},
		{
			Value:           -5,
			Expected: []byte{36},
		},
		{
			Value:           -4,
			Expected: []byte{35},
		},
		{
			Value:           -3,
			Expected: []byte{34},
		},
		{
			Value:           -2,
			Expected: []byte{33},
		},
		{
			Value:           -1,
			Expected: []byte{32},
		},
		{
			Value:           0,
			Expected: []byte{0},
		},
		{
			Value:           1,
			Expected: []byte{1},
		},
		{
			Value:           2,
			Expected: []byte{2},
		},
		{
			Value:           3,
			Expected: []byte{3},
		},
		{
			Value:           4,
			Expected: []byte{4},
		},
		{
			Value:           5,
			Expected: []byte{5},
		},
		{
			Value:           6,
			Expected: []byte{6},
		},
		{
			Value:           7,
			Expected: []byte{7},
		},
		{
			Value:           8,
			Expected: []byte{8},
		},
		{
			Value:           9,
			Expected: []byte{9},
		},
		{
			Value:           10,
			Expected: []byte{10},
		},
		{
			Value:           11,
			Expected: []byte{11},
		},
		{
			Value:           12,
			Expected: []byte{12},
		},
		{
			Value:           13,
			Expected: []byte{13},
		},
		{
			Value:           14,
			Expected: []byte{14},
		},
		{
			Value:           15,
			Expected: []byte{15},
		},
		{
			Value:           16,
			Expected: []byte{16},
		},
		{
			Value:           17,
			Expected: []byte{17},
		},
		{
			Value:           18,
			Expected: []byte{18},
		},
		{
			Value:           19,
			Expected: []byte{19},
		},
		{
			Value:           20,
			Expected: []byte{20},
		},
		{
			Value:           21,
			Expected: []byte{21},
		},
		{
			Value:           22,
			Expected: []byte{22},
		},
		{
			Value:           23,
			Expected: []byte{23},
		},
		{
			Value:               24,
			Expected: []byte{24, 24},
		},
		{
			Value:               25,
			Expected: []byte{24, 25},
		},
		{
			Value:               26,
			Expected: []byte{24, 26},
		},
		{
			Value:               27,
			Expected: []byte{24, 27},
		},
		{
			Value:               28,
			Expected: []byte{24, 28},
		},
		{
			Value:               29,
			Expected: []byte{24, 29},
		},
		{
			Value:               30,
			Expected: []byte{24, 30},
		},
		{
			Value:               31,
			Expected: []byte{24, 31},
		},
		{
			Value:               32,
			Expected: []byte{24, 32},
		},
		{
			Value:               33,
			Expected: []byte{24, 33},
		},
		{
			Value:               34,
			Expected: []byte{24, 34},
		},
		{
			Value:               35,
			Expected: []byte{24, 35},
		},
		{
			Value:               36,
			Expected: []byte{24, 36},
		},
		{
			Value:               37,
			Expected: []byte{24, 37},
		},
		{
			Value:               39,
			Expected: []byte{24, 39},
		},
		{
			Value:               40,
			Expected: []byte{24, 40},
		},
		{
			Value:               41,
			Expected: []byte{24, 41},
		},
		{
			Value:               42,
			Expected: []byte{24, 42},
		},
		{
			Value:               43,
			Expected: []byte{24, 43},
		},
		{
			Value:               44,
			Expected: []byte{24, 44},
		},
		{
			Value:               45,
			Expected: []byte{24, 45},
		},
		{
			Value:               46,
			Expected: []byte{24, 46},
		},
		{
			Value:               47,
			Expected: []byte{24, 47},
		},
		{
			Value:               48,
			Expected: []byte{24, 48},
		},
		{
			Value:               49,
			Expected: []byte{24, 49},
		},
		{
			Value:               50,
			Expected: []byte{24, 50},
		},
		{
			Value:               51,
			Expected: []byte{24, 51},
		},
		{
			Value:               52,
			Expected: []byte{24, 52},
		},
		{
			Value:               53,
			Expected: []byte{24, 53},
		},
		{
			Value:               54,
			Expected: []byte{24, 54},
		},
		{
			Value:               55,
			Expected: []byte{24, 55},
		},
		{
			Value:               56,
			Expected: []byte{24, 56},
		},
		{
			Value:               57,
			Expected: []byte{24, 57},
		},
		{
			Value:               58,
			Expected: []byte{24, 58},
		},
		{
			Value:               59,
			Expected: []byte{24, 59},
		},
		{
			Value:               60,
			Expected: []byte{24, 60},
		},
		{
			Value:               61,
			Expected: []byte{24, 61},
		},
		{
			Value:               62,
			Expected: []byte{24, 62},
		},
		{
			Value:               63,
			Expected: []byte{24, 63},
		},
		{
			Value:               64,
			Expected: []byte{24, 64},
		},
		{
			Value:               65,
			Expected: []byte{24, 65},
		},
		{
			Value:               66,
			Expected: []byte{24, 66},
		},
		{
			Value:               67,
			Expected: []byte{24, 67},
		},
		{
			Value:               68,
			Expected: []byte{24, 68},
		},
		{
			Value:               69,
			Expected: []byte{24, 69},
		},
		{
			Value:               70,
			Expected: []byte{24, 70},
		},
		{
			Value:               71,
			Expected: []byte{24, 71},
		},
		{
			Value:               72,
			Expected: []byte{24, 72},
		},
		{
			Value:               73,
			Expected: []byte{24, 73},
		},
		{
			Value:               74,
			Expected: []byte{24, 74},
		},
		{
			Value:               75,
			Expected: []byte{24, 75},
		},
		{
			Value:               76,
			Expected: []byte{24, 76},
		},
		{
			Value:               77,
			Expected: []byte{24, 77},
		},
		{
			Value:               78,
			Expected: []byte{24, 78},
		},
		{
			Value:               79,
			Expected: []byte{24, 79},
		},
		{
			Value:               80,
			Expected: []byte{24, 80},
		},
		{
			Value:               81,
			Expected: []byte{24, 81},
		},
		{
			Value:               82,
			Expected: []byte{24, 82},
		},
		{
			Value:               83,
			Expected: []byte{24, 83},
		},
		{
			Value:               84,
			Expected: []byte{24, 84},
		},
		{
			Value:               85,
			Expected: []byte{24, 85},
		},
		{
			Value:               86,
			Expected: []byte{24, 86},
		},
		{
			Value:               87,
			Expected: []byte{24, 87},
		},
		{
			Value:               88,
			Expected: []byte{24, 88},
		},
		{
			Value:               89,
			Expected: []byte{24, 89},
		},
		{
			Value:               90,
			Expected: []byte{24, 90},
		},
		{
			Value:               91,
			Expected: []byte{24, 91},
		},
		{
			Value:               92,
			Expected: []byte{24, 92},
		},
		{
			Value:               93,
			Expected: []byte{24, 93},
		},
		{
			Value:               94,
			Expected: []byte{24, 94},
		},
		{
			Value:               95,
			Expected: []byte{24, 95},
		},
		{
			Value:               96,
			Expected: []byte{24, 96},
		},
		{
			Value:               97,
			Expected: []byte{24, 97},
		},
		{
			Value:               98,
			Expected: []byte{24, 98},
		},
		{
			Value:               99,
			Expected: []byte{24, 99},
		},
		{
			Value:               100,
			Expected: []byte{24, 100},
		},
		{
			Value:               101,
			Expected: []byte{24, 101},
		},
		{
			Value:               102,
			Expected: []byte{24, 102},
		},
		{
			Value:               103,
			Expected: []byte{24, 103},
		},
		{
			Value:               104,
			Expected: []byte{24, 104},
		},
		{
			Value:               105,
			Expected: []byte{24, 105},
		},
		{
			Value:               106,
			Expected: []byte{24, 106},
		},
		{
			Value:               107,
			Expected: []byte{24, 107},
		},
		{
			Value:               108,
			Expected: []byte{24, 108},
		},
		{
			Value:               109,
			Expected: []byte{24, 109},
		},
		{
			Value:               110,
			Expected: []byte{24, 110},
		},
		{
			Value:               111,
			Expected: []byte{24, 111},
		},
		{
			Value:               112,
			Expected: []byte{24, 112},
		},
		{
			Value:               113,
			Expected: []byte{24, 113},
		},
		{
			Value:               114,
			Expected: []byte{24, 114},
		},
		{
			Value:               115,
			Expected: []byte{24, 115},
		},
		{
			Value:               116,
			Expected: []byte{24, 116},
		},
		{
			Value:               117,
			Expected: []byte{24, 117},
		},
		{
			Value:               118,
			Expected: []byte{24, 118},
		},
		{
			Value:               119,
			Expected: []byte{24, 119},
		},
		{
			Value:               120,
			Expected: []byte{24, 120},
		},
		{
			Value:               121,
			Expected: []byte{24, 121},
		},
		{
			Value:               122,
			Expected: []byte{24, 122},
		},
		{
			Value:               123,
			Expected: []byte{24, 123},
		},
		{
			Value:               124,
			Expected: []byte{24, 124},
		},
		{
			Value:               125,
			Expected: []byte{24, 125},
		},
		{
			Value:               126,
			Expected: []byte{24, 126},
		},
		{
			Value:               127,
			Expected: []byte{24, 127},
		},
	}

	for testNumber, test := range tests {

		actual, err := int8s.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %d (0x%02X) (%#08b)", test.Value, test.Value, test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				t.Logf("VALUE: %d (0x%02X) (%#08b)", test.Value, test.Value, test.Value)
				continue
			}
		}
	}
}
