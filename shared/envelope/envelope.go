package envelope

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// ExecutionStatus represents the lifecycle state of an execution
type ExecutionStatus string

const (
	StatusPending   ExecutionStatus = "pending"
	StatusExecuted  ExecutionStatus = "executed"
	StatusAnchored  ExecutionStatus = "anchored"
	StatusFailed    ExecutionStatus = "failed"
)

// ExecutionEnvelope is the canonical format for all BHIV system executions.
// It is the single unit that flows through: execution → hash → signing → L1 anchor.
type ExecutionEnvelope struct {
	ExecutionID     string          `json:"execution_id"`
	IntentID        string          `json:"intent_id"`
	SystemID        string          `json:"system_id"`
	AgentID         string          `json:"agent_id"`
	InputHash       [32]byte        `json:"input_hash"`
	OutputHash      [32]byte        `json:"output_hash"`
	Decision        string          `json:"decision"`
	ConfidenceScore float64         `json:"confidence_score"`
	PolicyID        string          `json:"policy_id"`
	DatasetIDs      []string        `json:"dataset_ids"`
	Timestamp       int64           `json:"timestamp"`
	EnvironmentHash [32]byte        `json:"environment_hash"`
	ExecutionStatus ExecutionStatus `json:"execution_status"`
	AnchorID        string          `json:"anchor_id,omitempty"`
}

// NewEnvelope creates a new execution envelope with a deterministic execution_id
// derived from systemID + agentID + intentID — no timestamps, no randomness.
func NewEnvelope(systemID, agentID, intentID string) *ExecutionEnvelope {
	execID := fmt.Sprintf("%x", sha256.Sum256([]byte(systemID+"|"+agentID+"|"+intentID)))
	return &ExecutionEnvelope{
		ExecutionID:     execID,
		IntentID:        intentID,
		SystemID:        systemID,
		AgentID:         agentID,
		Timestamp:       0,
		ConfidenceScore: 0.0,
		DatasetIDs:      make([]string, 0),
		ExecutionStatus: StatusPending,
	}
}

// Hash computes a deterministic SHA-256 hash of the envelope.
// AnchorID is excluded from the hash to avoid circular dependency.
func (e *ExecutionEnvelope) Hash() [32]byte {
	type hashable struct {
		ExecutionID     string          `json:"execution_id"`
		IntentID        string          `json:"intent_id"`
		SystemID        string          `json:"system_id"`
		AgentID         string          `json:"agent_id"`
		InputHash       [32]byte        `json:"input_hash"`
		OutputHash      [32]byte        `json:"output_hash"`
		Decision        string          `json:"decision"`
		ConfidenceScore float64         `json:"confidence_score"`
		PolicyID        string          `json:"policy_id"`
		DatasetIDs      []string        `json:"dataset_ids"`
		Timestamp       int64           `json:"timestamp"`
		EnvironmentHash [32]byte        `json:"environment_hash"`
		ExecutionStatus ExecutionStatus `json:"execution_status"`
	}
	h := hashable{
		ExecutionID: e.ExecutionID, IntentID: e.IntentID, SystemID: e.SystemID,
		AgentID: e.AgentID, InputHash: e.InputHash, OutputHash: e.OutputHash,
		Decision: e.Decision, ConfidenceScore: e.ConfidenceScore, PolicyID: e.PolicyID,
		DatasetIDs: e.DatasetIDs, Timestamp: e.Timestamp, EnvironmentHash: e.EnvironmentHash,
		ExecutionStatus: e.ExecutionStatus,
	}
	data, _ := json.Marshal(h)
	return sha256.Sum256(data)
}

// MarkAnchored updates the envelope status and records the anchor ID.
func (e *ExecutionEnvelope) MarkAnchored(anchorID string) {
	e.AnchorID = anchorID
	e.ExecutionStatus = StatusAnchored
}

// Validate ensures the envelope is complete and valid before anchoring.
func (e *ExecutionEnvelope) Validate() error {
	if e.ExecutionID == "" {
		return fmt.Errorf("execution_id required")
	}
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