package anchor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"bhiv-sovereign-chain/shared/envelope"
	"bhiv-sovereign-chain/shared/hashing"
)

// AnchorRequest represents an execution anchoring request
type AnchorRequest struct {
	Envelope        *envelope.ExecutionEnvelope `json:"envelope"`
	AgentSig        *hashing.SignedHash          `json:"agent_sig"`
	EnforcementSig  *hashing.SignedHash          `json:"enforcement_sig"`
}

// AnchorResponse represents the anchoring response
type AnchorResponse struct {
	AnchorID     string `json:"anchor_id"`
	EnvelopeHash string `json:"envelope_hash"`
	Success      bool   `json:"success"`
	Message      string `json:"message"`
}

// AnchorExecutionHandler handles POST /anchor_execution
// Flow: envelope → hash → signatures → L1 anchor
func AnchorExecutionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req AnchorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Envelope.Validate(); err != nil {
		writeErrorResponse(w, fmt.Sprintf("Invalid envelope: %v", err), http.StatusBadRequest)
		return
	}

	if req.AgentSig == nil || req.EnforcementSig == nil {
		writeErrorResponse(w, "agent_sig and enforcement_sig required", http.StatusBadRequest)
		return
	}

	envelopeHash := req.Envelope.Hash()

	var parentHash [32]byte
	anchorID, err := Submit(envelopeHash, parentHash, time.Now().Unix(), req.AgentSig, req.EnforcementSig)
	if err != nil {
		writeErrorResponse(w, fmt.Sprintf("Anchor submission failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AnchorResponse{
		AnchorID:     fmt.Sprintf("%x", anchorID),
		EnvelopeHash: fmt.Sprintf("%x", envelopeHash),
		Success:      true,
		Message:      "Execution anchored successfully",
	})
}

func writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(AnchorResponse{Success: false, Message: message})
}

func StartAnchorAPI(port string) error {
	http.HandleFunc("/anchor_execution", AnchorExecutionHandler)
	fmt.Printf("Anchor API listening on port %s\n", port)
	return http.ListenAndServe(":"+port, nil)
}