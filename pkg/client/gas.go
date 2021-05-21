package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
)

// https://docs.near.org/docs/develop/front-end/rpc#gas-price
func (c *Client) GasPriceView(ctx context.Context, block block.BlockCharacteristic) (res GasPrice, err error) {
	_, err = c.doRPC(ctx, &res, "gas_price", nil, blockIDArrayParams(block))

	return
}
