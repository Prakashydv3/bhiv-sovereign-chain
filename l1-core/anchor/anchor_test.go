package anchor_test

import (
	"crypto/sha256"
	"testing"
	"time"

	"bhiv-sovereign-chain/l1-core/anchor"
	"bhiv-sovereign-chain/shared/hashing"
)

func h(s string) [32]byte { return sha256.Sum256([]byte(s)) }

func makeSigners(t *testing.T, stateHash [32]byte) (*hashing.SignedHash, *hashing.SignedHash) {
	agent, err := hashing.NewSigner("agent-001")
	if err != nil {
		t.Fatal(err)
	}
	enforcement, err := hashing.NewSigner("enforcement-001")
	if err != nil {
		t.Fatal(err)
	}
	// Sign the stateHash directly — not re-hashed
	agentSig, err := agent.SignHashRaw(stateHash)
	if err != nil {
		t.Fatal(err)
	}
	enforcementSig, err := enforcement.SignHashRaw(stateHash)
	if err != nil {
		t.Fatal(err)
	}
	return agentSig, enforcementSig
}

func TestSubmitRejectsZeroHash(t *testing.T) {
	agentSig, enforcementSig := makeSigners(t, [32]byte{})
	_, err := anchor.Submit([32]byte{}, [32]byte{}, time.Now().Unix(), agentSig, enforcementSig)
	if err == nil {
		t.Fatal("expected error for zero stateHash")
	}
}

func TestSubmitDeterministicAnchorID(t *testing.T) {
	sh := h("state-X")
	ph := h("parent-X")
	ts := int64(1700000000)
	agentSig, enforcementSig := makeSigners(t, sh)

	id, err := anchor.Submit(sh, ph, ts, agentSig, enforcementSig)
	if err != nil {
		t.Fatal(err)
	}

	combined := append(sh[:], ph[:]...)
	expected := sha256.Sum256(combined)
	if id != expected {
		t.Fatalf("anchorID mismatch: got %x want %x", id, expected)
	}
}

func TestVerifyAnchorPassAndFail(t *testing.T) {
	sh := h("state-Y")
	ph := h("parent-Y")
	agentSig, enforcementSig := makeSigners(t, sh)

	id, err := anchor.Submit(sh, ph, 1700000001, agentSig, enforcementSig)
	if err != nil {
		t.Fatal(err)
	}

	ok, err := anchor.VerifyAnchor(id, sh, ph)
	if err != nil || !ok {
		t.Fatal("expected verify to pass")
	}

	ok, _ = anchor.VerifyAnchor(id, h("wrong"), ph)
	if ok {
		t.Fatal("expected verify to fail on wrong hash")
	}
}

func TestQueryMissingReturnsError(t *testing.T) {
	_, err := anchor.Query(h("nonexistent"))
	if err == nil {
		t.Fatal("expected error for missing anchor")
	}
}

func TestSubmitRejectsInvalidSignature(t *testing.T) {
	sh := h("state-Z")
	ph := h("parent-Z")
	// Sign a different hash — simulates tampered envelope
	agentSig, enforcementSig := makeSigners(t, h("different-data"))
	_, err := anchor.Submit(sh, ph, 1700000002, agentSig, enforcementSig)
	if err == nil {
		t.Fatal("expected error for invalid signature")
	}
}
