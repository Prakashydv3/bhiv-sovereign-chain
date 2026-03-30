# Authority Boundary Report

**DAY 2 - Task c: Authority Boundary Verification**

## What You're Testing

The node rejects invalid operations (you can't cheat the blockchain).

## Commands to Run

### Test 1: Try to send transaction without enough balance

```javascript
// Create a new account (has 0 balance)
const newAccount = personal.newAccount("test")
console.log("New account:", newAccount)

// Check balance (should be 0)
eth.getBalance(newAccount)
// Output: 0

// Try to send 10 ETH (which we don't have)
personal.unlockAccount(newAccount, "test")
eth.sendTransaction({
  from: newAccount,
  to: eth.accounts[0],
  value: web3.toWei(10, "ether")
})
```

**Expected: ERROR - "insufficient funds"**

### Test 2: Try to send transaction with invalid signature

```javascript
// Try to send a raw transaction with garbage data
eth.sendRawTransaction("0x1234567890")
```

**Expected: ERROR - "invalid transaction"**

### Test 3: Try to query non-existent block

```javascript
// Try to get block 999999 (doesn't exist)
eth.getBlock(999999)
```

**Expected: null (not an error, just returns null)**

### Test 4: Try to modify state directly (impossible via RPC)

```javascript
// There's no RPC method to directly change account balance
// This proves the node enforces transaction-based state changes only
```

## Rejection Log Required for Submission

**Test Results:**
```
Test 1: Send without balance
Command: eth.sendTransaction({from: 0x..., value: 10 ETH})
Result: Error: insufficient funds for gas * price + value
Status: ✓ REJECTED (correct)

Test 2: Invalid transaction
Command: eth.sendRawTransaction("0x1234567890")
Result: Error: invalid transaction
Status: ✓ REJECTED (correct)

Test 3: Non-existent block
Command: eth.getBlock(999999)
Result: null
Status: ✓ HANDLED (correct)
```

Save to `authority-boundary-proof.txt`

## What This Proves
✓ Invalid transactions are rejected
✓ Insufficient balance is checked
✓ Invalid signatures are rejected
✓ State can only be modified through valid transactions
✓ Node enforces protocol rules
