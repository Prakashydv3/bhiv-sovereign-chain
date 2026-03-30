package replay

import (
	"fmt"

	"bhiv-sovereign-chain/integration/gurukul"
	"bhiv-sovereign-chain/shared/envelope"
)

// ReplayEngine verifies determinism by re-executing from a captured execution event.
// Replay is always driven by the original envelope — not raw inputs alone.
type ReplayEngine struct {
	ttsProcessor *gurukul.TTSProcessor
}

func NewReplayEngine() *ReplayEngine {
	return &ReplayEngine{ttsProcessor: gurukul.NewTTSProcessor()}
}

// ReplayFromEnvelope re-executes the system identified in the envelope and verifies
// that the result is bit-for-bit identical. Mismatch is a hard failure.
func (r *ReplayEngine) ReplayFromEnvelope(original *envelope.ExecutionEnvelope, input interface{}) (*envelope.ExecutionEnvelope, error) {
	switch original.SystemID {
	case "gurukul-tts":
		return r.replayTTS(original, input)
	default:
		return nil, fmt.Errorf("no replay handler for system: %s", original.SystemID)
	}
}

func (r *ReplayEngine) replayTTS(original *envelope.ExecutionEnvelope, input interface{}) (*envelope.ExecutionEnvelope, error) {
	req, ok := input.(gurukul.TTSRequest)
	if !ok {
		return nil, fmt.Errorf("expected gurukul.TTSRequest")
	}
	// Ensure intent_id matches the original envelope so execution_id is identical
	req.IntentID = original.IntentID

	result, _, err := r.ttsProcessor.Execute(req)
	if err != nil {
		return nil, fmt.Errorf("replay execution failed: %v", err)
	}

	return result, r.verify(original, result)
}

// verify checks all determinism-critical fields. Any mismatch is a hard failure.
func (r *ReplayEngine) verify(original, replayed *envelope.ExecutionEnvelope) error {
	if original.ExecutionID != replayed.ExecutionID {
		return fmt.Errorf("execution_id mismatch: %s vs %s", original.ExecutionID, replayed.ExecutionID)
	}
	if original.InputHash != replayed.InputHash {
		return fmt.Errorf("input_hash mismatch: %x vs %x", original.InputHash, replayed.InputHash)
	}
	if original.OutputHash != replayed.OutputHash {
		return fmt.Errorf("output_hash mismatch: %x vs %x", original.OutputHash, replayed.OutputHash)
	}
	if original.Decision != replayed.Decision {
		return fmt.Errorf("decision mismatch: %s vs %s", original.Decision, replayed.Decision)
	}
	if original.ConfidenceScore != replayed.ConfidenceScore {
		return fmt.Errorf("confidence_score mismatch: %.3f vs %.3f", original.ConfidenceScore, replayed.ConfidenceScore)
	}
	return nil
}

// ReplayChain replays a sequence of execution events and returns the final hash.
// Fails immediately on the first determinism violation.
func (r *ReplayEngine) ReplayChain(events []ReplayEvent) ([32]byte, error) {
	var finalHash [32]byte
	for i, event := range events {
		result, err := r.ReplayFromEnvelope(event.OriginalEnvelope, event.Input)
		if err != nil {
			return [32]byte{}, fmt.Errorf("chain replay failed at step %d: %v", i, err)
		}
		finalHash = result.Hash()
	}
	return finalHash, nil
}

// ReplayEvent pairs a captured envelope with the input that produced it.
type ReplayEvent struct {
	OriginalEnvelope *envelope.ExecutionEnvelope
	Input            interface{}
}