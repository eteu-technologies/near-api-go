package key

import "errors"

type PublicKeyType string

const (
	RawPublicKeyTypeED25519 byte = iota
	RawPublicKeyTypeSECP256K1
)

const (
	PublicKeyTypeED25519   PublicKeyType = "ed25519"
	PublicKeyTypeSECP256K1 PublicKeyType = "secp256k1"
)

var (
	ErrInvalidPublicKey     = errors.New("invalid public key")
	ErrInvalidPublicKeyType = errors.New("invalid public key type")

	// nolint: deadcode,varcheck,unused
	publicKeyTypes = map[byte]PublicKeyType{
		RawPublicKeyTypeED25519:   PublicKeyTypeED25519,
		RawPublicKeyTypeSECP256K1: PublicKeyTypeSECP256K1,
	}
	reverseKeyTypeMapping = map[string]byte{
		string(PublicKeyTypeED25519):   RawPublicKeyTypeED25519,
		string(PublicKeyTypeSECP256K1): RawPublicKeyTypeSECP256K1,
	}
)
