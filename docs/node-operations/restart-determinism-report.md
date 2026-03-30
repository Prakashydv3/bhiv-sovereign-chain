# Restart Determinism Report

**DAY 1 - Task c: Deterministic Restart Test**

## What You're Testing

When you restart the node, all blocks remain exactly the same (same hashes).

## Commands to Run

### Step 1: Node is running and producing blocks
```bash
# Node should be running from previous task
# In dev mode, it auto-produces blocks
```

### Step 2: Check current state (in geth console)
```javascript
// Check how many blocks exist
eth.blockNumber
// Example output: 15

// Save block 10's hash
const block10Hash = eth.getBlock(10).hash
console.log("Block 10 hash:", block10Hash)
// Example: "0xabc123..."

// Save block 10's state root
const block10StateRoot = eth.getBlock(10).stateRoot
console.log("Block 10 state root:", block10StateRoot)
```

**COPY THESE VALUES - YOU'LL COMPARE THEM AFTER RESTART!**

### Step 3: Stop the node
```bash
# In the geth console, type:
exit

# Or press Ctrl+C in the node terminal
```

### Step 4: Restart the node
```bash
.\build\bin\geth.exe --datadir .\node-data --dev --http --http.api eth,net,web3
```

### Step 5: Attach again and verify
```bash
# In new terminal
.\build\bin\geth.exe attach http://localhost:8545
```

```javascript
// Check block number (should be >= 15)
eth.blockNumber

// Check block 10's hash - MUST BE IDENTICAL
eth.getBlock(10).hash
// Compare with the hash you saved earlier

// Check block 10's state root - MUST BE IDENTICAL
eth.getBlock(10).stateRoot
// Compare with the state root you saved earlier
```

## Hash Comparison Required for Submission

**Before Restart:**
```
Block 10 Hash: 0xabc123...
Block 10 State Root: 0xdef456...
```

**After Restart:**
```
Block 10 Hash: 0xabc123...  ← MUST MATCH
Block 10 State Root: 0xdef456...  ← MUST MATCH
```

Save this comparison to `restart-comparison.txt`

## What This Proves
✓ Restart doesn't change historical blocks
✓ Block hashes remain identical
✓ State roots remain identical
✓ No corruption from restart
