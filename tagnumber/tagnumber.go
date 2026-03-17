package tagnumber

// Well-known CBOR tag numbers from IETF RFC 8949 and the IANA CBOR Tags registry.
const (
	// RFC 8949 Section 3.4, Table 3.
	DateTime         uint64 = 0     // RFC 3339 date/time string
	Epoch            uint64 = 1     // Epoch-based date/time (seconds since 1970-01-01T00:00Z)
	UnsignedBignum   uint64 = 2     // Unsigned bignum (byte string, big-endian)
	NegativeBignum   uint64 = 3     // Negative bignum (value = -1 - bignum)
	DecimalFraction  uint64 = 4     // Decimal fraction [exponent, mantissa] (m * 10^e)
	Bigfloat         uint64 = 5     // Bigfloat [exponent, mantissa] (m * 2^e)

	// RFC 8949 Section 3.4, Table 4.
	ExpectedBase64URL uint64 = 21   // Expected conversion to base64url encoding
	ExpectedBase64    uint64 = 22   // Expected conversion to base64 encoding
	ExpectedBase16    uint64 = 23   // Expected conversion to base16 encoding
	EncodedCBOR       uint64 = 24   // Encoded CBOR data item (byte string)

	// RFC 8949 Section 3.4, Table 5.
	URI              uint64 = 32    // URI (RFC 3986)
	Base64URL        uint64 = 33    // base64url-encoded data
	Base64           uint64 = 34    // base64-encoded data
	Regexp           uint64 = 35    // Regular expression (UTF-8 string)
	MIMEMessage      uint64 = 36    // MIME message (RFC 2045)

	// IANA CBOR Tags registry.
	BinaryUUID       uint64 = 37    // Binary UUID (16-byte byte string)
	Identifier       uint64 = 39    // Identifier (value is an ID, not for arithmetic)
	IPv4             uint64 = 52    // IPv4 address with optional prefix
	IPv6             uint64 = 54    // IPv6 address with optional prefix
	DaysSinceEpoch   uint64 = 100   // Days since 1970-01-01
	EmbeddedJSON     uint64 = 262   // Embedded JSON object (byte string)
	HexString        uint64 = 263   // Hexadecimal string (byte string, display as hex)
	JSONNumber       uint64 = 284   // JSON numeric value as text string
	FullDate         uint64 = 1004  // RFC 3339 full-date string
	CBORLD           uint64 = 51997 // CBOR-LD (compressed JSON-LD)

	// RFC 8949 Section 3.4.6.
	SelfDescribedCBOR uint64 = 55799 // Self-described CBOR
)
