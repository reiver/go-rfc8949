# HACKING go-rfc8949

## Overview

Low-level CBOR (Concise Binary Object Representation) primitives for the Go programming-language (golang), implementing [IETF RFC 8949](https://datatracker.ietf.org/doc/html/rfc8949).
This is not a full CBOR encoder/decoder — it provides the building blocks used by the higher-level [go-cbor](http://github.com/reiver/go-cbor) package.

## Key Design Patterns

* Each package is intentionally small and single-purpose — one type or concept per package.
* Marshal functions return `([]byte, error)` and build output using `initialbyte` constants directly.
* Values 0–23 are encoded inline in the initial byte; larger values use the additional info field to signal 1/2/4/8 following bytes.
* Negative integers use CBOR's "value = -1 - argument" encoding (major type 1).
