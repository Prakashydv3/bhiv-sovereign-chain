# Failure Recovery Report

**DAY 1 - Task e: Failure Simulation**

## What You're Testing

The node can recover from a crash (forced shutdown) without losing data.

## Commands to Run

### Step 1: Node is running
```bash
# Make sure node is running
# Check it's producing blocks
```

### Step 2: Record current state (in geth console)
```javascript
const lastBlock = eth.blockNumber
const lastHash = eth.getBlock(lastBlock).hash

console.log("Last Block:", lastBlock)
console.log("Last Hash:", lastHash)
```

**SAVE THESE VALUES!**

### Step 3: Force kill the node (simulate crash)
```bash
# Open a NEW terminal (not the geth console)
# Force kill the geth process
taskkill /F /IM geth.exe
```

**This simulates a power failure or crash!**

### Step 4: Restart the node
```bash
.\build\bin\geth.exe --datadir .\node-data --dev --http --http.api eth,net,web3
```

**Watch the startup logs - you should see:**
```
WARN [timestamp] Unclean shutdown detected
INFO [timestamp] Repairing blockchain database
INFO [timestamp] Loaded most recent local header
```

### Step 5: Verify recovery (attach console)
```bash
.\build\bin\geth.exe attach http://localhost:8545
```

```javascript
// Check the last block is still there
eth.blockNumber
// Should be >= the lastBlock you saved

// Verify the block hash is still correct
eth.getBlock(lastBlock).hash
// Should match the hash you saved

// Try to query all blocks - none should be missing
for (let i = 0; i <= lastBlock; i++) {
  if (eth.getBlock(i) === null) {
    console.log("ERROR: Block", i, "is missing!")
  }
}
console.log("All blocks verified!")
```

## Recovery Log Required for Submission

**Before Crash:**
```
Last Block: 30
Last Hash: 0xabc123...
```

**After Recovery:**
```
Last Block: 30 or higher
Block 30 Hash: 0xabc123...  ← MUST MATCH
All blocks 0-30: Present and valid
```

**Startup Log showing:**
```
WARN Unclean shutdown detected
INFO Repairing blockchain database
INFO Loaded most recent local header
```

Save to `crash-recovery-proof.txt`

## What This Proves
✓ Node recovers from crash
✓ No data loss
✓ No corruption
✓ Automatic repair works
