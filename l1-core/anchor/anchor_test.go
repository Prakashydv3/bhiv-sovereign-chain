package anchor_test

import (
	"crypto/sha256"
	"testing"
	"time"

	"bhiv-sovereign-chain/l1-core/anchor"
)

func h(s string) [32]byte { return sha256.Sum256([]byte(s)) }

func TestSubmitRejectsZeroHash(t *testing.T) {
	_, err := anchor.Submit([32]byte{}, [32]byte{}, time.Now().Unix())
	if err == nil {
		t.Fatal("expected error for zero stateHash")
	}
}

func TestSubmitDeterministicAnchorID(t *testing.T) {
	sh := h("state-X")
	ph := h("parent-X")
	ts := int64(1700000000)

	id, err := anchor.Submit(sh, ph, ts)
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

	id, err := anchor.Submit(sh, ph, 1700000001)
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
