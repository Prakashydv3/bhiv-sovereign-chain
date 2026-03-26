// Package anchor implements the L1 anchor registry.
// Anchors are reference-only — no state mutation after submission.
package anchor

import (
	"errors"
	"fmt"

	"bhiv-sovereign-chain/shared/hashing"
)

// Record is a single immutable anchor entry on L1.
type Record struct {
	AnchorID   [32]byte
	StateHash  [32]byte
	ParentHash [32]byte
	Timestamp  int64
}

// registry is the in-memory store (replaced by on-chain contract in production).
var registry = map[[32]byte]Record{}

// Submit stores a new anchor. Returns anchorID = SHA-256(stateHash || parentHash).
// stateHash must be non-zero. Duplicate anchors are rejected.
func Submit(stateHash [32]byte, parentHash [32]byte, timestamp int64) ([32]byte, error) {
	if stateHash == ([32]byte{}) {
		return [32]byte{}, errors.New("stateHash must not be zero")
	}

	anchorID := hashing.CombineHashes(stateHash, parentHash)

	if _, exists := registry[anchorID]; exists {
		return anchorID, fmt.Errorf("anchor %x already exists", anchorID)
	}

	registry[anchorID] = Record{
		AnchorID:   anchorID,
		StateHash:  stateHash,
		ParentHash: parentHash,
		Timestamp:  timestamp,
	}
	return anchorID, nil
}

// Query retrieves an anchor by ID. Returns error if not found.
func Query(anchorID [32]byte) (Record, error) {
	r, ok := registry[anchorID]
	if !ok {
		return Record{}, fmt.Errorf("anchor %x not found", anchorID)
	}
	return r, nil
}

// VerifyAnchor checks that the stored stateHash and parentHash match expectations.
func VerifyAnchor(anchorID [32]byte, expectedState [32]byte, expectedParent [32]byte) (bool, error) {
	r, err := Query(anchorID)
	if err != nil {
		return false, err
	}
	return r.StateHash == expectedState && r.ParentHash == expectedParent, nil
}
