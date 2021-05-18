package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eteu-technologies/near-api-go/client/block"
	"github.com/eteu-technologies/near-api-go/jsonrpc"
)

// https://docs.near.org/docs/develop/front-end/rpc#block-details
func (c *Client) BlockDetails(ctx context.Context, block block.BlockCharacteristic) (resp BlockView, err error) {
	var res jsonrpc.JSONRPCResponse
	res, err = c.doRPC(ctx, "block", block, map[string]interface{}{})

	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf("%s", string(*res.Error))
		return
	}

	if err = json.Unmarshal(res.Result, &resp); err != nil {
		return
	}

	return
}

// TODO: decode resposne
// https://docs.near.org/docs/develop/front-end/rpc#changes-in-block
func (c *Client) BlockChanges(ctx context.Context, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "EXPERIMENTAL_changes_in_block", block, map[string]interface{}{})

	return
}
