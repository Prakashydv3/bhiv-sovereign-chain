// Package l2 defines the canonical L2 state structure and state root computation.
package l2

import (
	"encoding/json"
	"fmt"
	"sort"

	"bhiv-sovereign-chain/shared/hashing"
)

// StateEntry is a single key-value record in L2 state.
type StateEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Snapshot is the canonical L2 state at a given block height.
// All fields are required for a valid state root.
type Snapshot struct {
	ChainID   string       `json:"chain_id"`
	Height    uint64       `json:"height"`
	Timestamp int64        `json:"timestamp"`
	Entries   []StateEntry `json:"entries"`
}

// StateRoot computes a deterministic SHA-256 root hash from a Snapshot.
// Entries are sorted by key before hashing — order-independence is guaranteed.
func StateRoot(s Snapshot) ([32]byte, error) {
	if s.ChainID == "" {
		return [32]byte{}, fmt.Errorf("chain_id must not be empty")
	}
	if s.Height == 0 {
		return [32]byte{}, fmt.Errorf("height must be greater than zero")
	}

	sorted := make([]StateEntry, len(s.Entries))
	copy(sorted, s.Entries)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Key < sorted[j].Key
	})

	canonical := struct {
		ChainID   string       `json:"chain_id"`
		Height    uint64       `json:"height"`
		Timestamp int64        `json:"timestamp"`
		Entries   []StateEntry `json:"entries"`
	}{s.ChainID, s.Height, s.Timestamp, sorted}

	data, err := json.Marshal(canonical)
	if err != nil {
		return [32]byte{}, fmt.Errorf("marshal: %w", err)
	}
	return hashing.Sum(data), nil
}
