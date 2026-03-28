package envelope

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// ExecutionEnvelope defines the canonical format for all BHIV executions
type ExecutionEnvelope struct {
	IntentID         string   `json:"intent_id"`
	SystemID         string   `json:"system_id"`
	AgentID          string   `json:"agent_id"`
	InputHash        [32]byte `json:"input_hash"`
	OutputHash       [32]byte `json:"output_hash"`
	Decision         string   `json:"decision"`
	ConfidenceScore  float64  `json:"confidence_score"`
	PolicyID         string   `json:"policy_id"`
	DatasetIDs       []string `json:"dataset_ids"`
	Timestamp        int64    `json:"timestamp"`
	EnvironmentHash  [32]byte `json:"environment_hash"`
}

// NewEnvelope creates a new execution envelope
func NewEnvelope(systemID, agentID string) *ExecutionEnvelope {
	return &ExecutionEnvelope{
		IntentID:        generateIntentID(),
		SystemID:        systemID,
		AgentID:         agentID,
		Timestamp:       0, // Set explicitly for determinism
		ConfidenceScore: 0.0,
		DatasetIDs:      make([]string, 0),
	}
}

// Hash computes deterministic hash of the envelope
func (e *ExecutionEnvelope) Hash() [32]byte {
	data, _ := json.Marshal(e)
	return sha256.Sum256(data)
}

// Validate ensures envelope is complete and valid
func (e *ExecutionEnvelope) Validate() error {
	if e.SystemID == "" {
		return fmt.Errorf("system_id required")
	}
	if e.AgentID == "" {
		return fmt.Errorf("agent_id required")
	}
	if e.Decision == "" {
		return fmt.Errorf("decision required")
	}
	if e.ConfidenceScore < 0 || e.ConfidenceScore > 1 {
		return fmt.Errorf("confidence_score must be 0-1")
	}
	return nil
}

func generateIntentID() string {
	return "intent_deterministic_001" // Fixed for determinism
}