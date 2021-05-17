package key

import (
	"crypto/ed25519"
	"fmt"
	"strings"

	"github.com/mr-tron/base58"
)

type KeyPair struct {
	Type PublicKeyType

	PublicKey  Base58PublicKey
	PrivateKey ed25519.PrivateKey
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
		return kp, ErrInvalidPrivateKeyType
	}

	// TODO
	if keyType == RawPublicKeyTypeSECP256K1 {
		return kp, fmt.Errorf("SECP256K1 is not supported yet")
	}

	decoded, err := base58.Decode(encodedKey)
	if err != nil {
		return kp, fmt.Errorf("failed to decode private key: %w", err)
	}

	if len(decoded) != ed25519.PrivateKeySize {
		return kp, ErrInvalidPrivateKey
	}

	kp.PrivateKey = ed25519.PrivateKey(decoded)
	pubKey := kp.PrivateKey[32:] // See ed25519.Public()

	kp.Type = publicKeyTypes[keyType]
	if pubKey, err := WrapRawKey(kp.Type, pubKey); err != nil {
		return kp, err
	} else {
		kp.PublicKey = pubKey.ToBase58PublicKey()
	}

	return
}

func (kp *KeyPair) Sign(data []byte) (sig Signature) {
	sigType := reverseKeyTypeMapping[string(kp.Type)]

	switch sigType {
	//case RawPublicKeyTypeSECP256K1:
	case RawPublicKeyTypeED25519:
		sig = NewSignatureED25519(ed25519.Sign(kp.PrivateKey, data))
	}
	return
}
