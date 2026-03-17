package ipv6s_test

import (
	"testing"

	"bytes"
	"net/netip"

	"github.com/reiver/go-rfc8949/types/tags/ipv6s"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value    netip.Addr
		Expected []byte
	}{
		// Loopback (::1).
		{
			Value: netip.MustParseAddr("::1"),
			Expected: []byte{
				0xd8, 0x36,
				0x50,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		},



		// All zeros (::).
		{
			Value: netip.MustParseAddr("::"),
			Expected: []byte{
				0xd8, 0x36,
				0x50,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},



		// 2001:db8::1
		{
			Value: netip.MustParseAddr("2001:db8::1"),
			Expected: []byte{
				0xd8, 0x36,
				0x50,
				0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		},



		// All 0xFF.
		{
			Value: netip.MustParseAddr("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"),
			Expected: []byte{
				0xd8, 0x36,
				0x50,
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			},
		},



		// IPv4-mapped IPv6 (::ffff:192.168.1.1).
		{
			Value: netip.MustParseAddr("::ffff:192.168.1.1"),
			Expected: []byte{
				0xd8, 0x36,
				0x50,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0xff, 0xff, 0xc0, 0xa8, 0x01, 0x01,
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := ipv6s.Marshal(test.Value)

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
		// Pure IPv4.
		{
			Value: netip.MustParseAddr("192.168.1.1"),
		},
	}

	for testNumber, test := range tests {

		_, err := ipv6s.Marshal(test.Value)

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
		// 2001:db8::/32 — trailing zeros truncated to 4 bytes.
		{
			Value: netip.MustParsePrefix("2001:db8::/32"),
			Expected: []byte{
				0xd8, 0x36,
				0x82,
				0x18, 0x20,
				0x44, 0x20, 0x01, 0x0d, 0xb8,
			},
		},



		// ::1/128 — full 16 bytes, no truncation (last byte is 0x01).
		{
			Value: netip.MustParsePrefix("::1/128"),
			Expected: []byte{
				0xd8, 0x36,
				0x82,
				0x18, 0x80,
				0x50,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		},



		// ::/0 — all zeros truncated to empty byte string.
		{
			Value: netip.MustParsePrefix("::/0"),
			Expected: []byte{
				0xd8, 0x36,
				0x82,
				0x00,
				0x40,
			},
		},



		// 2001:db8:1234::/48 — truncated to 6 bytes.
		{
			Value: netip.MustParsePrefix("2001:db8:1234::/48"),
			Expected: []byte{
				0xd8, 0x36,
				0x82,
				0x18, 0x30,
				0x46, 0x20, 0x01, 0x0d, 0xb8, 0x12, 0x34,
			},
		},



		// fe80::/10 — truncated to 1 byte (0xfe, then 0x80 is truncated... wait no).
		// fe80:: = fe 80 00 00 ... — truncated to 2 bytes (0xfe, 0x80).
		{
			Value: netip.MustParsePrefix("fe80::/10"),
			Expected: []byte{
				0xd8, 0x36,
				0x82,
				0x0a,
				0x42, 0xfe, 0x80,
			},
		},
	}

	for testNumber, test := range tests {

		actual, err := ipv6s.MarshalPrefix(test.Value)

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
		// Pure IPv4 prefix.
		{
			Value: netip.MustParsePrefix("192.168.1.0/24"),
		},
	}

	for testNumber, test := range tests {

		_, err := ipv6s.MarshalPrefix(test.Value)

		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("VALUE: %s", test.Value)
			continue
		}
	}
}
