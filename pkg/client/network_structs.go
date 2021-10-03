package client

import (
	"encoding/json"
	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
	"github.com/eteu-technologies/near-api-go/pkg/types/signature"
)

// NetworkInfo holds network information
type NetworkInfo struct {
	ActivePeers         []FullPeerInfo  `json:"active_peers"`
	NumActivePeers      uint            `json:"num_active_peers"`
	PeerMaxCount        uint32          `json:"peer_max_count"`
	HighestHeightPeers  []FullPeerInfo  `json:"highest_height_peers"`
	SentBytesPerSec     uint64          `json:"sent_bytes_per_sec"`
	ReceivedBytesPerSec uint64          `json:"received_bytes_per_sec"`
	KnownProducers      []KnownProducer `json:"known_producers"`
	MetricRecorder      MetricRecorder  `json:"metric_recorder"`
	PeerCounter         uint            `json:"peer_counter"`
}

type FullPeerInfo struct {
	PeerInfo  PeerInfo      `json:"peer_info"`
	ChainInfo PeerChainInfo `json:"chain_info"`
	EdgeInfo  EdgeInfo      `json:"edge_info"`
}

// PeerInfo holds peer information
type PeerInfo struct {
	ID        key.PeerID       `json:"id"`
	Addr      *string          `json:"addr"`
	AccountID *types.AccountID `json:"account_id"`
}

// PeerChainInfo contains peer chain information. This is derived from PeerCHainInfoV2 in nearcore
type PeerChainInfo struct {
	// Chain Id and hash of genesis block.
	GenesisID GenesisID `json:"genesis_id"`
	// Last known chain height of the peer.
	Height types.BlockHeight `json:"height"`
	// Shards that the peer is tracking.
	TrackedShards []types.ShardID `json:"tracked_shards"`
	// Denote if a node is running in archival mode or not.
	Archival bool `json:"archival"`
}

// EdgeInfo contains information that will be ultimately used to create a new edge. It contains nonce proposed for the edge with signature from peer.
type EdgeInfo struct {
	Nonce     types.Nonce         `json:"nonce"`
	Signature signature.Signature `json:"signature"`
}

// KnownProducer is basically PeerInfo, but AccountID is known
type KnownProducer struct {
	AccountID types.AccountID `json:"account_id"`
	Addr      *string         `json:"addr"`
	PeerID    key.PeerID      `json:"peer_id"`
}

// TODO: chain/network/src/recorder.rs
type MetricRecorder = json.RawMessage

type GenesisID struct {
	// Chain Id
	ChainID string `json:"chain_id"`
	// Hash of genesis block
	Hash hash.CryptoHash `json:"hash"`
}
