package signature

import (
	"crypto/ed25519"
)

type Signature []byte

func NewSignatureED25519(data []byte) Signature {
	buf := make([]byte, 1+ed25519.SignatureSize)
	buf[0] = RawSignatureTypeED25519
	copy(buf[1:], data[0:ed25519.SignatureSize])
	return buf
}

func (s Signature) Type() byte {
	return s[0]
}

func (s Signature) Value() []byte {
	return s[1:]
}

func NewSignatureSECP256K1(data []byte) Signature {
	sign := make([]byte, 1+len(data))
	sign[0] = RawSignatureTypeSECP256K1
	copy(sign[1:], data)
	return sign
}
