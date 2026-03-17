package datetimes

import (
	"time"

	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a time.Time as a CBOR tagged RFC 3339 date/time string (tag 0, major type 6).
//
// The result is a tag 0 wrapping a CBOR text string containing the time
// formatted as RFC 3339. Sub-second precision is preserved when non-zero.
func Marshal(value time.Time) ([]byte, error) {
	return tags.Marshal(tagnumber.DateTime, value.Format(time.RFC3339Nano))
}
