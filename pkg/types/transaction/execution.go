package transaction

import (
	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
)

type ExcecutionOutcome struct {
	Logs        []string
	ReceiptIds  []hash.CryptoHash
	GasBurnt    types.Gas
	TokensBurnt types.Balance
	ExecutorID  types.AccountID
	Status      ExecutionStatus
}
