package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#gas-price
func (c *Client) GasPriceView(ctx context.Context, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "gas_price", nil, blockIDArrayParams(block))

	return
}
