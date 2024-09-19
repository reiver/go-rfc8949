package int16s_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/int16s"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value int16
		Expected []byte
	}{
		{
			Value:                 -259,
			Expected: []byte{57, 1, 259-1-256},
		},
		{
			Value:                 -258,
			Expected: []byte{57, 1, 258-1-256},
		},
		{
			Value:                 -257,
			Expected: []byte{57, 1, 257-1-256},
		},









		{
			Value:              -256,
			Expected: []byte{56, 256-1},
		},
		{
			Value:              -255,
			Expected: []byte{56, 255-1},
		},
		{
			Value:              -254,
			Expected: []byte{56, 254-1},
		},
		{
			Value:              -253,
			Expected: []byte{56, 253-1},
		},
		{
			Value:              -252,
			Expected: []byte{56, 252-1},
		},
		{
			Value:              -251,
			Expected: []byte{56, 251-1},
		},
		{
			Value:              -250,
			Expected: []byte{56, 250-1},
		},
		{
			Value:              -249,
			Expected: []byte{56, 249-1},
		},
		{
			Value:              -248,
			Expected: []byte{56, 248-1},
		},



		{
			Value:              -230,
			Expected: []byte{56, 230-1},
		},



		{
			Value:              -220,
			Expected: []byte{56, 220-1},
		},



		{
			Value:              -210,
			Expected: []byte{56, 210-1},
		},



		{
			Value:              -200,
			Expected: []byte{56, 200-1},
		},



		{
			Value:              -190,
			Expected: []byte{56, 190-1},
		},



		{
			Value:              -180,
			Expected: []byte{56, 180-1},
		},



		{
			Value:              -170,
			Expected: []byte{56, 170-1},
		},



		{
			Value:              -160,
			Expected: []byte{56, 160-1},
		},



		{
			Value:              -150,
			Expected: []byte{56, 150-1},
		},



		{
			Value:              -140,
			Expected: []byte{56, 140-1},
		},



		{
			Value:              -139,
			Expected: []byte{56, 139-1},
		},
		{
			Value:              -138,
			Expected: []byte{56, 138-1},
		},
		{
			Value:              -137,
			Expected: []byte{56, 137-1},
		},
		{
			Value:              -136,
			Expected: []byte{56, 136-1},
		},
		{
			Value:              -135,
			Expected: []byte{56, 135-1},
		},
		{
			Value:              -134,
			Expected: []byte{56, 134-1},
		},
		{
			Value:              -133,
			Expected: []byte{56, 133-1},
		},
		{
			Value:              -132,
			Expected: []byte{56, 132-1},
		},
		{
			Value:              -131,
			Expected: []byte{56, 131-1},
		},
		{
			Value:              -130,
			Expected: []byte{56, 130-1},
		},
		{
			Value:              -129,
			Expected: []byte{56, 129-1},
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
			Value: 1000,
			Expected: []byte{0x19,0x03,0xe8},
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
			Value: -1000,
			Expected: []byte{0x39,0x03,0xe7},
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









		{
			Value:                  256,
			Expected: []byte{25, 1, 256-256},
		},
		{
			Value:                  257,
			Expected: []byte{25, 1, 257-256},
		},
		{
			Value:                  258,
			Expected: []byte{25, 1, 258-256},
		},
		{
			Value:                  259,
			Expected: []byte{25, 1, 259-256},
		},
		{
			Value:                  260,
			Expected: []byte{25, 1, 260-256},
		},
		{
			Value:                  261,
			Expected: []byte{25, 1, 261-256},
		},
		{
			Value:                  262,
			Expected: []byte{25, 1, 262-256},
		},
		{
			Value:                  263,
			Expected: []byte{25, 1, 263-256},
		},
		{
			Value:                  264,
			Expected: []byte{25, 1, 264-256},
		},
		{
			Value:                  265,
			Expected: []byte{25, 1, 265-256},
		},
		{
			Value:                  266,
			Expected: []byte{25, 1, 266-256},
		},
		{
			Value:                  267,
			Expected: []byte{25, 1, 267-256},
		},
		{
			Value:                  268,
			Expected: []byte{25, 1, 268-256},
		},
		{
			Value:                  269,
			Expected: []byte{25, 1, 269-256},
		},
		{
			Value:                  270,
			Expected: []byte{25, 1, 270-256},
		},
		{
			Value:                  271,
			Expected: []byte{25, 1, 271-256},
		},
		{
			Value:                  272,
			Expected: []byte{25, 1, 272-256},
		},
		{
			Value:                  273,
			Expected: []byte{25, 1, 273-256},
		},
		{
			Value:                  274,
			Expected: []byte{25, 1, 274-256},
		},
		{
			Value:                  275,
			Expected: []byte{25, 1, 275-256},
		},
		{
			Value:                  276,
			Expected: []byte{25, 1, 276-256},
		},
		{
			Value:                  277,
			Expected: []byte{25, 1, 277-256},
		},
		{
			Value:                  278,
			Expected: []byte{25, 1, 278-256},
		},
		{
			Value:                  279,
			Expected: []byte{25, 1, 279-256},
		},
		{
			Value:                  280,
			Expected: []byte{25, 1, 280-256},
		},
		{
			Value:                  281,
			Expected: []byte{25, 1, 281-256},
		},
		{
			Value:                  282,
			Expected: []byte{25, 1, 282-256},
		},
		{
			Value:                  283,
			Expected: []byte{25, 1, 283-256},
		},
		{
			Value:                  284,
			Expected: []byte{25, 1, 284-256},
		},
		{
			Value:                  285,
			Expected: []byte{25, 1, 285-256},
		},
		{
			Value:                  286,
			Expected: []byte{25, 1, 286-256},
		},
		{
			Value:                  287,
			Expected: []byte{25, 1, 287-256},
		},
		{
			Value:                  288,
			Expected: []byte{25, 1, 288-256},
		},
		{
			Value:                  289,
			Expected: []byte{25, 1, 289-256},
		},
		{
			Value:                  290,
			Expected: []byte{25, 1, 290-256},
		},
		{
			Value:                  291,
			Expected: []byte{25, 1, 291-256},
		},
		{
			Value:                  292,
			Expected: []byte{25, 1, 292-256},
		},
		{
			Value:                  293,
			Expected: []byte{25, 1, 293-256},
		},
		{
			Value:                  294,
			Expected: []byte{25, 1, 294-256},
		},
		{
			Value:                  295,
			Expected: []byte{25, 1, 295-256},
		},
		{
			Value:                  296,
			Expected: []byte{25, 1, 296-256},
		},
		{
			Value:                  297,
			Expected: []byte{25, 1, 297-256},
		},
		{
			Value:                  298,
			Expected: []byte{25, 1, 298-256},
		},
		{
			Value:                  299,
			Expected: []byte{25, 1, 299-256},
		},
		{
			Value:                  300,
			Expected: []byte{25, 1, 300-256},
		},
		{
			Value:                  301,
			Expected: []byte{25, 1, 301-256},
		},
		{
			Value:                  302,
			Expected: []byte{25, 1, 302-256},
		},
		{
			Value:                  303,
			Expected: []byte{25, 1, 303-256},
		},
		{
			Value:                  304,
			Expected: []byte{25, 1, 304-256},
		},
		{
			Value:                  305,
			Expected: []byte{25, 1, 305-256},
		},
		{
			Value:                  306,
			Expected: []byte{25, 1, 306-256},
		},
		{
			Value:                  307,
			Expected: []byte{25, 1, 307-256},
		},
		{
			Value:                  308,
			Expected: []byte{25, 1, 308-256},
		},
		{
			Value:                  309,
			Expected: []byte{25, 1, 309-256},
		},
		{
			Value:                  310,
			Expected: []byte{25, 1, 310-256},
		},
		{
			Value:                  311,
			Expected: []byte{25, 1, 311-256},
		},
		{
			Value:                  312,
			Expected: []byte{25, 1, 312-256},
		},
		{
			Value:                  313,
			Expected: []byte{25, 1, 313-256},
		},
		{
			Value:                  314,
			Expected: []byte{25, 1, 314-256},
		},
		{
			Value:                  315,
			Expected: []byte{25, 1, 315-256},
		},
		{
			Value:                  316,
			Expected: []byte{25, 1, 316-256},
		},
		{
			Value:                  317,
			Expected: []byte{25, 1, 317-256},
		},
		{
			Value:                  318,
			Expected: []byte{25, 1, 318-256},
		},
		{
			Value:                  319,
			Expected: []byte{25, 1, 319-256},
		},
		{
			Value:                  320,
			Expected: []byte{25, 1, 320-256},
		},
		{
			Value:                  321,
			Expected: []byte{25, 1, 321-256},
		},
		{
			Value:                  322,
			Expected: []byte{25, 1, 322-256},
		},
		{
			Value:                  323,
			Expected: []byte{25, 1, 323-256},
		},
		{
			Value:                  324,
			Expected: []byte{25, 1, 324-256},
		},
		{
			Value:                  325,
			Expected: []byte{25, 1, 325-256},
		},
		{
			Value:                  326,
			Expected: []byte{25, 1, 326-256},
		},
		{
			Value:                  327,
			Expected: []byte{25, 1, 327-256},
		},
		{
			Value:                  328,
			Expected: []byte{25, 1, 328-256},
		},
		{
			Value:                  329,
			Expected: []byte{25, 1, 329-256},
		},
		{
			Value:                  330,
			Expected: []byte{25, 1, 330-256},
		},
		{
			Value:                  331,
			Expected: []byte{25, 1, 331-256},
		},
		{
			Value:                  332,
			Expected: []byte{25, 1, 332-256},
		},
		{
			Value:                  333,
			Expected: []byte{25, 1, 333-256},
		},
		{
			Value:                  334,
			Expected: []byte{25, 1, 334-256},
		},
		{
			Value:                  335,
			Expected: []byte{25, 1, 335-256},
		},
		{
			Value:                  336,
			Expected: []byte{25, 1, 336-256},
		},
		{
			Value:                  337,
			Expected: []byte{25, 1, 337-256},
		},
		{
			Value:                  338,
			Expected: []byte{25, 1, 338-256},
		},
		{
			Value:                  339,
			Expected: []byte{25, 1, 339-256},
		},
		{
			Value:                  340,
			Expected: []byte{25, 1, 340-256},
		},
		{
			Value:                  341,
			Expected: []byte{25, 1, 341-256},
		},
		{
			Value:                  342,
			Expected: []byte{25, 1, 342-256},
		},
		{
			Value:                  343,
			Expected: []byte{25, 1, 343-256},
		},
		{
			Value:                  344,
			Expected: []byte{25, 1, 344-256},
		},
		{
			Value:                  345,
			Expected: []byte{25, 1, 345-256},
		},
		{
			Value:                  346,
			Expected: []byte{25, 1, 346-256},
		},
		{
			Value:                  347,
			Expected: []byte{25, 1, 347-256},
		},
		{
			Value:                  348,
			Expected: []byte{25, 1, 348-256},
		},
		{
			Value:                  349,
			Expected: []byte{25, 1, 349-256},
		},
		{
			Value:                  350,
			Expected: []byte{25, 1, 350-256},
		},
		{
			Value:                  351,
			Expected: []byte{25, 1, 351-256},
		},
		{
			Value:                  352,
			Expected: []byte{25, 1, 352-256},
		},
		{
			Value:                  353,
			Expected: []byte{25, 1, 353-256},
		},
		{
			Value:                  354,
			Expected: []byte{25, 1, 354-256},
		},
		{
			Value:                  355,
			Expected: []byte{25, 1, 355-256},
		},
		{
			Value:                  356,
			Expected: []byte{25, 1, 356-256},
		},
		{
			Value:                  357,
			Expected: []byte{25, 1, 357-256},
		},
		{
			Value:                  358,
			Expected: []byte{25, 1, 358-256},
		},
		{
			Value:                  359,
			Expected: []byte{25, 1, 359-256},
		},
		{
			Value:                  360,
			Expected: []byte{25, 1, 360-256},
		},
		{
			Value:                  361,
			Expected: []byte{25, 1, 361-256},
		},
		{
			Value:                  362,
			Expected: []byte{25, 1, 362-256},
		},
		{
			Value:                  363,
			Expected: []byte{25, 1, 363-256},
		},
		{
			Value:                  364,
			Expected: []byte{25, 1, 364-256},
		},
		{
			Value:                  365,
			Expected: []byte{25, 1, 365-256},
		},
		{
			Value:                  366,
			Expected: []byte{25, 1, 366-256},
		},
		{
			Value:                  367,
			Expected: []byte{25, 1, 367-256},
		},
		{
			Value:                  368,
			Expected: []byte{25, 1, 368-256},
		},
		{
			Value:                  369,
			Expected: []byte{25, 1, 369-256},
		},
		{
			Value:                  370,
			Expected: []byte{25, 1, 370-256},
		},
		{
			Value:                  371,
			Expected: []byte{25, 1, 371-256},
		},
		{
			Value:                  372,
			Expected: []byte{25, 1, 372-256},
		},
		{
			Value:                  373,
			Expected: []byte{25, 1, 373-256},
		},
		{
			Value:                  374,
			Expected: []byte{25, 1, 374-256},
		},
		{
			Value:                  375,
			Expected: []byte{25, 1, 375-256},
		},
		{
			Value:                  376,
			Expected: []byte{25, 1, 376-256},
		},
		{
			Value:                  377,
			Expected: []byte{25, 1, 377-256},
		},
		{
			Value:                  378,
			Expected: []byte{25, 1, 378-256},
		},
		{
			Value:                  379,
			Expected: []byte{25, 1, 379-256},
		},
		{
			Value:                  380,
			Expected: []byte{25, 1, 380-256},
		},
		{
			Value:                  381,
			Expected: []byte{25, 1, 381-256},
		},
		{
			Value:                  382,
			Expected: []byte{25, 1, 382-256},
		},
		{
			Value:                  383,
			Expected: []byte{25, 1, 383-256},
		},
		{
			Value:                  384,
			Expected: []byte{25, 1, 384-256},
		},
		{
			Value:                  385,
			Expected: []byte{25, 1, 385-256},
		},
		{
			Value:                  386,
			Expected: []byte{25, 1, 386-256},
		},
		{
			Value:                  387,
			Expected: []byte{25, 1, 387-256},
		},
		{
			Value:                  388,
			Expected: []byte{25, 1, 388-256},
		},
		{
			Value:                  389,
			Expected: []byte{25, 1, 389-256},
		},
		{
			Value:                  390,
			Expected: []byte{25, 1, 390-256},
		},
		{
			Value:                  391,
			Expected: []byte{25, 1, 391-256},
		},
		{
			Value:                  392,
			Expected: []byte{25, 1, 392-256},
		},
		{
			Value:                  393,
			Expected: []byte{25, 1, 393-256},
		},
		{
			Value:                  394,
			Expected: []byte{25, 1, 394-256},
		},
		{
			Value:                  395,
			Expected: []byte{25, 1, 395-256},
		},
		{
			Value:                  396,
			Expected: []byte{25, 1, 396-256},
		},
		{
			Value:                  397,
			Expected: []byte{25, 1, 397-256},
		},
		{
			Value:                  398,
			Expected: []byte{25, 1, 398-256},
		},
		{
			Value:                  399,
			Expected: []byte{25, 1, 399-256},
		},
		{
			Value:                  400,
			Expected: []byte{25, 1, 400-256},
		},
		{
			Value:                  401,
			Expected: []byte{25, 1, 401-256},
		},
		{
			Value:                  402,
			Expected: []byte{25, 1, 402-256},
		},
		{
			Value:                  403,
			Expected: []byte{25, 1, 403-256},
		},
		{
			Value:                  404,
			Expected: []byte{25, 1, 404-256},
		},
		{
			Value:                  405,
			Expected: []byte{25, 1, 405-256},
		},
		{
			Value:                  406,
			Expected: []byte{25, 1, 406-256},
		},
		{
			Value:                  407,
			Expected: []byte{25, 1, 407-256},
		},
		{
			Value:                  408,
			Expected: []byte{25, 1, 408-256},
		},
		{
			Value:                  409,
			Expected: []byte{25, 1, 409-256},
		},
		{
			Value:                  410,
			Expected: []byte{25, 1, 410-256},
		},
		{
			Value:                  411,
			Expected: []byte{25, 1, 411-256},
		},
		{
			Value:                  412,
			Expected: []byte{25, 1, 412-256},
		},
		{
			Value:                  413,
			Expected: []byte{25, 1, 413-256},
		},
		{
			Value:                  414,
			Expected: []byte{25, 1, 414-256},
		},
		{
			Value:                  415,
			Expected: []byte{25, 1, 415-256},
		},
		{
			Value:                  416,
			Expected: []byte{25, 1, 416-256},
		},
		{
			Value:                  417,
			Expected: []byte{25, 1, 417-256},
		},
		{
			Value:                  418,
			Expected: []byte{25, 1, 418-256},
		},
		{
			Value:                  419,
			Expected: []byte{25, 1, 419-256},
		},
		{
			Value:                  420,
			Expected: []byte{25, 1, 420-256},
		},
		{
			Value:                  421,
			Expected: []byte{25, 1, 421-256},
		},
		{
			Value:                  422,
			Expected: []byte{25, 1, 422-256},
		},
		{
			Value:                  423,
			Expected: []byte{25, 1, 423-256},
		},
		{
			Value:                  424,
			Expected: []byte{25, 1, 424-256},
		},
		{
			Value:                  425,
			Expected: []byte{25, 1, 425-256},
		},
		{
			Value:                  426,
			Expected: []byte{25, 1, 426-256},
		},
		{
			Value:                  427,
			Expected: []byte{25, 1, 427-256},
		},
		{
			Value:                  428,
			Expected: []byte{25, 1, 428-256},
		},
		{
			Value:                  429,
			Expected: []byte{25, 1, 429-256},
		},
		{
			Value:                  430,
			Expected: []byte{25, 1, 430-256},
		},
		{
			Value:                  431,
			Expected: []byte{25, 1, 431-256},
		},
		{
			Value:                  432,
			Expected: []byte{25, 1, 432-256},
		},
		{
			Value:                  433,
			Expected: []byte{25, 1, 433-256},
		},
		{
			Value:                  434,
			Expected: []byte{25, 1, 434-256},
		},
		{
			Value:                  435,
			Expected: []byte{25, 1, 435-256},
		},
		{
			Value:                  436,
			Expected: []byte{25, 1, 436-256},
		},
		{
			Value:                  437,
			Expected: []byte{25, 1, 437-256},
		},
		{
			Value:                  438,
			Expected: []byte{25, 1, 438-256},
		},
		{
			Value:                  439,
			Expected: []byte{25, 1, 439-256},
		},
		{
			Value:                  440,
			Expected: []byte{25, 1, 440-256},
		},
		{
			Value:                  441,
			Expected: []byte{25, 1, 441-256},
		},
		{
			Value:                  442,
			Expected: []byte{25, 1, 442-256},
		},
		{
			Value:                  443,
			Expected: []byte{25, 1, 443-256},
		},
		{
			Value:                  444,
			Expected: []byte{25, 1, 444-256},
		},
		{
			Value:                  445,
			Expected: []byte{25, 1, 445-256},
		},
		{
			Value:                  446,
			Expected: []byte{25, 1, 446-256},
		},
		{
			Value:                  447,
			Expected: []byte{25, 1, 447-256},
		},
		{
			Value:                  448,
			Expected: []byte{25, 1, 448-256},
		},
		{
			Value:                  449,
			Expected: []byte{25, 1, 449-256},
		},
		{
			Value:                  450,
			Expected: []byte{25, 1, 450-256},
		},
		{
			Value:                  451,
			Expected: []byte{25, 1, 451-256},
		},
		{
			Value:                  452,
			Expected: []byte{25, 1, 452-256},
		},
		{
			Value:                  453,
			Expected: []byte{25, 1, 453-256},
		},
		{
			Value:                  454,
			Expected: []byte{25, 1, 454-256},
		},
		{
			Value:                  455,
			Expected: []byte{25, 1, 455-256},
		},
		{
			Value:                  456,
			Expected: []byte{25, 1, 456-256},
		},
		{
			Value:                  457,
			Expected: []byte{25, 1, 457-256},
		},
		{
			Value:                  458,
			Expected: []byte{25, 1, 458-256},
		},
		{
			Value:                  459,
			Expected: []byte{25, 1, 459-256},
		},
		{
			Value:                  460,
			Expected: []byte{25, 1, 460-256},
		},
		{
			Value:                  461,
			Expected: []byte{25, 1, 461-256},
		},
		{
			Value:                  462,
			Expected: []byte{25, 1, 462-256},
		},
		{
			Value:                  463,
			Expected: []byte{25, 1, 463-256},
		},
		{
			Value:                  464,
			Expected: []byte{25, 1, 464-256},
		},
		{
			Value:                  465,
			Expected: []byte{25, 1, 465-256},
		},
		{
			Value:                  466,
			Expected: []byte{25, 1, 466-256},
		},
		{
			Value:                  467,
			Expected: []byte{25, 1, 467-256},
		},
		{
			Value:                  468,
			Expected: []byte{25, 1, 468-256},
		},
		{
			Value:                  469,
			Expected: []byte{25, 1, 469-256},
		},
		{
			Value:                  470,
			Expected: []byte{25, 1, 470-256},
		},
		{
			Value:                  471,
			Expected: []byte{25, 1, 471-256},
		},
		{
			Value:                  472,
			Expected: []byte{25, 1, 472-256},
		},
		{
			Value:                  473,
			Expected: []byte{25, 1, 473-256},
		},
		{
			Value:                  474,
			Expected: []byte{25, 1, 474-256},
		},
		{
			Value:                  475,
			Expected: []byte{25, 1, 475-256},
		},
		{
			Value:                  476,
			Expected: []byte{25, 1, 476-256},
		},
		{
			Value:                  477,
			Expected: []byte{25, 1, 477-256},
		},
		{
			Value:                  478,
			Expected: []byte{25, 1, 478-256},
		},
		{
			Value:                  479,
			Expected: []byte{25, 1, 479-256},
		},
		{
			Value:                  480,
			Expected: []byte{25, 1, 480-256},
		},
		{
			Value:                  481,
			Expected: []byte{25, 1, 481-256},
		},
		{
			Value:                  482,
			Expected: []byte{25, 1, 482-256},
		},
		{
			Value:                  483,
			Expected: []byte{25, 1, 483-256},
		},
		{
			Value:                  484,
			Expected: []byte{25, 1, 484-256},
		},
		{
			Value:                  485,
			Expected: []byte{25, 1, 485-256},
		},
		{
			Value:                  486,
			Expected: []byte{25, 1, 486-256},
		},
		{
			Value:                  487,
			Expected: []byte{25, 1, 487-256},
		},
		{
			Value:                  488,
			Expected: []byte{25, 1, 488-256},
		},
		{
			Value:                  489,
			Expected: []byte{25, 1, 489-256},
		},
		{
			Value:                  490,
			Expected: []byte{25, 1, 490-256},
		},
		{
			Value:                  491,
			Expected: []byte{25, 1, 491-256},
		},
		{
			Value:                  492,
			Expected: []byte{25, 1, 492-256},
		},
		{
			Value:                  493,
			Expected: []byte{25, 1, 493-256},
		},
		{
			Value:                  494,
			Expected: []byte{25, 1, 494-256},
		},
		{
			Value:                  495,
			Expected: []byte{25, 1, 495-256},
		},
		{
			Value:                  496,
			Expected: []byte{25, 1, 496-256},
		},
		{
			Value:                  497,
			Expected: []byte{25, 1, 497-256},
		},
		{
			Value:                  498,
			Expected: []byte{25, 1, 498-256},
		},
		{
			Value:                  499,
			Expected: []byte{25, 1, 499-256},
		},
		{
			Value:                  501,
			Expected: []byte{25, 1, 501-256},
		},
		{
			Value:                  502,
			Expected: []byte{25, 1, 502-256},
		},
		{
			Value:                  503,
			Expected: []byte{25, 1, 503-256},
		},
		{
			Value:                  504,
			Expected: []byte{25, 1, 504-256},
		},
		{
			Value:                  505,
			Expected: []byte{25, 1, 505-256},
		},
		{
			Value:                  506,
			Expected: []byte{25, 1, 506-256},
		},
		{
			Value:                  507,
			Expected: []byte{25, 1, 507-256},
		},
		{
			Value:                  508,
			Expected: []byte{25, 1, 508-256},
		},
		{
			Value:                  509,
			Expected: []byte{25, 1, 509-256},
		},
		{
			Value:                  510,
			Expected: []byte{25, 1, 510-256},
		},
		{
			Value:                  511,
			Expected: []byte{25, 1, 511-256},
		},









		{
			Value:                  512,
			Expected: []byte{25, 2, 512-(256*2)},
		},
		{
			Value:                  513,
			Expected: []byte{25, 2, 513-(256*2)},
		},
		{
			Value:                  514,
			Expected: []byte{25, 2, 514-(256*2)},
		},
		{
			Value:                  515,
			Expected: []byte{25, 2, 515-(256*2)},
		},
		{
			Value:                  516,
			Expected: []byte{25, 2, 516-(256*2)},
		},
		{
			Value:                  517,
			Expected: []byte{25, 2, 517-(256*2)},
		},



		{
			Value:                  765,
			Expected: []byte{25, 2, 765-(256*2)},
		},
		{
			Value:                  766,
			Expected: []byte{25, 2, 766-(256*2)},
		},
		{
			Value:                  767,
			Expected: []byte{25, 2, 767-(256*2)},
		},









		{
			Value:                  768,
			Expected: []byte{25, 3, 768-(256*3)},
		},
		{
			Value:                  769,
			Expected: []byte{25, 3, 769-(256*3)},
		},



		{
			Value:                  1022,
			Expected: []byte{25, 3, 1022-(256*3)},
		},
		{
			Value:                  1023,
			Expected: []byte{25, 3, 1023-(256*3)},
		},









		{
			Value:                  1024,
			Expected: []byte{25, 4, 1024-(256*4)},
		},
		{
			Value:                  1025,
			Expected: []byte{25, 4, 1025-(256*4)},
		},



		{
			Value:                  1278,
			Expected: []byte{25, 4, 1278-(256*4)},
		},
		{
			Value:                  1279,
			Expected: []byte{25, 4, 1279-(256*4)},
		},









		{
			Value:                  1280,
			Expected: []byte{25, 5, 1280-(256*5)},
		},
		{
			Value:                  1281,
			Expected: []byte{25, 5, 1281-(256*5)},
		},



		{
			Value:                  1534,
			Expected: []byte{25, 5, 1534-(256*5)},
		},
		{
			Value:                  1535,
			Expected: []byte{25, 5, 1535-(256*5)},
		},









		{
			Value:                  1536,
			Expected: []byte{25, 6, 1536-(256*6)},
		},
		{
			Value:                  1537,
			Expected: []byte{25, 6, 1537-(256*6)},
		},



		{
			Value:                  1790,
			Expected: []byte{25, 6, 1790-(256*6)},
		},
		{
			Value:                  1791,
			Expected: []byte{25, 6, 1791-(256*6)},
		},









		{
			Value:                  1792,
			Expected: []byte{25, 7, 1792-(256*7)},
		},
		{
			Value:                  1793,
			Expected: []byte{25, 7, 1793-(256*7)},
		},









		{
			Value:                    25600,
			Expected: []byte{25, 100, 25600-(256*100)},
		},









				
	}

	for testNumber, test := range tests {

		actual, err := int16s.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %d (0x%02X) (%#016b)", test.Value, test.Value, test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				t.Logf("VALUE: %d (0x%02X) (%#016b)", test.Value, test.Value, test.Value)
				continue
			}
		}
	}
}
