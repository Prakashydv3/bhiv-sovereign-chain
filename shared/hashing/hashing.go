// Package hashing provides deterministic SHA-256 hashing primitives.
// All functions are pure — same input always produces same output.
package hashing

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Sum returns SHA-256 of raw bytes.
func Sum(data []byte) [32]byte {
	return sha256.Sum256(data)
}

// SumHex returns hex-encoded SHA-256 of raw bytes.
func SumHex(data []byte) string {
	h := sha256.Sum256(data)
	return fmt.Sprintf("%x", h)
}

// SumJSON returns SHA-256 of the canonical JSON encoding of v.
// Field order is determined by struct definition — callers must use structs, not maps.
func SumJSON(v any) ([32]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return [32]byte{}, fmt.Errorf("marshal: %w", err)
	}
	return sha256.Sum256(data), nil
}

// CombineHashes returns SHA-256(a || b) — used for anchor ID derivation.
func CombineHashes(a, b [32]byte) [32]byte {
	combined := append(a[:], b[:]...)
	return sha256.Sum256(combined)
}
