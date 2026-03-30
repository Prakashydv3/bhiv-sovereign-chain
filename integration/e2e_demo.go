package main

import (
	"fmt"
	"log"
	"time"

	"bhiv-sovereign-chain/integration/gurukul"
	"bhiv-sovereign-chain/l1-core/anchor"
	"bhiv-sovereign-chain/shared/hashing"
	"bhiv-sovereign-chain/shared/replay"
	"bhiv-sovereign-chain/shared/tracing"
)

func main() {
	fmt.Println("=== BHIV Sovereign Chain — End-to-End Demo ===")

	tracer := tracing.NewTracer("./logs")
	replayEngine := replay.NewReplayEngine()
	ttsProcessor := gurukul.NewTTSProcessor()
	agentSigner, _ := hashing.NewSigner("agent-001")

	// --- Step 1: Execution Tap (Gurukul TTS) ---
	fmt.Println("\n[1] Execution tap — Gurukul TTS")
	request := gurukul.TTSRequest{
		IntentID: "bhiv-intent-tts-001",
		Text:     "Hello BHIV sovereign chain",
		Voice:    "neural-voice-1",
		Language: "en",
		Speed:    1.0,
	}

	tracer.LogInput("pending", request.IntentID, ttsProcessor.SystemID)

	env, ttsResponse, err := ttsProcessor.Execute(request)
	if err != nil {
		log.Fatal(err)
	}
	tracer.LogEnvelopeCreated(env)
	fmt.Printf("  execution_id    : %s\n", env.ExecutionID)
	fmt.Printf("  intent_id       : %s\n", env.IntentID)
	fmt.Printf("  system_id       : %s\n", env.SystemID)
	fmt.Printf("  decision        : %s\n", env.Decision)
	fmt.Printf("  confidence      : %.2f\n", env.ConfidenceScore)
	fmt.Printf("  execution_status: %s\n", env.ExecutionStatus)

	// --- Step 2: Hash Generation ---
	fmt.Println("\n[2] Hash generation")
	envelopeHash := env.Hash()
	tracer.LogHashGenerated(env, envelopeHash)
	fmt.Printf("  envelope_hash   : %x\n", envelopeHash)

	// --- Step 3: Signing ---
	fmt.Println("\n[3] Agent signing")
	signature, err := agentSigner.SignHash(envelopeHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  signer_id       : %s\n", agentSigner.SignerID())
	fmt.Printf("  signature       : %x\n", signature[:16])

	// --- Step 4: L1 Anchor ---
	fmt.Println("\n[4] L1 anchoring")
	var parentHash [32]byte
	anchorID, err := anchor.Submit(envelopeHash, parentHash, time.Now().Unix())
	if err != nil {
		log.Fatal(err)
	}
	anchorIDHex := fmt.Sprintf("%x", anchorID)
	env.MarkAnchored(anchorIDHex)
	tracer.LogAnchored(env, anchorIDHex)
	fmt.Printf("  anchor_id       : %s\n", anchorIDHex)
	fmt.Printf("  execution_status: %s\n", env.ExecutionStatus)

	// --- Step 5: Anchor Verification ---
	fmt.Println("\n[5] Anchor verification")
	ok, err := anchor.VerifyAnchor(anchorID, envelopeHash, parentHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  anchor_status   : verified=%v\n", ok)

	// --- Step 6: Deterministic Replay (tied to real execution event) ---
	fmt.Println("\n[6] Deterministic replay from execution event")
	replayResult, err := replayEngine.ReplayFromEnvelope(env, request)
	if err != nil {
		log.Fatal(err)
	}
	tracer.LogReplayVerified(env.ExecutionID, env.IntentID, env.SystemID)
	fmt.Printf("  original_hash   : %x\n", envelopeHash)
	fmt.Printf("  replay_hash     : %x\n", replayResult.Hash())
	fmt.Printf("  deterministic   : %v\n", envelopeHash == replayResult.Hash())

	// --- Step 7: Observability trace ---
	fmt.Println("\n[7] Execution trace (execution_id scoped)")
	logs, err := tracer.GetTrace(env.ExecutionID)
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range logs {
		fmt.Printf("  stage=%-20s execution_status=%-10s anchor_status=%s\n",
			l.Stage, l.ExecutionStatus, l.AnchorStatus)
	}

	fmt.Println("\n=== Demo Complete ===")
	fmt.Printf("TTS: %d bytes, %.1fs duration\n", len(ttsResponse.AudioData), ttsResponse.Duration)
}
