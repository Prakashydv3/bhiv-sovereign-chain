# Anchor System Design

## Overview
The anchor system provides the L1 truth layer for BHIV sovereign chain. It stores immutable references to execution state roots without mutating the underlying state.

## Architecture

### Core Principle
**Anchors are reference-only** - they point to state but never modify it.

### Components

#### 1. Anchor Record
```go
type Record struct {
    AnchorID   [32]byte  // SHA-256(stateHash || parentHash)
    StateHash  [32]byte  // Hash of the anchored state
    ParentHash [32]byte  // Previous anchor (zero for genesis)
    Timestamp  int64     // Unix timestamp
}
```

#### 2. Registry
- In-memory store (production: on-chain contract)
- Immutable once written
- Duplicate prevention

### Key Functions

#### Submit()
```go
func Submit(stateHash [32]byte, parentHash [32]byte, timestamp int64) ([32]byte, error)
```
- **Purpose**: Store new anchor reference
- **Returns**: Deterministic anchorID
- **Validation**: stateHash must be non-zero
- **Immutability**: Rejects duplicates

#### VerifyAnchor()
```go
func VerifyAnchor(anchorID [32]byte, expectedState [32]byte, expectedParent [32]byte) (bool, error)
```
- **Purpose**: Verify anchor integrity
- **Validation**: Checks stored vs expected hashes
- **Usage**: Replay verification, audit trails

#### Query()
```go
func Query(anchorID [32]byte) (Record, error)
```
- **Purpose**: Retrieve anchor by ID
- **Returns**: Complete anchor record
- **Error**: If anchor not found

## Integration Points

### L2 Execution Layer
```go
// L2 computes state root
stateRoot := l2.StateRoot(snapshot)

// Anchor on L1
anchorID, err := anchor.Submit(stateRoot, parentHash, timestamp)
```

### Execution Envelopes
```go
// Envelope hash becomes state root
envelopeHash := envelope.Hash()

// Anchor envelope hash
anchorID, err := anchor.Submit(envelopeHash, parentHash, timestamp)
```

## Deterministic Properties

### AnchorID Generation
- `anchorID = SHA-256(stateHash || parentHash)`
- Same inputs → same anchorID
- Enables deterministic verification

### Immutability Guarantees
- Once anchored, never modified
- Duplicate submissions rejected
- Historical integrity preserved

## Security Model

### Reference-Only Design
- Anchors don't contain executable logic
- No state mutation capabilities
- Pure reference storage

### Cryptographic Integrity
- SHA-256 hashing for anchor IDs
- Deterministic verification
- Tamper-evident structure

## Production Considerations

### On-Chain Deployment
- Replace in-memory registry with smart contract
- Gas optimization for anchor submissions
- Event emission for indexing

### Scalability
- Batch anchor submissions
- State root compression
- Historical pruning strategies

## Usage Examples

### Genesis Anchor
```go
var parentHash [32]byte // Zero hash for genesis
anchorID, err := anchor.Submit(genesisStateRoot, parentHash, timestamp)
```

### Chain Continuation
```go
// Use previous anchor as parent
anchorID, err := anchor.Submit(newStateRoot, previousAnchorID, timestamp)
```

### Verification
```go
// Verify anchor integrity
ok, err := anchor.VerifyAnchor(anchorID, expectedState, expectedParent)
if !ok {
    // Anchor verification failed
}
```

## Error Handling

### Common Errors
- `stateHash must not be zero` - Invalid state hash
- `anchor already exists` - Duplicate submission
- `anchor not found` - Query for non-existent anchor

### Recovery Strategies
- Validate inputs before submission
- Handle duplicate gracefully
- Implement retry logic for queries

## Integration Status
✅ **Moved to**: `/l1-core/anchor/`  
✅ **Reference-only**: No state mutation  
✅ **VerifyAnchor()**: Implemented and tested  
✅ **Deterministic**: Same inputs → same outputs