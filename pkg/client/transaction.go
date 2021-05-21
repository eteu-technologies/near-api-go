package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
)

// https://docs.near.org/docs/api/rpc#send-transaction-async
func (c *Client) RPCTransactionSend(ctx context.Context, signedTxnBase64 string) (resp hash.CryptoHash, err error) {
	_, err = c.doRPC(ctx, &resp, "broadcast_tx_async", nil, []string{signedTxnBase64})

	return
}

// https://docs.near.org/docs/api/rpc#send-transaction-await
func (c *Client) RPCTransactionSendAwait(ctx context.Context, signedTxnBase64 string) (resp FinalExecutionOutcomeView, err error) {
	_, err = c.doRPC(ctx, &resp, "broadcast_tx_commit", nil, []string{signedTxnBase64})

	return
}

// https://docs.near.org/docs/api/rpc#transaction-status
func (c *Client) TransactionStatus(ctx context.Context, tx hash.CryptoHash, sender types.AccountID) (resp FinalExecutionOutcomeView, err error) {
	_, err = c.doRPC(ctx, &resp, "tx", nil, []string{
		tx.String(), sender,
	})

	return
}

// https://docs.near.org/docs/api/rpc#transaction-status-with-receipts
func (c *Client) TransactionStatusWithReceipts(ctx context.Context, tx hash.CryptoHash, sender types.AccountID) (resp FinalExecutionOutcomeWithReceiptView, err error) {
	_, err = c.doRPC(ctx, &resp, "EXPERIMENTAL_tx_status", nil, []string{
		tx.String(), sender,
	})

	return
}
