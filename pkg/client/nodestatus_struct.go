package client

import (
	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
	"time"
)

type NodeStatus struct {
	// Binary version
	Version NodeVersion `json:"version"`
	// Unique chain id.
	ChainID string `json:"chain_id"`
	// Currently active protocol version.
	ProtocolVersion uint32 `json:"protocol_version"`
	// Latest protocol version that this client supports.
	LatestProtocolVersion uint32 `json:"latest_protocol_version"`
	// Address for RPC server.
	RPCAddr string `json:"rpc_addr"`
	// Current epoch validators.
	Validators []ValidatorInfo `json:"validators"`
	// Sync status of the node.
	SyncInfo StatusSyncInfo `json:"sync_info"`
	// Validator id of the node
	ValidatorAccountID *types.AccountID `json:"validator_account_id"`
}

type NodeVersion struct {
	Version string `json:"version"`
	Build   string `json:"build"`
}

type ValidatorInfo struct {
	AccountID types.AccountID `json:"account_id"`
	Slashed   bool            `json:"is_slashed"`
}

type StatusSyncInfo struct {
	LatestBlockHash   hash.CryptoHash   `json:"latest_block_hash"`
	LatestBlockHeight types.BlockHeight `json:"latest_block_height"`
	LatestBlockTime   time.Time         `json:"latest_block_time"`
	Syncing           bool              `json:"syncing"`
}

type ValidatorsResponse struct {
	CurrentValidators []CurrentEpochValidatorInfo `json:"current_validator"`
}

type CurrentEpochValidatorInfo struct {
	ValidatorInfo
	PublicKey         key.PublicKey   `json:"public_key"`
	Stake             types.Balance   `json:"stake"`
	Shards            []types.ShardID `json:"shards"`
	NumProducedBlocks types.NumBlocks `json:"num_produced_blocks"`
	NumExpectedBlocks types.NumBlocks `json:"num_expected_blocks"`
}
