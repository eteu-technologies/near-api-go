package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eteu-technologies/near-api-go/pkg/jsonrpc"
	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
)

// https://docs.near.org/docs/develop/front-end/rpc#send-transaction-async
func (c *Client) RPCTransactionSend(ctx context.Context, signedTxnBase64 string) (resp hash.CryptoHash, err error) {
	var res jsonrpc.JSONRPCResponse
	res, err = c.doRPC(ctx, "broadcast_tx_async", nil, []string{signedTxnBase64})

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

// https://docs.near.org/docs/develop/front-end/rpc#send-transaction-await
func (c *Client) RPCTransactionSendAwait(ctx context.Context, signedTxnBase64 string) (resp FinalExecutionOutcomeView, err error) {
	var res jsonrpc.JSONRPCResponse
	res, err = c.doRPC(ctx, "broadcast_tx_commit", nil, []string{signedTxnBase64})

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

// https://docs.near.org/docs/develop/front-end/rpc#transaction-status
func (c *Client) TransactionStatus(ctx context.Context, tx hash.CryptoHash, sender types.AccountID) (resp FinalExecutionOutcomeView, err error) {
	var res jsonrpc.JSONRPCResponse
	res, err = c.doRPC(ctx, "tx", nil, []string{
		tx.String(), sender,
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

// https://docs.near.org/docs/develop/front-end/rpc#transaction-status-with-receipts
func (c *Client) TransactionStatusWithReceipts(ctx context.Context, tx hash.CryptoHash, sender types.AccountID) (resp FinalExecutionOutcomeWithReceiptView, err error) {
	var res jsonrpc.JSONRPCResponse
	res, err = c.doRPC(ctx, "EXPERIMENTAL_tx_status", nil, []string{
		tx.String(), sender,
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
