# L2 Execution Framework Design

## Overview
The L2 execution framework provides the structure for processing BHIV executions and computing deterministic state roots that anchor to L1.

## Architecture

### Core Flow
```
L2 Execution → State Root → L1 Anchor
```

### Design Principle
**Structure-only implementation** - defines the framework without full BHIV logic.

## Components

### 1. L2Chain Struct
```go
type L2Chain struct {
    ChainID   string
    Height    uint64
    Timestamp int64
    Entries   []StateEntry
}
```

### 2. StateEntry
```go
type StateEntry struct {
    Key   string
    Value string
}
```

### 3. Snapshot
```go
type Snapshot struct {
    ChainID   string
    Height    uint64
    Timestamp int64
    Entries   []StateEntry
}
```

## Key Functions

### StateRoot()
```go
func StateRoot(snapshot Snapshot) ([32]byte, error)
```
- **Purpose**: Compute deterministic state root from snapshot
- **Algorithm**: SHA-256 of sorted key-value pairs
- **Determinism**: Same snapshot → same state root
- **Usage**: Anchoring point for L1

### BuildSnapshot()
```go
func BuildSnapshot(chainID string, entries []StateEntry) Snapshot
```
- **Purpose**: Create execution snapshot
- **Timestamp**: Current Unix time
- **Height**: Incremental block height
- **Entries**: State key-value pairs

## State Management

### Deterministic Ordering
```go
// Sort entries by key for consistent hashing
sort.Slice(entries, func(i, j int) bool {
    return entries[i].Key < entries[j].Key
})
```

### Hash Computation
```go
// Combine all entries into single hash
var combined []byte
for _, entry := range sortedEntries {
    combined = append(combined, []byte(entry.Key+":"+entry.Value+"|")...)
}
stateRoot := sha256.Sum256(combined)
```

## Integration Points

### Execution Envelopes
```go
// Convert envelope to L2 state entries
entries := []l2.StateEntry{
    {Key: "envelope:hash", Value: fmt.Sprintf("%x", envelope.Hash())},
    {Key: "envelope:decision", Value: envelope.Decision},
    {Key: "envelope:confidence", Value: fmt.Sprintf("%.3f", envelope.ConfidenceScore)},
}

// Build snapshot
snapshot := l2.BuildSnapshot("bhiv-l2", entries)

// Compute state root
stateRoot, err := l2.StateRoot(snapshot)
```

### L1 Anchoring
```go
// State root becomes anchor input
anchorID, err := anchor.Submit(stateRoot, parentHash, snapshot.Timestamp)
```

## Execution Simulation

### Agent State Example
```go
snapshot := l2.Snapshot{
    ChainID: "bhiv-l2-001",
    Height:  1,
    Entries: []l2.StateEntry{
        {Key: "agent:karma", Value: "100"},
        {Key: "agent:status", Value: "active"},
        {Key: "execution:count", Value: "42"},
    },
}
```

### System State Example
```go
snapshot := l2.Snapshot{
    ChainID: "bhiv-l2-gurukul",
    Height:  5,
    Entries: []l2.StateEntry{
        {Key: "tts:requests", Value: "1337"},
        {Key: "tts:success_rate", Value: "0.95"},
        {Key: "system:version", Value: "v2.1"},
    },
}
```

## Deterministic Properties

### State Root Computation
- **Sorting**: Keys sorted alphabetically
- **Encoding**: Consistent string format
- **Hashing**: SHA-256 for determinism
- **Reproducible**: Same input → same output

### Timestamp Handling
- Unix seconds precision
- Set at snapshot creation
- Preserved through state root computation

## Future Extensions

### Full BHIV Logic Integration
```go
// Future: Process actual BHIV executions
func ProcessBHIVExecution(intent Intent) (Snapshot, error) {
    // Execute BHIV logic
    // Generate state changes
    // Return execution snapshot
}
```

### Multi-System Support
```go
// Future: Support multiple execution systems
func ProcessSystemExecution(systemID string, input interface{}) (Snapshot, error) {
    switch systemID {
    case "gurukul":
        return processGurukul(input)
    case "marine":
        return processMarine(input)
    // ... other systems
    }
}
```

## Testing Strategy

### Determinism Tests
```go
func TestStateRootDeterminism(t *testing.T) {
    snapshot := createTestSnapshot()
    
    root1, _ := StateRoot(snapshot)
    root2, _ := StateRoot(snapshot)
    
    assert.Equal(t, root1, root2, "State root must be deterministic")
}
```

### Ordering Tests
```go
func TestKeyOrdering(t *testing.T) {
    // Test that key order doesn't affect state root
    entries1 := []StateEntry{{"b", "2"}, {"a", "1"}}
    entries2 := []StateEntry{{"a", "1"}, {"b", "2"}}
    
    root1, _ := StateRoot(Snapshot{Entries: entries1})
    root2, _ := StateRoot(Snapshot{Entries: entries2})
    
    assert.Equal(t, root1, root2, "Key order must not affect state root")
}
```

## Current Implementation

### File Location
- **Implementation**: `/l2-execution/l2.go`
- **Tests**: `/l2-execution/l2_test.go`

### Status
✅ **L2Chain struct**: Defined  
✅ **StateRoot()**: Implemented with deterministic hashing  
✅ **Snapshot creation**: Working  
✅ **L1 integration**: State root anchoring functional  
⏳ **Full BHIV logic**: Deferred (structure-only as specified)

## Integration Example

```go
// Complete L2 → L1 flow
snapshot := l2.Snapshot{
    ChainID: "bhiv-l2-001",
    Height:  1,
    Entries: []l2.StateEntry{
        {Key: "agent:karma", Value: "100"},
    },
}

// Compute state root
stateRoot, err := l2.StateRoot(snapshot)

// Anchor on L1
var parentHash [32]byte
anchorID, err := anchor.Submit(stateRoot, parentHash, snapshot.Timestamp)

// Verify
ok, err := anchor.VerifyAnchor(anchorID, stateRoot, parentHash)
```

This provides the **structure** for L2 execution without implementing full BHIV logic, exactly as specified in Phase 4.