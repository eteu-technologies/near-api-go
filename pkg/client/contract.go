package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
	"github.com/eteu-technologies/near-api-go/pkg/types"
)

// https://docs.near.org/docs/api/rpc#view-contract-state
func (c *Client) ContractViewState(ctx context.Context, accountID types.AccountID, prefixBase64 string, block block.BlockCharacteristic) (res ViewStateResult, err error) {
	_, err = c.doRPC(ctx, &res, "query", block, map[string]interface{}{
		"request_type":  "view_state",
		"account_id":    accountID,
		"prefix_base64": prefixBase64,
	})

	return
}

// TODO: decode response
// https://docs.near.org/docs/api/rpc#view-contract-state-changes
func (c *Client) ContractViewStateChanges(ctx context.Context, accountIDs []types.AccountID, keyPrefixBase64 string, block block.BlockCharacteristic) (res jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type":      "data_changes",
		"account_ids":       accountIDs,
		"key_prefix_base64": keyPrefixBase64,
	})

	return
}

// TODO: decode response
// https://docs.near.org/docs/api/rpc#view-contract-code-changes
func (c *Client) ContractViewCodeChanges(ctx context.Context, accountIDs []types.AccountID, block block.BlockCharacteristic) (res jsonrpc.Response, err error) {
	res, err = c.doRPC(ctx, nil, "EXPERIMENTAL_changes", block, map[string]interface{}{
		"changes_type": "contract_code_changes",
		"account_ids":  accountIDs,
	})

	return
}

// https://docs.near.org/docs/api/rpc#call-a-contract-function
func (c *Client) ContractViewCallFunction(ctx context.Context, accountID, methodName, argsBase64 string, block block.BlockCharacteristic) (res CallResult, err error) {
	_, err = c.doRPC(ctx, &res, "query", block, map[string]interface{}{
		"request_type": "call_function",
		"account_id":   accountID,
		"method_name":  methodName,
		"args_base64":  argsBase64,
	})

	return
}
