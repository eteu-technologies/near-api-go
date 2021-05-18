package client

import (
	"encoding/json"
	"github.com/eteu-technologies/near-api-go/types"
	"github.com/eteu-technologies/near-api-go/types/hash"
	"github.com/eteu-technologies/near-api-go/types/key"
)

type ChallengesResult = []SlashedValidator

type SlashedValidator struct {
	AccountID    types.AccountID `json:"account_id"`
	IsDoubleSign bool            `json:"is_double_sign"`
}

// ValidatorStake is based on ValidatorStakeV1 struct in nearcore
type ValidatorStakeView struct {
	AccountID types.AccountID `json:"account_id"`
	PublicKey key.PublicKey   `json:"public_key"`
	Stake     types.Balance   `json:"stake"`
}

type BlockView struct {
	Author types.AccountID   `json:"author"`
	Header BlockHeaderView   `json:"header"`
	Chunks []ChunkHeaderView `json:"chunks"`
}

type BlockHeaderView struct {
	Height                types.BlockHeight    `json:"height"`
	EpochID               hash.CryptoHash      `json:"epoch_id"`
	NextEpochID           hash.CryptoHash      `json:"next_epoch_id"`
	Hash                  hash.CryptoHash      `json:"hash"`
	PrevHash              hash.CryptoHash      `json:"prev_hash"`
	PrevStateRoot         hash.CryptoHash      `json:"prev_state_root"`
	ChunkReceiptsRoot     hash.CryptoHash      `json:"chunk_receipts_root"`
	ChunkHeadersRoot      hash.CryptoHash      `json:"chunk_headers_root"`
	ChunkTxRoot           hash.CryptoHash      `json:"chunk_tx_root"`
	OutcomeRoot           hash.CryptoHash      `json:"outcome_root"`
	ChunksIncluded        uint64               `json:"chunks_included"`
	ChallengesRoot        hash.CryptoHash      `json:"challenges_root"`
	Timestamp             uint64               `json:"timestamp"`         // milliseconds
	TimestampNanosec      string               `json:"timestamp_nanosec"` // nanoseconds, uint128
	RandomValue           hash.CryptoHash      `json:"random_value"`
	ValidatorProposals    []ValidatorStakeView `json:"validator_proposals"`
	ChunkMask             []bool               `json:"chunk_mask"`
	GasPrice              types.Balance        `json:"gas_price"`
	RentPaid              types.Balance        `json:"rent_paid"`        // NOTE: deprecated - 2021-05-14
	ValidatorReward       types.Balance        `json:"validator_reward"` // NOTE: deprecated - 2021-05-14
	TotalSupply           types.Balance        `json:"total_supply"`
	ChallengesResult      ChallengesResult     `json:"challenges_result"`
	LastFinalBlock        hash.CryptoHash      `json:"last_final_block"`
	LastDSFinalBlock      hash.CryptoHash      `json:"last_ds_final_block"`
	NextBPHash            hash.CryptoHash      `json:"next_bp_hash"`
	BlockMerkleRoot       hash.CryptoHash      `json:"block_merkle_root"`
	Approvals             []string             `json:"approvals"` // TODO: array of nullable ed25519 signatures
	Signature             string               `json:"signature"` // TODO: ed25519 signature
	LatestProtocolVersion uint64               `json:"latest_protocol_version"`
}

type ChunkHeaderView struct {
	ChunkHash            hash.CryptoHash      `json:"chunk_hash"`
	PrevBlockHash        hash.CryptoHash      `json:"prev_block_hash"`
	OutcomeRoot          hash.CryptoHash      `json:"outcome_root"`
	PrevStateRoot        json.RawMessage      `json:"prev_state_root"` // TODO: needs a type!
	EncodedMerkleRoot    hash.CryptoHash      `json:"encoded_merkle_root"`
	EncodedLength        uint64               `json:"encoded_length"`
	HeightCreated        types.BlockHeight    `json:"height_created"`
	HeightIncluded       types.BlockHeight    `json:"height_included"`
	ShardID              types.ShardID        `json:"shard_id"`
	GasUsed              types.Gas            `json:"gas_used"`
	GasLimit             types.Gas            `json:"gas_limit"`
	RentPaid             types.Balance        `json:"rent_paid"`        // TODO: deprecated
	ValidatorReward      types.Balance        `json:"validator_reward"` // TODO: deprecated
	BalanceBurnt         types.Balance        `json:"balance_burnt"`
	OutgoingReceiptsRoot hash.CryptoHash      `json:"outgoing_receipts_root"`
	TxRoot               hash.CryptoHash      `json:"tx_root"`
	ValidatorProposals   []ValidatorStakeView `json:"validator_proposals"`
	Signature            json.RawMessage      `json:"signature"` // TODO: needs a type!
}
