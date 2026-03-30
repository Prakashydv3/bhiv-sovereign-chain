# L1 Minimal Truth Layer Design

## Overview
L1 is the canonical truth layer for BHIV. It stores immutable state root anchors only — no execution logic, no validator system, no complex consensus.

## Core Principle
**L1 = reference storage only. It never executes. It never mutates.**

## Components

### Genesis Block
The first anchor in the chain. Parent hash is zero.
```go
var parentHash [32]byte // zero hash
anchorID, err := anchor.Submit(genesisStateRoot, parentHash, timestamp)
```

### Block Struct (via Anchor Record)
```go
type Record struct {
    AnchorID   [32]byte  // SHA-256(stateHash || parentHash)
    StateHash  [32]byte  // Hash of anchored state
    ParentHash [32]byte  // Previous anchor (zero = genesis)
    Timestamp  int64     // Unix timestamp
}
```

### Deterministic Hashing
- AnchorID = `SHA-256(stateHash || parentHash)`
- Same inputs always produce same AnchorID
- Implemented via `hashing.CombineHashes()`

### State Root Registry
- In-memory map: `map[[32]byte]Record`
- Production replacement: on-chain smart contract
- Immutable: once written, never modified

## Key Functions

### InitializeGenesis()
```go
// Equivalent — submit with zero parent hash
var parentHash [32]byte
anchorID, err := anchor.Submit(genesisStateRoot, parentHash, timestamp)
```

### AnchorStateRoot(hash)
```go
// Submit any state root to L1
anchorID, err := anchor.Submit(stateHash, parentHash, timestamp)
```

### VerifyAnchor()
```go
// Verify stored anchor matches expected values
ok, err := anchor.VerifyAnchor(anchorID, expectedState, expectedParent)
```

## What L1 Does NOT Have
- No validator system
- No consensus mechanism
- No execution logic
- No state mutation after submission
- No complex block headers

## Determinism Guarantee
- `AnchorID = SHA-256(stateHash || parentHash)` — fully deterministic
- Duplicate submissions rejected — integrity preserved
- Zero-hash stateHash rejected — prevents empty anchors

## File Location
- **Implementation**: `l1-core/anchor/anchor.go`
- **HTTP API**: `l1-core/anchor/api.go`
- **Tests**: `l1-core/anchor/anchor_test.go`
