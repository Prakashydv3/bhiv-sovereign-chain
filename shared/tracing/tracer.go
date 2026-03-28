package tracing

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"bhiv-sovereign-chain/shared/envelope"
)

// TraceEntry represents a single trace log entry
type TraceEntry struct {
	Timestamp    int64                        `json:"timestamp"`
	IntentID     string                       `json:"intent_id"`
	Stage        string                       `json:"stage"`
	Input        interface{}                  `json:"input,omitempty"`
	Envelope     *envelope.ExecutionEnvelope `json:"envelope,omitempty"`
	Hash         string                       `json:"hash,omitempty"`
	AnchorTxID   string                       `json:"anchor_tx_id,omitempty"`
	Success      bool                         `json:"success"`
	ErrorMessage string                       `json:"error_message,omitempty"`
}

// Tracer handles execution tracing
type Tracer struct {
	logDir string
}

// NewTracer creates a new tracer
func NewTracer(logDir string) *Tracer {
	os.MkdirAll(logDir, 0755)
	return &Tracer{logDir: logDir}
}

// LogInput logs the initial input
func (t *Tracer) LogInput(intentID string, input interface{}) error {
	entry := TraceEntry{
		Timestamp: time.Now().Unix(),
		IntentID:  intentID,
		Stage:     "input",
		Input:     input,
		Success:   true,
	}
	return t.writeEntry(entry)
}

// LogEnvelope logs the execution envelope
func (t *Tracer) LogEnvelope(env *envelope.ExecutionEnvelope) error {
	entry := TraceEntry{
		Timestamp: time.Now().Unix(),
		IntentID:  env.IntentID,
		Stage:     "envelope",
		Envelope:  env,
		Hash:      fmt.Sprintf("%x", env.Hash()),
		Success:   true,
	}
	return t.writeEntry(entry)
}

// LogAnchor logs the anchor transaction
func (t *Tracer) LogAnchor(intentID, anchorTxID string) error {
	entry := TraceEntry{
		Timestamp:  time.Now().Unix(),
		IntentID:   intentID,
		Stage:      "anchor",
		AnchorTxID: anchorTxID,
		Success:    true,
	}
	return t.writeEntry(entry)
}

// LogError logs an error at any stage
func (t *Tracer) LogError(intentID, stage, errorMsg string) error {
	entry := TraceEntry{
		Timestamp:    time.Now().Unix(),
		IntentID:     intentID,
		Stage:        stage,
		Success:      false,
		ErrorMessage: errorMsg,
	}
	return t.writeEntry(entry)
}

// writeEntry writes a trace entry to the log file
func (t *Tracer) writeEntry(entry TraceEntry) error {
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

// GetTrace retrieves all trace entries for an intent
func (t *Tracer) GetTrace(intentID string) ([]TraceEntry, error) {
	var entries []TraceEntry
	
	// Search through all log files (simplified - in production would index by intent)
	files, err := filepath.Glob(filepath.Join(t.logDir, "trace_*.jsonl"))
	if err != nil {
		return nil, err
	}

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			continue
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		for decoder.More() {
			var entry TraceEntry
			if err := decoder.Decode(&entry); err != nil {
				continue
			}
			if entry.IntentID == intentID {
				entries = append(entries, entry)
			}
		}
	}

	return entries, nil
}