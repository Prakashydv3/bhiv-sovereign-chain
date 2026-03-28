# Execution Envelope Specification

## Overview
The ExecutionEnvelope is the canonical format for all BHIV system executions. It ensures deterministic processing and traceability across all components.

## Structure
```go
type ExecutionEnvelope struct {
    IntentID         string   // Unique execution identifier
    SystemID         string   // Source system (gurukul, marine, etc)
    AgentID          string   // Executing agent identifier
    InputHash        [32]byte // SHA256 of input data
    OutputHash       [32]byte // SHA256 of output data
    Decision         string   // Execution decision/result
    ConfidenceScore  float64  // Decision confidence (0-1)
    PolicyID         string   // Applied policy identifier
    DatasetIDs       []string // Referenced datasets
    Timestamp        int64    // Unix timestamp
    EnvironmentHash  [32]byte // Environment state hash
}
```

## Usage
1. Create envelope: `NewEnvelope(systemID, agentID)`
2. Set input/output hashes
3. Record decision and confidence
4. Validate: `envelope.Validate()`
5. Hash: `envelope.Hash()`

## Integration Points
- L2 execution layer consumes envelopes
- L1 anchors envelope hashes
- Replay system uses envelopes for determinism
- All BHIV systems must produce valid envelopes

## Determinism Requirements
- JSON marshaling must be consistent
- Hash computation must be reproducible
- Timestamp precision: Unix seconds
- All fields must be populated for valid envelope