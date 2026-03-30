# State Persistence Report

**DAY 1 - Task d: State Persistence Validation**

## What You're Testing

The blockchain state (account balances, contract data) persists correctly across restarts.

## Commands to Run

### Step 1: Create an account (in geth console)
```javascript
// Create a new account
personal.newAccount("password123")
// Output: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"

// Check the account balance (should be 0)
eth.getBalance("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb")
// Output: 0
```

### Step 2: Send some ETH to the account
```javascript
// In dev mode, eth.accounts[0] has unlimited ETH
eth.sendTransaction({
  from: eth.accounts[0],
  to: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  value: web3.toWei(10, "ether")
})

// Wait a moment for the transaction to be mined
// Check balance again
eth.getBalance("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb")
// Output: 10000000000000000000 (10 ETH in wei)
```

**SAVE THIS BALANCE!**

### Step 3: Record current state
```javascript
const accountAddress = "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
const balanceBefore = eth.getBalance(accountAddress)
const blockNumber = eth.blockNumber

console.log("Account:", accountAddress)
console.log("Balance:", balanceBefore.toString())
console.log("Block Number:", blockNumber)
```

### Step 4: Restart node
```bash
# Exit console
exit

# Stop node (Ctrl+C)
# Restart node
.\build\bin\geth.exe --datadir .\node-data --dev --http --http.api eth,net,web3
```

### Step 5: Verify state persisted
```bash
# Attach again
.\build\bin\geth.exe attach http://localhost:8545
```

```javascript
// Check the balance - MUST BE IDENTICAL
eth.getBalance("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb")
// Should still be: 10000000000000000000
```

## State Comparison Required for Submission

**Before Restart:**
```
Account: 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb
Balance: 10000000000000000000 wei (10 ETH)
Block Number: 25
```

**After Restart:**
```
Account: 0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb
Balance: 10000000000000000000 wei (10 ETH)  ← MUST MATCH
Block Number: 25 or higher
```

Save to `state-persistence-proof.txt`

## What This Proves
✓ Account balances persist across restarts
✓ State is not corrupted
✓ Database writes are durable
