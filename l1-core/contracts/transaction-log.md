# Transaction Log - Anchor Creation

## Overview
This document demonstrates the transaction flow for creating anchors on the blockchain.

---

## Transaction 1: Root Registry Snapshot

### Input Parameters
```
artifactType: "registry_snapshot"
artifactHash: 0xa1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2
parentHash: 0x0000000000000000000000000000000000000000000000000000000000000000
schemaVersion: 1
```

### Transaction Details
```
Transaction Hash: 0x1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1c2d3e4f5a6b7c8d9e0f1a2b
Block Number: 12345678
Gas Used: 85432
Timestamp: 1704067200 (2024-01-01 00:00:00 UTC)
Sender: 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb1
```

### Generated Anchor
```
anchorId: 0x7d3e4f2a1b8c9d6e5f4a3b2c1d0e9f8a7b6c5d4e3f2a1b0c9d8e7f6a5b4c3d2e
```

### Event Emitted
```
AnchorCreated(
  anchorId: 0x7d3e4f2a1b8c9d6e5f4a3b2c1d0e9f8a7b6c5d4e3f2a1b0c9d8e7f6a5b4c3d2e,
  artifactType: "registry_snapshot",
  artifactHash: 0xa1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2,
  parentHash: 0x0000000000000000000000000000000000000000000000000000000000000000,
  timestampUtc: 1704067200,
  schemaVersion: 1
)
```

### Status
 **SUCCESS** - Root anchor created

---

## Transaction 2: Second Registry Snapshot (Linked)

### Input Parameters
```
artifactType: "registry_snapshot"
artifactHash: 0xb2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3
parentHash: 0x7d3e4f2a1b8c9d6e5f4a3b2c1d0e9f8a7b6c5d4e3f2a1b0c9d8e7f6a5b4c3d2e
schemaVersion: 1
```

### Transaction Details
```
Transaction Hash: 0x2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1c2d3e4f5a6b7c8d9e0f1a2b3c
Block Number: 12346789
Gas Used: 87654
Timestamp: 1704153600 (2024-01-02 00:00:00 UTC)
Sender: 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb1
```

### Generated Anchor
```
anchorId: 0x8e4f5a3b2c9d0e1f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f
```

### Event Emitted
```
AnchorCreated(
  anchorId: 0x8e4f5a3b2c9d0e1f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f,
  artifactType: "registry_snapshot",
  artifactHash: 0xb2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3,
  parentHash: 0x7d3e4f2a1b8c9d6e5f4a3b2c1d0e9f8a7b6c5d4e3f2a1b0c9d8e7f6a5b4c3d2e,
  timestampUtc: 1704153600,
  schemaVersion: 1
)
```

### Status
 **SUCCESS** - Linked anchor created (parent: 0x7d3e...)

---

## Transaction 3: Replay Proof (Independent)

### Input Parameters
```
artifactType: "replay_proof"
artifactHash: 0xc3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4
parentHash: 0x0000000000000000000000000000000000000000000000000000000000000000
schemaVersion: 1
```

### Transaction Details
```
Transaction Hash: 0x3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1c2d3e4f5a6b7c8d9e0f1a2b3c4d
Block Number: 12347890
Gas Used: 86123
Timestamp: 1704240000 (2024-01-03 00:00:00 UTC)
Sender: 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb1
```

### Generated Anchor
```
anchorId: 0x9f5a6b4c3d0e1f2a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a
```

### Event Emitted
```
AnchorCreated(
  anchorId: 0x9f5a6b4c3d0e1f2a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a,
  artifactType: "replay_proof",
  artifactHash: 0xc3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4,
  parentHash: 0x0000000000000000000000000000000000000000000000000000000000000000,
  timestampUtc: 1704240000,
  schemaVersion: 1
)
```

### Status
 **SUCCESS** - Independent anchor created (no parent)

---

## Transaction 4: Projection Root (Cross-Reference)

### Input Parameters
```
artifactType: "projection_root"
artifactHash: 0xd4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5
parentHash: 0x8e4f5a3b2c9d0e1f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f
schemaVersion: 1
```

### Transaction Details
```
Transaction Hash: 0x4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1c2d3e4f5a6b7c8d9e0f1a2b3c4d5e
Block Number: 12348901
Gas Used: 88765
Timestamp: 1704326400 (2024-01-04 00:00:00 UTC)
Sender: 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb1
```

### Generated Anchor
```
anchorId: 0xa0b6c7d5e4f1a2b3c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b
```

### Event Emitted
```
AnchorCreated(
  anchorId: 0xa0b6c7d5e4f1a2b3c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b,
  artifactType: "projection_root",
  artifactHash: 0xd4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3d4e5,
  parentHash: 0x8e4f5a3b2c9d0e1f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f,
  timestampUtc: 1704326400,
  schemaVersion: 1
)
```

### Status
 **SUCCESS** - Cross-referenced anchor created (parent: 0x8e4f...)

---

## Chain State Summary

### Anchor Count: 4

### Anchor Relationships
```
Registry Chain:
  0x7d3e... (root)
    └─> 0x8e4f... (child)
          └─> 0xa0b6... (projection root reference)

Independent:
  0x9f5a... (replay proof, no parent)
```

### Storage Verification
All anchors are immutable and retrievable via getAnchor(anchorId).

### Event Log Verification
All AnchorCreated events are indexed by anchorId for efficient querying.

---

## Hash Verification Example

### Verify Anchor 0x8e4f...

**Step 1**: Retrieve anchor from contract
```solidity
getAnchor(0x8e4f5a3b2c9d0e1f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f)
```

**Step 2**: Compare artifactHash with off-chain artifact
```
On-chain:  0xb2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3
Off-chain: 0xb2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2c3
Result:  MATCH
```

**Step 3**: Verify parent linkage
```
parentHash: 0x7d3e4f2a1b8c9d6e5f4a3b2c1d0e9f8a7b6c5d4e3f2a1b0c9d8e7f6a5b4c3d2e
Parent exists:  YES
Parent is valid:  YES
```

**Step 4**: Verify timestamp
```
Timestamp: 1704153600 (2024-01-02 00:00:00 UTC)
Reasonable:  YES
```

**Conclusion**: Anchor is valid and immutable.

---

## Gas Cost Analysis

| Operation | Gas Used | Notes |
|-----------|----------|-------|
| Create anchor (no parent) | ~85,000 | Root anchor |
| Create anchor (with parent) | ~87,000 | Linked anchor |
| Read anchor | ~3,000 | View function |

---

## Deployment Information

```
Contract Address: 0x1234567890abcdef1234567890abcdef12345678
Network: [Ethereum/Polygon/etc.]
Deployer: 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb1
Deployment Block: 12345000
Deployment Timestamp: 1704000000
```

---

## Notes

- All anchors are **immutable** after creation
- No update or delete functions exist
- Parent linking is **structural only** (no semantic validation)
- Contract performs **no interpretation** of artifact types or content
- Off-chain system remains **sovereign**
