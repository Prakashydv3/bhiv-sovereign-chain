# Replay Safety Model

## Transaction Ordering Preservation
- Transactions ordered within blocks deterministically
- Blocks ordered by height sequentially
- Nonce-based ordering per account

## Duplicate Transaction Handling
- Transaction hash tracking
- Nonce checking prevents same-account duplicates
- Block inclusion tracking prevents cross-block duplicates

## Hash Integrity Preservation
- Transaction hash: SHA-256 of transaction data
- Block hash: SHA-256 of block header
- State root hash: Merkle tree root of all state
- Any change breaks hash chain

## Replay Attack Prevention
- Nonce system: Sequential counter per account
- Transaction hash registry: Store executed hashes
- Chain ID: Prevents cross-chain replay
- Expiration timestamps: Limits replay window

## Nondeterminism Leak Points
- System time usage
- Random number generation
- Map iteration order
- Floating point arithmetic
- External data sources
- Concurrency issues
- Network timing
- Hardware differences