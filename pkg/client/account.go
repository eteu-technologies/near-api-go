package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#view-account
func (c *Client) AccountView(ctx context.Context, accountID string, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "query", block, map[string]interface{}{
		"request_type": "view_account",
		"account_id":   accountID,
	})

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#view-account-changes
func (c *Client) AccountViewChanges(ctx context.Context, accountIDs []string, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "account_changes",
		"account_ids":  accountIDs,
	})

	return
}
