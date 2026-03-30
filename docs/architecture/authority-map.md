# Authority Map

## Mission
Cartograph where authority lives in the system. Document every decision point that can affect the ledger. No interpretation. Only structural mapping.

---

## Authority Hierarchy Layers

### Layer 1: Validator Authority (Consensus Delegation)

**Who decides**: The validator set

**What they control**:
- Block proposal (who creates the next block)
- Block validation (accepting or rejecting proposed blocks)
- Transaction ordering (sequence within a block)
- Block finality (when a block becomes irreversible)

**Execution flow**:
1. Validator is selected → Duty to propose block
2. Block is proposed → Other validators receive it
3. Validators validate → Yes/No decision
4. Consensus threshold reached → Block finalized
5. Block added to chain → Authority fulfilled

**Authority scope**:
- CAN: Decide block validity
- CAN: Order transactions
- CAN: Propose fork (by selecting different chain)
- CANNOT: Violate consensus rules
- CANNOT: Revert finalized blocks
- CANNOT: Create invalid blocks (rejected by other validators)

**Invariant**: Validator authority is bound by consensus algorithm
- Proof of Stake: Authority = stake weight
- Proof of Work: Authority = computational power
- Authority delegation: Authority = cryptographic signature verification

---

### Layer 2: Supply Authority (Token Minting)

**Who decides**: The consensus protocol (reward mechanism)

**What they control**:
- When new tokens are created
- How many tokens are created
- Which account receives tokens
- When tokens are destroyed (burning rules)

**Execution flow**:
1. Block is produced → Reward is triggered
2. Reward calculation → How many new tokens?
3. Recipient selection → Which validator, which account?
4. State mutation → New tokens added to supply
5. Supply ledger updated → Authority fulfilled

**Authority scope**:
- CAN: Create tokens per protocol rule
- CAN: Distribute to designated recipients
- CAN: Apply inflation schedule
- CANNOT: Create tokens outside protocol
- CANNOT: Distribute to unexpected accounts
- CANNOT: Change supply in finalized blocks

**Invariant**: Supply authority is algorithmic, not discretionary
- Only the protocol defines new supply
- No validator can create extra tokens
- No authority can retroactively change minting

---

### Layer 3: State Mutation Authority (Ledger Changes)

**Who decides**: Transaction execution engine

**What they control**:
- When account balances change
- When nonce increments
- When contract state updates
- When stake is locked/released

**Execution flow**:
1. Transaction is validated → Structure is correct
2. Signature verified → Authority is proven
3. State transition logic executes → Changes applied
4. Invariants checked → Changes are valid
5. State committed → Authority fulfilled

**Authority scope**:
- CAN: Modify accounts controlled by transaction signer
- CAN: Execute contract code (smart contracts)
- CAN: Transfer value from authorized sender
- CANNOT: Modify accounts without authorization
- CANNOT: Violate fund integrity
- CANNOT: Execute without valid signature

**Invariant**: State mutation requires cryptographic proof
- Transaction is signed by account holder
- Signature is verified before state change
- No bypass to authorization

---

### Layer 4: Fork Authority (Chain Selection)

**Who decides**: Node's fork choice rule

**What they control**:
- Which chain is considered "canonical"
- When to follow a competing chain
- When to reject an alternative fork
- Chain finality threshold

**Execution flow**:
1. Node receives block → Multiple chains possible
2. Fork choice evaluated → Which is "best"?
3. Chain selection logic → Canonical chain selected
4. State committed from selected chain → Authority fulfilled
5. Other chains abandoned → Authority decision complete

**Authority scope**:
- CAN: Evaluate fork weight (stake, work, etc.)
- CAN: Select best chain per algorithm
- CAN: Reorg unfinalized chain
- CANNOT: Finalize arbitrary blocks
- CANNOT: Skip consensus validation
- CANNOT: Violate finality rules

**Invariant**: Fork authority is algorithmic, based on chain weight
- No node can arbitrarily choose chains
- Fork choice must follow protocol
- Finality prevents fork choice on old blocks

---

### Layer 5: Network Authority (Message Propagation)

**Who decides**: Node gossip protocol

**What they control**:
- Which blocks are broadcast
- Which transactions are relayed
- Message validity checking
- Timestamp acceptance windows

**Execution flow**:
1. Node creates/receives message → Must decide if relay
2. Message validation → Is it well-formed?
3. Propagation decision → Broadcast or drop?
4. Message distributed → Other nodes receive
5. Authority fulfilled → Network carries message

**Authority scope**:
- CAN: Validate message format
- CAN: Check message signatures
- CAN: Drop invalid messages
- CAN: Rate limit senders
- CANNOT: Modify message content
- CANNOT: Accept invalid signatures
- CANNOT: Inject false messages

**Invariant**: Network authority is protective, not creative
- Network relays valid messages
- Network rejects invalid messages
- No node creates messages for others

---

## Authority Execution Gates

### Gate 1: Consensus Gate
**Function**: ValidateBlock() → Accept/Reject

**Who controls entry**: Consensus protocol
**What can pass**: Only blocks that satisfy consensus rules
**What is blocked**: Invalid signatures, unauthorized proposers, malformed structures

### Gate 2: State Mutation Gate
**Function**: ApplyTransaction() → Success/Fail

