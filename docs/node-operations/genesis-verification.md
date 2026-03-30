# Genesis Verification

**DAY 1 - Task b: Ledger Initialization Verification**

## What You're Verifying

The genesis block (block 0) is created correctly and has a consistent hash.

## Commands to Run

### 1. Start node (if not running)
```bash
cd d:\Go Lang\MP\Task5\go-ethereum
.\build\bin\geth.exe --datadir .\node-data --dev --http --http.api eth,net,web3
```

### 2. In a NEW terminal, attach to the node
```bash
.\build\bin\geth.exe attach http://localhost:8545
```

### 3. Query the genesis block
```javascript
eth.getBlock(0)
```

**Expected Output:**
```javascript
{
  difficulty: 1,
  extraData: "0x...",
  gasLimit: 11500000,
  gasUsed: 0,
  hash: "0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0",
  miner: "0x0000000000000000000000000000000000000000",
  number: 0,
  parentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
  stateRoot: "0x...",
  timestamp: 0,
  transactions: []
}
```

### 4. Save the genesis hash
```javascript
eth.getBlock(0).hash
```

**Copy this hash - you'll compare it after restart!**

## Terminal Log Required for Submission

**In the geth console, run:**
```javascript
console.log(JSON.stringify(eth.getBlock(0), null, 2))
```

Copy the output and save to `genesis-block.txt`

## What This Proves
✓ Genesis block exists
✓ Block number is 0
✓ Parent hash is all zeros (no parent)
✓ Genesis hash is deterministic
