package maps

import (
	"github.com/reiver/go-rfc8949/internal/marshalitem"
)

func Marshal(value map[any]any) ([]byte, error) {
	return marshalitem.MarshalMap(value)
}

func MarshalStringKeys(value map[string]any) ([]byte, error) {
	return marshalitem.MarshalMapStringKeys(value)
}
