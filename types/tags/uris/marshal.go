package uris

import (
	"fmt"
	"net/url"

	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a *url.URL as a CBOR tagged URI (tag 32, major type 6).
//
// The result is a tag 32 wrapping a CBOR text string containing the URI.
//
// This function does not validate that the *url.URL represents a valid
// URI per RFC 3986 beyond what Go's net/url package already provides.
// It is the caller's responsibility to ensure the URL is well-formed.
//
// A nil *url.URL returns an error.
func Marshal(value *url.URL) ([]byte, error) {
	if nil == value {
		return nil, fmt.Errorf("rfc8949: nil *url.URL")
	}

	return tags.Marshal(tagnumber.URI, value.String())
}
