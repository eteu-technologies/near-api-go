package transaction

import (
	"crypto/ed25519"

	"github.com/eteu-technologies/near-api-go/types/hash"
	"github.com/eteu-technologies/near-api-go/types/key"
)

type SignedTransaction struct {
	Transaction Transaction
	Signature   key.Signature

	SerializedTransaction []byte          `borsh_skip:"true"`
	hash                  hash.CryptoHash `borsh_skip:"true"`
	size                  int             `borsh_skip:"true"`
}

func NewSignedTransaction(signingKey ed25519.PrivateKey, transaction Transaction) (stxn SignedTransaction, err error) {
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

func (st *SignedTransaction) Verify(pubKey key.PublicKey) (ok bool, err error) {
	var txnHash hash.CryptoHash
	txnHash, _, err = st.Transaction.Hash()
	if err != nil {
		return
	}

	return pubKey.Verify(txnHash[:], st.Signature)
}
