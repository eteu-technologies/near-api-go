package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/eteu-technologies/near-api-go/client"
	"github.com/eteu-technologies/near-api-go/types"
	"github.com/eteu-technologies/near-api-go/types/action"
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

	spew.Dump(res)
	fmt.Println()

	log.Printf("tx id: https://explorer.testnet.near.org/transactions/%s", res.Transaction.Hash)
}
