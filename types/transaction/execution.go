package transaction

import (
	"github.com/eteu-technologies/near-api-go/types"
	"github.com/eteu-technologies/near-api-go/types/hash"
)

type ExcecutionOutcome struct {
	Logs        []string
	ReceiptIds  []hash.CryptoHash
	GasBurnt    types.Gas
	TokensBurnt types.Balance
	ExecutorID  types.AccountID
	Status      ExecutionStatus
}
