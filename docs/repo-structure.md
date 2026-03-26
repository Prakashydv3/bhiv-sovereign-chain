# BHIV Sovereign Chain — Repo Structure

## Module
`bhiv-sovereign-chain`

## Directory Layout

```
bhiv-sovereign-chain/
├── go.mod                          # Single unified Go module
│
├── shared/
│   └── hashing/
│       └── hashing.go              # Deterministic SHA-256 primitives (Sum, SumHex, SumJSON, CombineHashes)
│
├── l1-core/
│   └── anchor/
│       ├── anchor.go               # L1 anchor registry: Submit, Query, VerifyAnchor
│       └── anchor_test.go          # Anchor determinism + verification tests
│
├── l2-execution/
│   ├── l2.go                       # L2 Snapshot struct + StateRoot computation
│   └── l2_test.go                  # Determinism, order-independence, field-sensitivity tests
│
├── integration/
│   └── e2e_demo.go                 # End-to-end: L2 snapshot → state root → L1 anchor → verify
│
├── tools/
│   └── artifact-hash/
│       └── main.go                 # CLI: hash a canonical artifact JSON file
│
├── agent-layer/                    # Reserved — Sankalp / Abhishek integration
│
└── docs/
    └── repo-structure.md           # This file
```

## Merged Sources

| Prior Repo | What Was Merged | Destination |
|---|---|---|
| `New folder/sovereign-anchor-system/anchor-client/` | `client.go`, `verifier.go`, `client_test.go` | `l1-core/anchor/` |
| `New folder/sovereign-anchor-system/artifact-tools/` | `artifact_schema.go`, `hash_generator.go`, `l2_state.go`, `determinism_test.go` | `l2-execution/`, `shared/hashing/` |
| `L1L2 Anchoring Bridge Preparation/anchor-client/contract/` | `contract.go` (anchor store logic) | `l1-core/anchor/anchor.go` |
| `L1L2 Anchoring Bridge Preparation/artifact-tools/` | `main.go` (canonical JSON hashing) | `tools/artifact-hash/main.go` |
| `L2_blockchain application/artifact-hash-generator/` | `main.go` (artifact struct + hash) | `tools/artifact-hash/main.go` |
| `New folder/sovereign-anchor-system/examples/` | `e2e_anchor_example.go` | `integration/e2e_demo.go` |

## Duplication Removed

- Three separate `Artifact` struct definitions → unified in `tools/artifact-hash`
- Three separate anchor store implementations → single `l1-core/anchor`
- Two separate hashing utilities → single `shared/hashing`
- Simulation-only contract (file-backed JSON store) → in-memory registry with clear production swap point

## Build

```bash
go build ./...
go test ./...
```

## Determinism Contract

- All hashing uses `crypto/sha256` — no randomness, no timestamps in hash inputs
- L2 state entries are sorted by key before hashing — insertion order is irrelevant
- Struct field order defines JSON serialisation order — maps are never used for canonical data

## Integration Points (Future)

| Owner | System | Entry Point |
|---|---|---|
| Siddhesh Narkar | Bucket (storage + logs) | `l1-core/anchor` — anchor records |
| Sankalp | Intelligence (IR/CET) | `agent-layer/` |
| Raj Prajapati | Gurukul + execution flows | `l2-execution/` |
| Abhishek | Enforcement / decision layer | `agent-layer/` |
