package client

import (
	"context"
)

// https://docs.near.org/docs/api/rpc#network-info
func (c *Client) NetworkInfo(ctx context.Context) (res NetworkInfo, err error) {
	_, err = c.doRPC(ctx, &res, "network_info", nil, []string{})

	return
}
