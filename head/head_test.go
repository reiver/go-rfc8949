package head_test

import (
	"testing"

	"github.com/reiver/go-rfc8949/head"
	"github.com/reiver/go-rfc8949/majortype"
)

func TestDecode(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected head.Head
		ExpectedMinimal bool
	}{
		// RFC 8949 Appendix A examples.
		{
			// 0
			Data:     []byte{0x00},
			Expected: head.Head{MajorType: majortype.UnsignedInteger, AdditionalInformation: 0, Argument: 0, Length: 1},
			ExpectedMinimal: true,
		},
		{
			// 23
			Data:     []byte{0x17},
			Expected: head.Head{MajorType: majortype.UnsignedInteger, AdditionalInformation: 23, Argument: 23, Length: 1},
			ExpectedMinimal: true,
		},
		{
			// 24
			Data:     []byte{0x18, 0x18},
			Expected: head.Head{MajorType: majortype.UnsignedInteger, AdditionalInformation: 24, Argument: 24, Length: 2},
			ExpectedMinimal: true,
		},
		{
			// 1000
			Data:     []byte{0x19, 0x03, 0xe8},
			Expected: head.Head{MajorType: majortype.UnsignedInteger, AdditionalInformation: 25, Argument: 1000, Length: 3},
			ExpectedMinimal: true,
		},
		{
			// 1000000
			Data:     []byte{0x1a, 0x00, 0x0f, 0x42, 0x40},
			Expected: head.Head{MajorType: majortype.UnsignedInteger, AdditionalInformation: 26, Argument: 1000000, Length: 5},
			ExpectedMinimal: true,
		},
		{
			// 18446744073709551615
			Data:     []byte{0x1b, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			Expected: head.Head{MajorType: majortype.UnsignedInteger, AdditionalInformation: 27, Argument: 18446744073709551615, Length: 9},
			ExpectedMinimal: true,
		},
		{
			// -1000
			Data:     []byte{0x39, 0x03, 0xe7},
			Expected: head.Head{MajorType: majortype.NegativeInteger, AdditionalInformation: 25, Argument: 999, Length: 3},
			ExpectedMinimal: true,
		},
		{
			// "IETF" (only the head)
			Data:     []byte{0x64, 0x49, 0x45, 0x54, 0x46},
			Expected: head.Head{MajorType: majortype.TextString, AdditionalInformation: 4, Argument: 4, Length: 1},
			ExpectedMinimal: true,
		},
		{
			// 1(1363896240) (only the head)
			Data:     []byte{0xc1, 0x1a, 0x51, 0x4b, 0x67, 0xb0},
			Expected: head.Head{MajorType: majortype.Tagged, AdditionalInformation: 1, Argument: 1, Length: 1},
			ExpectedMinimal: true,
		},
		{
			// 1.1 (only the head)
			Data:     []byte{0xfb, 0x3f, 0xf1, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a},
			Expected: head.Head{MajorType: majortype.MajorType7, AdditionalInformation: 27, Argument: 0x3ff199999999999a, Length: 9},
			ExpectedMinimal: true,
		},
		{
			// null
			Data:     []byte{0xf6},
			Expected: head.Head{MajorType: majortype.MajorType7, AdditionalInformation: 22, Argument: 22, Length: 1},
			ExpectedMinimal: true,
		},
		{
			// (_ h'0102', h'030405') — an indefinite-length byte string (only the head)
			Data:     []byte{0x5f, 0x42, 0x01, 0x02, 0x43, 0x03, 0x04, 0x05, 0xff},
			Expected: head.Head{MajorType: majortype.ByteString, AdditionalInformation: 31, Argument: 0, Length: 1},
			ExpectedMinimal: true,
		},



		// Not minimal (well-formed, but not the preferred serialization).
		{
			// 0, in 2 bytes
			Data:     []byte{0x18, 0x00},
			Expected: head.Head{MajorType: majortype.UnsignedInteger, AdditionalInformation: 24, Argument: 0, Length: 2},
			ExpectedMinimal: false,
		},
		{
			// 255, in 3 bytes
			Data:     []byte{0x19, 0x00, 0xff},
			Expected: head.Head{MajorType: majortype.UnsignedInteger, AdditionalInformation: 25, Argument: 255, Length: 3},
			ExpectedMinimal: false,
		},
		{
			// 65535, in 5 bytes
			Data:     []byte{0x1a, 0x00, 0x00, 0xff, 0xff},
			Expected: head.Head{MajorType: majortype.UnsignedInteger, AdditionalInformation: 26, Argument: 65535, Length: 5},
			ExpectedMinimal: false,
		},
		{
			// 4294967295, in 9 bytes
			Data:     []byte{0x1b, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff},
			Expected: head.Head{MajorType: majortype.UnsignedInteger, AdditionalInformation: 27, Argument: 4294967295, Length: 9},
			ExpectedMinimal: false,
		},
		{
			// simple value 20 (false), but in 2 bytes
			Data:     []byte{0xf8, 0x14},
			Expected: head.Head{MajorType: majortype.MajorType7, AdditionalInformation: 24, Argument: 20, Length: 2},
			ExpectedMinimal: false,
		},
	}

	for testNumber, test := range tests {

		actual, err := head.Decode(test.Data)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA: %#v", test.Data)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual head is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("DATA: %#v", test.Data)
				continue
			}
		}

		{
			expected := test.ExpectedMinimal
			actual   := actual.IsMinimal()

			if expected != actual {
				t.Errorf("For test #%d, the actual is-minimal is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("DATA: %#v", test.Data)
				continue
			}
		}
	}
}

func TestDecode_errors(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected error
	}{
		{
			Data:     nil,
			Expected: head.ErrEmpty,
		},
		{
			Data:     []byte{},
			Expected: head.ErrEmpty,
		},



		{
			Data:     []byte{0x18},
			Expected: head.ErrUnexpectedEnd,
		},
		{
			Data:     []byte{0x19, 0x03},
			Expected: head.ErrUnexpectedEnd,
		},
		{
			Data:     []byte{0x1a, 0x00, 0x0f, 0x42},
			Expected: head.ErrUnexpectedEnd,
		},
		{
			Data:     []byte{0x1b, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			Expected: head.ErrUnexpectedEnd,
		},



		{
			Data:     []byte{0x1c},
			Expected: head.ErrReservedAdditionalInformation,
		},
		{
			Data:     []byte{0x5d},
			Expected: head.ErrReservedAdditionalInformation,
		},
		{
			Data:     []byte{0xfe},
			Expected: head.ErrReservedAdditionalInformation,
		},
	}

	for testNumber, test := range tests {

		_, err := head.Decode(test.Data)

		{
			expected := test.Expected
			actual   := err

			if expected != actual {
				t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED: %v", expected)
				t.Logf("ACTUAL:   %v", actual)
				t.Logf("DATA: %#v", test.Data)
				continue
			}
		}
	}
}
