package client

import (
	"encoding/json"

	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
	"github.com/eteu-technologies/near-api-go/pkg/types/signature"
)

type ChunkView struct {
	Author       types.AccountID         `json:"author"`
	Header       ChunkHeaderView         `json:"header"`
	Transactions []SignedTransactionView `json:"transactions"`
	Receipts     []ReceiptView           `json:"receipts"`
}

type ChunkHeaderView struct {
	ChunkHash            hash.CryptoHash           `json:"chunk_hash"`
	PrevBlockHash        hash.CryptoHash           `json:"prev_block_hash"`
	OutcomeRoot          hash.CryptoHash           `json:"outcome_root"`
	PrevStateRoot        json.RawMessage           `json:"prev_state_root"` // TODO: needs a type!
	EncodedMerkleRoot    hash.CryptoHash           `json:"encoded_merkle_root"`
	EncodedLength        uint64                    `json:"encoded_length"`
	HeightCreated        types.BlockHeight         `json:"height_created"`
	HeightIncluded       types.BlockHeight         `json:"height_included"`
	ShardID              types.ShardID             `json:"shard_id"`
	GasUsed              types.Gas                 `json:"gas_used"`
	GasLimit             types.Gas                 `json:"gas_limit"`
	RentPaid             types.Balance             `json:"rent_paid"`        // TODO: deprecated
	ValidatorReward      types.Balance             `json:"validator_reward"` // TODO: deprecated
	BalanceBurnt         types.Balance             `json:"balance_burnt"`
	OutgoingReceiptsRoot hash.CryptoHash           `json:"outgoing_receipts_root"`
	TxRoot               hash.CryptoHash           `json:"tx_root"`
	ValidatorProposals   []ValidatorStakeView      `json:"validator_proposals"`
	Signature            signature.Base58Signature `json:"signature"`
}
