package int8s

import (
	"github.com/reiver/go-rfc8949/initialbyte"
	"github.com/reiver/go-rfc8949/types/uint8s"
)

func Marshal(value int8) ([]byte, error) {

	if 0 <= value {
		return uint8s.Marshal(uint8(value))
	}

	if -24 <= value {
		switch value {
		case                                               -24:
			return []byte{initialbyte.NegativeIntegerNeg24}, nil
		case                                               -23:
			return []byte{initialbyte.NegativeIntegerNeg23}, nil
		case                                               -22:
			return []byte{initialbyte.NegativeIntegerNeg22}, nil
		case                                               -21:
			return []byte{initialbyte.NegativeIntegerNeg21}, nil
		case                                               -20:
			return []byte{initialbyte.NegativeIntegerNeg20}, nil
		case                                               -19:
			return []byte{initialbyte.NegativeIntegerNeg19}, nil
		case                                               -18:
			return []byte{initialbyte.NegativeIntegerNeg18}, nil
		case                                               -17:
			return []byte{initialbyte.NegativeIntegerNeg17}, nil
		case                                               -16:
			return []byte{initialbyte.NegativeIntegerNeg16}, nil
		case                                               -15:
			return []byte{initialbyte.NegativeIntegerNeg15}, nil
		case                                               -14:
			return []byte{initialbyte.NegativeIntegerNeg14}, nil
		case                                               -13:
			return []byte{initialbyte.NegativeIntegerNeg13}, nil
		case                                               -12:
			return []byte{initialbyte.NegativeIntegerNeg12}, nil
		case                                               -11:
			return []byte{initialbyte.NegativeIntegerNeg11}, nil
		case                                               -10:
			return []byte{initialbyte.NegativeIntegerNeg10}, nil
		case                                               -9:
			return []byte{initialbyte.NegativeIntegerNeg9}, nil
		case                                               -8:
			return []byte{initialbyte.NegativeIntegerNeg8}, nil
		case                                               -7:
			return []byte{initialbyte.NegativeIntegerNeg7}, nil
		case                                               -6:
			return []byte{initialbyte.NegativeIntegerNeg6}, nil
		case                                               -5:
			return []byte{initialbyte.NegativeIntegerNeg5}, nil
		case                                               -4:
			return []byte{initialbyte.NegativeIntegerNeg4}, nil
		case                                               -3:
			return []byte{initialbyte.NegativeIntegerNeg3}, nil
		case                                               -2:
			return []byte{initialbyte.NegativeIntegerNeg2}, nil
		case                                               -1:
			return []byte{initialbyte.NegativeIntegerNeg1}, nil
		}
	}

	return []byte{initialbyte.Int8, byte((value + 1)*-1)}, nil
}

