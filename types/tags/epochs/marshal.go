package epochs

import (
	"time"

	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a time.Time as a CBOR tagged epoch-based date/time (tag 1, major type 6).
//
// The result is a tag 1 wrapping a CBOR integer containing the number of
// seconds since 1970-01-01T00:00Z (Unix epoch). Dates before 1970 produce
// negative integers.
//
// Sub-second precision is truncated. Use types/tags/datetimes (tag 0) to
// preserve nanosecond precision via the RFC 3339 text format.
func Marshal(value time.Time) ([]byte, error) {
	return tags.Marshal(tagnumber.Epoch, value.Unix())
}
