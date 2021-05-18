package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/eteu-technologies/borsh-go"
	"github.com/eteu-technologies/near-api-go/client"
	"github.com/eteu-technologies/near-api-go/client/block"
	"github.com/eteu-technologies/near-api-go/types"
	"github.com/eteu-technologies/near-api-go/types/action"
	"github.com/eteu-technologies/near-api-go/types/hash"
	"github.com/eteu-technologies/near-api-go/types/key"
	"github.com/eteu-technologies/near-api-go/types/transaction"
)

var (
	accID       = "mikroskeem.testnet"
	secretKey   = os.Getenv("NEAR_PRIV_KEY")
	targetAccID = "dev-1621263077598-74843909627468"
)

func main() {
	keyPair, err := key.NewBase58KeyPair(secretKey)
	if err != nil {
		log.Fatal("failed to load private key: ", err)
	}

	addr := "https://rpc.testnet.near.org"

	rpc, err := client.NewClient(addr)
	if err != nil {
		log.Fatal("failed to create rpc client: ", err)
	}

	log.Printf("near network: %s", rpc.NetworkAddr())

	ctx := context.Background()

	// Query this key
	accessKeyView, err := rpc.AccessKeyView(ctx, accID, keyPair.PublicKey, block.FinalityFinal())
	if err != nil {
		log.Fatal("failed to query access key: ", err)
	}

	nonce := accessKeyView.Nonce

	// Query latest block
	blockRes, err := rpc.BlockDetails(ctx, block.FinalityFinal())
	if err != nil {
		log.Fatal("failed to query latest block: ", err)
	}

	type SlashedValidator struct {
		AccountID    types.AccountID `json:"account_id"`
		IsDoubleSign bool            `json:"is_double_sign"`
	}

	type ChallengesResult []SlashedValidator

	// ValidatorStake is based on ValidatorStakeV1 struct in nearcore
	type ValidatorStake struct {
		AccountID types.AccountID `json:"account_id"`
		PublicKey key.PublicKey   `json:"public_key"`
	}

	var blockDetails struct {
		Header struct {
			Hash hash.CryptoHash `json:"hash"`
		} `json:"header"`
	}

	if blockRes.Error != nil {
		log.Fatal("failed to query latest block: ", string(*blockRes.Error))
	} else if err := json.Unmarshal(blockRes.Result, &blockDetails); err != nil {
		log.Fatal("failed to deserialize block details response: ", err)
	}

	blockHash := blockDetails.Header.Hash
	log.Println("latest block hash: ", blockHash)

	// Create a transaction
	txn := transaction.Transaction{
		SignerID:   accID,
		PublicKey:  keyPair.PublicKey.ToPublicKey(),
		Nonce:      nonce + 1,
		ReceiverID: targetAccID,
		BlockHash:  blockHash,
		Actions: []action.Action{
			action.NewFunctionCall("increment", nil, types.DefaultFunctionCallGas, types.NEARToYocto(0)),
		},
	}

	// Sign the transaction
	signedTxn, err := transaction.NewSignedTransaction(keyPair, txn)
	if err != nil {
		log.Fatal("failed to create signed txn: ", err)
	}

	stxnSerialized, err := borsh.Serialize(signedTxn)
	if err != nil {
		log.Fatal("failed to serialize txn: ", err)
	}

	// Try to verify the signature
	if ok, err := signedTxn.Verify(keyPair.PublicKey.ToPublicKey()); err != nil {
		log.Fatalf("failed to verify payload: %s", err)
	} else if !ok {
		log.Fatalf("failed to verify payload: %s", "invalid signature")
	}

	// Try to parse txn
	var txn2 transaction.Transaction
	if err := borsh.Deserialize(&txn2, signedTxn.SerializedTransaction); err != nil {
		log.Fatal("failed to deserialize txn: ", err)
	}

	stxnBlob := base64.StdEncoding.EncodeToString(stxnSerialized)
	res, err := rpc.RPCTransactionSendAwait(ctx, stxnBlob)
	if err != nil {
		log.Fatal("failed to do txn: ", err)
	}

	spew.Dump(res)
}
