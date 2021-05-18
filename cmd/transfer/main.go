package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/eteu-technologies/near-api-go/client"
	"github.com/eteu-technologies/near-api-go/types"
	"github.com/eteu-technologies/near-api-go/types/action"
	"github.com/eteu-technologies/near-api-go/types/hash"
	"github.com/eteu-technologies/near-api-go/types/key"
	"github.com/eteu-technologies/near-api-go/types/transaction"
)

var (
	// accID       = "node0"
	// secretKey   = "ed25519:3D4YudUQRE39Lc4JHghuB5WM8kbgDDa34mnrEP5DdTApVH81af7e2dWgNPEaiQfdJnZq1CNPp5im4Rg5b733oiMP"
	// targetAccID = "node1"
	accID       = "mikroskeem.testnet"
	secretKey   = os.Getenv("NEAR_PRIV_KEY")
	targetAccID = "mikroskeem2.testnet"
)

func main() {
	keyPair, err := key.NewBase58KeyPair(secretKey)
	if err != nil {
		log.Fatal("failed to load private key: ", err)
	}

	//addr := "http://127.0.0.1:3030"
	addr := "https://rpc.testnet.near.org"

	rpc, err := client.NewClient(addr)
	if err != nil {
		log.Fatal("failed to create rpc client: ", err)
	}

	log.Printf("near network: %s", rpc.NetworkAddr())

	ctx := client.ContextWithKeyPair(context.Background(), keyPair)
	txn := transaction.Transaction{
		SignerID:   accID,
		ReceiverID: targetAccID,
		Actions: []action.Action{
			action.NewTransfer(types.NEARToYocto(1).Div64(1000)),
		},
	}

	res, err := rpc.TransactionSendAwait(ctx, txn, client.WithLatestBlock())
	if err != nil {
		log.Fatal("failed to do txn: ", err)
	}

	type Status struct {
		SuccessValue     string          `json:"SuccessValue"`
		SuccessReceiptID string          `json:"SuccessReceiptId"`
		Failure          json.RawMessage `json:"Failure"` // TODO
	}

	type ExecutionOutcomeView struct {
		Logs        []string          `json:"logs"`
		ReceiptIDs  []hash.CryptoHash `json:"receipt_ids"`
		GasBurnt    types.Gas         `json:"gas_burnt"`
		TokensBurnt types.Balance     `json:"tokens_burnt"`
		ExecutorID  types.AccountID   `json:"executor_id"`
		Status      Status            `json:"status"`
	}

	type MerklePathItem struct {
		Hash      hash.CryptoHash `json:"hash"`
		Direction string          `json:"direction"` // TODO: enum type, either 'Left' or 'Right'
	}

	type MerklePath []MerklePathItem

	type ExecutionOutcomeWithIdView struct {
		Proof     MerklePath           `json:"proof"`
		BlockHash hash.CryptoHash      `json:"block_hash"`
		ID        hash.CryptoHash      `json:"id"`
		Outcome   ExecutionOutcomeView `json:"outcome"`
	}

	var txnRes struct {
		Status      Status `json:"status"`
		Transaction struct {
			SignerID   types.AccountID `json:"signer_id"`
			PublicKey  string          `json:"public_key"`
			Nonce      types.Nonce     `json:"nonce"`
			ReceiverID types.AccountID `json:"receiver_id"`
			Actions    []action.Action `json:"actions"`
			Signature  string          `json:"signature"`
			Hash       hash.CryptoHash `json:"hash"`
		} `json:"transaction"`
		TransactionOutcome ExecutionOutcomeWithIdView   `json:"transaction_outcome"`
		ReceiptsOutcome    []ExecutionOutcomeWithIdView `json:"receipts_outcome"`
	}

	if res.Error != nil {
		fmt.Println(string(*res.Error))
		return
	}

	if err := json.Unmarshal(res.Result, &txnRes); err != nil {
		log.Fatal("failed to parse txn result: ", err)
	}

	spew.Dump(txnRes)
}
