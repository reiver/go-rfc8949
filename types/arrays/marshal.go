package arrays

import (
	"github.com/reiver/go-rfc8949/internal/marshalitem"
)

func Marshal(value []any) ([]byte, error) {
	return marshalitem.MarshalArray(value)
}
