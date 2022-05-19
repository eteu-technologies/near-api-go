package client

import (
	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
	"github.com/eteu-technologies/near-api-go/pkg/types/signature"
)

type ChallengesResult = []SlashedValidator

type SlashedValidator struct {
	AccountID    types.AccountID `json:"account_id"`
	IsDoubleSign bool            `json:"is_double_sign"`
}

// ValidatorStakeView is based on ValidatorStakeV1 struct in nearcore
type ValidatorStakeView struct {
	AccountID types.AccountID     `json:"account_id"`
	PublicKey key.Base58PublicKey `json:"public_key"`
	Stake     types.Balance       `json:"stake"`
}

type BlockView struct {
	Author types.AccountID   `json:"author"`
	Header BlockHeaderView   `json:"header"`
	Chunks []ChunkHeaderView `json:"chunks"`
}

type BlockHeaderView struct {
	Height                types.BlockHeight            `json:"height"`
	EpochID               hash.CryptoHash              `json:"epoch_id"`
	NextEpochID           hash.CryptoHash              `json:"next_epoch_id"`
	Hash                  hash.CryptoHash              `json:"hash"`
	PrevHash              hash.CryptoHash              `json:"prev_hash"`
	PrevStateRoot         hash.CryptoHash              `json:"prev_state_root"`
	ChunkReceiptsRoot     hash.CryptoHash              `json:"chunk_receipts_root"`
	ChunkHeadersRoot      hash.CryptoHash              `json:"chunk_headers_root"`
	ChunkTxRoot           hash.CryptoHash              `json:"chunk_tx_root"`
	OutcomeRoot           hash.CryptoHash              `json:"outcome_root"`
	ChunksIncluded        uint64                       `json:"chunks_included"`
	ChallengesRoot        hash.CryptoHash              `json:"challenges_root"`
	Timestamp             uint64                       `json:"timestamp"`         // milliseconds
	TimestampNanosec      types.TimeNanos              `json:"timestamp_nanosec"` // nanoseconds, uint128
	RandomValue           hash.CryptoHash              `json:"random_value"`
	ValidatorProposals    []ValidatorStakeView         `json:"validator_proposals"`
	ChunkMask             []bool                       `json:"chunk_mask"`
	GasPrice              types.Balance                `json:"gas_price"`
	RentPaid              types.Balance                `json:"rent_paid"`        // NOTE: deprecated - 2021-05-14
	ValidatorReward       types.Balance                `json:"validator_reward"` // NOTE: deprecated - 2021-05-14
	TotalSupply           types.Balance                `json:"total_supply"`
	ChallengesResult      ChallengesResult             `json:"challenges_result"`
	LastFinalBlock        hash.CryptoHash              `json:"last_final_block"`
	LastDSFinalBlock      hash.CryptoHash              `json:"last_ds_final_block"`
	NextBPHash            hash.CryptoHash              `json:"next_bp_hash"`
	BlockMerkleRoot       hash.CryptoHash              `json:"block_merkle_root"`
	Approvals             []*signature.Base58Signature `json:"approvals"`
	Signature             signature.Base58Signature    `json:"signature"`
	LatestProtocolVersion uint64                       `json:"latest_protocol_version"`
}
