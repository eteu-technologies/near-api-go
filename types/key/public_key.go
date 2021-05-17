package key

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/mr-tron/base58"
)

// TODO: SECP256K1
type PublicKey [33]byte

func (p PublicKey) Hash() string {
	return hex.EncodeToString(p[1:])
}

func (p PublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(p[:]))
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

	*p = PublicKey{}
	copy(p[:], dec)
	return nil
}

func PublicKeyFromBytes(b []byte) (pk PublicKey, err error) {
	f := b[0]
	l := len(b) - 1
	switch f {
	case RawPublicKeyTypeED25519:
		if l != ed25519.PublicKeySize {
			return pk, ErrInvalidPublicKey
		}
		copy(pk[:], b)
		return
	case RawPublicKeyTypeSECP256K1:
		// TODO!
		return pk, fmt.Errorf("SECP256K1 is not supported yet")
	}

	return pk, ErrInvalidPublicKeyType
}

func WrapED25519(key ed25519.PublicKey) PublicKey {
	var buf PublicKey
	buf[0] = RawPublicKeyTypeED25519
	copy(buf[1:], key[0:ed25519.PublicKeySize])
	return buf
}
