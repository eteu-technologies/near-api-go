package transaction

import (
	"crypto/ed25519"

	"github.com/eteu-technologies/near-api-go/types/hash"
)

type SignedTransaction struct {
	Transaction Transaction
	Signature   Signature

	SerializedTransaction []byte          `borsh_skip:"true"`
	hash                  hash.CryptoHash `borsh_skip:"true"`
	size                  int             `borsh_skip:"true"`
}

func NewSignedTransaction(transaction Transaction, signingKey ed25519.PrivateKey) (stxn SignedTransaction, err error) {
	stxn.Transaction = transaction
	stxn.hash, stxn.SerializedTransaction, stxn.Signature, err = transaction.HashAndSign(signingKey)
	if err != nil {
		return
	}

	stxn.size = len(stxn.SerializedTransaction)
	return
}

func (st *SignedTransaction) Hash() hash.CryptoHash {
	return st.hash
}

func (st *SignedTransaction) Size() int {
	return st.size
}
