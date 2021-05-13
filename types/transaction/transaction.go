package transaction

import (
	"crypto/ed25519"

	"github.com/eteu-technologies/borsh-go"

	"github.com/eteu-technologies/near-api-go/types"
	. "github.com/eteu-technologies/near-api-go/types/action"
	. "github.com/eteu-technologies/near-api-go/types/hash"
)

// NOTE: jsonrpc params -> something in that
type RPCBroadcastTransactionRequest struct {
	SignedTransaction SignedTransaction
}

type SignedTransaction struct {
	Transaction Transaction
	Signature   Signature

	SerializedTransaction []byte     `borsh_skip:"true"`
	hash                  CryptoHash `borsh_skip:"true"`
	size                  int        `borsh_skip:"true"`
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

func (st *SignedTransaction) Hash() CryptoHash {
	return st.hash
}

func (st *SignedTransaction) Size() int {
	return st.size
}

type Transaction struct {
	SignerID   string
	PublicKey  PublicKey
	Nonce      uint64
	ReceiverID string
	BlockHash  CryptoHash
	Actions    []Action
}

func (t *Transaction) HashAndSign(key ed25519.PrivateKey) (CryptoHash, []byte, Signature, error) {
	// Serialize into Borsh
	serialized, err := borsh.Serialize(*t)
	if err != nil {
		return CryptoHash{}, nil, Signature{}, err
	}

	// Compute hash
	hash := NewCryptoHash(serialized) // XXX: what?

	// Sign
	signature := ed25519.Sign(key, hash[:])
	return hash, serialized, NewSignatureED25519(signature), nil
}

type ExcecutionOutcome struct {
	Logs        []string // TODO: LogEntry type
	ReceiptIds  []CryptoHash
	GasBurnt    types.Gas
	TokensBurnt types.Balance
	ExecutorID  types.AccountID
	Status      interface{} // TODO: ExecutionStatus
}

// TODO: SECP256K1
type Signature [1 + ed25519.SignatureSize]byte

func NewSignatureED25519(data []byte) Signature {
	var buf [65]byte

	bbuf := []byte{0x0}
	bbuf = append(bbuf, data...)

	copy(buf[:], bbuf)
	return buf
}

// TODO: SECP256K1
type PublicKey [33]byte

func PublicKeyFromED25519Key(key ed25519.PublicKey) PublicKey {
	var buf [33]byte

	bbuf := []byte{0x0}
	bbuf = append(bbuf, key...)

	copy(buf[:], bbuf)
	return buf
}
