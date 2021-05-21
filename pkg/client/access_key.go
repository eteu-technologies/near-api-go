package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
)

// https://docs.near.org/docs/api/rpc#view-access-key
func (c *Client) AccessKeyView(ctx context.Context, accountID types.AccountID, publicKey key.Base58PublicKey, block block.BlockCharacteristic) (resp AccessKeyView, err error) {
	_, err = c.doRPC(ctx, &resp, "query", block, map[string]interface{}{
		"request_type": "view_access_key",
		"account_id":   accountID,
		"public_key":   publicKey,
	})

	return
}

// https://docs.near.org/docs/api/rpc#view-access-key-list
func (c *Client) AccessKeyViewList(ctx context.Context, accountID types.AccountID, block block.BlockCharacteristic) (resp AccessKeyList, err error) {
	_, err = c.doRPC(ctx, &resp, "query", block, map[string]interface{}{
		"request_type": "view_access_key_list",
		"account_id":   accountID,
	})

	return
}

// TODO: decode response
// https://docs.near.org/docs/api/rpc#view-access-key-changes-single
func (c *Client) AccessKeyViewChanges(ctx context.Context, accountID types.AccountID, publicKey key.Base58PublicKey, block block.BlockCharacteristic) (res jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "single_access_key_changes",
		"keys": map[string]interface{}{
			"account_id": accountID,
			"public_key": publicKey,
		},
	})

	return
}

// TODO: decode response
// https://docs.near.org/docs/api/rpc#view-access-key-changes-all
func (c *Client) AccessKeyViewChangesAll(ctx context.Context, accountIDs []types.AccountID, block block.BlockCharacteristic) (res jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "all_access_key_changes",
		"account_ids":  accountIDs,
	})

	return
}
