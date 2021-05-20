package client

import (
	"encoding/json"

	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
)

type AccessKey struct {
	Nonce      types.Nonce     `json:"nonce"`
	Permission json.RawMessage `json:"permission"` // TODO: conditional decoding
}

type AccessKeyView struct {
	AccessKey
	BlockHeight uint64          `json:"block_height"`
	BlockHash   hash.CryptoHash `json:"block_hash"`
}

type AccessKeyViewInfo struct {
	PublicKey key.Base58PublicKey `json:"public_key"`
	AccessKey AccessKey           `json:"access_key"`
}

type AccessKeyList struct {
	Keys []AccessKeyViewInfo `json:"keys"`
}
