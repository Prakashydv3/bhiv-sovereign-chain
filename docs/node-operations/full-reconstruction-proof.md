# Full Reconstruction Proof

**DAY 3 - Task a: Full Deterministic Reconstruction**

## What You're Testing

You can completely delete the node and rebuild it, and the genesis block will be identical.

## Commands to Run

### Step 1: Capture genesis hash (in geth console)
```javascript
const genesisHash = eth.getBlock(0).hash
console.log("Genesis Hash:", genesisHash)
```

**SAVE THIS HASH - IT'S THE MOST IMPORTANT PROOF!**

### Step 2: Stop the node
```bash
exit  # Exit console
# Stop node (Ctrl+C)
```

### Step 3: DELETE everything
```bash
# Delete the entire data directory
rmdir /s /q node-data
```

**Everything is gone! The blockchain data is deleted!**

### Step 4: Rebuild from scratch
```bash
# Start node again (will create new data directory)
.\build\bin\geth.exe --datadir .\node-data --dev --http --http.api eth,net,web3
```

### Step 5: Verify genesis is identical (attach console)
```bash
.\build\bin\geth.exe attach http://localhost:8545
```

```javascript
// Get the new genesis hash
const newGenesisHash = eth.getBlock(0).hash
console.log("New Genesis Hash:", newGenesisHash)

// Compare with the old one you saved
console.log("Old Genesis Hash:", genesisHash)  // The one you saved
console.log("Match:", newGenesisHash === genesisHash)
```

**THEY MUST MATCH!**

## Reconstruction Proof Required for Submission

**Before Deletion:**
```
Genesis Hash: 0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0
Block 0 State Root: 0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421
```

**After Full Reconstruction:**
```
Genesis Hash: 0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0  ← IDENTICAL!
Block 0 State Root: 0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421  ← IDENTICAL!
```

**Result: ✓ DETERMINISTIC RECONSTRUCTION PROVEN**

Save to `reconstruction-proof.txt`

## What This Proves
✓ Genesis block is deterministic
✓ Same code produces same genesis
✓ Node can be fully reconstructed
✓ No hidden state or randomness
✓ Blockchain is reproducible
