package uint8s_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/uint8s"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value uint8
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
		{
			Value:               128,
			Expected: []byte{24, 128},
		},
		{
			Value:               129,
			Expected: []byte{24, 129},
		},
		{
			Value:               130,
			Expected: []byte{24, 130},
		},
		{
			Value:               131,
			Expected: []byte{24, 131},
		},
		{
			Value:               132,
			Expected: []byte{24, 132},
		},
		{
			Value:               133,
			Expected: []byte{24, 133},
		},
		{
			Value:               134,
			Expected: []byte{24, 134},
		},
		{
			Value:               135,
			Expected: []byte{24, 135},
		},
		{
			Value:               136,
			Expected: []byte{24, 136},
		},
		{
			Value:               137,
			Expected: []byte{24, 137},
		},
		{
			Value:               138,
			Expected: []byte{24, 138},
		},
		{
			Value:               139,
			Expected: []byte{24, 139},
		},
		{
			Value:               140,
			Expected: []byte{24, 140},
		},
		{
			Value:               141,
			Expected: []byte{24, 141},
		},
		{
			Value:               142,
			Expected: []byte{24, 142},
		},
		{
			Value:               143,
			Expected: []byte{24, 143},
		},
		{
			Value:               144,
			Expected: []byte{24, 144},
		},
		{
			Value:               145,
			Expected: []byte{24, 145},
		},
		{
			Value:               146,
			Expected: []byte{24, 146},
		},
		{
			Value:               147,
			Expected: []byte{24, 147},
		},
		{
			Value:               148,
			Expected: []byte{24, 148},
		},
		{
			Value:               149,
			Expected: []byte{24, 149},
		},
		{
			Value:               150,
			Expected: []byte{24, 150},
		},
		{
			Value:               151,
			Expected: []byte{24, 151},
		},
		{
			Value:               152,
			Expected: []byte{24, 152},
		},
		{
			Value:               153,
			Expected: []byte{24, 153},
		},
		{
			Value:               154,
			Expected: []byte{24, 154},
		},
		{
			Value:               155,
			Expected: []byte{24, 155},
		},
		{
			Value:               156,
			Expected: []byte{24, 156},
		},
		{
			Value:               157,
			Expected: []byte{24, 157},
		},
		{
			Value:               158,
			Expected: []byte{24, 158},
		},
		{
			Value:               159,
			Expected: []byte{24, 159},
		},
		{
			Value:               160,
			Expected: []byte{24, 160},
		},
		{
			Value:               161,
			Expected: []byte{24, 161},
		},
		{
			Value:               162,
			Expected: []byte{24, 162},
		},
		{
			Value:               163,
			Expected: []byte{24, 163},
		},
		{
			Value:               164,
			Expected: []byte{24, 164},
		},
		{
			Value:               165,
			Expected: []byte{24, 165},
		},
		{
			Value:               166,
			Expected: []byte{24, 166},
		},
		{
			Value:               167,
			Expected: []byte{24, 167},
		},
		{
			Value:               168,
			Expected: []byte{24, 168},
		},
		{
			Value:               169,
			Expected: []byte{24, 169},
		},
		{
			Value:               170,
			Expected: []byte{24, 170},
		},
		{
			Value:               171,
			Expected: []byte{24, 171},
		},
		{
			Value:               172,
			Expected: []byte{24, 172},
		},
		{
			Value:               173,
			Expected: []byte{24, 173},
		},
		{
			Value:               174,
			Expected: []byte{24, 174},
		},
		{
			Value:               175,
			Expected: []byte{24, 175},
		},
		{
			Value:               176,
			Expected: []byte{24, 176},
		},
		{
			Value:               177,
			Expected: []byte{24, 177},
		},
		{
			Value:               178,
			Expected: []byte{24, 178},
		},
		{
			Value:               179,
			Expected: []byte{24, 179},
		},
		{
			Value:               180,
			Expected: []byte{24, 180},
		},
		{
			Value:               181,
			Expected: []byte{24, 181},
		},
		{
			Value:               182,
			Expected: []byte{24, 182},
		},
		{
			Value:               183,
			Expected: []byte{24, 183},
		},
		{
			Value:               184,
			Expected: []byte{24, 184},
		},
		{
			Value:               185,
			Expected: []byte{24, 185},
		},
		{
			Value:               186,
			Expected: []byte{24, 186},
		},
		{
			Value:               187,
			Expected: []byte{24, 187},
		},
		{
			Value:               188,
			Expected: []byte{24, 188},
		},
		{
			Value:               189,
			Expected: []byte{24, 189},
		},
		{
			Value:               190,
			Expected: []byte{24, 190},
		},
		{
			Value:               191,
			Expected: []byte{24, 191},
		},
		{
			Value:               192,
			Expected: []byte{24, 192},
		},
		{
			Value:               193,
			Expected: []byte{24, 193},
		},
		{
			Value:               194,
			Expected: []byte{24, 194},
		},
		{
			Value:               195,
			Expected: []byte{24, 195},
		},
		{
			Value:               196,
			Expected: []byte{24, 196},
		},
		{
			Value:               197,
			Expected: []byte{24, 197},
		},
		{
			Value:               198,
			Expected: []byte{24, 198},
		},
		{
			Value:               199,
			Expected: []byte{24, 199},
		},
		{
			Value:               200,
			Expected: []byte{24, 200},
		},
		{
			Value:               201,
			Expected: []byte{24, 201},
		},
		{
			Value:               202,
			Expected: []byte{24, 202},
		},
		{
			Value:               203,
			Expected: []byte{24, 203},
		},
		{
			Value:               204,
			Expected: []byte{24, 204},
		},
		{
			Value:               205,
			Expected: []byte{24, 205},
		},
		{
			Value:               206,
			Expected: []byte{24, 206},
		},
		{
			Value:               207,
			Expected: []byte{24, 207},
		},
		{
			Value:               208,
			Expected: []byte{24, 208},
		},
		{
			Value:               209,
			Expected: []byte{24, 209},
		},
		{
			Value:               210,
			Expected: []byte{24, 210},
		},
		{
			Value:               211,
			Expected: []byte{24, 211},
		},
		{
			Value:               212,
			Expected: []byte{24, 212},
		},
		{
			Value:               213,
			Expected: []byte{24, 213},
		},
		{
			Value:               214,
			Expected: []byte{24, 214},
		},
		{
			Value:               215,
			Expected: []byte{24, 215},
		},
		{
			Value:               216,
			Expected: []byte{24, 216},
		},
		{
			Value:               217,
			Expected: []byte{24, 217},
		},
		{
			Value:               218,
			Expected: []byte{24, 218},
		},
		{
			Value:               219,
			Expected: []byte{24, 219},
		},
		{
			Value:               220,
			Expected: []byte{24, 220},
		},
		{
			Value:               221,
			Expected: []byte{24, 221},
		},
		{
			Value:               222,
			Expected: []byte{24, 222},
		},
		{
			Value:               223,
			Expected: []byte{24, 223},
		},
		{
			Value:               224,
			Expected: []byte{24, 224},
		},
		{
			Value:               225,
			Expected: []byte{24, 225},
		},
		{
			Value:               226,
			Expected: []byte{24, 226},
		},
		{
			Value:               227,
			Expected: []byte{24, 227},
		},
		{
			Value:               228,
			Expected: []byte{24, 228},
		},
		{
			Value:               229,
			Expected: []byte{24, 229},
		},
		{
			Value:               230,
			Expected: []byte{24, 230},
		},
		{
			Value:               231,
			Expected: []byte{24, 231},
		},
		{
			Value:               232,
			Expected: []byte{24, 232},
		},
		{
			Value:               233,
			Expected: []byte{24, 233},
		},
		{
			Value:               234,
			Expected: []byte{24, 234},
		},
		{
			Value:               235,
			Expected: []byte{24, 235},
		},
		{
			Value:               236,
			Expected: []byte{24, 236},
		},
		{
			Value:               237,
			Expected: []byte{24, 237},
		},
		{
			Value:               238,
			Expected: []byte{24, 238},
		},
		{
			Value:               239,
			Expected: []byte{24, 239},
		},
		{
			Value:               240,
			Expected: []byte{24, 240},
		},
		{
			Value:               241,
			Expected: []byte{24, 241},
		},
		{
			Value:               242,
			Expected: []byte{24, 242},
		},
		{
			Value:               243,
			Expected: []byte{24, 243},
		},
		{
			Value:               244,
			Expected: []byte{24, 244},
		},
		{
			Value:               245,
			Expected: []byte{24, 245},
		},
		{
			Value:               246,
			Expected: []byte{24, 246},
		},
		{
			Value:               247,
			Expected: []byte{24, 247},
		},
		{
			Value:               248,
			Expected: []byte{24, 248},
		},
		{
			Value:               249,
			Expected: []byte{24, 249},
		},
		{
			Value:               250,
			Expected: []byte{24, 250},
		},
		{
			Value:               251,
			Expected: []byte{24, 251},
		},
		{
			Value:               252,
			Expected: []byte{24, 252},
		},
		{
			Value:               253,
			Expected: []byte{24, 253},
		},
		{
			Value:               254,
			Expected: []byte{24, 254},
		},
		{
			Value:               255,
			Expected: []byte{24, 255},
		},
	}

	for testNumber, test := range tests {

		actual, err := uint8s.Marshal(test.Value)

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
