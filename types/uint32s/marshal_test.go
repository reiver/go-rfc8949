package uint32s_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-rfc8949/types/uint32s"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value uint32
		Expected []byte
	}{
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









		{
			Value:                    51200,
			Expected: []byte{25, 200, 51200-(256*200)},
		},









		{
			Value:                    64000,
			Expected: []byte{25, 250, 64000-(256*250)},
		},









		{
			Value:                    65020,
			Expected: []byte{25, 253, 65020-(256*253)},
		},
		{
			Value:                    65021,
			Expected: []byte{25, 253, 65021-(256*253)},
		},
		{
			Value:                    65022,
			Expected: []byte{25, 253, 65022-(256*253)},
		},
		{
			Value:                    65023,
			Expected: []byte{25, 253, 65023-(256*253)},
		},









		{
			Value:                    65024,
			Expected: []byte{25, 254, 65024-(256*254)},
		},
		{
			Value:                    65025,
			Expected: []byte{25, 254, 65025-(256*254)},
		},
		{
			Value:                    65026,
			Expected: []byte{25, 254, 65026-(256*254)},
		},
		{
			Value:                    65027,
			Expected: []byte{25, 254, 65027-(256*254)},
		},
		{
			Value:                    65028,
			Expected: []byte{25, 254, 65028-(256*254)},
		},
		{
			Value:                    65029,
			Expected: []byte{25, 254, 65029-(256*254)},
		},
		{
			Value:                    65030,
			Expected: []byte{25, 254, 65030-(256*254)},
		},
		{
			Value:                    65031,
			Expected: []byte{25, 254, 65031-(256*254)},
		},
		{
			Value:                    65032,
			Expected: []byte{25, 254, 65032-(256*254)},
		},
		{
			Value:                    65033,
			Expected: []byte{25, 254, 65033-(256*254)},
		},
		{
			Value:                    65034,
			Expected: []byte{25, 254, 65034-(256*254)},
		},
		{
			Value:                    65035,
			Expected: []byte{25, 254, 65035-(256*254)},
		},
		{
			Value:                    65036,
			Expected: []byte{25, 254, 65036-(256*254)},
		},
		{
			Value:                    65037,
			Expected: []byte{25, 254, 65037-(256*254)},
		},
		{
			Value:                    65038,
			Expected: []byte{25, 254, 65038-(256*254)},
		},
		{
			Value:                    65039,
			Expected: []byte{25, 254, 65039-(256*254)},
		},
		{
			Value:                    65040,
			Expected: []byte{25, 254, 65040-(256*254)},
		},
		{
			Value:                    65041,
			Expected: []byte{25, 254, 65041-(256*254)},
		},
		{
			Value:                    65042,
			Expected: []byte{25, 254, 65042-(256*254)},
		},
		{
			Value:                    65043,
			Expected: []byte{25, 254, 65043-(256*254)},
		},
		{
			Value:                    65044,
			Expected: []byte{25, 254, 65044-(256*254)},
		},
		{
			Value:                    65045,
			Expected: []byte{25, 254, 65045-(256*254)},
		},
		{
			Value:                    65046,
			Expected: []byte{25, 254, 65046-(256*254)},
		},
		{
			Value:                    65047,
			Expected: []byte{25, 254, 65047-(256*254)},
		},
		{
			Value:                    65048,
			Expected: []byte{25, 254, 65048-(256*254)},
		},
		{
			Value:                    65049,
			Expected: []byte{25, 254, 65049-(256*254)},
		},
		{
			Value:                    65050,
			Expected: []byte{25, 254, 65050-(256*254)},
		},
		{
			Value:                    65051,
			Expected: []byte{25, 254, 65051-(256*254)},
		},
		{
			Value:                    65052,
			Expected: []byte{25, 254, 65052-(256*254)},
		},
		{
			Value:                    65053,
			Expected: []byte{25, 254, 65053-(256*254)},
		},
		{
			Value:                    65054,
			Expected: []byte{25, 254, 65054-(256*254)},
		},
		{
			Value:                    65055,
			Expected: []byte{25, 254, 65055-(256*254)},
		},
		{
			Value:                    65056,
			Expected: []byte{25, 254, 65056-(256*254)},
		},
		{
			Value:                    65057,
			Expected: []byte{25, 254, 65057-(256*254)},
		},
		{
			Value:                    65058,
			Expected: []byte{25, 254, 65058-(256*254)},
		},
		{
			Value:                    65059,
			Expected: []byte{25, 254, 65059-(256*254)},
		},

		{
			Value:                    65260,
			Expected: []byte{25, 254, 65260-(256*254)},
		},
		{
			Value:                    65261,
			Expected: []byte{25, 254, 65261-(256*254)},
		},
		{
			Value:                    65262,
			Expected: []byte{25, 254, 65262-(256*254)},
		},
		{
			Value:                    65263,
			Expected: []byte{25, 254, 65263-(256*254)},
		},
		{
			Value:                    65264,
			Expected: []byte{25, 254, 65264-(256*254)},
		},
		{
			Value:                    65265,
			Expected: []byte{25, 254, 65265-(256*254)},
		},
		{
			Value:                    65266,
			Expected: []byte{25, 254, 65266-(256*254)},
		},
		{
			Value:                    65267,
			Expected: []byte{25, 254, 65267-(256*254)},
		},
		{
			Value:                    65268,
			Expected: []byte{25, 254, 65268-(256*254)},
		},
		{
			Value:                    65269,
			Expected: []byte{25, 254, 65269-(256*254)},
		},
		{
			Value:                    65270,
			Expected: []byte{25, 254, 65270-(256*254)},
		},
		{
			Value:                    65271,
			Expected: []byte{25, 254, 65271-(256*254)},
		},
		{
			Value:                    65272,
			Expected: []byte{25, 254, 65272-(256*254)},
		},
		{
			Value:                    65273,
			Expected: []byte{25, 254, 65273-(256*254)},
		},
		{
			Value:                    65274,
			Expected: []byte{25, 254, 65274-(256*254)},
		},
		{
			Value:                    65275,
			Expected: []byte{25, 254, 65275-(256*254)},
		},
		{
			Value:                    65276,
			Expected: []byte{25, 254, 65276-(256*254)},
		},
		{
			Value:                    65277,
			Expected: []byte{25, 254, 65277-(256*254)},
		},
		{
			Value:                    65278,
			Expected: []byte{25, 254, 65278-(256*254)},
		},
		{
			Value:                    65279,
			Expected: []byte{25, 254, 65279-(256*254)},
		},









		{
			Value:                    65280,
			Expected: []byte{25, 255, 65280-(256*255)},
		},
		{
			Value:                    65281,
			Expected: []byte{25, 255, 65281-(256*255)},
		},
		{
			Value:                    65282,
			Expected: []byte{25, 255, 65282-(256*255)},
		},
		{
			Value:                    65283,
			Expected: []byte{25, 255, 65283-(256*255)},
		},
		{
			Value:                    65284,
			Expected: []byte{25, 255, 65284-(256*255)},
		},
		{
			Value:                    65285,
			Expected: []byte{25, 255, 65285-(256*255)},
		},
		{
			Value:                    65286,
			Expected: []byte{25, 255, 65286-(256*255)},
		},
		{
			Value:                    65287,
			Expected: []byte{25, 255, 65287-(256*255)},
		},
		{
			Value:                    65288,
			Expected: []byte{25, 255, 65288-(256*255)},
		},
		{
			Value:                    65289,
			Expected: []byte{25, 255, 65289-(256*255)},
		},
		{
			Value:                    65290,
			Expected: []byte{25, 255, 65290-(256*255)},
		},
		{
			Value:                    65291,
			Expected: []byte{25, 255, 65291-(256*255)},
		},
		{
			Value:                    65292,
			Expected: []byte{25, 255, 65292-(256*255)},
		},
		{
			Value:                    65293,
			Expected: []byte{25, 255, 65293-(256*255)},
		},
		{
			Value:                    65294,
			Expected: []byte{25, 255, 65294-(256*255)},
		},
		{
			Value:                    65295,
			Expected: []byte{25, 255, 65295-(256*255)},
		},
		{
			Value:                    65296,
			Expected: []byte{25, 255, 65296-(256*255)},
		},
		{
			Value:                    65297,
			Expected: []byte{25, 255, 65297-(256*255)},
		},
		{
			Value:                    65298,
			Expected: []byte{25, 255, 65298-(256*255)},
		},
		{
			Value:                    65299,
			Expected: []byte{25, 255, 65299-(256*255)},
		},
		{
			Value:                    65300,
			Expected: []byte{25, 255, 65300-(256*255)},
		},



		{
			Value:                    65310,
			Expected: []byte{25, 255, 65310-(256*255)},
		},



		{
			Value:                    65320,
			Expected: []byte{25, 255, 65320-(256*255)},
		},



		{
			Value:                    65330,
			Expected: []byte{25, 255, 65330-(256*255)},
		},



		{
			Value:                    65340,
			Expected: []byte{25, 255, 65340-(256*255)},
		},



		{
			Value:                    65350,
			Expected: []byte{25, 255, 65350-(256*255)},
		},



		{
			Value:                    65360,
			Expected: []byte{25, 255, 65360-(256*255)},
		},



		{
			Value:                    65370,
			Expected: []byte{25, 255, 65370-(256*255)},
		},



		{
			Value:                    65380,
			Expected: []byte{25, 255, 65380-(256*255)},
		},



		{
			Value:                    65390,
			Expected: []byte{25, 255, 65390-(256*255)},
		},




		{
			Value:                    65400,
			Expected: []byte{25, 255, 65400-(256*255)},
		},
		{
			Value:                    65401,
			Expected: []byte{25, 255, 65401-(256*255)},
		},
		{
			Value:                    65402,
			Expected: []byte{25, 255, 65402-(256*255)},
		},
		{
			Value:                    65403,
			Expected: []byte{25, 255, 65403-(256*255)},
		},
		{
			Value:                    65404,
			Expected: []byte{25, 255, 65404-(256*255)},
		},
		{
			Value:                    65405,
			Expected: []byte{25, 255, 65405-(256*255)},
		},
		{
			Value:                    65406,
			Expected: []byte{25, 255, 65406-(256*255)},
		},
		{
			Value:                    65407,
			Expected: []byte{25, 255, 65407-(256*255)},
		},
		{
			Value:                    65408,
			Expected: []byte{25, 255, 65408-(256*255)},
		},
		{
			Value:                    65409,
			Expected: []byte{25, 255, 65409-(256*255)},
		},
		{
			Value:                    65410,
			Expected: []byte{25, 255, 65410-(256*255)},
		},
		{
			Value:                    65411,
			Expected: []byte{25, 255, 65411-(256*255)},
		},
		{
			Value:                    65412,
			Expected: []byte{25, 255, 65412-(256*255)},
		},
		{
			Value:                    65413,
			Expected: []byte{25, 255, 65413-(256*255)},
		},
		{
			Value:                    65414,
			Expected: []byte{25, 255, 65414-(256*255)},
		},
		{
			Value:                    65415,
			Expected: []byte{25, 255, 65415-(256*255)},
		},
		{
			Value:                    65416,
			Expected: []byte{25, 255, 65416-(256*255)},
		},
		{
			Value:                    65417,
			Expected: []byte{25, 255, 65417-(256*255)},
		},
		{
			Value:                    65418,
			Expected: []byte{25, 255, 65418-(256*255)},
		},
		{
			Value:                    65419,
			Expected: []byte{25, 255, 65419-(256*255)},
		},
		{
			Value:                    65420,
			Expected: []byte{25, 255, 65420-(256*255)},
		},
		{
			Value:                    65421,
			Expected: []byte{25, 255, 65421-(256*255)},
		},
		{
			Value:                    65422,
			Expected: []byte{25, 255, 65422-(256*255)},
		},
		{
			Value:                    65423,
			Expected: []byte{25, 255, 65423-(256*255)},
		},
		{
			Value:                    65424,
			Expected: []byte{25, 255, 65424-(256*255)},
		},
		{
			Value:                    65425,
			Expected: []byte{25, 255, 65425-(256*255)},
		},
		{
			Value:                    65426,
			Expected: []byte{25, 255, 65426-(256*255)},
		},
		{
			Value:                    65427,
			Expected: []byte{25, 255, 65427-(256*255)},
		},
		{
			Value:                    65428,
			Expected: []byte{25, 255, 65428-(256*255)},
		},
		{
			Value:                    65429,
			Expected: []byte{25, 255, 65429-(256*255)},
		},
		{
			Value:                    65430,
			Expected: []byte{25, 255, 65430-(256*255)},
		},
		{
			Value:                    65431,
			Expected: []byte{25, 255, 65431-(256*255)},
		},
		{
			Value:                    65432,
			Expected: []byte{25, 255, 65432-(256*255)},
		},
		{
			Value:                    65433,
			Expected: []byte{25, 255, 65433-(256*255)},
		},
		{
			Value:                    65434,
			Expected: []byte{25, 255, 65434-(256*255)},
		},
		{
			Value:                    65435,
			Expected: []byte{25, 255, 65435-(256*255)},
		},
		{
			Value:                    65436,
			Expected: []byte{25, 255, 65436-(256*255)},
		},
		{
			Value:                    65437,
			Expected: []byte{25, 255, 65437-(256*255)},
		},
		{
			Value:                    65438,
			Expected: []byte{25, 255, 65438-(256*255)},
		},
		{
			Value:                    65439,
			Expected: []byte{25, 255, 65439-(256*255)},
		},
		{
			Value:                    65440,
			Expected: []byte{25, 255, 65440-(256*255)},
		},
		{
			Value:                    65441,
			Expected: []byte{25, 255, 65441-(256*255)},
		},
		{
			Value:                    65442,
			Expected: []byte{25, 255, 65442-(256*255)},
		},
		{
			Value:                    65443,
			Expected: []byte{25, 255, 65443-(256*255)},
		},
		{
			Value:                    65444,
			Expected: []byte{25, 255, 65444-(256*255)},
		},
		{
			Value:                    65445,
			Expected: []byte{25, 255, 65445-(256*255)},
		},
		{
			Value:                    65446,
			Expected: []byte{25, 255, 65446-(256*255)},
		},
		{
			Value:                    65447,
			Expected: []byte{25, 255, 65447-(256*255)},
		},
		{
			Value:                    65448,
			Expected: []byte{25, 255, 65448-(256*255)},
		},
		{
			Value:                    65449,
			Expected: []byte{25, 255, 65449-(256*255)},
		},
		{
			Value:                    65450,
			Expected: []byte{25, 255, 65450-(256*255)},
		},
		{
			Value:                    65451,
			Expected: []byte{25, 255, 65451-(256*255)},
		},
		{
			Value:                    65452,
			Expected: []byte{25, 255, 65452-(256*255)},
		},
		{
			Value:                    65453,
			Expected: []byte{25, 255, 65453-(256*255)},
		},
		{
			Value:                    65454,
			Expected: []byte{25, 255, 65454-(256*255)},
		},
		{
			Value:                    65455,
			Expected: []byte{25, 255, 65455-(256*255)},
		},
		{
			Value:                    65456,
			Expected: []byte{25, 255, 65456-(256*255)},
		},
		{
			Value:                    65457,
			Expected: []byte{25, 255, 65457-(256*255)},
		},
		{
			Value:                    65458,
			Expected: []byte{25, 255, 65458-(256*255)},
		},
		{
			Value:                    65459,
			Expected: []byte{25, 255, 65459-(256*255)},
		},
		{
			Value:                    65460,
			Expected: []byte{25, 255, 65460-(256*255)},
		},
		{
			Value:                    65461,
			Expected: []byte{25, 255, 65461-(256*255)},
		},
		{
			Value:                    65462,
			Expected: []byte{25, 255, 65462-(256*255)},
		},
		{
			Value:                    65463,
			Expected: []byte{25, 255, 65463-(256*255)},
		},
		{
			Value:                    65464,
			Expected: []byte{25, 255, 65464-(256*255)},
		},
		{
			Value:                    65465,
			Expected: []byte{25, 255, 65465-(256*255)},
		},
		{
			Value:                    65466,
			Expected: []byte{25, 255, 65466-(256*255)},
		},
		{
			Value:                    65467,
			Expected: []byte{25, 255, 65467-(256*255)},
		},
		{
			Value:                    65468,
			Expected: []byte{25, 255, 65468-(256*255)},
		},
		{
			Value:                    65469,
			Expected: []byte{25, 255, 65469-(256*255)},
		},
		{
			Value:                    65470,
			Expected: []byte{25, 255, 65470-(256*255)},
		},
		{
			Value:                    65471,
			Expected: []byte{25, 255, 65471-(256*255)},
		},
		{
			Value:                    65472,
			Expected: []byte{25, 255, 65472-(256*255)},
		},
		{
			Value:                    65473,
			Expected: []byte{25, 255, 65473-(256*255)},
		},
		{
			Value:                    65474,
			Expected: []byte{25, 255, 65474-(256*255)},
		},
		{
			Value:                    65475,
			Expected: []byte{25, 255, 65475-(256*255)},
		},
		{
			Value:                    65476,
			Expected: []byte{25, 255, 65476-(256*255)},
		},
		{
			Value:                    65477,
			Expected: []byte{25, 255, 65477-(256*255)},
		},
		{
			Value:                    65478,
			Expected: []byte{25, 255, 65478-(256*255)},
		},
		{
			Value:                    65479,
			Expected: []byte{25, 255, 65479-(256*255)},
		},
		{
			Value:                    65480,
			Expected: []byte{25, 255, 65480-(256*255)},
		},
		{
			Value:                    65481,
			Expected: []byte{25, 255, 65481-(256*255)},
		},
		{
			Value:                    65482,
			Expected: []byte{25, 255, 65482-(256*255)},
		},
		{
			Value:                    65483,
			Expected: []byte{25, 255, 65483-(256*255)},
		},
		{
			Value:                    65484,
			Expected: []byte{25, 255, 65484-(256*255)},
		},
		{
			Value:                    65485,
			Expected: []byte{25, 255, 65485-(256*255)},
		},
		{
			Value:                    65486,
			Expected: []byte{25, 255, 65486-(256*255)},
		},
		{
			Value:                    65487,
			Expected: []byte{25, 255, 65487-(256*255)},
		},
		{
			Value:                    65488,
			Expected: []byte{25, 255, 65488-(256*255)},
		},
		{
			Value:                    65489,
			Expected: []byte{25, 255, 65489-(256*255)},
		},
		{
			Value:                    65490,
			Expected: []byte{25, 255, 65490-(256*255)},
		},
		{
			Value:                    65491,
			Expected: []byte{25, 255, 65491-(256*255)},
		},
		{
			Value:                    65492,
			Expected: []byte{25, 255, 65492-(256*255)},
		},
		{
			Value:                    65493,
			Expected: []byte{25, 255, 65493-(256*255)},
		},
		{
			Value:                    65494,
			Expected: []byte{25, 255, 65494-(256*255)},
		},
		{
			Value:                    65495,
			Expected: []byte{25, 255, 65495-(256*255)},
		},
		{
			Value:                    65496,
			Expected: []byte{25, 255, 65496-(256*255)},
		},
		{
			Value:                    65497,
			Expected: []byte{25, 255, 65497-(256*255)},
		},
		{
			Value:                    65498,
			Expected: []byte{25, 255, 65498-(256*255)},
		},
		{
			Value:                    65499,
			Expected: []byte{25, 255, 65499-(256*255)},
		},
		{
			Value:                    65500,
			Expected: []byte{25, 255, 65500-(256*255)},
		},
		{
			Value:                    65501,
			Expected: []byte{25, 255, 65501-(256*255)},
		},
		{
			Value:                    65502,
			Expected: []byte{25, 255, 65502-(256*255)},
		},
		{
			Value:                    65503,
			Expected: []byte{25, 255, 65503-(256*255)},
		},
		{
			Value:                    65504,
			Expected: []byte{25, 255, 65504-(256*255)},
		},
		{
			Value:                    65505,
			Expected: []byte{25, 255, 65505-(256*255)},
		},
		{
			Value:                    65506,
			Expected: []byte{25, 255, 65506-(256*255)},
		},
		{
			Value:                    65507,
			Expected: []byte{25, 255, 65507-(256*255)},
		},
		{
			Value:                    65508,
			Expected: []byte{25, 255, 65508-(256*255)},
		},

		{
			Value:                    65509,
			Expected: []byte{25, 255, 65509-(256*255)},
		},
		{
			Value:                    65510,
			Expected: []byte{25, 255, 65510-(256*255)},
		},
		{
			Value:                    65511,
			Expected: []byte{25, 255, 65511-(256*255)},
		},
		{
			Value:                    65512,
			Expected: []byte{25, 255, 65512-(256*255)},
		},
		{
			Value:                    65513,
			Expected: []byte{25, 255, 65513-(256*255)},
		},
		{
			Value:                    65514,
			Expected: []byte{25, 255, 65514-(256*255)},
		},
		{
			Value:                    65515,
			Expected: []byte{25, 255, 65515-(256*255)},
		},
		{
			Value:                    65516,
			Expected: []byte{25, 255, 65516-(256*255)},
		},
		{
			Value:                    65517,
			Expected: []byte{25, 255, 65517-(256*255)},
		},
		{
			Value:                    65518,
			Expected: []byte{25, 255, 65518-(256*255)},
		},
		{
			Value:                    65519,
			Expected: []byte{25, 255, 65519-(256*255)},
		},
		{
			Value:                    65520,
			Expected: []byte{25, 255, 65520-(256*255)},
		},
		{
			Value:                    65521,
			Expected: []byte{25, 255, 65521-(256*255)},
		},
		{
			Value:                    65522,
			Expected: []byte{25, 255, 65522-(256*255)},
		},
		{
			Value:                    65523,
			Expected: []byte{25, 255, 65523-(256*255)},
		},
		{
			Value:                    65524,
			Expected: []byte{25, 255, 65524-(256*255)},
		},
		{
			Value:                    65525,
			Expected: []byte{25, 255, 65525-(256*255)},
		},
		{
			Value:                    65526,
			Expected: []byte{25, 255, 65526-(256*255)},
		},
		{
			Value:                    65527,
			Expected: []byte{25, 255, 65527-(256*255)},
		},
		{
			Value:                    65528,
			Expected: []byte{25, 255, 65528-(256*255)},
		},
		{
			Value:                    65529,
			Expected: []byte{25, 255, 65529-(256*255)},
		},
		{
			Value:                    65530,
			Expected: []byte{25, 255, 65530-(256*255)},
		},
		{
			Value:                    65531,
			Expected: []byte{25, 255, 65531-(256*255)},
		},
		{
			Value:                    65532,
			Expected: []byte{25, 255, 65532-(256*255)},
		},
		{
			Value:                    65533,
			Expected: []byte{25, 255, 65533-(256*255)},
		},
		{
			Value:                    65534,
			Expected: []byte{25, 255, 65534-(256*255)},
		},
		{
			Value:                    65535,
			Expected: []byte{25, 255, 65535-(256*255)},
		},









		{
			Value:                        65536,
			Expected: []byte{26, 0, 1, 0, 65536-(256*256)},
		},
		{
			Value:                        65537,
			Expected: []byte{26, 0, 1, 0, 65537-(256*256)},
		},
		{
			Value:                        65538,
			Expected: []byte{26, 0, 1, 0, 65538-(256*256)},
		},
		{
			Value:                        65539,
			Expected: []byte{26, 0, 1, 0, 65539-(256*256)},
		},
		{
			Value:                        65540,
			Expected: []byte{26, 0, 1, 0, 65540-(256*256)},
		},
		{
			Value:                        65541,
			Expected: []byte{26, 0, 1, 0, 65541-(256*256)},
		},
		{
			Value:                        65542,
			Expected: []byte{26, 0, 1, 0, 65542-(256*256)},
		},
		{
			Value:                        65543,
			Expected: []byte{26, 0, 1, 0, 65543-(256*256)},
		},
		{
			Value:                        65544,
			Expected: []byte{26, 0, 1, 0, 65544-(256*256)},
		},
		{
			Value:                        65545,
			Expected: []byte{26, 0, 1, 0, 65545-(256*256)},
		},
		{
			Value:                        65546,
			Expected: []byte{26, 0, 1, 0, 65546-(256*256)},
		},
		{
			Value:                        65547,
			Expected: []byte{26, 0, 1, 0, 65547-(256*256)},
		},
		{
			Value:                        65548,
			Expected: []byte{26, 0, 1, 0, 65548-(256*256)},
		},
		{
			Value:                        65549,
			Expected: []byte{26, 0, 1, 0, 65549-(256*256)},
		},
		{
			Value:                        65550,
			Expected: []byte{26, 0, 1, 0, 65550-(256*256)},
		},
		{
			Value:                        65551,
			Expected: []byte{26, 0, 1, 0, 65551-(256*256)},
		},
		{
			Value:                        65552,
			Expected: []byte{26, 0, 1, 0, 65552-(256*256)},
		},
		{
			Value:                        65553,
			Expected: []byte{26, 0, 1, 0, 65553-(256*256)},
		},
		{
			Value:                        65554,
			Expected: []byte{26, 0, 1, 0, 65554-(256*256)},
		},
		{
			Value:                        65555,
			Expected: []byte{26, 0, 1, 0, 65555-(256*256)},
		},
		{
			Value:                        65556,
			Expected: []byte{26, 0, 1, 0, 65556-(256*256)},
		},
		{
			Value:                        65557,
			Expected: []byte{26, 0, 1, 0, 65557-(256*256)},
		},
		{
			Value:                        65558,
			Expected: []byte{26, 0, 1, 0, 65558-(256*256)},
		},
		{
			Value:                        65559,
			Expected: []byte{26, 0, 1, 0, 65559-(256*256)},
		},
		{
			Value:                        65560,
			Expected: []byte{26, 0, 1, 0, 65560-(256*256)},
		},
		{
			Value:                        65561,
			Expected: []byte{26, 0, 1, 0, 65561-(256*256)},
		},
		{
			Value:                        65562,
			Expected: []byte{26, 0, 1, 0, 65562-(256*256)},
		},
		{
			Value:                        65563,
			Expected: []byte{26, 0, 1, 0, 65563-(256*256)},
		},
		{
			Value:                        65564,
			Expected: []byte{26, 0, 1, 0, 65564-(256*256)},
		},
		{
			Value:                        65565,
			Expected: []byte{26, 0, 1, 0, 65565-(256*256)},
		},
		{
			Value:                        65566,
			Expected: []byte{26, 0, 1, 0, 65566-(256*256)},
		},
		{
			Value:                        65567,
			Expected: []byte{26, 0, 1, 0, 65567-(256*256)},
		},
		{
			Value:                        65568,
			Expected: []byte{26, 0, 1, 0, 65568-(256*256)},
		},
		{
			Value:                        65569,
			Expected: []byte{26, 0, 1, 0, 65569-(256*256)},
		},
		{
			Value:                        65570,
			Expected: []byte{26, 0, 1, 0, 65570-(256*256)},
		},
		{
			Value:                        65571,
			Expected: []byte{26, 0, 1, 0, 65571-(256*256)},
		},
		{
			Value:                        65572,
			Expected: []byte{26, 0, 1, 0, 65572-(256*256)},
		},
		{
			Value:                        65573,
			Expected: []byte{26, 0, 1, 0, 65573-(256*256)},
		},
		{
			Value:                        65574,
			Expected: []byte{26, 0, 1, 0, 65574-(256*256)},
		},
		{
			Value:                        65575,
			Expected: []byte{26, 0, 1, 0, 65575-(256*256)},
		},
		{
			Value:                        65576,
			Expected: []byte{26, 0, 1, 0, 65576-(256*256)},
		},
		{
			Value:                        65577,
			Expected: []byte{26, 0, 1, 0, 65577-(256*256)},
		},
		{
			Value:                        65578,
			Expected: []byte{26, 0, 1, 0, 65578-(256*256)},
		},
		{
			Value:                        65579,
			Expected: []byte{26, 0, 1, 0, 65579-(256*256)},
		},
		{
			Value:                        65580,
			Expected: []byte{26, 0, 1, 0, 65580-(256*256)},
		},
		{
			Value:                        65581,
			Expected: []byte{26, 0, 1, 0, 65581-(256*256)},
		},
		{
			Value:                        65582,
			Expected: []byte{26, 0, 1, 0, 65582-(256*256)},
		},
		{
			Value:                        65583,
			Expected: []byte{26, 0, 1, 0, 65583-(256*256)},
		},
		{
			Value:                        65584,
			Expected: []byte{26, 0, 1, 0, 65584-(256*256)},
		},
		{
			Value:                        65585,
			Expected: []byte{26, 0, 1, 0, 65585-(256*256)},
		},
		{
			Value:                        65586,
			Expected: []byte{26, 0, 1, 0, 65586-(256*256)},
		},
		{
			Value:                        65587,
			Expected: []byte{26, 0, 1, 0, 65587-(256*256)},
		},
		{
			Value:                        65588,
			Expected: []byte{26, 0, 1, 0, 65588-(256*256)},
		},
		{
			Value:                        65589,
			Expected: []byte{26, 0, 1, 0, 65589-(256*256)},
		},
		{
			Value:                        65590,
			Expected: []byte{26, 0, 1, 0, 65590-(256*256)},
		},
		{
			Value:                        65591,
			Expected: []byte{26, 0, 1, 0, 65591-(256*256)},
		},
		{
			Value:                        65592,
			Expected: []byte{26, 0, 1, 0, 65592-(256*256)},
		},
		{
			Value:                        65593,
			Expected: []byte{26, 0, 1, 0, 65593-(256*256)},
		},
		{
			Value:                        65594,
			Expected: []byte{26, 0, 1, 0, 65594-(256*256)},
		},
		{
			Value:                        65595,
			Expected: []byte{26, 0, 1, 0, 65595-(256*256)},
		},
		{
			Value:                        65596,
			Expected: []byte{26, 0, 1, 0, 65596-(256*256)},
		},
		{
			Value:                        65597,
			Expected: []byte{26, 0, 1, 0, 65597-(256*256)},
		},
		{
			Value:                        65598,
			Expected: []byte{26, 0, 1, 0, 65598-(256*256)},
		},
		{
			Value:                        65599,
			Expected: []byte{26, 0, 1, 0, 65599-(256*256)},
		},
		{
			Value:                        65600,
			Expected: []byte{26, 0, 1, 0, 65600-(256*256)},
		},



		{
			Value:                        65610,
			Expected: []byte{26, 0, 1, 0, 65610-(256*256)},
		},



		{
			Value:                        65620,
			Expected: []byte{26, 0, 1, 0, 65620-(256*256)},
		},



		{
			Value:                        65630,
			Expected: []byte{26, 0, 1, 0, 65630-(256*256)},
		},



		{
			Value:                        65640,
			Expected: []byte{26, 0, 1, 0, 65640-(256*256)},
		},



		{
			Value:                        65650,
			Expected: []byte{26, 0, 1, 0, 65650-(256*256)},
		},



		{
			Value:                        65660,
			Expected: []byte{26, 0, 1, 0, 65660-(256*256)},
		},



		{
			Value:                        65670,
			Expected: []byte{26, 0, 1, 0, 65670-(256*256)},
		},
		{
			Value:                        65671,
			Expected: []byte{26, 0, 1, 0, 65671-(256*256)},
		},
		{
			Value:                        65672,
			Expected: []byte{26, 0, 1, 0, 65672-(256*256)},
		},
		{
			Value:                        65673,
			Expected: []byte{26, 0, 1, 0, 65673-(256*256)},
		},
		{
			Value:                        65674,
			Expected: []byte{26, 0, 1, 0, 65674-(256*256)},
		},
		{
			Value:                        65675,
			Expected: []byte{26, 0, 1, 0, 65675-(256*256)},
		},
		{
			Value:                        65676,
			Expected: []byte{26, 0, 1, 0, 65676-(256*256)},
		},
		{
			Value:                        65677,
			Expected: []byte{26, 0, 1, 0, 65677-(256*256)},
		},
		{
			Value:                        65678,
			Expected: []byte{26, 0, 1, 0, 65678-(256*256)},
		},
		{
			Value:                        65679,
			Expected: []byte{26, 0, 1, 0, 65679-(256*256)},
		},
		{
			Value:                        65680,
			Expected: []byte{26, 0, 1, 0, 65680-(256*256)},
		},
		{
			Value:                        65681,
			Expected: []byte{26, 0, 1, 0, 65681-(256*256)},
		},
		{
			Value:                        65682,
			Expected: []byte{26, 0, 1, 0, 65682-(256*256)},
		},
		{
			Value:                        65683,
			Expected: []byte{26, 0, 1, 0, 65683-(256*256)},
		},
		{
			Value:                        65684,
			Expected: []byte{26, 0, 1, 0, 65684-(256*256)},
		},
		{
			Value:                        65685,
			Expected: []byte{26, 0, 1, 0, 65685-(256*256)},
		},
		{
			Value:                        65686,
			Expected: []byte{26, 0, 1, 0, 65686-(256*256)},
		},
		{
			Value:                        65687,
			Expected: []byte{26, 0, 1, 0, 65687-(256*256)},
		},
		{
			Value:                        65688,
			Expected: []byte{26, 0, 1, 0, 65688-(256*256)},
		},
		{
			Value:                        65789,
			Expected: []byte{26, 0, 1, 0, 65789-(256*256)},
		},
		{
			Value:                        65790,
			Expected: []byte{26, 0, 1, 0, 65790-(256*256)},
		},
		{
			Value:                        65791,
			Expected: []byte{26, 0, 1, 0, 65791-(256*256)},
		},









		{
			Value:                        65792,
			Expected: []byte{26, 0, 1, 1, 65792-(256*257)},
		},
		{
			Value:                        65793,
			Expected: []byte{26, 0, 1, 1, 65793-(256*257)},
		},
		{
			Value:                        65794,
			Expected: []byte{26, 0, 1, 1, 65794-(256*257)},
		},









		{
			Value:                         76800,
			Expected: []byte{26, 0, 1, 44, 76800-(256*300)},
		},
		{
			Value:                         76801,
			Expected: []byte{26, 0, 1, 44, 76801-(256*300)},
		},









		{
			Value:                          102400,
			Expected: []byte{26, 0, 1, 144, 102400-(256*400)},
		},
		{
			Value:                          102401,
			Expected: []byte{26, 0, 1, 144, 102401-(256*400)},
		},









		{
			Value:                          128000,
			Expected: []byte{26, 0, 1, 244, 128000-(256*500)},
		},
		{
			Value:                          128001,
			Expected: []byte{26, 0, 1, 244, 128001-(256*500)},
		},









		{
			Value:         256000,
			Expected: []byte{26,
				byte( (256000 & 0xff_00_00_00) >> (8*3) ),
				byte( (256000 & 0x00_ff_00_00) >> (8*2) ),
				byte( (256000 & 0x00_00_ff_00) >> (8*1) ),
				byte( (256000 & 0x00_00_00_ff)          ),
			},
		},









		{
			Value:         512000,
			Expected: []byte{26,
				byte( (512000 & 0xff_00_00_00) >> (8*3) ),
				byte( (512000 & 0x00_ff_00_00) >> (8*2) ),
				byte( (512000 & 0x00_00_ff_00) >> (8*1) ),
				byte( (512000 & 0x00_00_00_ff)          ),
			},
		},









		{
			Value:         2560000,
			Expected: []byte{26,
				byte( (2560000 & 0xff_00_00_00) >> (8*3) ),
				byte( (2560000 & 0x00_ff_00_00) >> (8*2) ),
				byte( (2560000 & 0x00_00_ff_00) >> (8*1) ),
				byte( (2560000 & 0x00_00_00_ff)          ),
			},
		},









		{
			Value:         25600000,
			Expected: []byte{26,
				byte( (25600000 & 0xff_00_00_00) >> (8*3) ),
				byte( (25600000 & 0x00_ff_00_00) >> (8*2) ),
				byte( (25600000 & 0x00_00_ff_00) >> (8*1) ),
				byte( (25600000 & 0x00_00_00_ff)          ),
			},
		},









		{
			Value:         256000000,
			Expected: []byte{26,
				byte( (256000000 & 0xff_00_00_00) >> (8*3) ),
				byte( (256000000 & 0x00_ff_00_00) >> (8*2) ),
				byte( (256000000 & 0x00_00_ff_00) >> (8*1) ),
				byte( (256000000 & 0x00_00_00_ff)          ),
			},
		},









		{
			Value:         2560000000,
			Expected: []byte{26,
				byte( (2560000000 & 0xff_00_00_00) >> (8*3) ),
				byte( (2560000000 & 0x00_ff_00_00) >> (8*2) ),
				byte( (2560000000 & 0x00_00_ff_00) >> (8*1) ),
				byte( (2560000000 & 0x00_00_00_ff)          ),
			},
		},









		{
			Value:         3210000000,
			Expected: []byte{26,
				byte( (3210000000 & 0xff_00_00_00) >> (8*3) ),
				byte( (3210000000 & 0x00_ff_00_00) >> (8*2) ),
				byte( (3210000000 & 0x00_00_ff_00) >> (8*1) ),
				byte( (3210000000 & 0x00_00_00_ff)          ),
			},
		},









		{
			Value:         4293157860,
			Expected: []byte{26,
				byte( (4293157860 & 0xff_00_00_00) >> (8*3) ),
				byte( (4293157860 & 0x00_ff_00_00) >> (8*2) ),
				byte( (4293157860 & 0x00_00_ff_00) >> (8*1) ),
				byte( (4293157860 & 0x00_00_00_ff)          ),
			},
		},



		{
			Value:         4294967295,
			Expected: []byte{26,
				byte( (4294967295 & 0xff_00_00_00) >> (8*3) ),
				byte( (4294967295 & 0x00_ff_00_00) >> (8*2) ),
				byte( (4294967295 & 0x00_00_ff_00) >> (8*1) ),
				byte( (4294967295 & 0x00_00_00_ff)          ),
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := uint32s.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %d (0x%02X) (%#032b)", test.Value, test.Value, test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				t.Logf("VALUE: %d (0x%02X) (%#032b)", test.Value, test.Value, test.Value)
				continue
			}
		}
	}
}
