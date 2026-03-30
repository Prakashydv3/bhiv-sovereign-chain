# Peer Connectivity Report

**DAY 2 - Task a: Peer Connectivity Observation**

## What You're Testing

How the node connects to other nodes (peers) on the network.

## Note About Dev Mode

In `--dev` mode, networking is disabled. For this task, we'll observe the peer system without actually connecting to mainnet.

## Commands to Run

### Step 1: Start node with networking enabled (testnet)
```bash
# Stop dev mode node
# Start in testnet mode (Sepolia)
.\build\bin\geth.exe --datadir .\node-data-sepolia --sepolia --http --http.api eth,net,web3,admin
```

**This will start syncing with Sepolia testnet**

### Step 2: Check peer connections (in geth console)
```bash
# Attach to the node
.\build\bin\geth.exe attach http://localhost:8545
```

```javascript
// Check how many peers are connected
net.peerCount
// Example: 5

// List all connected peers
admin.peers
// Shows array of peer objects with IDs and addresses

// Check sync status
eth.syncing
// Shows current sync progress
```

### Step 3: Observe block synchronization
```javascript
// Watch blocks being downloaded
eth.blockNumber
// Will increase as blocks sync

// Check sync progress
eth.syncing
// Shows: currentBlock, highestBlock, startingBlock
```

## Peer Observation Log Required for Submission

**Peer Status:**
```
Peer Count: 5
Sample Peer: {
  id: "abc123...",
  name: "Geth/v1.13.0",
  network: {
    remoteAddress: "192.168.1.100:30303"
  }
}
```

**Sync Status:**
```
Current Block: 1000
Highest Block: 5000000
Syncing: true
```

Save to `peer-connectivity-log.txt`

## What This Proves
✓ Node can discover peers
✓ Node can connect to peers
✓ Block synchronization works
✓ Network protocol is functional

## Alternative (If you don't want to sync testnet)

Just document the peer system in dev mode:
```javascript
// In dev mode
net.peerCount  // Returns 0 (no peers in dev mode)
admin.nodeInfo  // Shows your node's info
```

This proves you understand the peer system even if not actively using it.
