package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/client/block"
	"github.com/eteu-technologies/near-api-go/jsonrpc"
	"github.com/eteu-technologies/near-api-go/types/key"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#view-access-key
func (c *Client) AccessKeyView(ctx context.Context, accountID string, publicKey key.Base58PublicKey, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "query", block, map[string]interface{}{
		"request_type": "view_access_key",
		"account_id":   accountID,
		"public_key":   publicKey,
	})

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#view-access-key-list
func (c *Client) AccessKeyViewList(ctx context.Context, accountID string, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "query", block, map[string]interface{}{
		"request_type": "view_access_key_list",
		"account_id":   accountID,
	})

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#view-access-key-changes-single
func (c *Client) AccessKeyViewChanges(ctx context.Context, accountID string, publicKey key.Base58PublicKey, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "single_access_key_changes",
		"keys": map[string]interface{}{
			"account_id": accountID,
			"public_key": publicKey,
		},
	})

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#view-access-key-changes-all
func (c *Client) AccessKeyViewChangesAll(ctx context.Context, accountIDs []string, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "all_access_key_changes",
		"account_ids":  accountIDs,
	})

	return
}
