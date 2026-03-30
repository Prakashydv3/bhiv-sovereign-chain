# Observability Map

**DAY 2 - Task d: Observability Mapping**

## What You're Documenting

How to monitor the node's health and status.

## Key Signals to Monitor

### 1. Logs (in node terminal)

| Log Message | What It Means | Action |
|-------------|---------------|--------|
| "Starting Geth" | Node is starting | Wait for "HTTP server started" |
| "HTTP server started" | Node is ready | Can now connect |
| "Imported new chain segment" | Blocks being added | Normal operation |
| "Database closed" | Node shutting down | Clean shutdown |
| "Unclean shutdown detected" | Node crashed before | Will auto-repair |

### 2. Metrics (via geth console)

| Command | What It Shows | Healthy Value |
|---------|---------------|---------------|
| `eth.blockNumber` | Current block height | Increasing over time |
| `net.peerCount` | Number of connected peers | > 0 (if not dev mode) |
| `eth.syncing` | Sync status | `false` (when synced) |
| `eth.mining` | Mining status | `true` (in dev mode) |
| `admin.nodeInfo` | Node information | Shows version, network |

### 3. Block Events

```javascript
// Watch for new blocks
eth.blockNumber  // Check periodically

// If number increases → Node is working
// If number stuck → Node may be frozen
```

## Monitoring Commands to Run

```javascript
// Health check script
function healthCheck() {
  console.log("=== Node Health Check ===")
  console.log("Block Number:", eth.blockNumber)
  console.log("Peer Count:", net.peerCount)
  console.log("Syncing:", eth.syncing)
  console.log("Mining:", eth.mining)
  console.log("========================")
}

// Run it
healthCheck()
```

## Observability Log Required for Submission

**Sample Health Check Output:**
```
=== Node Health Check ===
Block Number: 45
Peer Count: 0 (dev mode)
Syncing: false
Mining: true
========================

Status: ✓ HEALTHY
- Blocks are being produced
- Node is responsive
- No errors in logs
```

Save to `observability-log.txt`

## What This Proves
✓ You can monitor node health
✓ You understand key metrics
✓ You can detect problems
✓ You know what "healthy" looks like
