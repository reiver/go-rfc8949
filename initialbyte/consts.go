package initialbyte

import (
	"github.com/reiver/go-rfc8949/majortype"
)

const (
	UnsignedInteger0  byte = majortype.UnsignedInteger | 0
	UnsignedInteger1  byte = majortype.UnsignedInteger | 1
	UnsignedInteger2  byte = majortype.UnsignedInteger | 2
	UnsignedInteger3  byte = majortype.UnsignedInteger | 3
	UnsignedInteger4  byte = majortype.UnsignedInteger | 4
	UnsignedInteger5  byte = majortype.UnsignedInteger | 5
	UnsignedInteger6  byte = majortype.UnsignedInteger | 6
	UnsignedInteger7  byte = majortype.UnsignedInteger | 7
	UnsignedInteger8  byte = majortype.UnsignedInteger | 8
	UnsignedInteger9  byte = majortype.UnsignedInteger | 9
	UnsignedInteger10 byte = majortype.UnsignedInteger | 10
	UnsignedInteger11 byte = majortype.UnsignedInteger | 11
	UnsignedInteger12 byte = majortype.UnsignedInteger | 12
	UnsignedInteger13 byte = majortype.UnsignedInteger | 13
	UnsignedInteger14 byte = majortype.UnsignedInteger | 14
	UnsignedInteger15 byte = majortype.UnsignedInteger | 15
	UnsignedInteger16 byte = majortype.UnsignedInteger | 16
	UnsignedInteger17 byte = majortype.UnsignedInteger | 17
	UnsignedInteger18 byte = majortype.UnsignedInteger | 18
	UnsignedInteger19 byte = majortype.UnsignedInteger | 19
	UnsignedInteger20 byte = majortype.UnsignedInteger | 20
	UnsignedInteger21 byte = majortype.UnsignedInteger | 21
	UnsignedInteger22 byte = majortype.UnsignedInteger | 22
	UnsignedInteger23 byte = majortype.UnsignedInteger | 23

	Uint8             byte = majortype.UnsignedInteger | 24
	Uint16            byte = majortype.UnsignedInteger | 25
	Uint32            byte = majortype.UnsignedInteger | 26
	Uint64            byte = majortype.UnsignedInteger | 27

	NegativeIntegerNeg1  byte = majortype.NegativeInteger | (1  - 1)
	NegativeIntegerNeg2  byte = majortype.NegativeInteger | (2  - 1)
	NegativeIntegerNeg3  byte = majortype.NegativeInteger | (3  - 1)
	NegativeIntegerNeg4  byte = majortype.NegativeInteger | (4  - 1)
	NegativeIntegerNeg5  byte = majortype.NegativeInteger | (5  - 1)
	NegativeIntegerNeg6  byte = majortype.NegativeInteger | (6  - 1)
	NegativeIntegerNeg7  byte = majortype.NegativeInteger | (7  - 1)
	NegativeIntegerNeg8  byte = majortype.NegativeInteger | (8  - 1)
	NegativeIntegerNeg9  byte = majortype.NegativeInteger | (9  - 1)
	NegativeIntegerNeg10 byte = majortype.NegativeInteger | (10 - 1)
	NegativeIntegerNeg11 byte = majortype.NegativeInteger | (11 - 1)
	NegativeIntegerNeg12 byte = majortype.NegativeInteger | (12 - 1)
	NegativeIntegerNeg13 byte = majortype.NegativeInteger | (13 - 1)
	NegativeIntegerNeg14 byte = majortype.NegativeInteger | (14 - 1)
	NegativeIntegerNeg15 byte = majortype.NegativeInteger | (15 - 1)
	NegativeIntegerNeg16 byte = majortype.NegativeInteger | (16 - 1)
	NegativeIntegerNeg17 byte = majortype.NegativeInteger | (17 - 1)
	NegativeIntegerNeg18 byte = majortype.NegativeInteger | (18 - 1)
	NegativeIntegerNeg19 byte = majortype.NegativeInteger | (19 - 1)
	NegativeIntegerNeg20 byte = majortype.NegativeInteger | (20 - 1)
	NegativeIntegerNeg21 byte = majortype.NegativeInteger | (21 - 1)
	NegativeIntegerNeg22 byte = majortype.NegativeInteger | (22 - 1)
	NegativeIntegerNeg23 byte = majortype.NegativeInteger | (23 - 1)
	NegativeIntegerNeg24 byte = majortype.NegativeInteger | (24 - 1)

	Int8                 byte = majortype.NegativeInteger | (25 - 1)
	Int16                byte = majortype.NegativeInteger | (26 - 1)
	Int32                byte = majortype.NegativeInteger | (27 - 1)
	Int64                byte = majortype.NegativeInteger | (28 - 1)

	False     byte = majortype.SimpleValue | 20
	True      byte = majortype.SimpleValue | 21
	Null      byte = majortype.SimpleValue | 22
	Undefined byte = majortype.SimpleValue | 23
)
