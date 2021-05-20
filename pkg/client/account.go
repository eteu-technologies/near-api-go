package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
	"github.com/eteu-technologies/near-api-go/pkg/types"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#view-account
func (c *Client) AccountView(ctx context.Context, accountID types.AccountID, block block.BlockCharacteristic) (res jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "query", block, map[string]interface{}{
		"request_type": "view_account",
		"account_id":   accountID,
	})

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#view-account-changes
func (c *Client) AccountViewChanges(ctx context.Context, accountIDs []types.AccountID, block block.BlockCharacteristic) (res jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "account_changes",
		"account_ids":  accountIDs,
	})

	return
}
