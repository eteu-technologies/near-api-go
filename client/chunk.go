package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/jsonrpc"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#chunk-details
func (c *Client) ChunkDetails(ctx context.Context, chunkHash string) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "chunk", nil, []string{chunkHash})

	return
}
