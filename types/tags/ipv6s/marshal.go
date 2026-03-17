package ipv6s

import (
	"fmt"
	"net/netip"

	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a netip.Addr as a CBOR tagged IPv6 address (tag 54, major type 6).
//
// The result is a tag 54 wrapping a 16-byte CBOR byte string containing
// the raw address bytes.
//
// The address must be a valid IPv6 address. Pure IPv4 addresses are
// rejected — use tag 52 for those. IPv4-mapped IPv6 addresses (e.g.,
// ::ffff:192.168.1.1) are accepted.
func Marshal(value netip.Addr) ([]byte, error) {

	if !value.IsValid() {
		return nil, fmt.Errorf("rfc8949: invalid netip.Addr")
	}

	if value.Is4() {
		return nil, fmt.Errorf("rfc8949: IPv4 address passed to IPv6 marshal (tag 54); use tag 52 for IPv4")
	}

	addr := value.As16()

	return tags.Marshal(tagnumber.IPv6, addr[:])
}

// MarshalPrefix encodes a netip.Prefix as a CBOR tagged IPv6 address with
// prefix length (tag 54, major type 6).
//
// The result is a tag 54 wrapping a 2-element CBOR array [prefix-length, byte-string].
// Trailing zero bytes in the address are omitted per RFC 9164.
//
// The address must be a valid IPv6 prefix. Pure IPv4 prefixes are
// rejected — use tag 52 for those.
func MarshalPrefix(value netip.Prefix) ([]byte, error) {

	if !value.IsValid() {
		return nil, fmt.Errorf("rfc8949: invalid netip.Prefix")
	}

	if value.Addr().Is4() {
		return nil, fmt.Errorf("rfc8949: IPv4 prefix passed to IPv6 marshal (tag 54); use tag 52 for IPv4")
	}

	addr := value.Addr().As16()

	// Truncate trailing zero bytes per RFC 9164.
	truncated := addr[:]
	for len(truncated) > 0 && truncated[len(truncated)-1] == 0 {
		truncated = truncated[:len(truncated)-1]
	}

	return tags.Marshal(tagnumber.IPv6, []any{uint64(value.Bits()), truncated})
}
