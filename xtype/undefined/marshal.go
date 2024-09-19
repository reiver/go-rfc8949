package undefined

import (
	"github.com/reiver/go-rfc8949/initialbyte"
)

func Marshal() ([]byte, error) {
	return []byte{initialbyte.Undefined}, nil
}
