package key

import (
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
	ecdsa "github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"github.com/mr-tron/base58"

	"github.com/eteu-technologies/near-api-go/pkg/types/signature"
)

type KeyPair struct {
	Type PublicKeyType

	PublicKey  Base58PublicKey
	PrivateKey interface{}
}

type PrivateKeyType interface {
	ed25519.PrivateKey | secp256k1.PrivateKey
}

func GenerateKeyPair(keyType PublicKeyType, rand io.Reader) (kp KeyPair, err error) {
	if _, ok := reverseKeyTypeMapping[string(keyType)]; !ok {
		return kp, ErrInvalidKeyType
	}

	var rawPub PublicKey

	switch keyType {
	case KeyTypeED25519:
		var pub ed25519.PublicKey
		var priv ed25519.PrivateKey

		pub, priv, err = ed25519.GenerateKey(rand)
		if err != nil {
			return
		}

		rawPub, err = WrapRawKey(keyType, pub)
		if err != nil {
			return
		}

		kp = CreateKeyPair(keyType, rawPub.ToBase58PublicKey(), priv)
	case KeyTypeSECP256K1:
		var ephemeralPrivKey *secp256k1.PrivateKey
		ephemeralPrivKey, err = secp256k1.GeneratePrivateKey()
		if err != nil {
			return
		}

		rawPub, err = WrapRawKey(keyType, ephemeralPrivKey.PubKey().SerializeCompressed())
		if err != nil {
			return
		}

		kp = CreateKeyPair(keyType, rawPub.ToBase58PublicKey(), *ephemeralPrivKey)
	}

	return
}

func CreateKeyPair[P PrivateKeyType](keyType PublicKeyType, pub Base58PublicKey, priv P) KeyPair {
	return KeyPair{
		Type:       keyType,
		PublicKey:  pub,
		PrivateKey: priv,
	}
}

func NewBase58KeyPair(raw string) (kp KeyPair, err error) {
	split := strings.SplitN(raw, ":", 2)
	if len(split) != 2 {
		return kp, ErrInvalidPrivateKey
	}

	keyTypeRaw := split[0]
	encodedKey := split[1]

	keyType, ok := reverseKeyTypeMapping[keyTypeRaw]
	if !ok {
		return kp, ErrInvalidKeyType
	}

	var decoded []byte

	switch keyType {
	case RawKeyTypeED25519:
		decoded, err = base58.Decode(encodedKey)
		if err != nil {
			return kp, fmt.Errorf("failed to decode private key: %w", err)
		}

		if len(decoded) != ed25519.PrivateKeySize {
			return kp, ErrInvalidPrivateKey
		}

		var pubKey PublicKey

		theKeyType := keyTypes[keyType]
		privKey := ed25519.PrivateKey(decoded)
		pubKey, err = WrapRawKey(theKeyType, privKey[32:]) // See ed25519.Public()
		if err != nil {
			println("wraprawkey failed")
			return
		}

		kp = CreateKeyPair(theKeyType, pubKey.ToBase58PublicKey(), privKey)
	case RawKeyTypeSECP256K1:
		decoded, err = base58.Decode(encodedKey)
		if err != nil {
			return kp, fmt.Errorf("failed to decode private key: %w", err)
		}

		privateKey := secp256k1.PrivKeyFromBytes(decoded)
		ephemeralPubKey := privateKey.PubKey().SerializeCompressed()

		theKeyType := keyTypes[keyType]

		var pubKey PublicKey
		pubKey, err = WrapRawKey(theKeyType, ephemeralPubKey)
		if err != nil {
			println("wraprawkey failed")
			return
		}

		kp = CreateKeyPair(theKeyType, pubKey.ToBase58PublicKey(), *privateKey)
	}

	return
}

func (kp *KeyPair) Sign(data []byte) (sig signature.Signature) {
	sigType := reverseKeyTypeMapping[string(kp.Type)]

	switch sigType {
	case RawKeyTypeED25519:
		privateKey := kp.PrivateKey.(ed25519.PrivateKey)
		sig = signature.NewSignatureED25519(ed25519.Sign(privateKey, data))
	case RawKeyTypeSECP256K1:
		privateKey := kp.PrivateKey.(secp256k1.PrivateKey)
		sig = signature.NewSignatureSECP256K1(ecdsa.Sign(&privateKey, data).Serialize())
	}
	return
}

func (kp *KeyPair) PrivateEncoded() string {
	var encoded string

	switch kp.Type {
	case KeyTypeED25519:
		privateKey := kp.PrivateKey.(ed25519.PrivateKey)
		encoded = fmt.Sprintf("%s:%s", kp.Type, base58.Encode(privateKey))
	case KeyTypeSECP256K1:
		privateKey := kp.PrivateKey.(secp256k1.PrivateKey)
		encoded = fmt.Sprintf("%s:%s", kp.Type, base58.Encode(privateKey.Serialize()))
	}

	return encoded
}

func (kp *KeyPair) UnmarshalJSON(b []byte) (err error) {
	var s string
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	*kp, err = NewBase58KeyPair(s)
	return
}
