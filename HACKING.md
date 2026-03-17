# HACKING go-rfc8949

## Overview

Low-level CBOR (Concise Binary Object Representation) primitives for the Go programming-language (golang), implementing [IETF RFC 8949](https://datatracker.ietf.org/doc/html/rfc8949).
This is not a full CBOR encoder/decoder — it provides the building blocks used by the higher-level [go-cbor](http://github.com/reiver/go-cbor) package.

## Module

`github.com/reiver/go-rfc8949`

Go 1.22.4

No external dependencies.

## Commands

Run all tests:

```bash
go test ./...           # run all tests
```

Run tests for a single package:

```bash
go test ./prefix/       # run tests for a single package
```

Static analysis:

```
go vet ./...            # static analysis
```

## Architecture

The library models CBOR's byte-level encoding structure as separate Go packages:

- **`majortype/`** — The 8 CBOR major types (3 high bits of the initial byte). Constants (`MajorType0`–`MajorType7`) and type-check functions (`IsUnsignedInteger`, `IsByteString`, etc.). Major types are represented as `byte` values with the 3 high bits set.

- **`initialbyte/`** — Pre-computed initial byte constants combining major type + additional info (e.g., `UnsignedInteger0`–`UnsignedInteger23`, `Uint8`, `True`, `False`, `Null`, `Undefined`).

- **`initialbyte/additionalinfo/`** — Constants for the additional information field (low 5 bits): `Uint8=24`, `Uint16=25`, `Uint32=26`, `Uint64=27`.

- **`prefix/`** — `Prefix(majorType, length)` builds the variable-length CBOR prefix bytes (initial byte + 0/1/2/4/8 argument bytes) based on value size.

- **`tagnumber/`** — Constants for well-known CBOR tag numbers from RFC 8949 and the IANA CBOR Tags registry (e.g., `DateTime=0`, `Epoch=1`, `URI=32`, `BinaryUUID=37`, `IPv4=52`, `IPv6=54`, `SelfDescribedCBOR=55799`).

- **`types/`** — Per-type `Marshal()` functions that produce CBOR byte output:
  - `bools/`, `nils/`, `uint8s/`, `uint16s/`, `uint32s/`, `uint64s/`, `int8s/`, `int16s/`, `int32s/`, `int64s/`
  - `textstrings/` — CBOR text strings (major type 3) with UTF-8 validation.
  - `bytestrings/` — CBOR byte strings (major type 2).
  - `arrays/` — CBOR arrays.
  - `maps/` — CBOR maps with deterministic key ordering (RFC 8949 §4.2.1). Supports `map[any]any` and `map[string]any`.
  - `tags/` — CBOR tagged data items (major type 6). `Marshal(tagNumber, content)` wraps content in a tag.
    - `tags/datetimes/` — Tag 0: RFC 3339 date/time strings from `time.Time`.
    - `tags/epochs/` — Tag 1: Unix epoch seconds from `time.Time`.
    - `tags/ipv4s/` — Tag 52: IPv4 addresses and prefixes from `netip.Addr`/`netip.Prefix` (RFC 9164).
    - `tags/ipv6s/` — Tag 54: IPv6 addresses and prefixes from `netip.Addr`/`netip.Prefix` (RFC 9164).
    - `tags/uris/` — Tag 32: URIs from `*url.URL`.
    - `tags/uuids/` — Tag 37: Binary UUIDs from `[16]byte`.

- **`internal/marshalitem/`** — Internal marshaling dispatcher. Routes Go values to type-specific marshal functions, handles recursive collections (arrays, maps), and implements deterministic map key ordering.

- **`xtype/undefined/`** — Marshal for CBOR's "undefined" simple value (not present in Go's type system).

## Key Design Patterns

* Each package is intentionally small and single-purpose — one type or concept per package.
* Marshal functions return `([]byte, error)` and build output using `initialbyte` constants directly.
* Values 0–23 are encoded inline in the initial byte; larger values use the additional info field to signal 1/2/4/8 following bytes.
* Negative integers use CBOR's "value = -1 - argument" encoding (major type 1).
