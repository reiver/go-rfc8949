package initialbyte

import (
	"github.com/reiver/go-rfc8949/initialbyte/additionalinfo"
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

	Uint8             byte = majortype.UnsignedInteger | additionalinfo.Uint8
	Uint16            byte = majortype.UnsignedInteger | additionalinfo.Uint16
	Uint32            byte = majortype.UnsignedInteger | additionalinfo.Uint32
	Uint64            byte = majortype.UnsignedInteger | additionalinfo.Uint64

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

	Int8                 byte = majortype.NegativeInteger | additionalinfo.Uint8
	Int16                byte = majortype.NegativeInteger | additionalinfo.Uint16
	Int32                byte = majortype.NegativeInteger | additionalinfo.Uint32
	Int64                byte = majortype.NegativeInteger | additionalinfo.Uint64

	ByteStringLen0      byte = majortype.ByteString | 0
	ByteStringLen1      byte = majortype.ByteString | 1
	ByteStringLen2      byte = majortype.ByteString | 2
	ByteStringLen3      byte = majortype.ByteString | 3
	ByteStringLen4      byte = majortype.ByteString | 4
	ByteStringLen5      byte = majortype.ByteString | 5
	ByteStringLen6      byte = majortype.ByteString | 6
	ByteStringLen7      byte = majortype.ByteString | 7
	ByteStringLen8      byte = majortype.ByteString | 8
	ByteStringLen9      byte = majortype.ByteString | 9
	ByteStringLen10     byte = majortype.ByteString | 10
	ByteStringLen11     byte = majortype.ByteString | 11
	ByteStringLen12     byte = majortype.ByteString | 12
	ByteStringLen13     byte = majortype.ByteString | 13
	ByteStringLen14     byte = majortype.ByteString | 14
	ByteStringLen15     byte = majortype.ByteString | 15
	ByteStringLen16     byte = majortype.ByteString | 16
	ByteStringLen17     byte = majortype.ByteString | 17
	ByteStringLen18     byte = majortype.ByteString | 18
	ByteStringLen19     byte = majortype.ByteString | 19
	ByteStringLen20     byte = majortype.ByteString | 20
	ByteStringLen21     byte = majortype.ByteString | 21
	ByteStringLen22     byte = majortype.ByteString | 22
	ByteStringLen23     byte = majortype.ByteString | 23

	ByteStringLenUint8  byte = majortype.ByteString | additionalinfo.Uint8
	ByteStringLenUint16 byte = majortype.ByteString | additionalinfo.Uint16
	ByteStringLenUint32 byte = majortype.ByteString | additionalinfo.Uint32
	ByteStringLenUint64 byte = majortype.ByteString | additionalinfo.Uint64

	TextStringLen0      byte = majortype.TextString | 0
	TextStringLen1      byte = majortype.TextString | 1
	TextStringLen2      byte = majortype.TextString | 2
	TextStringLen3      byte = majortype.TextString | 3
	TextStringLen4      byte = majortype.TextString | 4
	TextStringLen5      byte = majortype.TextString | 5
	TextStringLen6      byte = majortype.TextString | 6
	TextStringLen7      byte = majortype.TextString | 7
	TextStringLen8      byte = majortype.TextString | 8
	TextStringLen9      byte = majortype.TextString | 9
	TextStringLen10     byte = majortype.TextString | 10
	TextStringLen11     byte = majortype.TextString | 11
	TextStringLen12     byte = majortype.TextString | 12
	TextStringLen13     byte = majortype.TextString | 13
	TextStringLen14     byte = majortype.TextString | 14
	TextStringLen15     byte = majortype.TextString | 15
	TextStringLen16     byte = majortype.TextString | 16
	TextStringLen17     byte = majortype.TextString | 17
	TextStringLen18     byte = majortype.TextString | 18
	TextStringLen19     byte = majortype.TextString | 19
	TextStringLen20     byte = majortype.TextString | 20
	TextStringLen21     byte = majortype.TextString | 21
	TextStringLen22     byte = majortype.TextString | 22
	TextStringLen23     byte = majortype.TextString | 23

	TextStringLenUint8  byte = majortype.TextString | additionalinfo.Uint8
	TextStringLenUint16 byte = majortype.TextString | additionalinfo.Uint16
	TextStringLenUint32 byte = majortype.TextString | additionalinfo.Uint32
	TextStringLenUint64 byte = majortype.TextString | additionalinfo.Uint64

	False     byte = majortype.SimpleValue | 20
	True      byte = majortype.SimpleValue | 21
	Null      byte = majortype.SimpleValue | 22
	Undefined byte = majortype.SimpleValue | 23
)
