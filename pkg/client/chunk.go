package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
)

// https://docs.near.org/docs/api/rpc#chunk-details
func (c *Client) ChunkDetails(ctx context.Context, chunkHash hash.CryptoHash) (res ChunkView, err error) {
	_, err = c.doRPC(ctx, &res, "chunk", nil, []string{chunkHash.String()})

	return
}
