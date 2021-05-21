package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
)

// https://docs.near.org/docs/api/rpc#block-details
func (c *Client) BlockDetails(ctx context.Context, block block.BlockCharacteristic) (resp BlockView, err error) {
	_, err = c.doRPC(ctx, &resp, "block", block, map[string]interface{}{})

	return
}

// TODO: decode resposne
// https://docs.near.org/docs/api/rpc#changes-in-block
func (c *Client) BlockChanges(ctx context.Context, block block.BlockCharacteristic) (res jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes_in_block", block, map[string]interface{}{})

	return
}
