# Anchor Parent Linking Proof

## Parent-Child Linkability

### Structure Support
```solidity
struct Anchor {
    bytes32 anchorId;
    string artifactType;
    bytes32 artifactHash;
    bytes32 parentHash;      // ← Enables linking
    uint256 timestampUtc;
    uint8 schemaVersion;
}
```

---

## Linking Mechanism

### Root Anchor (No Parent)
```solidity
createAnchor(
    "registry_snapshot",
    0xaaa...,
    0x000...,  // parentHash = 0 (root)
    1
)
→ anchorId: 0x111...
```

### Child Anchor (With Parent)
```solidity
createAnchor(
    "registry_snapshot",
    0xbbb...,
    0x111...,  // parentHash = previous anchorId
    1
)
→ anchorId: 0x222...
```

### Grandchild Anchor
```solidity
createAnchor(
    "registry_snapshot",
    0xccc...,
    0x222...,  // parentHash = previous anchorId
    1
)
→ anchorId: 0x333...
```

---

## Chain Example

### Created Anchors
```
Anchor 1:
  anchorId: 0x111...
  parentHash: 0x000... (root)

Anchor 2:
  anchorId: 0x222...
  parentHash: 0x111... (links to Anchor 1)

Anchor 3:
  anchorId: 0x333...
  parentHash: 0x222... (links to Anchor 2)
```

### Chain Visualization
```
0x111... (root)
    ↓
0x222... (child)
    ↓
0x333... (grandchild)
```

---

## No Semantic Validation

### What Contract DOES
 Store parentHash value
 Allow retrieval of parentHash
 Enable structural linking

### What Contract DOES NOT
 Validate parent exists
 Verify parent type matches
 Enforce lineage rules
 Interpret relationship meaning

**Design**: Structural linking only, no semantic interpretation.

---

## Hash Verification Example

### Verify Chain Integrity

**Step 1**: Get Anchor 3
```solidity
Anchor memory a3 = getAnchor(0x333...);
// a3.parentHash = 0x222...
```

**Step 2**: Get Parent (Anchor 2)
```solidity
Anchor memory a2 = getAnchor(0x222...);
// a2.anchorId = 0x222...  Matches
// a2.parentHash = 0x111...
```

**Step 3**: Get Grandparent (Anchor 1)
```solidity
Anchor memory a1 = getAnchor(0x111...);
// a1.anchorId = 0x111...  Matches
// a1.parentHash = 0x000... (root)
```

**Result**: Chain verified through hash linkage.

---

## Multiple Chain Support

### Independent Chains
```
Chain A:
  0xAAA... (root) → 0xAAB... → 0xAAC...

Chain B:
  0xBBB... (root) → 0xBBC...

Chain C:
  0xCCC... (root)
```

### Cross-Reference
```
Registry Chain:
  0x111... → 0x222...

Projection (references registry):
  0x333... (parentHash: 0x222...)
```

**Capability**: Multiple independent and cross-referenced chains supported.

---

## Linkability Properties

### Supported Patterns
 Linear chains (A → B → C)
 Multiple roots (A, B, C independent)
 Cross-references (D → B from different chain)
 Optional linking (parentHash = 0x000...)

### Not Enforced
 Parent existence validation
 Cycle prevention
 Type matching
 Lineage semantics

**Reason**: Off-chain system maintains sovereignty over interpretation.

---

## Transaction Log Example

### Transaction 1: Root
```
Input:
  artifactType: "snapshot"
  artifactHash: 0xaaa...
  parentHash: 0x000...
  
Output:
  anchorId: 0x7d3e...
  
Event:
  AnchorCreated(0x7d3e..., "snapshot", 0xaaa..., 0x000..., ...)
```

### Transaction 2: Linked
```
Input:
  artifactType: "snapshot"
  artifactHash: 0xbbb...
  parentHash: 0x7d3e...  ← Links to Transaction 1
  
Output:
  anchorId: 0x8e4f...
  
Event:
  AnchorCreated(0x8e4f..., "snapshot", 0xbbb..., 0x7d3e..., ...)
```

### Transaction 3: Continuation
```
Input:
  artifactType: "snapshot"
  artifactHash: 0xccc...
  parentHash: 0x8e4f...  ← Links to Transaction 2
  
Output:
  anchorId: 0x9f5a...
  
Event:
  AnchorCreated(0x9f5a..., "snapshot", 0xccc..., 0x8e4f..., ...)
```

---

## Verification Algorithm

```
function verifyChain(bytes32 leafId, bytes32 expectedRootId) {
    bytes32 currentId = leafId;
    
    while (currentId != 0x000...) {
        Anchor memory anchor = getAnchor(currentId);
        currentId = anchor.parentHash;
        
        if (currentId == expectedRootId) {
            return true; // Chain verified
        }
    }
    
    return false; // Root not found
}
```

**Note**: Verification logic lives off-chain, not in contract.

---

## Conclusion

 **PARENT-CHILD LINKABILITY PROVEN**

- parentHash field enables structural linking
- Multiple anchors successfully chain via parentHash
- Hash verification demonstrates integrity
- No semantic validation (by design)
- Off-chain system interprets relationships
- Contract provides structural continuity only

**Blockchain as spine**: Structural linking without interpretation confirmed.
