package transaction

import (
	"crypto/ed25519"

	"github.com/eteu-technologies/borsh-go"

	"github.com/eteu-technologies/near-api-go/types"
	. "github.com/eteu-technologies/near-api-go/types/action"
	. "github.com/eteu-technologies/near-api-go/types/hash"
	"github.com/eteu-technologies/near-api-go/types/key"
)

// NOTE: jsonrpc params -> something in that
type RPCBroadcastTransactionRequest struct {
	SignedTransaction SignedTransaction
}

type Transaction struct {
	SignerID   string
	PublicKey  key.PublicKey
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
	GasBurnt    uint64 // TODO: panic: reflect.Set: value of type uint64 is not assignable to type types.Gas
	TokensBurnt types.Balance
	ExecutorID  types.AccountID
	Status      interface{} // TODO: ExecutionStatus
}
