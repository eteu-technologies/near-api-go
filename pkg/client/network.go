package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#network-info
func (c *Client) NetworkInfo(ctx context.Context) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "network_info", nil, []string{})

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#general-validator-status
func (c *Client) NetworkStatusValidators(ctx context.Context) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "status", nil, []string{})

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#detailed-validator-status
func (c *Client) NetworkStatusValidatorsDetailed(ctx context.Context, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "validators", nil, blockIDArrayParams(block))

	return
}
