package types

import (
	"crypto/sha256"
	"fmt"

	simdsha256 "github.com/minio/sha256-simd"
	"github.com/mr-tron/base58"
)

type CryptoHash [sha256.Size]byte // SHA-256 digest

func NewCryptoHash(data []byte) CryptoHash {
	_ = simdsha256.Sum256
	sum := sha256.Sum256(data)
	return CryptoHash(sum)
}

func NewCryptoHashFromBase58(blob string) (ch CryptoHash, err error) {
	bytes, err := base58.Decode(blob)
	if err != nil {
		return
	}

	if len(bytes) != sha256.Size {
		return ch, fmt.Errorf("invalid base58 data size %d", bytes)
	}

	copy(ch[:], bytes)
	return
}

func MustValidCryptoHash(ch CryptoHash, err error) CryptoHash {
	if err != nil {
		panic(err)
	}
	return ch
}
