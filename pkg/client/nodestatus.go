package client

import (
	"context"
	"github.com/eteu-technologies/near-api-go/pkg/client/block"
)

// https://docs.near.org/docs/api/rpc#general-validator-status
func (c *Client) NodeStatusValidators(ctx context.Context) (res NodeStatus, err error) {
	_, err = c.doRPC(ctx, &res, "status", nil, []string{})

	return
}

// https://docs.near.org/docs/api/rpc#detailed-validator-status
func (c *Client) NodeStatusValidatorsDetailed(ctx context.Context, block block.BlockCharacteristic) (res ValidatorsResponse, err error) {
	_, err = c.doRPC(ctx, nil, "validators", nil, blockIDArrayParams(block))

	return
}
