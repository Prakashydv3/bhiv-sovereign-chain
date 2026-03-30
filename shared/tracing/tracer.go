package tracing

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"bhiv-sovereign-chain/shared/envelope"
)

// ExecutionLog is the structured observability record for every BHIV execution event.
// Every stage emits one of these — nothing is implicit.
type ExecutionLog struct {
	Timestamp       int64  `json:"timestamp"`
	ExecutionID     string `json:"execution_id"`
	IntentID        string `json:"intent_id"`
	SystemID        string `json:"system_id"`
	Stage           string `json:"stage"`             // input | envelope_created | hash_generated | anchored | replay_verified
	ExecutionStatus string `json:"execution_status"` // pending | executed | anchored | failed
	AnchorStatus    string `json:"anchor_status"`    // none | submitted | verified | failed
	Hash            string `json:"hash,omitempty"`
	AnchorID        string `json:"anchor_id,omitempty"`
	Success         bool   `json:"success"`
	Error           string `json:"error,omitempty"`
}

// Tracer writes structured execution logs to JSONL files.
type Tracer struct {
	logDir string
}

func NewTracer(logDir string) *Tracer {
	os.MkdirAll(logDir, 0755)
	return &Tracer{logDir: logDir}
}

// LogInput logs the raw input capture — the execution tap entry point.
func (t *Tracer) LogInput(executionID, intentID, systemID string) error {
	return t.write(ExecutionLog{
		Timestamp:       time.Now().Unix(),
		ExecutionID:     executionID,
		IntentID:        intentID,
		SystemID:        systemID,
		Stage:           "input",
		ExecutionStatus: "pending",
		AnchorStatus:    "none",
		Success:         true,
	})
}

// LogEnvelopeCreated logs the moment an execution envelope is produced.
func (t *Tracer) LogEnvelopeCreated(env *envelope.ExecutionEnvelope) error {
	return t.write(ExecutionLog{
		Timestamp:       time.Now().Unix(),
		ExecutionID:     env.ExecutionID,
		IntentID:        env.IntentID,
		SystemID:        env.SystemID,
		Stage:           "envelope_created",
		ExecutionStatus: string(env.ExecutionStatus),
		AnchorStatus:    "none",
		Success:         true,
	})
}

// LogHashGenerated logs the computed envelope hash.
func (t *Tracer) LogHashGenerated(env *envelope.ExecutionEnvelope, hash [32]byte) error {
	return t.write(ExecutionLog{
		Timestamp:       time.Now().Unix(),
		ExecutionID:     env.ExecutionID,
		IntentID:        env.IntentID,
		SystemID:        env.SystemID,
		Stage:           "hash_generated",
		ExecutionStatus: string(env.ExecutionStatus),
		AnchorStatus:    "none",
		Hash:            fmt.Sprintf("%x", hash),
		Success:         true,
	})
}

// LogAnchored logs the successful L1 anchor submission.
func (t *Tracer) LogAnchored(env *envelope.ExecutionEnvelope, anchorID string) error {
	return t.write(ExecutionLog{
		Timestamp:       time.Now().Unix(),
		ExecutionID:     env.ExecutionID,
		IntentID:        env.IntentID,
		SystemID:        env.SystemID,
		Stage:           "anchored",
		ExecutionStatus: "anchored",
		AnchorStatus:    "verified",
		AnchorID:        anchorID,
		Success:         true,
	})
}

// LogReplayVerified logs a successful deterministic replay.
func (t *Tracer) LogReplayVerified(executionID, intentID, systemID string) error {
	return t.write(ExecutionLog{
		Timestamp:       time.Now().Unix(),
		ExecutionID:     executionID,
		IntentID:        intentID,
		SystemID:        systemID,
		Stage:           "replay_verified",
		ExecutionStatus: "anchored",
		AnchorStatus:    "verified",
		Success:         true,
	})
}

// LogError logs a failure at any stage.
func (t *Tracer) LogError(executionID, intentID, systemID, stage, errMsg string) error {
	return t.write(ExecutionLog{
		Timestamp:       time.Now().Unix(),
		ExecutionID:     executionID,
		IntentID:        intentID,
		SystemID:        systemID,
		Stage:           stage,
		ExecutionStatus: "failed",
		AnchorStatus:    "failed",
		Success:         false,
		Error:           errMsg,
	})
}

// GetTrace retrieves all log entries for a given execution_id.
func (t *Tracer) GetTrace(executionID string) ([]ExecutionLog, error) {
	var entries []ExecutionLog
	files, err := filepath.Glob(filepath.Join(t.logDir, "trace_*.jsonl"))
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			continue
		}
		dec := json.NewDecoder(file)
		for dec.More() {
			var e ExecutionLog
			if err := dec.Decode(&e); err != nil {
				continue
			}
			if e.ExecutionID == executionID {
				entries = append(entries, e)
			}
		}
		file.Close()
	}
	return entries, nil
}

func (t *Tracer) write(entry ExecutionLog) error {
	filename := filepath.Join(t.logDir, fmt.Sprintf("trace_%s.jsonl",
		time.Unix(entry.Timestamp, 0).Format("2006-01-02")))
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	_, err = file.Write(append(data, '\n'))
	return err
}