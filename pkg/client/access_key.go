package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#view-access-key
func (c *Client) AccessKeyView(ctx context.Context, accountID types.AccountID, publicKey key.Base58PublicKey, block block.BlockCharacteristic) (resp AccessKeyView, err error) {
	var res jsonrpc.JSONRPCResponse
	res, err = c.doRPC(ctx, "query", block, map[string]interface{}{
		"request_type": "view_access_key",
		"account_id":   accountID,
		"public_key":   publicKey,
	})
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf("%s", string(*res.Error))
		return
	}

	if err = json.Unmarshal(res.Result, &resp); err != nil {
		return
	}

	return
}

// https://docs.near.org/docs/develop/front-end/rpc#view-access-key-list
func (c *Client) AccessKeyViewList(ctx context.Context, accountID types.AccountID, block block.BlockCharacteristic) (resp AccessKeyList, err error) {
	var res jsonrpc.JSONRPCResponse
	res, err = c.doRPC(ctx, "query", block, map[string]interface{}{
		"request_type": "view_access_key_list",
		"account_id":   accountID,
	})
	if err != nil {
		return
	}

	if res.Error != nil {
		err = fmt.Errorf("%s", string(*res.Error))
		return
	}

	if err = json.Unmarshal(res.Result, &resp); err != nil {
		return
	}

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#view-access-key-changes-single
func (c *Client) AccessKeyViewChanges(ctx context.Context, accountID types.AccountID, publicKey key.Base58PublicKey, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
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
func (c *Client) AccessKeyViewChangesAll(ctx context.Context, accountIDs []types.AccountID, block block.BlockCharacteristic) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "all_access_key_changes",
		"account_ids":  accountIDs,
	})

	return
}
