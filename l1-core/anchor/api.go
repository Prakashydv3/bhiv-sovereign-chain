package anchor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"bhiv-sovereign-chain/shared/envelope"
)

// AnchorRequest represents an execution anchoring request
type AnchorRequest struct {
	Envelope  *envelope.ExecutionEnvelope `json:"envelope"`
	Signature []byte                       `json:"signature"`
	SignerID  string                       `json:"signer_id"`
}

// AnchorResponse represents the anchoring response
type AnchorResponse struct {
	AnchorID    string `json:"anchor_id"`
	EnvelopeHash string `json:"envelope_hash"`
	Success     bool   `json:"success"`
	Message     string `json:"message"`
}

// AnchorExecutionHandler handles POST /anchor_execution
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

	// Validate envelope
	if err := req.Envelope.Validate(); err != nil {
		writeErrorResponse(w, fmt.Sprintf("Invalid envelope: %v", err), http.StatusBadRequest)
		return
	}

	// Compute envelope hash
	envelopeHash := req.Envelope.Hash()

	// Verify signature (simplified - in production would verify against known signers)
	if len(req.Signature) == 0 {
		writeErrorResponse(w, "Signature required", http.StatusBadRequest)
		return
	}

	// Submit to L1 anchor (using envelope hash as state root)
	var parentHash [32]byte // Genesis for now
	anchorID, err := Submit(envelopeHash, parentHash, time.Now().Unix())
	if err != nil {
		writeErrorResponse(w, fmt.Sprintf("Anchor submission failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Success response
	response := AnchorResponse{
		AnchorID:     fmt.Sprintf("%x", anchorID),
		EnvelopeHash: fmt.Sprintf("%x", envelopeHash),
		Success:      true,
		Message:      "Execution anchored successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// writeErrorResponse writes an error response
func writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := AnchorResponse{
		Success: false,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// StartAnchorAPI starts the anchor HTTP API server
func StartAnchorAPI(port string) error {
	http.HandleFunc("/anchor_execution", AnchorExecutionHandler)
	fmt.Printf("Anchor API listening on port %s\n", port)
	return http.ListenAndServe(":"+port, nil)
}