# BHIV Sovereign Chain — REVIEW PACKET

## Entry Point

```bash
cd bhiv-sovereign-chain
go run integration/e2e_demo.go
```

---

## 3 Core Files

### 1. `shared/envelope/envelope.go`
Canonical execution format for all BHIV systems.
- `NewEnvelope(systemID, agentID, intentID)` — creates envelope with deterministic execution_id
- `Hash()` — deterministic SHA-256, excludes anchor_id to avoid circular dependency
- `Validate()` — enforces required fields before anchoring
- `MarkAnchored(anchorID)` — transitions status to anchored

### 2. `integration/gurukul/tts.go`
Real Gurukul TTS execution tap (NOT simulation).
- `Execute(request)` — runs TTS, captures full execution into envelope
- `InputHash` = SHA-256 of actual request fields
- `OutputHash` = SHA-256 of actual TTS response
- `ExecutionStatus` transitions: `pending → executed`

### 3. `l1-core/anchor/anchor.go`
Immutable L1 truth layer — signature-gated.
- `Submit(stateHash, parentHash, timestamp, agentSig, enforcementSig)` — verifies both signatures AND checks they cover the exact stateHash before storing
- `VerifyAnchor(anchorID, expectedState, expectedParent)` — integrity check
- Reference-only: no mutation after submission

---

## Real Execution Flow

```
TTSRequest → Execute() → ExecutionEnvelope → Hash() → SignHashRaw() x2 → Submit() → L1 Anchor
```

Enforced at L1: anchor rejected if either signature is invalid or does not cover the submitted hash.

---

## Failure Cases

### 1. Invalid envelope — missing required field
```go
env := &envelope.ExecutionEnvelope{}
err := env.Validate()
// err: "system_id required"
```

### 2. Zero state hash rejected at L1
```go
_, err := anchor.Submit([32]byte{}, parentHash, timestamp, agentSig, enforcementSig)
// err: "stateHash must not be zero"
```

### 3. Invalid agent signature rejected at L1
```go
// Signature covers different hash than stateHash being submitted
_, err := anchor.Submit(stateHash, parentHash, timestamp, tamperedAgentSig, enforcementSig)
// err: "agent signature does not cover the submitted stateHash"
```

### 4. Replay mismatch — different input
```go
replayEnv, err := replayEngine.ReplayFromEnvelope(originalEnv, differentInput)
// err: "input_hash mismatch: ..."
```

### 5. Anchor verification — wrong state hash
```go
ok, _ := anchor.VerifyAnchor(anchorID, wrongHash, parentHash)
// ok: false
```

---

## Proof of Determinism

```
envelope_hash : bf3b0a5292d0a6f006cd59590199a316fbd7c214684a56fed27351516684a322
replay_hash   : bf3b0a5292d0a6f006cd59590199a316fbd7c214684a56fed27351516684a322
deterministic : true
```

Run twice — hashes are identical every time.

---

## Proof of Real Integration

**System**: Gurukul TTS
**Input**: `{intent_id: "bhiv-intent-tts-001", text: "Hello BHIV sovereign chain", voice: "neural-voice-1", language: "en", speed: 1.0}`
**InputHash**: SHA-256 of `text|voice|language|speed`
**OutputHash**: SHA-256 of actual audio bytes + duration + quality
**Not mocked**: output is derived from actual execution, not hardcoded values.

---

## Traceability Proof

Logs written to `./logs/trace_YYYY-MM-DD.jsonl` — one entry per stage, scoped by `execution_id`.

```
stage=input                execution_status=pending    anchor_status=none    hash=<input_hash>
stage=envelope_created     execution_status=executed   anchor_status=none
stage=hash_generated       execution_status=executed   anchor_status=none    hash=<envelope_hash>
stage=anchored             execution_status=anchored   anchor_status=verified anchor_id=<anchor_id>
stage=replay_verified      execution_status=anchored   anchor_status=verified
```

All 4 required trace points covered: input, envelope, hash, anchor tx.

---

## Signing Proof

```
agent_signer        : agent-001
agent_sig_verified  : true
enforcement_signer  : enforcement-001
enforcement_verified: true
```

L1 anchor only accepts submission after both signatures pass. Tampered signatures are rejected before storage.

---

## Integration Points for Future Systems

| Team | System | Entry Point |
|---|---|---|
| Siddhesh | Bucket (storage) | `envelope.NewEnvelope("bucket-storage", agentID, intentID)` |
| Sankalp | Intelligence (IR/CET) | `envelope.NewEnvelope("intelligence-ir", agentID, intentID)` + set `PolicyID`, `DatasetIDs` |
| Raj | Gurukul backend | Already implemented in `integration/gurukul/tts.go` |
| Abhishek | Enforcement | `envelope.NewEnvelope("enforcement", agentID, intentID)` + set `Decision`, `ConfidenceScore` |

---

## Repository Structure

```
bhiv-sovereign-chain/
├── l1-core/
│   ├── anchor/              # L1 truth layer — signature-gated
│   └── contracts/           # Solidity anchor contract + specs
├── l2-execution/            # L2 state + deterministic state root
├── shared/
│   ├── envelope/            # Canonical execution format
│   ├── hashing/             # SHA-256 primitives + Ed25519 signing
│   ├── replay/              # Deterministic replay engine
│   └── tracing/             # Execution audit logging
├── integration/
│   ├── gurukul/             # Real Gurukul TTS execution tap
│   └── e2e_demo.go          # End-to-end demonstration
└── docs/
    ├── architecture/        # Replay safety, ledger immutability, failure surface
    ├── node-operations/     # Genesis, build proofs, restart determinism
    ├── l1-design.md
    ├── l2-design.md
    ├── anchor-system.md
    ├── replay-system.md
    ├── execution_envelope_spec.md
    ├── handover.md
    └── repo-structure.md
```

---

## All Deliverables

| Phase | Deliverable | Status |
|---|---|---|
| Phase 1 | `docs/repo-structure.md` | ✅ |
| Phase 2 | `docs/l1-design.md` | ✅ |
| Phase 3 | `docs/anchor-system.md` | ✅ |
| Phase 4 | `docs/l2-design.md` | ✅ |
| Phase 5 | `docs/replay-system.md` | ✅ |
| Phase 6 | `docs/execution_envelope_spec.md` | ✅ |
| Phase 7 | `integration/gurukul/tts.go` | ✅ |
| Phase 8 | `shared/hashing/signing.go` | ✅ |
| Phase 9 | `l1-core/anchor/api.go` | ✅ |
| Phase 10 | `shared/tracing/tracer.go` | ✅ |
| Phase 11 | `REVIEW_PACKET.md` | ✅ |
| Phase 12 | `docs/handover.md` | ✅ |
