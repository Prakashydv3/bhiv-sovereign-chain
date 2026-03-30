# Anchor Interface

## Boundary Definition

### Contract Responsibility
**ONLY**: Store immutable hash commitments

### Contract Does NOT
-   Validate artifact content
-   Enforce governance rules
-   Check authority/permissions
-   Interpret artifact types
-   Verify parent existence
-   Execute business logic
-   Manage tokens
-   Store registry data
-   Validate policies

---

## Interface Contract

### Write Interface
```solidity
function createAnchor(
    string memory _artifactType,
    bytes32 _artifactHash,
    bytes32 _parentHash,
    uint8 _schemaVersion
) external returns (bytes32 anchorId)
```

**Accepts**: Artifact metadata  
**Returns**: Deterministic anchor ID  
**Guarantees**: Immutable storage  
**Does NOT**: Validate semantics

### Read Interface
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

**Accepts**: Anchor ID  
**Returns**: Stored anchor data  
**Guarantees**: Unchanged data  
**Does NOT**: Interpret data

### Event Interface
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

**Purpose**: Off-chain indexing  
**Does NOT**: Trigger on-chain logic

---

## No Governance Coupling

### Absent Elements
```
✗ No owner/admin role
✗ No access control
✗ No permission checks
✗ No policy validation
✗ No voting mechanism
✗ No proposal system
✗ No authority registry
✗ No role management
```

### Why
Off-chain system maintains full sovereignty over:
- Who can anchor
- What can be anchored
- When anchoring is valid
- How artifacts are interpreted

**Blockchain**: Storage spine only

---

## No Authority Assumptions

### Contract Does NOT Assume
-   Caller identity matters
-   Artifact type has meaning
-   Parent relationship is valid
-   Timestamp ordering is enforced
-   Schema version compatibility

### Contract ONLY Assumes
-  Caller provides valid parameters
-  Storage is permanent
-  Events are emitted

**Authority**: Remains entirely off-chain

---

## Chain-Agnostic Structure

### No Dependency On
```
✗ Consensus mechanism (PoW/PoS/etc)
✗ Block finality semantics
✗ Chain-specific features
✗ Network parameters
✗ Gas pricing models
✗ Validator sets
```

### Only Uses
```
 Basic storage (mapping)
 Basic types (bytes32, string, uint)
 Block timestamp (for ordering only)
 Events (for indexing)
```

**Portability**: Contract works on any EVM chain

---

## Replaceability Constraint

### Design Principle
**If blockchain is replaced, off-chain system remains sovereign.**

### Implementation
1. **No Chain Finality Assumptions**
   - Contract doesn't rely on finality semantics
   - Off-chain system decides when to trust anchors

2. **No Consensus Coupling**
   - No validator checks
   - No consensus parameter dependencies

3. **No Authority On-Chain**
   - All authority logic lives off-chain
   - Contract is pure storage

4. **Portable Structure**
   - Can migrate to different chain
   - Can run on multiple chains simultaneously
   - Off-chain system chooses which chain(s) to trust

### Migration Scenario
```
Current: Ethereum
Future: Polygon / Arbitrum / Custom Chain

Action Required:
1. Deploy same contract on new chain
2. Off-chain system points to new contract
3. Continue anchoring

Off-chain system: UNCHANGED
Registry logic: UNCHANGED
Authority model: UNCHANGED
```

**Result**: Blockchain is replaceable infrastructure

---

## Sovereignty Boundaries

### Off-Chain (Sovereign)
```
 Registry storage
 Policy validation
 Authority checks
 Governance logic
 Business rules
 Artifact interpretation
 Parent validation
 Type semantics
```

### On-Chain (Spine)
```
 Hash commitment storage
 Immutability guarantee
 Temporal ordering
 Event emission
```

**Separation**: Complete

---

## Interface Guarantees

### What Contract Guarantees
1. Anchors are immutable after creation
2. Anchors are retrievable by ID
3. Events are emitted for indexing
4. Storage is append-only
5. No mutation paths exist

### What Contract Does NOT Guarantee
1. Artifact validity
2. Parent existence
3. Type correctness
4. Authority compliance
5. Business rule enforcement

**Boundary**: Clear and minimal

---

## Anti-Patterns Avoided

###   Governance Creep
```solidity
// AVOIDED:
modifier onlyGovernor() { ... }
function updatePolicy() onlyGovernor { ... }
```

###   Authority Checks
```solidity
// AVOIDED:
require(authorizedUsers[msg.sender], "Not authorized");
```

###   Semantic Validation
```solidity
// AVOIDED:
require(artifactType == "valid_type", "Invalid type");
require(parentExists(parentHash), "Parent not found");
```

###   Execution Logic
```solidity
// AVOIDED:
function executeAnchor(bytes32 id) { ... }
function processArtifact(bytes32 hash) { ... }
```

**Result**: Pure storage contract

---

## Chain Replacement Readiness

### Checklist
-  No chain-specific features used
-  No finality assumptions made
-  No consensus dependencies
-  No authority on-chain
-  Portable data structure
-  Off-chain system sovereign
-  Contract is pure infrastructure

### Replacement Process
```
1. Deploy contract on new chain
2. Update off-chain configuration
3. Continue operations
```

**Downtime**: Zero (off-chain system unchanged)

---

## Final Interface Summary

### Minimal Surface
- 1 write function (createAnchor)
- 1 read function (getAnchor)
- 1 event (AnchorCreated)
- 0 admin functions
- 0 governance functions
- 0 execution functions

### Clear Boundaries
- Storage: On-chain
- Logic: Off-chain
- Authority: Off-chain
- Interpretation: Off-chain

### Replaceability
- Chain-agnostic design
- No consensus coupling
- Off-chain sovereignty preserved
- Infrastructure is replaceable

---

## Conclusion

 **INTERFACE BOUNDARIES DEFINED**

The anchor contract is:
- Minimal (3 functions total)
- Immutable (append-only)
- Agnostic (no semantic logic)
- Replaceable (no chain coupling)
- Sovereign (off-chain authority)

**Blockchain role**: Spine, not brain.  
**Off-chain role**: Brain, memory, execution.  
**Separation**: Complete and enforced.
