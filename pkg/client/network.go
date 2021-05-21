package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
)

// https://docs.near.org/docs/api/rpc#network-info
func (c *Client) NetworkInfo(ctx context.Context) (res NetworkInfo, err error) {
	_, err = c.doRPC(ctx, &res, "network_info", nil, []string{})

	return
}

// https://docs.near.org/docs/api/rpc#general-validator-status
func (c *Client) NetworkStatusValidators(ctx context.Context) (res StatusResponse, err error) {
	_, err = c.doRPC(ctx, &res, "status", nil, []string{})

	return
}

// https://docs.near.org/docs/api/rpc#detailed-validator-status
func (c *Client) NetworkStatusValidatorsDetailed(ctx context.Context, block block.BlockCharacteristic) (res ValidatorsResponse, err error) {
	_, err = c.doRPC(ctx, nil, "validators", nil, blockIDArrayParams(block))

	return
}
