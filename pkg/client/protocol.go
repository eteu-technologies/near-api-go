package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
)

// https://docs.near.org/api/rpc/protocol#protocol-config
func (c *Client) ProtocolConfig(ctx context.Context, block block.BlockCharacteristic) (res map[string]interface{}, err error) {
	_, err = c.doRPC(ctx, &res, "EXPERIMENTAL_protocol_config", block, map[string]interface{}{})

	return
}
