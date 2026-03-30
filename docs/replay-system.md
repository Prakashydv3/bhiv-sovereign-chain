# Replay System Design

## Overview
The replay system guarantees determinism across BHIV executions. Given the same input, `ReplayChain()` must always produce the same final hash. Any mismatch is a hard failure.

## Core Principle
**Same input → same envelope → same hash. Always. No exceptions.**

## Components

### ReplayEngine
```go
type ReplayEngine struct {
    ttsProcessor *gurukul.TTSProcessor
}
```
Holds references to all system processors needed for re-execution.

### ReplayInput
```go
type ReplayInput struct {
    OriginalEnvelope *envelope.ExecutionEnvelope
    Input            interface{}
}
```
Pairs the original envelope with the raw input used to produce it.

## Key Functions

### ReplayExecution()
Re-executes a single execution and verifies determinism against the original envelope.
```go
func (r *ReplayEngine) ReplayExecution(
    originalEnvelope *envelope.ExecutionEnvelope,
    input interface{},
) (*envelope.ExecutionEnvelope, error)
```
- Routes to correct system processor by `SystemID`
- Re-executes with same input
- Verifies: `InputHash`, `OutputHash`, `Decision`, `ConfidenceScore`
- Returns error on any mismatch

### ReplayChain()
Replays a sequence of executions and returns the final hash.
```go
func (r *ReplayEngine) ReplayChain(executions []ReplayInput) ([32]byte, error)
```
- Replays each execution in order
- Fails immediately on first mismatch
- Returns final envelope hash of last execution

## Determinism Verification

### Fields Checked
| Field | Mismatch Behaviour |
|---|---|
| `InputHash` | Hard fail — input changed |
| `OutputHash` | Hard fail — output changed |
| `Decision` | Hard fail — logic changed |
| `ConfidenceScore` | Hard fail — scoring changed |

### Failure Case
```go
replayEnv, err := replayEngine.ReplayExecution(originalEnv, differentInput)
// err: "input hash mismatch: original=abc123..., replay=def456..."
```

## Determinism Requirements for System Processors
For a system to be replayable it must:
1. Use fixed timestamps (not `time.Now()`)
2. Use fixed intent IDs (not random/nano-based)
3. Produce identical output for identical input
4. Use no external state that can change between runs

## Adding New Systems to Replay
```go
// In ReplayExecution(), add a new case:
case "your-system-id":
    return r.replayYourSystem(originalEnv, input)
```

## Proven Output
```
Original hash: 847fca8e0b59edeef456ebd5e3234d3ea4bfbf3d3bdf3a9c16d17d626fa15fd4
Replay hash:   847fca8e0b59edeef456ebd5e3234d3ea4bfbf3d3bdf3a9c16d17d626fa15fd4
Deterministic: true
```

## File Location
- **Implementation**: `shared/replay/replay.go`
