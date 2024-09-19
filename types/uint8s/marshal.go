package uint8s

import (
	"github.com/reiver/go-rfc8949/initialbyte"
)

func Marshal(value uint8) ([]byte, error) {

	if value <= 23 {
		switch value {
		case                                             0:
			return []byte{initialbyte.UnsignedInteger0}, nil
		case                                             1:
			return []byte{initialbyte.UnsignedInteger1}, nil
		case                                             2:
			return []byte{initialbyte.UnsignedInteger2}, nil
		case                                             3:
			return []byte{initialbyte.UnsignedInteger3}, nil
		case                                             4:
			return []byte{initialbyte.UnsignedInteger4}, nil
		case                                             5:
			return []byte{initialbyte.UnsignedInteger5}, nil
		case                                             6:
			return []byte{initialbyte.UnsignedInteger6}, nil
		case                                             7:
			return []byte{initialbyte.UnsignedInteger7}, nil
		case                                             8:
			return []byte{initialbyte.UnsignedInteger8}, nil
		case                                             9:
			return []byte{initialbyte.UnsignedInteger9}, nil
		case                                             10:
			return []byte{initialbyte.UnsignedInteger10}, nil
		case                                             11:
			return []byte{initialbyte.UnsignedInteger11}, nil
		case                                             12:
			return []byte{initialbyte.UnsignedInteger12}, nil
		case                                             13:
			return []byte{initialbyte.UnsignedInteger13}, nil
		case                                             14:
			return []byte{initialbyte.UnsignedInteger14}, nil
		case                                             15:
			return []byte{initialbyte.UnsignedInteger15}, nil
		case                                             16:
			return []byte{initialbyte.UnsignedInteger16}, nil
		case                                             17:
			return []byte{initialbyte.UnsignedInteger17}, nil
		case                                             18:
			return []byte{initialbyte.UnsignedInteger18}, nil
		case                                             19:
			return []byte{initialbyte.UnsignedInteger19}, nil
		case                                             20:
			return []byte{initialbyte.UnsignedInteger20}, nil
		case                                             21:
			return []byte{initialbyte.UnsignedInteger21}, nil
		case                                             22:
			return []byte{initialbyte.UnsignedInteger22}, nil
		case                                             23:
			return []byte{initialbyte.UnsignedInteger23}, nil
		}
	}

	return []byte{initialbyte.Uint8, byte(value)}, nil
}