**Who controls entry**: Cryptographic signature verification
**What can pass**: Transactions with valid signatures
**What is blocked**: Unsigned transactions, invalid senders

### Gate 3: Supply Gate
**Function**: MintReward() → NewTokens

**Who controls entry**: Consensus protocol (automatic)
**What can pass**: Only programmed reward amounts
**What is blocked**: Arbitrary minting outside protocol

### Gate 4: Finality Gate
**Function**: FinalizeBlock() → Irreversible

**Who controls entry**: Consensus finality rules
**What can pass**: Blocks meeting finality threshold
**What is blocked**: Reverting finalized blocks

### Gate 5: Network Gate
**Function**: PropagateMessage() → Relay/Drop

**Who controls entry**: Message validation rules
**What can pass**: Valid, well-formed messages
**What is blocked**: Invalid, malformed messages

---

## Authority Separation Model

```
┌────────────────────────────────────────┐
│   VALIDATOR SET (Consensus Authority)  │
│   ├─ Propose blocks                    │
│   ├─ Validate blocks                   │
│   └─ Finalize blocks                   │
└────────────────────────────────────────┘
                    ↓
┌────────────────────────────────────────┐
│   PROTOCOL ENGINE (Supply Authority)   │
│   ├─ Calculate rewards                 │
│   ├─ Mint tokens                       │
│   └─ Update supply                     │
└────────────────────────────────────────┘
                    ↓
┌────────────────────────────────────────┐
│   STATE MACHINE (Mutation Authority)   │
│   ├─ Apply transactions                │
│   ├─ Update accounts                   │
│   └─ Execute contracts                 │
└────────────────────────────────────────┘
                    ↓
┌────────────────────────────────────────┐
│   FORK CHOICE (Chain Authority)        │
│   ├─ Evaluate competing chains         │
│   ├─ Select canonical chain            │
│   └─ Manage reorg logic                │
└────────────────────────────────────────┘
                    ↓
┌────────────────────────────────────────┐
│   NETWORK LAYER (Propagation Auth)     │
│   ├─ Validate messages                 │
│   ├─ Relay valid messages              │
│   └─ Drop invalid messages             │
└────────────────────────────────────────┘
```

---

## Single Source of Truth Per Domain

### Domain 1: Consensus
**Single Authority**: Validator set (weighted by stake or power)
**Bypass**: None (by design)
**Risk**: Validator collusion (mitigated by decentralization)

### Domain 2: Supply
**Single Authority**: Protocol algorithm
**Bypass**: None (hardcoded in consensus)
**Risk**: Protocol bug (mitigated by review)

### Domain 3: State Validity
**Single Authority**: Cryptographic signatures
**Bypass**: None (cryptographically impossible)
**Risk**: Key compromise (mitigated by key management)

### Domain 4: Chain Selection
**Single Authority**: Fork choice rule
**Bypass**: None (client implementation enforces)
**Risk**: Bug in fork choice (mitigated by testing)

### Domain 5: Message Validity
**Single Authority**: Message validation rules
**Bypass**: None (peer rejects invalid messages)
**Risk**: DoS through invalid spam (mitigated by rate limiting)

---

## Authority Flow Under Normal Operation

```
User creates transaction
        ↓
Signed by user key (User authority)
        ↓
Submitted to mempool (Network gate validates signature)
        ↓
Validator selected (Consensus authority)
        ↓
Block proposed with transaction (Validator authority)
        ↓
Other validators validate (Consensus authority)
        ↓
Consensus threshold reached (Consensus authority)
        ↓
Block finalized (Finality authority)
        ↓
Rewards minted (Supply authority)
        ↓
State updated (State mutation authority)
        ↓
Block added to chain (Chain authority confirms)
        ↓
Transaction complete (All authorities aligned)
```

---

## Authority Under Failure Conditions

### Failure: Validator produces invalid block

**Authority check sequence**:
1. Other validators receive block
2. Consensus authority validates → FAILS
3. Block rejected by gate
4. Invalid validator punished (if Proof of Stake)
5. Block never finalized
6. Chain remains secure

### Failure: Double-spend attempt

**Authority check sequence**:
1. First transaction submitted → Signature verified (State authority)
2. First transaction in block → Consensus authority confirms
3. Second transaction submitted → Signature verified
4. Second transaction evaluated → Account nonce already incremented
5. Second transaction fails at state gate
6. Ledger secure

### Failure: Network partition

**Authority check sequence**:
1. One partition continues consensus (Consensus authority)
2. Other partition continues consensus (Consensus authority)
3. Both forks valid under their authority
4. Network rejoins
5. Fork choice authority selects canonical fork (based on weight)
6. Losing fork abandoned
7. Majority chain remains secure

### Failure: Reward minting bug

**Authority check sequence**:
1. Bug in minting logic → Too many tokens created
2. Block produced with invalid supply
3. Validators apply state mutation authority
4. Invariant check: Supply does not match expected
5. Block rejected at consensus gate
6. Bug prevented by verification layer

---

## Summary: Authority is Delegated, Not Absolute

- **Validator authority** is bounded by consensus rules
- **Supply authority** is bounded by protocol algorithm
- **State authority** is bounded by cryptographic proof
- **Chain authority** is bounded by fork choice rules
- **Network authority** is bounded by message validation

No single party has absolute control. Authority is separated across layers. Each layer has a gate. No gate can be bypassed.

