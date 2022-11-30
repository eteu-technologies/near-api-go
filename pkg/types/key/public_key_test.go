package key

import (
	"testing"
)

func TestED25519Key(t *testing.T) {
	expected := `ed25519:DcA2MzgpJbrUATQLLceocVckhhAqrkingax4oJ9kZ847`

	parsed, err := NewBase58PublicKey(expected)
	if err != nil {
		t.Errorf("failed to parse public key: %s", err)
	}

	if s := parsed.String(); s != expected {
		t.Errorf("%s != %s", s, expected)
	}
}

func TestED25519Key_Base58_And_Back(t *testing.T) {
	expected := `ed25519:3xCFas58RKvD5UpF9GqvEb6q9rvgfbEJPhLf85zc4HpC`

	parsed, err := NewBase58PublicKey(expected)
	if err != nil {
		t.Errorf("failed to parse public key: %s", err)
	}

	publicKey := parsed.ToPublicKey()
	converted := publicKey.ToBase58PublicKey()

	if s := converted.String(); s != expected {
		t.Errorf("%s != %s", s, expected)
	}
}

func TestED25519UnmarshalJSON(t *testing.T) {
	expected := `ed25519:DcA2MzgpJbrUATQLLceocVckhhAqrkingax4oJ9kZ847`
	parsed, err := NewBase58PublicKey(expected)
	if err != nil {
		t.Errorf("failed to parse public key: %s", err)
	}

	var parsed2 Base58PublicKey
	err = parsed2.UnmarshalJSON([]byte(`"` + expected + `"`))
	if err != nil {
		t.Errorf("failed to parse public key: %s", err)
	}

	if s := parsed2.String(); s != expected {
		t.Errorf("%s != %s", s, expected)
	}

	if parsed2.Type != parsed.Type {
		t.Errorf("parsed2 != parsed")
	}

	if parsed2.Value != parsed.Value {
		t.Errorf("parsed2 != parsed")
	}
}

func TestSECP256k1UnmarshalJSON(t *testing.T) {
	expected := `secp256k1:23URfhHiWFYsFArc5nLrmj8qDMXXrgF2iU39Dod3cXpBu` // Public key
	parsed, err := NewBase58PublicKey(expected)
	if err != nil {
		t.Errorf("failed to parse public key: %s", err)
	}

	var parsed2 Base58PublicKey
	err = parsed2.UnmarshalJSON([]byte(`"` + expected + `"`))
	if err != nil {
		t.Errorf("failed to parse public key: %s", err)
	}

	if s := parsed2.String(); s != expected {
		t.Errorf("%s != %s", s, expected)
	}

	if parsed2.Type != parsed.Type {
		t.Errorf("parsed2 != parsed")
	}

	if parsed2.Value != parsed.Value {
		t.Errorf("parsed2 != parsed")
	}
}

func TestPublicKeyHash(t *testing.T) {
	raw := "ed25519:2MDRrkKRTXFPuMXkcKm39KzLQznuaCAybKKYKie4j26k8S2Nth8SvDyWxfBbFk8MC1svEJbuekRAUpnDRSFXdd9s" // Private key in base58
	expectedHash := "a7a56191a40b5586bee21fb0e4cd711b5b70dadc02b7486f62bbf1b9b3e51992"

	keyPair, err := NewBase58KeyPair(raw)
	if err != nil {
		t.Errorf("failed to create key pair: %s", err)
	}

	hash := keyPair.PublicKey.pk.Hash()
	if hash != expectedHash {
		t.Errorf("%s != %s", hash, expectedHash)
	}
}

func TestPublicKeyString(t *testing.T) {
	raw := "ed25519:2MDRrkKRTXFPuMXkcKm39KzLQznuaCAybKKYKie4j26k8S2Nth8SvDyWxfBbFk8MC1svEJbuekRAUpnDRSFXdd9s" // Private key in base58
	expectedString := "ed25519:CHRMGVtFYyJ1uPWCpne8WRDEhJgaRGTa1akXUuDCfEhF"

	keyPair, err := NewBase58KeyPair(raw)
	if err != nil {
		t.Errorf("failed to create key pair: %s", err)
	}

	hash := keyPair.PublicKey.pk.String()
	if hash != expectedString {
		t.Errorf("%s != %s", hash, expectedString)
	}
}
