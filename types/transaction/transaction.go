package transaction

import (
	"github.com/eteu-technologies/borsh-go"

	"github.com/eteu-technologies/near-api-go/types"
	"github.com/eteu-technologies/near-api-go/types/action"
	"github.com/eteu-technologies/near-api-go/types/hash"
	"github.com/eteu-technologies/near-api-go/types/key"
)

type Transaction struct {
	SignerID   types.AccountID
	PublicKey  key.PublicKey
	Nonce      types.Nonce
	ReceiverID types.AccountID
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

func (t Transaction) HashAndSign(keyPair key.KeyPair) (txnHash hash.CryptoHash, serialized []byte, sig key.Signature, err error) {
	txnHash, serialized, err = t.Hash()
	if err != nil {
		return
	}

	sig = keyPair.Sign(txnHash[:])
	return
}
