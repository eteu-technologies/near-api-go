package key

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"testing"
)

// ------------------------------------------------
// Tests for ED25519

// TestGenerateKeyPairED25519 tests the generation of a key pair.
func TestGenerateKeyPairED25519(t *testing.T) {
	keyPair, err := GenerateKeyPair(KeyTypeED25519, rand.Reader)
	if err != nil {
		t.Errorf("failed to generate key pair: %s", err)
	}

	if keyPair.Type != KeyTypeED25519 {
		t.Errorf("invalid key type: %s", keyPair.Type)
	}

	if keyPair.PublicKey.Value == "" {
		t.Errorf("public key is nil")
	}

	if len(keyPair.PublicKey.pk.Value()) != ed25519.PublicKeySize {
		t.Errorf("public key is not valid,  %d != %d", len(keyPair.PublicKey.pk), ed25519.PublicKeySize)
	}

	if keyPair.PrivateKey == nil {
		t.Errorf("private key is nil")
	}
}

// TestNewBase58KeyPairED25519 tests the creation of a key pair from a base58 encoded string.
func TestNewBase58KeyPairED25519(t *testing.T) {
	raw := "ed25519:2MDRrkKRTXFPuMXkcKm39KzLQznuaCAybKKYKie4j26k8S2Nth8SvDyWxfBbFk8MC1svEJbuekRAUpnDRSFXdd9s" // Private key in base58
	expectedPubliKey := "CHRMGVtFYyJ1uPWCpne8WRDEhJgaRGTa1akXUuDCfEhF"

	keyPair, err := NewBase58KeyPair(raw)
	if err != nil {
		t.Errorf("failed to create key pair: %s", err)
	}

	if keyPair.Type != KeyTypeED25519 {
		t.Errorf("invalid key type: %s", keyPair.Type)
	}

	if keyPair.PublicKey.Value != expectedPubliKey {
		t.Errorf("public key is not valid: %s", keyPair.PublicKey.Value)
	}

	if keyPair.PrivateKey == nil {
		t.Errorf("private key is nil")
	}

	if keyPair.PrivateEncoded() != raw {
		t.Errorf("private key is not valid: %s", keyPair.PrivateEncoded())
	}
}

// TestSignAndVerifyED25519 tests the signing and verification of a message.
func TestSignAndVerifyED25519(t *testing.T) {
	keyPair, err := GenerateKeyPair(KeyTypeED25519, rand.Reader)
	if err != nil {
		t.Errorf("failed to generate key pair: %s", err)
	}

	message := []byte("Hello World")
	signature := keyPair.Sign(message)

	ok, err := keyPair.PublicKey.pk.Verify(message, signature)
	if err != nil {
		t.Errorf("failed to verify signature: %s", err)
	}

	if !ok {
		t.Errorf("signature is not valid")
	}
}

// TestUnmarshalJSONED25519 tests the unmarshalling of a key pair from a JSON string.
func TestUnmarshalJSONED25519(t *testing.T) {
	keyPair, err := GenerateKeyPair(KeyTypeED25519, rand.Reader)
	if err != nil {
		t.Errorf("failed to generate key pair: %s", err)
	}

	keyPairJSON, err := json.Marshal(keyPair.PrivateEncoded())
	if err != nil {
		t.Errorf("failed to marshal key pair: %s", err)
	}

	var keyPair2 KeyPair
	err = keyPair2.UnmarshalJSON(keyPairJSON)
	if err != nil {
		t.Errorf("failed to unmarshal key pair: %s", err)
	}

	if keyPair2.Type != KeyTypeED25519 {
		t.Errorf("invalid key type: %s", keyPair2.Type)
	}

	if keyPair2.PublicKey.Value != keyPair.PublicKey.Value {
		t.Errorf("public key is not valid: %s", keyPair2.PublicKey.Value)
	}

	if keyPair2.PrivateEncoded() != keyPair.PrivateEncoded() {
		t.Errorf("private key is not valid: %s", keyPair2.PrivateEncoded())
	}
}

// ------------------------------------------------
// Tests for secp256k1

// TestGenerateKeyPair tests the generation of a key pair.
func TestGenerateKeyPairSECP256k1(t *testing.T) {
	keyPair, err := GenerateKeyPair(KeyTypeSECP256K1, rand.Reader)
	if err != nil {
		t.Errorf("failed to generate key pair: %s", err)
	}

	if keyPair.Type != KeyTypeSECP256K1 {
		t.Errorf("invalid key type: %s", keyPair.Type)
	}

	if keyPair.PublicKey.Value == "" {
		t.Errorf("public key is nil")
	}

	if keyPair.PrivateKey == nil {
		t.Errorf("private key is nil")
	}
}

// TestNewBase58KeyPair tests the creation of a key pair from a base58 encoded string.
func TestNewBase58KeyPairSECP256k1(t *testing.T) {
	raw := "secp256k1:3aq6RcztvhMw8PMRbUgyechLS9rpNETDAHFqip3Zb4cb"     // Private key in base58
	expectedPubliKey := "23URfhHiWFYsFArc5nLrmj8qDMXXrgF2iU39Dod3cXpBu" // Public key

	keyPair, err := NewBase58KeyPair(raw)
	if err != nil {
		t.Errorf("failed to create key pair: %s", err)
	}

	if keyPair.Type != KeyTypeSECP256K1 {
		t.Errorf("invalid key type: %s", keyPair.Type)
	}

	if keyPair.PublicKey.Value != expectedPubliKey {
		t.Errorf("public key is not valid: %s", keyPair.PublicKey.Value)
	}

	if keyPair.PrivateKey == nil {
		t.Errorf("private key is nil")
	}

	if keyPair.PrivateEncoded() != raw {
		t.Errorf("private key is not valid: %s", keyPair.PrivateEncoded())
	}
}

// TestSignAndVerify tests the signing and verification of a message.
func TestSignAndVerifySECP256k1(t *testing.T) {
	keyPair, err := GenerateKeyPair(KeyTypeSECP256K1, rand.Reader)
	if err != nil {
		t.Errorf("failed to generate key pair: %s", err)
	}

	message := []byte("Hello World")
	signature := keyPair.Sign(message)

	ok, err := keyPair.PublicKey.pk.Verify(message, signature)
	if err != nil {
		t.Errorf("failed to verify signature: %s", err)
	}

	if !ok {
		t.Errorf("signature is not valid")
	}
}

// TestUnmarshalJSONECP256k1 tests the unmarshalling of a key pair from a JSON string.
func TestUnmarshalJSONECP256k1(t *testing.T) {
	keyPair, err := GenerateKeyPair(KeyTypeSECP256K1, rand.Reader)
	if err != nil {
		t.Errorf("failed to generate key pair: %s", err)
	}

	keyPairJSON, err := json.Marshal(keyPair.PrivateEncoded())
	if err != nil {
		t.Errorf("failed to marshal key pair: %s", err)
	}

	var keyPair2 KeyPair
	err = keyPair2.UnmarshalJSON(keyPairJSON)
	if err != nil {
		t.Errorf("failed to unmarshal key pair: %s", err)
	}

	if keyPair2.Type != KeyTypeSECP256K1 {
		t.Errorf("invalid key type: %s", keyPair2.Type)
	}

	if keyPair2.PublicKey.Value != keyPair.PublicKey.Value {
		t.Errorf("public key is not valid: %s", keyPair2.PublicKey.Value)
	}

	if keyPair2.PrivateEncoded() != keyPair.PrivateEncoded() {
		t.Errorf("private key is not valid: %s", keyPair2.PrivateEncoded())
	}
}

// ------------------------------------------------
