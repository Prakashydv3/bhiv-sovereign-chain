# Failure Surface Model

## Node Crash

### What Breaks
- Crashed node stops producing blocks
- Crashed node stops validating
- Local mempool lost
- Pending transactions lost

### What Degrades
- Network has fewer validators
- Block production may slow
- Transaction propagation reduced
- Sync capacity reduced

### What Must Never Halt
- Other nodes continue
- Consensus continues
- Block production continues
- Network remains operational

### Recovery
- Restart node
- Resync from peers
- Rejoin validator set
- Resume normal operation

## Partial Network Partition

### What Breaks
- Network splits into groups
- Groups cannot communicate
- Each group continues independently
- Temporary fork created

### What Degrades
- Consensus in minority partition
- Block finalization may pause
- Transaction confirmation delayed
- Network throughput reduced

### What Must Never Halt
- Majority partition continues
- Block production in majority
- State consistency maintained
- Consensus in majority

### Recovery
- Network heals automatically
- Minority adopts majority chain
- Fork resolved via fork choice
- Full consensus restored

## Replay Attack Scenario

### What Breaks
- Nothing (if properly defended)
- Attack is detected and rejected

### What Degrades
- Network bandwidth (attack traffic)
- Node CPU (validation overhead)
- Mempool space (spam)

### What Must Never Halt
- Transaction validation
- Nonce checking
- Signature verification
- Block production

### Defense
- Nonce system prevents replay
- Transaction hash tracking
- Signature verification
- Mempool filtering

## Duplicate Transaction Broadcast

### What Breaks
- Nothing (duplicates filtered)

### What Degrades
- Network bandwidth
- Mempool efficiency
- Propagation speed

### What Must Never Halt
- Transaction processing
- Duplicate detection
- Block production
- Network operation

### Defense
- Transaction hash deduplication
- Mempool duplicate filtering
- Nonce verification
- Efficient propagation

## Block Propagation Delay

### What Breaks
- Nothing (delay is tolerated)

### What Degrades
- Block confirmation time
- Network synchronization
- Fork probability increases
- User experience

### What Must Never Halt
- Block validation
- Consensus mechanism
- State transitions
- Network operation

### Mitigation
- Efficient gossip protocol
- Block compression
- Parallel propagation
- Redundant paths

## Fork Scenario

### What Breaks
- Temporary chain split
- Conflicting blocks exist
- State diverges temporarily

### What Degrades
- Transaction finality delayed
- Network uncertainty
- User confusion
- Validator coordination

### What Must Never Halt
- Fork choice rule execution
- Block validation
- State consistency per chain
- Eventual convergence

### Resolution
- Fork choice rule selects winner
- Losing chain abandoned
- State converges
- Normal operation resumes

## Reward Mint Failure

### What Breaks
- Validator doesn't receive reward
- Supply calculation affected
- Economic incentive broken

### What Degrades
- Validator motivation
- Network security
- Economic model

### What Must Never Halt
- Block production
- Transaction processing
- Network consensus
- State transitions

### Recovery
- Identify failure cause
- Fix minting logic
- Verify supply integrity
- Resume normal rewards