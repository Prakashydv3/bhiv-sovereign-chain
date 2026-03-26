// artifact-hash-generator: CLI tool to hash a canonical L2 artifact JSON file.
// Usage: go run ./tools/artifact-hash <artifact.json>
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"bhiv-sovereign-chain/shared/hashing"
)

// Artifact is the canonical L2 state root artifact.
// Field order is the serialisation contract — do not reorder.
type Artifact struct {
	ChainID              string `json:"chain_id"`
	BlockHeight          uint64 `json:"block_height"`
	Timestamp            int64  `json:"timestamp"`
	ApplicationStateRoot string `json:"application_state_root"`
	RegistrySnapshot     string `json:"registry_snapshot"`
	ProjectionLogRoot    string `json:"projection_log_root"`
	ReplayProofHash      string `json:"replay_proof_hash"`
	ParentAnchorID       string `json:"parent_anchor_id"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: artifact-hash <artifact.json>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	var a Artifact
	if err := json.Unmarshal(data, &a); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing artifact: %v\n", err)
		os.Exit(1)
	}

	canonical, err := json.Marshal(a)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshalling: %v\n", err)
		os.Exit(1)
	}

	h := hashing.Sum(canonical)
	fmt.Printf("canonical_input : %s\n", canonical)
	fmt.Printf("artifact_hash   : %x\n", h)
}
