package client

import (
	"encoding/json"

	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/action"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
	"github.com/eteu-technologies/near-api-go/pkg/types/signature"
)

type TransactionStatus struct {
	SuccessValue     string          `json:"SuccessValue"`
	SuccessReceiptID string          `json:"SuccessReceiptId"`
	Failure          json.RawMessage `json:"Failure"` // TODO
}

type SignedTransactionView struct {
	SignerID   types.AccountID           `json:"signer_id"`
	PublicKey  key.Base58PublicKey       `json:"public_key"`
	Nonce      types.Nonce               `json:"nonce"`
	ReceiverID types.AccountID           `json:"receiver_id"`
	Actions    []action.Action           `json:"actions"`
	Signature  signature.Base58Signature `json:"signature"`
	Hash       hash.CryptoHash           `json:"hash"`
}

type FinalExecutionOutcomeView struct {
	Status             TransactionStatus            `json:"status"`
	Transaction        SignedTransactionView        `json:"transaction"`
	TransactionOutcome ExecutionOutcomeWithIdView   `json:"transaction_outcome"`
	ReceiptsOutcome    []ExecutionOutcomeWithIdView `json:"receipts_outcome"`
}

type FinalExecutionOutcomeWithReceiptView struct {
	FinalExecutionOutcomeView
	Receipts []ReceiptView `json:"receipts"`
}

type ReceiptView struct {
	PredecessorId types.AccountID `json:"predecessor_id"`
	ReceiverID    types.AccountID `json:"receiver_id"`
	ReceiptID     hash.CryptoHash `jsom:"receipt_id"`
	Receipt       json.RawMessage `json:"receipt"` // TODO: needs a type!
}

type ExecutionOutcomeView struct {
	Logs        []string          `json:"logs"`
	ReceiptIDs  []hash.CryptoHash `json:"receipt_ids"`
	GasBurnt    types.Gas         `json:"gas_burnt"`
	TokensBurnt types.Balance     `json:"tokens_burnt"`
	ExecutorID  types.AccountID   `json:"executor_id"`
	Status      TransactionStatus `json:"status"`
}

type MerklePathItem struct {
	Hash      hash.CryptoHash `json:"hash"`
	Direction string          `json:"direction"` // TODO: enum type, either 'Left' or 'Right'
}

type MerklePath = []MerklePathItem

type ExecutionOutcomeWithIdView struct {
	Proof     MerklePath           `json:"proof"`
	BlockHash hash.CryptoHash      `json:"block_hash"`
	ID        hash.CryptoHash      `json:"id"`
	Outcome   ExecutionOutcomeView `json:"outcome"`
}
