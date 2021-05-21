package client

import (
	"context"
)

// https://docs.near.org/docs/api/rpc#genesis-config
func (c *Client) GenesisConfig(ctx context.Context) (res map[string]interface{}, err error) {
	_, err = c.doRPC(ctx, &res, "EXPERIMENTAL_genesis_config", nil, nil)

	return
}
