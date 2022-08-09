package client

import (
	"context"
)

// https://docs.near.org/api/rpc/protocol#protocol-config
func (c *Client) ProtocolConfig(ctx context.Context) (res map[string]interface{}, err error) {
	_, err = c.doRPC(ctx, &res, "EXPERIMENTAL_protocol_config", nil, nil)

	return
}
