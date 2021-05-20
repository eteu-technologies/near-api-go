package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#genesis-config
func (c *Client) GenesisConfig(ctx context.Context) (res jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_genesis_config", nil, nil)

	return
}
