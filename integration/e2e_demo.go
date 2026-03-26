// Package main demonstrates the end-to-end flow:
// L2 Snapshot → StateRoot → L1 Anchor → Verify
package main

import (
	"fmt"
	"log"
	"time"

	"bhiv-sovereign-chain/l1-core/anchor"
	"bhiv-sovereign-chain/l2-execution"
)

func main() {
	// Step 1: Build L2 snapshot
	snapshot := l2.Snapshot{
		ChainID:   "bhiv-l2-001",
		Height:    1,
		Timestamp: time.Now().Unix(),
		Entries: []l2.StateEntry{
			{Key: "agent:karma", Value: "100"},
			{Key: "agent:status", Value: "active"},
		},
	}

	// Step 2: Compute deterministic state root
	stateRoot, err := l2.StateRoot(snapshot)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("state_root : %x\n", stateRoot)

	// Step 3: Anchor state root on L1 (genesis — parentHash is zero)
	var parentHash [32]byte
	anchorID, err := anchor.Submit(stateRoot, parentHash, snapshot.Timestamp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("anchor_id  : %x\n", anchorID)

	// Step 4: Verify anchor
	ok, err := anchor.VerifyAnchor(anchorID, stateRoot, parentHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("verified   : %v\n", ok)
}
