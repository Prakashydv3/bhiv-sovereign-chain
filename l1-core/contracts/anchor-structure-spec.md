# Anchor Structure Specification

## Overview
This document defines the canonical structure of an anchor entry in the minimal anchor contract.

## Anchor Structure

```solidity
struct Anchor {
    bytes32 anchorId;
    string artifactType;
    bytes32 artifactHash;
    bytes32 parentHash;
    uint256 timestampUtc;
    uint8 schemaVersion;
}
```

## Field Definitions

### anchorId
- **Type**: `bytes32`
- **Required**: Yes (auto-generated)
- **Description**: Deterministic identifier derived from content and transaction context
- **Generation**: `keccak256(artifactType + artifactHash + parentHash + timestamp + schemaVersion + sender)`
- **Purpose**: Unique identification of anchor entry

### artifactType
- **Type**: `string`
- **Required**: Yes
- **Description**: Type classification of the anchored artifact
- **Examples**: `"registry_snapshot"`, `"replay_proof"`, `"projection_root"`
- **Validation**: Must be non-empty string
- **Purpose**: Structural classification only (no semantic interpretation)

### artifactHash
- **Type**: `bytes32`
- **Required**: Yes
- **Description**: Cryptographic hash of the artifact being anchored
- **Validation**: Must be non-zero
- **Purpose**: Immutable commitment to off-chain artifact

### parentHash
- **Type**: `bytes32`
- **Required**: No (optional)
- **Description**: Reference to parent anchor's anchorId for chain linking
- **Default**: `bytes32(0)` for root anchors
- **Purpose**: Structural parent-child relationship (no semantic validation)

### timestampUtc
- **Type**: `uint256`
- **Required**: Yes (auto-generated)
- **Description**: Unix timestamp (seconds since epoch) when anchor was created
- **Source**: `block.timestamp`
- **Purpose**: Temporal ordering of anchors

### schemaVersion
- **Type**: `uint8`
- **Required**: Yes
- **Description**: Version of the anchor schema structure
- **Range**: 1-255
- **Purpose**: Future-proofing for schema evolution

## Constraints

### Immutability
- Once created, anchors CANNOT be modified
- No update operations exist
- No delete operations exist
- Append-only structure

### No Semantic Logic
- Contract does NOT validate artifact content
- Contract does NOT interpret artifact_type meaning
- Contract does NOT verify parent_hash validity
- Contract does NOT enforce business rules

### Minimal Surface
- No additional metadata fields
- No expansion beyond defined structure
- No governance logic
- No authority checks
- No execution logic

## Interface

### Write Operation
```solidity
function createAnchor(
    string memory _artifactType,
    bytes32 _artifactHash,
    bytes32 _parentHash,
    uint8 _schemaVersion
) external returns (bytes32)
```

### Read Operation
```solidity
function getAnchor(bytes32 _anchorId) external view returns (
    bytes32 anchorId,
    string memory artifactType,
    bytes32 artifactHash,
    bytes32 parentHash,
    uint256 timestampUtc,
    uint8 schemaVersion
)
```

### Event
```solidity
event AnchorCreated(
    bytes32 indexed anchorId,
    string artifactType,
    bytes32 artifactHash,
    bytes32 parentHash,
    uint256 timestampUtc,
    uint8 schemaVersion
)
```

## Design Principles

1. **Blockchain as Spine**: Storage only, no interpretation
2. **Off-Chain Sovereignty**: All logic remains off-chain
3. **Immutable Commitment**: Hash-based verification
4. **Structural Continuity**: Parent linking without semantic validation
5. **Chain Agnostic**: No dependency on consensus internals

## Usage Pattern

```
Off-Chain System → Generate Artifact Hash → Create Anchor → Emit Event
                                                ↓
                                        Immutable Storage
```

## Non-Goals

-  Governance logic
-  Registry storage
-  Policy validation
-  Authority checks
-  Execution logic
-  Token mechanics
-  Business logic

## Schema Version History

- **Version 1**: Initial minimal anchor structure
