package types

import (
	"crypto/ed25519"
	"math/big"

	"github.com/near/borsh-go"
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
	GasBurnt    Gas
	TokensBurnt Balance
	ExecutorID  AccountID
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

// Account identifier. Provides access to user's state.
type AccountID string

// Gas is a type for storing amounts of gas.
type Gas uint64

type Balance big.Int

// Nonce for transactions.
type Nonce uint64

// ExecutionStatus
// - Unknown; The execution is pending or unknown.
// - Failure(TxExecutionError); The execution has failed with the given execution error.
// - SuccessValue([]byte); The final action succeeded and returned some value or an empty vec.
// - SucessReceiptId(CryptoHash); The final action of the receipt returned a promise or the signed transaction was converted to a receipt. Contains the receipt_id of the generated receipt.
