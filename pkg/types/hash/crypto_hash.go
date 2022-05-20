package hash

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/mr-tron/base58"
)

// CryptoHash is a wrapper for SHA-256 digest byte array.
// Note that nearcore also defines MerkleHash as an alias, but it's omitted from this project.
type CryptoHash [sha256.Size]byte

func (c *CryptoHash) UnmarshalJSON(b []byte) (err error) {
	var s string
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	if *c, err = NewCryptoHashFromBase58(s); err != nil {
		return
	}

	return nil
}

func (c CryptoHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c CryptoHash) String() string {
	return base58.Encode(c[:])
}

func NewCryptoHash(data []byte) CryptoHash {
	return CryptoHash(sha256.Sum256(data))
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

func MustCryptoHashFromBase58(blob string) CryptoHash {
	if hash, err := NewCryptoHashFromBase58(blob); err != nil {
		panic(err)
	} else {
		return hash
	}
}
