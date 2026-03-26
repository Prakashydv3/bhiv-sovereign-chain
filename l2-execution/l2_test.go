package l2_test

import (
	"testing"

	"bhiv-sovereign-chain/l2-execution"
)

var base = l2.Snapshot{
	ChainID:   "bhiv-l2-001",
	Height:    1,
	Timestamp: 1700000000,
	Entries: []l2.StateEntry{
		{Key: "agent:karma", Value: "100"},
		{Key: "agent:status", Value: "active"},
	},
}

func TestStateRootDeterminism(t *testing.T) {
	h1, err := l2.StateRoot(base)
	if err != nil {
		t.Fatal(err)
	}
	h2, _ := l2.StateRoot(base)
	h3, _ := l2.StateRoot(base)

	if h1 != h2 || h2 != h3 {
		t.Fatalf("non-deterministic: %x %x %x", h1, h2, h3)
	}
}

func TestStateRootEntryOrderIndependence(t *testing.T) {
	reversed := l2.Snapshot{
		ChainID:   base.ChainID,
		Height:    base.Height,
		Timestamp: base.Timestamp,
		Entries: []l2.StateEntry{
			{Key: "agent:status", Value: "active"},
			{Key: "agent:karma", Value: "100"},
		},
	}

	h1, _ := l2.StateRoot(base)
	h2, _ := l2.StateRoot(reversed)

	if h1 != h2 {
		t.Fatalf("entry order changed hash: %x vs %x", h1, h2)
	}
}

func TestStateRootFieldSensitivity(t *testing.T) {
	baseHash, _ := l2.StateRoot(base)

	variants := []l2.Snapshot{
		{ChainID: "bhiv-l2-002", Height: 1, Timestamp: 1700000000, Entries: base.Entries},
		{ChainID: "bhiv-l2-001", Height: 2, Timestamp: 1700000000, Entries: base.Entries},
		{ChainID: "bhiv-l2-001", Height: 1, Timestamp: 1700000001, Entries: base.Entries},
	}

	for _, v := range variants {
		h, _ := l2.StateRoot(v)
		if h == baseHash {
			t.Fatalf("field change did not change hash: %+v", v)
		}
	}
}

func TestStateRootRejectsInvalid(t *testing.T) {
	_, err := l2.StateRoot(l2.Snapshot{ChainID: "", Height: 1, Timestamp: 1700000000})
	if err == nil {
		t.Fatal("expected error for empty chain_id")
	}

	_, err = l2.StateRoot(l2.Snapshot{ChainID: "bhiv-l2-001", Height: 0, Timestamp: 1700000000})
	if err == nil {
		t.Fatal("expected error for zero height")
	}
}
