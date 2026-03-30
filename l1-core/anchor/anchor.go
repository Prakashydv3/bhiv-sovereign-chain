package anchor

import (
	"errors"
	"fmt"

	"bhiv-sovereign-chain/shared/hashing"
)

// Record is a single immutable anchor entry on L1.
type Record struct {
	AnchorID          [32]byte
	StateHash         [32]byte
	ParentHash        [32]byte
	Timestamp         int64
	AgentSignerID     string
	EnforcementSignerID string
}

// registry is the in-memory store (replaced by on-chain contract in production).
var registry = map[[32]byte]Record{}

// Submit stores a new anchor only after both agent and enforcement signatures are verified.
// Flow: envelope → hash → signatures → L1 anchor
// stateHash must be non-zero. Both signatures must be valid. Duplicates are rejected.
func Submit(
	stateHash [32]byte,
	parentHash [32]byte,
	timestamp int64,
	agentSig *hashing.SignedHash,
	enforcementSig *hashing.SignedHash,
) ([32]byte, error) {
	if stateHash == ([32]byte{}) {
		return [32]byte{}, errors.New("stateHash must not be zero")
	}

	// Verify agent signature — anchor is rejected if invalid
	if err := agentSig.Verify(); err != nil {
		return [32]byte{}, fmt.Errorf("agent signature invalid: %w", err)
	}
	if agentSig.Hash != stateHash {
		return [32]byte{}, fmt.Errorf("agent signature does not cover the submitted stateHash")
	}

	// Verify enforcement signature — anchor is rejected if invalid
	if err := enforcementSig.Verify(); err != nil {
		return [32]byte{}, fmt.Errorf("enforcement signature invalid: %w", err)
	}
	if enforcementSig.Hash != stateHash {
		return [32]byte{}, fmt.Errorf("enforcement signature does not cover the submitted stateHash")
	}

	anchorID := hashing.CombineHashes(stateHash, parentHash)

	if _, exists := registry[anchorID]; exists {
		return anchorID, fmt.Errorf("anchor %x already exists", anchorID)
	}

	registry[anchorID] = Record{
		AnchorID:            anchorID,
		StateHash:           stateHash,
		ParentHash:          parentHash,
		Timestamp:           timestamp,
		AgentSignerID:       agentSig.SignerID,
		EnforcementSignerID: enforcementSig.SignerID,
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
