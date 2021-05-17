package transaction

import (
	"github.com/eteu-technologies/near-api-go/types"
	"github.com/eteu-technologies/near-api-go/types/hash"
)

type ExcecutionOutcome struct {
	Logs        []string
	ReceiptIds  []hash.CryptoHash
	GasBurnt    uint64 // TODO: panic: reflect.Set: value of type uint64 is not assignable to type types.Gas
	TokensBurnt types.Balance
	ExecutorID  types.AccountID
	Status      ExecutionStatus
}
