package hashing

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
)

// Signer handles deterministic signing operations
type Signer struct {
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
	signerID   string
}

// NewSigner creates a new signer with generated keys
func NewSigner(signerID string) (*Signer, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	
	return &Signer{
		privateKey: priv,
		publicKey:  pub,
		signerID:   signerID,
	}, nil
}

// SignHash signs a hash with the private key
func (s *Signer) SignHash(hash [32]byte) ([]byte, error) {
	signature := ed25519.Sign(s.privateKey, hash[:])
	return signature, nil
}

// VerifySignature verifies a signature against a hash
func (s *Signer) VerifySignature(hash [32]byte, signature []byte) bool {
	return ed25519.Verify(s.publicKey, hash[:], signature)
}

// PublicKeyHex returns the public key as hex string
func (s *Signer) PublicKeyHex() string {
	return hex.EncodeToString(s.publicKey)
}

// SignerID returns the signer identifier
func (s *Signer) SignerID() string {
	return s.signerID
}

// SignedHash represents a hash with its signature
type SignedHash struct {
	Hash      [32]byte `json:"hash"`
	Signature []byte   `json:"signature"`
	SignerID  string   `json:"signer_id"`
	PublicKey string   `json:"public_key"`
}

// SignHashRaw signs an already-computed hash directly without re-hashing.
// Use this when the caller has already computed SHA-256 (e.g. envelope.Hash()).
func (s *Signer) SignHashRaw(hash [32]byte) (*SignedHash, error) {
	signature := ed25519.Sign(s.privateKey, hash[:])
	return &SignedHash{
		Hash:      hash,
		Signature: signature,
		SignerID:  s.signerID,
		PublicKey: s.PublicKeyHex(),
	}, nil
}

// CreateSignedHash computes SHA-256 of data then signs it.
func (s *Signer) CreateSignedHash(data []byte) (*SignedHash, error) {
	hash := sha256.Sum256(data)
	signature, err := s.SignHash(hash)
	if err != nil {
		return nil, err
	}
	
	return &SignedHash{
		Hash:      hash,
		Signature: signature,
		SignerID:  s.signerID,
		PublicKey: s.PublicKeyHex(),
	}, nil
}

// Verify verifies the signed hash
func (sh *SignedHash) Verify() error {
	pubKeyBytes, err := hex.DecodeString(sh.PublicKey)
	if err != nil {
		return fmt.Errorf("invalid public key: %v", err)
	}
	
	if !ed25519.Verify(pubKeyBytes, sh.Hash[:], sh.Signature) {
		return fmt.Errorf("signature verification failed")
	}
	
	return nil
}