package ipv4s

import (
	"fmt"
	"net/netip"

	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a netip.Addr as a CBOR tagged IPv4 address (tag 52, major type 6).
//
// The result is a tag 52 wrapping a 4-byte CBOR byte string containing
// the raw address bytes.
//
// The address must be a valid IPv4 address. Pure IPv6 addresses are
// rejected — use tag 54 for those. IPv4-mapped IPv6 addresses (e.g.,
// ::ffff:192.168.1.1) are accepted and unmapped to their IPv4 form.
func Marshal(value netip.Addr) ([]byte, error) {

	if !value.IsValid() {
		return nil, fmt.Errorf("rfc8949: invalid netip.Addr")
	}

	value = value.Unmap()

	if !value.Is4() {
		return nil, fmt.Errorf("rfc8949: IPv6 address passed to IPv4 marshal (tag 52); use tag 54 for IPv6")
	}

	addr := value.As4()

	return tags.Marshal(tagnumber.IPv4, addr[:])
}

// MarshalPrefix encodes a netip.Prefix as a CBOR tagged IPv4 address with
// prefix length (tag 52, major type 6).
//
// The result is a tag 52 wrapping a 2-element CBOR array [prefix-length, byte-string].
// Trailing zero bytes in the address are omitted per RFC 9164.
//
// The address must be a valid IPv4 prefix. Pure IPv6 prefixes are
// rejected — use tag 54 for those. IPv4-mapped IPv6 prefixes are
// accepted and unmapped to their IPv4 form.
func MarshalPrefix(value netip.Prefix) ([]byte, error) {

	if !value.IsValid() {
		return nil, fmt.Errorf("rfc8949: invalid netip.Prefix")
	}

	unmapped := value.Addr().Unmap()

	if !unmapped.Is4() {
		return nil, fmt.Errorf("rfc8949: IPv6 prefix passed to IPv4 marshal (tag 52); use tag 54 for IPv6")
	}

	addr := unmapped.As4()

	// Truncate trailing zero bytes per RFC 9164.
	truncated := addr[:]
	for len(truncated) > 0 && truncated[len(truncated)-1] == 0 {
		truncated = truncated[:len(truncated)-1]
	}

	return tags.Marshal(tagnumber.IPv4, []any{uint64(value.Bits()), truncated})
}
