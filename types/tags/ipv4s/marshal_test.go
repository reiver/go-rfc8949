package ipv4s_test

import (
	"testing"

	"bytes"
	"net/netip"

	"github.com/reiver/go-rfc8949/types/tags/ipv4s"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    netip.Addr
		Expected []byte
	}{
		// 192.168.1.1
		{
			Value: netip.MustParseAddr("192.168.1.1"),
			Expected: []byte{
				0xd8, 0x34,
				0x44,
				0xc0, 0xa8, 0x01, 0x01,
			},
		},



		// 0.0.0.0
		{
			Value: netip.MustParseAddr("0.0.0.0"),
			Expected: []byte{
				0xd8, 0x34,
				0x44,
				0x00, 0x00, 0x00, 0x00,
			},
		},



		// 255.255.255.255
		{
			Value: netip.MustParseAddr("255.255.255.255"),
			Expected: []byte{
				0xd8, 0x34,
				0x44,
				0xff, 0xff, 0xff, 0xff,
			},
		},



		// 127.0.0.1 (loopback)
		{
			Value: netip.MustParseAddr("127.0.0.1"),
			Expected: []byte{
				0xd8, 0x34,
				0x44,
				0x7f, 0x00, 0x00, 0x01,
			},
		},



		// 10.0.0.1
		{
			Value: netip.MustParseAddr("10.0.0.1"),
			Expected: []byte{
				0xd8, 0x34,
				0x44,
				0x0a, 0x00, 0x00, 0x01,
			},
		},



		// IPv4-mapped IPv6 accepted and unmapped.
		{
			Value: netip.MustParseAddr("::ffff:192.168.1.1"),
			Expected: []byte{
				0xd8, 0x34,
				0x44,
				0xc0, 0xa8, 0x01, 0x01,
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := ipv4s.Marshal(test.Value)

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

func TestMarshalErrors(t *testing.T) {

	tests := []struct{
		Value netip.Addr
	}{
		// Invalid (zero value).
		{
			Value: netip.Addr{},
		},
		// Pure IPv6.
		{
			Value: netip.MustParseAddr("2001:db8::1"),
		},
	}

	for testNumber, test := range tests {

		_, err := ipv4s.Marshal(test.Value)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %s", test.Value)
			continue
		}
	}
}

func TestMarshalPrefix(t *testing.T) {

	tests := []struct{
		Value    netip.Prefix
		Expected []byte
	}{
		// 192.168.0.0/16 — truncated to 2 bytes.
		{
			Value: netip.MustParsePrefix("192.168.0.0/16"),
			Expected: []byte{
				0xd8, 0x34,
				0x82,
				0x10,
				0x42, 0xc0, 0xa8,
			},
		},



		// 10.0.0.0/8 — truncated to 1 byte.
		{
			Value: netip.MustParsePrefix("10.0.0.0/8"),
			Expected: []byte{
				0xd8, 0x34,
				0x82,
				0x08,
				0x41, 0x0a,
			},
		},



		// 0.0.0.0/0 — all zeros truncated to empty byte string.
		{
			Value: netip.MustParsePrefix("0.0.0.0/0"),
			Expected: []byte{
				0xd8, 0x34,
				0x82,
				0x00,
				0x40,
			},
		},



		// 192.168.1.1/32 — full 4 bytes, no truncation.
		{
			Value: netip.MustParsePrefix("192.168.1.1/32"),
			Expected: []byte{
				0xd8, 0x34,
				0x82,
				0x18, 0x20,
				0x44, 0xc0, 0xa8, 0x01, 0x01,
			},
		},



		// 172.16.0.0/12 — truncated to 2 bytes.
		{
			Value: netip.MustParsePrefix("172.16.0.0/12"),
			Expected: []byte{
				0xd8, 0x34,
				0x82,
				0x0c,
				0x42, 0xac, 0x10,
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := ipv4s.MarshalPrefix(test.Value)

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

func TestMarshalPrefixErrors(t *testing.T) {

	tests := []struct{
		Value netip.Prefix
	}{
		// Invalid (zero value).
		{
			Value: netip.Prefix{},
		},
		// Pure IPv6 prefix.
		{
			Value: netip.MustParsePrefix("2001:db8::/32"),
		},
	}

	for testNumber, test := range tests {

		_, err := ipv4s.MarshalPrefix(test.Value)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %s", test.Value)
			continue
		}
	}
}
