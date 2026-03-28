package main

import (
	"fmt"
	"log"
	"time"

	"bhiv-sovereign-chain/l1-core/anchor"
	"bhiv-sovereign-chain/integration/gurukul"
	"bhiv-sovereign-chain/shared/hashing"
	"bhiv-sovereign-chain/shared/replay"
	"bhiv-sovereign-chain/shared/tracing"
)

func main() {
	fmt.Println("=== BHIV Sovereign Chain Day 2 Demo ===")
	
	// Initialize components
	tracer := tracing.NewTracer("./logs")
	replayEngine := replay.NewReplayEngine()
	ttsProcessor := gurukul.NewTTSProcessor()
	agentSigner, _ := hashing.NewSigner("agent-001")

	// Step 1: Real System Execution (Gurukul TTS)
	fmt.Println("\n1. Executing real TTS request...")
	ttsRequest := gurukul.TTSRequest{
		Text:     "Hello BHIV sovereign chain",
		Voice:    "neural-voice-1",
		Language: "en",
		Speed:    1.0,
	}

	// Log input
	tracer.LogInput("demo-intent-001", ttsRequest)

	// Execute TTS and get envelope
	envelope, ttsResponse, err := ttsProcessor.ProcessTTS(ttsRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Envelope created: %s\n", envelope.IntentID)
	fmt.Printf("Decision: %s (confidence: %.2f)\n", envelope.Decision, envelope.ConfidenceScore)

	// Log envelope
	tracer.LogEnvelope(envelope)

	// Step 2: Sign envelope
	fmt.Println("\n2. Signing execution envelope...")
	envelopeHash := envelope.Hash()
	signature, err := agentSigner.SignHash(envelopeHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Envelope hash: %x\n", envelopeHash)
	fmt.Printf("Signature: %x\n", signature[:16]) // Show first 16 bytes

	// Step 3: Anchor on L1
	fmt.Println("\n3. Anchoring execution on L1...")
	var parentHash [32]byte
	anchorID, err := anchor.Submit(envelopeHash, parentHash, time.Now().Unix())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Anchor ID: %x\n", anchorID)

	// Log anchor
	tracer.LogAnchor(envelope.IntentID, fmt.Sprintf("%x", anchorID))

	// Step 4: Verify anchor
	fmt.Println("\n4. Verifying anchor...")
	ok, err := anchor.VerifyAnchor(anchorID, envelopeHash, parentHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Anchor verified: %v\n", ok)

	// Step 5: Deterministic Replay
	fmt.Println("\n5. Testing deterministic replay...")
	replayEnv, err := replayEngine.ReplayExecution(envelope, ttsRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Replay successful: %s\n", replayEnv.IntentID)
	fmt.Printf("Original hash: %x\n", envelope.Hash())
	fmt.Printf("Replay hash:   %x\n", replayEnv.Hash())
	fmt.Printf("Deterministic: %v\n", envelope.Hash() == replayEnv.Hash())

	// Step 6: Show traceability
	fmt.Println("\n6. Execution trace:")
	traceEntries, err := tracer.GetTrace(envelope.IntentID)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range traceEntries {
		fmt.Printf("  %s: %s (success: %v)\n", 
			time.Unix(entry.Timestamp, 0).Format("15:04:05"), 
			entry.Stage, entry.Success)
	}

	fmt.Println("\n=== Demo Complete ===")
	fmt.Printf("TTS Response: %d bytes audio, %.1fs duration\n", 
		len(ttsResponse.AudioData), ttsResponse.Duration)
}
