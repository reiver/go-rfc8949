package arrays

import (
	"github.com/reiver/go-rfc8949/internal/marshalitem"
)

func Marshal[T any](value []T) ([]byte, error) {
	return marshalitem.MarshalArray(value)
}
