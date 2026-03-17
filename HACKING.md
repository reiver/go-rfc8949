# HACKING go-rfc8949

## Overview

Low-level CBOR (Concise Binary Object Representation) primitives for the Go programming-language (golang), implementing [IETF RFC 8949](https://datatracker.ietf.org/doc/html/rfc8949).
This is not a full CBOR encoder/decoder — it provides the building blocks used by the higher-level [go-cbor](http://github.com/reiver/go-cbor) package.

## Architecture

The library models CBOR's byte-level encoding structure as separate Go packages:

- **`majortype/`** — The 8 CBOR major types (3 high bits of the initial byte). Constants (`MajorType0`–`MajorType7`) and type-check functions (`IsUnsignedInteger`, `IsByteString`, etc.). Major types are represented as `byte` values with the 3 high bits set.

- **`initialbyte/`** — Pre-computed initial byte constants combining major type + additional info (e.g., `UnsignedInteger0`–`UnsignedInteger23`, `Uint8`, `True`, `False`, `Null`, `Undefined`).

- **`initialbyte/additionalinfo/`** — Constants for the additional information field (low 5 bits): `Uint8=24`, `Uint16=25`, `Uint32=26`, `Uint64=27`.

- **`prefix/`** — `Prefix(majorType, length)` builds the variable-length CBOR prefix bytes (initial byte + 0/1/2/4/8 argument bytes) based on value size.

- **`types/`** — Per-type `Marshal()` functions that produce CBOR byte output:
  - `bools/`, `nils/`, `uint8s/`, `uint16s/`, `uint32s/`, `uint64s/`, `int8s/`, `int16s/`, `int32s/`, `int64s/`

- **`xtype/undefined/`** — Marshal for CBOR's "undefined" simple value (not present in Go's type system).

## Key Design Patterns

* Each package is intentionally small and single-purpose — one type or concept per package.
* Marshal functions return `([]byte, error)` and build output using `initialbyte` constants directly.
* Values 0–23 are encoded inline in the initial byte; larger values use the additional info field to signal 1/2/4/8 following bytes.
* Negative integers use CBOR's "value = -1 - argument" encoding (major type 1).
