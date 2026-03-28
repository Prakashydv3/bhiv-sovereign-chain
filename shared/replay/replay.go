package replay

import (
	"fmt"

	"bhiv-sovereign-chain/shared/envelope"
	"bhiv-sovereign-chain/integration/gurukul"
)

// ReplayEngine handles deterministic execution replay
type ReplayEngine struct {
	ttsProcessor *gurukul.TTSProcessor
}

// NewReplayEngine creates a new replay engine
func NewReplayEngine() *ReplayEngine {
	return &ReplayEngine{
		ttsProcessor: gurukul.NewTTSProcessor(),
	}
}

// ReplayExecution replays an execution and verifies determinism
func (r *ReplayEngine) ReplayExecution(originalEnvelope *envelope.ExecutionEnvelope, input interface{}) (*envelope.ExecutionEnvelope, error) {
	// Determine system type and replay accordingly
	switch originalEnvelope.SystemID {
	case "gurukul-tts":
		return r.replayTTS(originalEnvelope, input)
	default:
		return nil, fmt.Errorf("unsupported system: %s", originalEnvelope.SystemID)
	}
}

// replayTTS replays a TTS execution
func (r *ReplayEngine) replayTTS(originalEnv *envelope.ExecutionEnvelope, input interface{}) (*envelope.ExecutionEnvelope, error) {
	ttsRequest, ok := input.(gurukul.TTSRequest)
	if !ok {
		return nil, fmt.Errorf("invalid TTS input type")
	}

	// Re-execute TTS
	replayEnv, _, err := r.ttsProcessor.ProcessTTS(ttsRequest)
	if err != nil {
		return nil, fmt.Errorf("replay execution failed: %v", err)
	}

	// Verify determinism
	if err := r.verifyDeterminism(originalEnv, replayEnv); err != nil {
		return nil, err
	}

	return replayEnv, nil
}

// verifyDeterminism checks if replay matches original execution
func (r *ReplayEngine) verifyDeterminism(original, replay *envelope.ExecutionEnvelope) error {
	// Check critical fields for determinism
	if original.InputHash != replay.InputHash {
		return fmt.Errorf("input hash mismatch: original=%x, replay=%x", 
			original.InputHash, replay.InputHash)
	}

	if original.OutputHash != replay.OutputHash {
		return fmt.Errorf("output hash mismatch: original=%x, replay=%x", 
			original.OutputHash, replay.OutputHash)
	}

	if original.Decision != replay.Decision {
		return fmt.Errorf("decision mismatch: original=%s, replay=%s", 
			original.Decision, replay.Decision)
	}

	if original.ConfidenceScore != replay.ConfidenceScore {
		return fmt.Errorf("confidence score mismatch: original=%.3f, replay=%.3f", 
			original.ConfidenceScore, replay.ConfidenceScore)
	}

	return nil
}

// ReplayChain replays a sequence of executions
func (r *ReplayEngine) ReplayChain(executions []ReplayInput) ([32]byte, error) {
	var finalHash [32]byte
	
	for i, exec := range executions {
		replayEnv, err := r.ReplayExecution(exec.OriginalEnvelope, exec.Input)
		if err != nil {
			return [32]byte{}, fmt.Errorf("replay failed at step %d: %v", i, err)
		}
		
		// Update final hash with this execution
		finalHash = replayEnv.Hash()
	}
	
	return finalHash, nil
}

// ReplayInput represents input for replay
type ReplayInput struct {
	OriginalEnvelope *envelope.ExecutionEnvelope
	Input            interface{}
}