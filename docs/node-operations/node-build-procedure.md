# Node Build Procedure

**DAY 1 - Task a: Environment Reconstruction**

## What You Did

You built the Geth blockchain node from source code.

## Commands You Ran

### 1. Navigate to go-ethereum folder
```bash
cd d:\Go Lang\MP\Task5\go-ethereum
```

### 2. Build the Geth binary
```bash
go build -o build\bin\geth.exe .\cmd\geth
```

### 3. Verify the build worked
```bash
.\build\bin\geth.exe version
```

**Expected Output:**
```
Geth
Version: 1.13.x-stable
Architecture: amd64
Go Version: go1.21.x
Operating System: windows
```

### 4. Start the node in dev mode
```bash
.\build\bin\geth.exe --datadir .\node-data --dev --http --http.api eth,net,web3
```

**What this does:**
- `--datadir .\node-data` = Store blockchain data here
- `--dev` = Developer mode (instant block mining)
- `--http` = Enable HTTP-RPC server
- `--http.api eth,net,web3` = Enable these APIs

## Terminal Log Required for Submission

**Save this output to a file:**
```bash
.\build\bin\geth.exe version > build-proof.txt
```

This proves you successfully built the node.

## What This Proves
✓ You can build blockchain node from source
✓ Binary is executable
✓ Node can start successfully
