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

func (t Transaction) HashAndSign(key ed25519.PrivateKey) (hash.CryptoHash, []byte, Signature, error) {
	// Serialize into Borsh
	serialized, err := borsh.Serialize(t)
	if err != nil {
		return hash.CryptoHash{}, nil, Signature{}, err
	}

	// Sign the transaction
	txnPayloadHash := hash.NewCryptoHash(serialized)
	signature := ed25519.Sign(key, txnPayloadHash[:])
	return txnPayloadHash, serialized, NewSignatureED25519(signature), nil
}
