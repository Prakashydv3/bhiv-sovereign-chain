package gurukul

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"bhiv-sovereign-chain/shared/envelope"
)

// TTSRequest represents a Gurukul text-to-speech execution request.
type TTSRequest struct {
	IntentID string  `json:"intent_id"`
	Text     string  `json:"text"`
	Voice    string  `json:"voice"`
	Language string  `json:"language"`
	Speed    float64 `json:"speed"`
}

// TTSResponse represents the TTS output.
type TTSResponse struct {
	AudioData []byte  `json:"audio_data"`
	Duration  float64 `json:"duration"`
	Quality   string  `json:"quality"`
}

// TTSProcessor is the non-blocking execution tap for the Gurukul TTS system.
// Every call to Execute() captures the full execution into an ExecutionEnvelope.
type TTSProcessor struct {
	SystemID string
	AgentID  string
}

func NewTTSProcessor() *TTSProcessor {
	return &TTSProcessor{SystemID: "gurukul-tts", AgentID: "tts-agent-001"}
}

// Execute is the execution tap: runs TTS and captures the result into an envelope.
// The envelope is the authoritative record of what happened — not the response alone.
func (p *TTSProcessor) Execute(request TTSRequest) (*envelope.ExecutionEnvelope, *TTSResponse, error) {
	// Create envelope — execution_id is deterministic from system+agent+intent
	env := envelope.NewEnvelope(p.SystemID, p.AgentID, request.IntentID)
	env.Timestamp = 1703123456 // Fixed for determinism; real system sets this from event time
	env.ExecutionStatus = envelope.StatusPending

	// Compute input hash from all request fields
	inputData := fmt.Sprintf("%s|%s|%s|%.2f", request.Text, request.Voice, request.Language, request.Speed)
	env.InputHash = sha256.Sum256([]byte(inputData))

	// Execute TTS
	response := &TTSResponse{
		AudioData: []byte(fmt.Sprintf("AUDIO_DATA_FOR_%s", strings.ToUpper(request.Text))),
		Duration:  float64(len(request.Text)) * 0.1,
		Quality:   "high",
	}

	// Compute output hash from actual response
	outputData := fmt.Sprintf("%x|%.2f|%s", response.AudioData, response.Duration, response.Quality)
	env.OutputHash = sha256.Sum256([]byte(outputData))

	// Record execution result
	env.Decision = "tts_generated"
	env.ConfidenceScore = calculateConfidence(request)
	env.PolicyID = "tts-policy-v1"
	env.DatasetIDs = []string{"voice-models-v2", "language-pack-en"}
	env.EnvironmentHash = sha256.Sum256([]byte("gurukul-tts-env-v1"))
	env.ExecutionStatus = envelope.StatusExecuted

	if err := env.Validate(); err != nil {
		return nil, nil, fmt.Errorf("envelope validation failed: %v", err)
	}

	return env, response, nil
}

func calculateConfidence(request TTSRequest) float64 {
	confidence := 0.8
	if len(request.Text) > 100 {
		confidence += 0.1
	}
	if request.Language == "en" {
		confidence += 0.05
	}
	if confidence > 1.0 {
		confidence = 1.0
	}
	return confidence
}