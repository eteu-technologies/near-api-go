package key

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"github.com/mr-tron/base58"

	"github.com/eteu-technologies/near-api-go/pkg/types/signature"
)

type PublicKey []byte

func (p PublicKey) Hash() string {
	return hex.EncodeToString(p[1:])
}

func (p PublicKey) TypeByte() byte {
	return p[0]
}

func (p PublicKey) Value() []byte {
	return p[1:]
}

func (p PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(p[:]))
}

func (p PublicKey) String() string {
	return fmt.Sprintf("%s:%s", keyTypes[p.TypeByte()], base58.Encode(p.Value()))
}

func (p *PublicKey) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	dec, err := base58.Decode(s)
	if err != nil {
		return err
	}

	*p = dec
	return nil
}

func (p *PublicKey) Verify(data []byte, signature signature.Signature) (ok bool, err error) {
	keyType := p.TypeByte()
	if signature.Type() != keyType {
		return false, fmt.Errorf("cannot verify signature type %d with key type %d", signature.Type(), p.TypeByte())
	}

	switch keyType {
	case RawKeyTypeED25519:
		ok = ed25519.Verify(ed25519.PublicKey(p.Value()), data, signature.Value())
	case RawKeyTypeSECP256K1:
		var pubKey *secp256k1.PublicKey
		pubKey, err = secp256k1.ParsePubKey(p.Value())
		if err != nil {
			return
		}

		var sign *ecdsa.Signature
		sign, err = ecdsa.ParseDERSignature(signature.Value())
		if err != nil {
			return
		}

		ok = sign.Verify(data, pubKey)
	}

	return
}

func (p *PublicKey) ToBase58PublicKey() Base58PublicKey {
	return Base58PublicKey{
		Type:  keyTypes[p.TypeByte()],
		Value: base58.Encode(p.Value()),
		pk:    *p,
	}
}

func PublicKeyFromBytes(b []byte) (pk PublicKey, err error) {
	f := b[0]
	l := len(b) - 1
	switch f {
	case RawKeyTypeED25519:
		if l != ed25519.PublicKeySize {
			return pk, ErrInvalidPublicKey
		}
		pk = b
		return
	case RawKeyTypeSECP256K1:
		if l != secp256k1.PubKeyBytesLenCompressed {
			return pk, ErrInvalidPublicKey
		}
		pk = b
		return
	}

	return pk, ErrInvalidKeyType
}

func WrapRawKey(keyType PublicKeyType, key []byte) (pk PublicKey, err error) {
	switch keyType {
	case KeyTypeED25519:
		if len(key) != ed25519.PublicKeySize {
			return pk, ErrInvalidPublicKey
		}

		pk = make([]byte, ed25519.PublicKeySize+1)
		pk[0] = RawKeyTypeED25519
		copy(pk[1:], key[0:ed25519.PublicKeySize])
		return
	case KeyTypeSECP256K1:
		if len(key) != secp256k1.PubKeyBytesLenCompressed {
			return pk, ErrInvalidPublicKey
		}

		pk = make([]byte, secp256k1.PubKeyBytesLenCompressed+1)
		pk[0] = RawKeyTypeSECP256K1
		copy(pk[1:], key[0:secp256k1.PubKeyBytesLenCompressed])
		return
	}

	return pk, ErrInvalidKeyType
}

func WrapED25519(key ed25519.PublicKey) PublicKey {
	if pk, err := WrapRawKey(KeyTypeED25519, key); err != nil {
		panic(err)
	} else {
		return pk
	}
}
