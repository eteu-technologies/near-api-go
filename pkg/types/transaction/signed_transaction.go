package transaction

import (
	"encoding/base64"

	"github.com/eteu-technologies/borsh-go"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
	"github.com/eteu-technologies/near-api-go/pkg/types/signature"
)

type SignedTransaction struct {
	Transaction Transaction
	Signature   signature.Signature

	SerializedTransaction []byte          `borsh_skip:"true"`
	hash                  hash.CryptoHash `borsh_skip:"true"`
	size                  int             `borsh_skip:"true"`
}

func NewSignedTransaction(keyPair key.KeyPair, transaction Transaction) (stxn SignedTransaction, err error) {
	stxn.Transaction = transaction
	stxn.hash, stxn.SerializedTransaction, stxn.Signature, err = transaction.HashAndSign(keyPair)
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

func (st SignedTransaction) Serialize() (serialized string, err error) {
	var blob []byte

	blob, err = borsh.Serialize(st)
	if err != nil {
		return
	}

	serialized = base64.StdEncoding.EncodeToString(blob)

	return
}
