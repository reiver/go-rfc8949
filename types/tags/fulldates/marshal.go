package fulldates

import (
	"time"

	"github.com/reiver/go-rfc8949/tagnumber"
	"github.com/reiver/go-rfc8949/types/tags"
)

// Marshal encodes a time.Time as a CBOR tagged RFC 3339 full-date string (tag 1004, major type 6).
//
// The result is tag 1004 wrapping a CBOR text string containing only
// the date portion formatted as "YYYY-MM-DD" (RFC 3339 full-date production).
// Time-of-day and timezone information is discarded.
//
// See: RFC 8943.
func Marshal(value time.Time) ([]byte, error) {
	return tags.Marshal(tagnumber.FullDate, value.Format(time.DateOnly))
}
