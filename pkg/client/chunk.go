package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#chunk-details
func (c *Client) ChunkDetails(ctx context.Context, chunkHash hash.CryptoHash) (res jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, "chunk", nil, []string{chunkHash.String()})

	return
}
