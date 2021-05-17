package transaction

import (
	"crypto/ed25519"

	"github.com/eteu-technologies/borsh-go"

	"github.com/eteu-technologies/near-api-go/types/action"
	"github.com/eteu-technologies/near-api-go/types/hash"
	"github.com/eteu-technologies/near-api-go/types/key"
)

type Transaction struct {
	SignerID   string
	PublicKey  key.PublicKey
	Nonce      uint64
	ReceiverID string
	BlockHash  hash.CryptoHash
	Actions    []action.Action
}

func (t Transaction) Hash() (txnHash hash.CryptoHash, serialized []byte, err error) {
	// Serialize into Borsh
	serialized, err = borsh.Serialize(t)
	if err != nil {
		return
	}
	txnHash = hash.NewCryptoHash(serialized)
	return
}

func (t Transaction) HashAndSign(privKey ed25519.PrivateKey) (txnHash hash.CryptoHash, serialized []byte, sig key.Signature, err error) {
	txnHash, serialized, err = t.Hash()
	if err != nil {
		return
	}

	sig = key.NewSignatureED25519(ed25519.Sign(privKey, txnHash[:]))
	return
}
