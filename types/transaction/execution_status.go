package transaction

import (
	"github.com/eteu-technologies/borsh-go"
	"github.com/eteu-technologies/near-api-go/types/hash"
)

type ExecutionStatus struct {
	Enum borsh.Enum `borsh_enum:"true"`

	// The execution is pending or unknown.
	Unknown struct{}

	// The execution has failed with the given execution error.
	Failure ExecutionStatusFailure

	// The final action succeeded and returned some value or an empty vec.
	SuccessValue []byte

	// The final action of the receipt returned a promise or the signed transaction was converted to a receipt. Contains the receipt_id of the generated receipt.
	SuccessReceiptID hash.CryptoHash
}

// TODO: core/primitives/src/errors.rs & existing errors package is beyond useless for this :(
type ExecutionStatusFailure struct {
	Enum borsh.Enum `borsh_enum:"true"`

	// TODO
	ActionError struct{}

	// TODO
	InvalidTxError struct{}
}
