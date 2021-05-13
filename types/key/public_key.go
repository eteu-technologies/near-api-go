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
	case byte(RawPublicKeyTypeED25519):
		if l != ed25519.PublicKeySize {
			return pk, ErrInvalidPublicKey
		}
		copy(pk[:], b)
		return
	case byte(RawPublicKeyTypeSECP256K1):
		// TODO!
		return pk, fmt.Errorf("SECP256K1 is not supported yet")
	}

	return pk, ErrInvalidPublicKeyType
}

func WrapED25519(key ed25519.PublicKey) PublicKey {
	var buf [33]byte

	bbuf := []byte{0x0}
	bbuf = append(bbuf, key...)

	copy(buf[:], bbuf)
	return buf
}
