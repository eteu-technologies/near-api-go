package transaction

import "crypto/ed25519"

// TODO: SECP256K1 support
type Signature [1 + ed25519.SignatureSize]byte

func NewSignatureED25519(data []byte) Signature {
	var buf Signature
	buf[0] = 0
	copy(buf[1:], data[0:ed25519.SignatureSize])
	return buf
}
