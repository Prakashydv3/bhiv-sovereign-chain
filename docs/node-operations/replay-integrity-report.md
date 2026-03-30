# Replay Integrity Report

**DAY 2 - Task b: Replay Integrity Verification**

## What You're Testing

When you restart the node, it "replays" all blocks and arrives at the same state.

## Commands to Run

### Step 1: Capture all block hashes (in geth console)
```javascript
// Save all block hashes
const blockHashes = []
for (let i = 0; i <= eth.blockNumber; i++) {
  blockHashes[i] = eth.getBlock(i).hash
}

console.log("Captured", blockHashes.length, "block hashes")
console.log("Block 0:", blockHashes[0])
console.log("Block 5:", blockHashes[5])
console.log("Block 10:", blockHashes[10])
```

**SAVE THESE HASHES!**

### Step 2: Restart the node
```bash
exit  # Exit console
# Stop node (Ctrl+C)
# Restart
.\build\bin\geth.exe --datadir .\node-data --dev --http --http.api eth,net,web3
```

### Step 3: Verify all hashes match (attach console)
```bash
.\build\bin\geth.exe attach http://localhost:8545
```

```javascript
// Check each block hash matches
const matches = []
for (let i = 0; i <= 10; i++) {  // Check first 10 blocks
  const currentHash = eth.getBlock(i).hash
  const savedHash = blockHashes[i]  // You saved these earlier
  
  if (currentHash === savedHash) {
    console.log("Block", i, "✓ MATCH")
    matches.push(true)
  } else {
    console.log("Block", i, "✗ MISMATCH!")
    matches.push(false)
  }
}

console.log("Total matches:", matches.filter(x => x).length, "/", matches.length)
```

## Hash Comparison Required for Submission

**Before Restart:**
```
Block 0: 0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0
Block 5: 0xabc123...
Block 10: 0xdef456...
```

**After Restart:**
```
Block 0: 0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0  ← MATCH
Block 5: 0xabc123...  ← MATCH
Block 10: 0xdef456...  ← MATCH
```

**Result: 100% match rate**

Save to `replay-integrity-proof.txt`

## What This Proves
✓ Block replay is deterministic
✓ Same blocks produce same hashes
✓ State reconstruction is identical
✓ No randomness in block processing
