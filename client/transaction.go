package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/jsonrpc"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#send-transaction-async
func (c *Client) TransactionSend(ctx context.Context, signedTxnBase64 string) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "broadcast_tx_async", nil, []string{signedTxnBase64})

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#send-transaction-await
func (c *Client) TransactionSendAwait(ctx context.Context, signedTxnBase64 string) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "broadcast_tx_commit", nil, []string{signedTxnBase64})

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#transaction-status
func (c *Client) TransactionStatus(ctx context.Context, txHash string, senderAccountID string) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "tx", nil, []string{
		txHash, senderAccountID,
	})

	return
}

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#transaction-status-with-receipts
func (c *Client) TransactionStatusWithReceipts(ctx context.Context, txHash string, senderAccountID string) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "EXPERIMENTAL_tx_status", nil, []string{
		txHash, senderAccountID,
	})

	return
}
