package uris_test

import (
	"testing"

	"bytes"
	"net/url"

	"github.com/reiver/go-rfc8949/types/tags/uris"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    *url.URL
		Expected []byte
	}{
		// RFC 8949 Appendix A example.
		{
			// 32("http://www.example.com")
			Value:    mustParseURL("http://www.example.com"),
			Expected: []byte{0xd8, 0x20, 0x76, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d},
		},



		// Various URI forms.
		{
			// 32("https://example.com")
			Value:    mustParseURL("https://example.com"),
			Expected: []byte{0xd8, 0x20, 0x73, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d},
		},
		{
			// 32("https://example.com/path?query=value#fragment")
			Value:    mustParseURL("https://example.com/path?query=value#fragment"),
			Expected: append(
				[]byte{0xd8, 0x20},
				append(
					[]byte{0x78, 0x2d},
					[]byte("https://example.com/path?query=value#fragment")...,
				)...,
			),
		},
		{
			// 32("urn:isbn:0451450523")
			Value:    mustParseURL("urn:isbn:0451450523"),
			Expected: append(
				[]byte{0xd8, 0x20},
				append(
					[]byte{0x73},
					[]byte("urn:isbn:0451450523")...,
				)...,
			),
		},
	}

	for testNumber, test := range tests {

		actual, err := uris.Marshal(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %s", test.Value)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual cbor-marshaled bytes is not what was expected.", testNumber)
				t.Logf("EXPECTED: (len=%d) %#v", len(expected), expected)
				t.Logf("ACTUAL:   (len=%d) %#v", len(actual),   actual)
				t.Logf("VALUE: %s", test.Value)
				continue
			}
		}
	}
}

func TestMarshalNil(t *testing.T) {

	_, err := uris.Marshal(nil)

	if nil == err {
		t.Errorf("Expected an error for nil *url.URL but did not get one.")
	}
}

func mustParseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if nil != err {
		panic(err)
	}
	return u
}
