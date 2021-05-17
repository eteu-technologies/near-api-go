package main

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/eteu-technologies/borsh-go"
	"github.com/eteu-technologies/near-api-go/types"
	"github.com/eteu-technologies/near-api-go/types/action"
	"github.com/eteu-technologies/near-api-go/types/hash"
	"github.com/eteu-technologies/near-api-go/types/key"
	"github.com/eteu-technologies/near-api-go/types/transaction"
	nearrpc "github.com/eteu-technologies/near-rpc-go"
	oldkey "github.com/eteu-technologies/near-rpc-go/key"
	"github.com/eteu-technologies/near-rpc-go/shim"
	"github.com/mr-tron/base58"
)

var (
	accID       = "mikroskeem.testnet"
	secretKey   = os.Getenv("NEAR_PRIV_KEY")
	targetAccID = "dev-1621263077598-74843909627468"
)

func loadPrivKey(key string) (ed25519.PrivateKey, ed25519.PublicKey, error) {
	split := strings.SplitN(key, ":", 2)
	if k := split[0]; k != "ed25519" {
		return nil, nil, fmt.Errorf("unsupported key %s", k)
	}

	seed, err := base58.FastBase58Decoding(split[1])
	if err != nil {
		return nil, nil, err
	}

	priv := ed25519.PrivateKey(seed)
	return priv, ed25519.PublicKey(priv[32:]), nil
}

func main() {
	privKey, pubKey, err := loadPrivKey(secretKey)
	if err != nil {
		log.Fatal("failed to load private key: ", err)
	}

	thePubK := oldkey.WrapED25519PubKey(pubKey)

	addr := "https://rpc.testnet.near.org"

	shim.ShimURL = addr

	rpc, err := nearrpc.NewClient(addr)
	if err != nil {
		log.Fatal("failed to create rpc client: ", err)
	}

	log.Printf("near network: %s", rpc.NetworkAddr)

	ctx := context.Background()

	// Query this key
	accessKey, err := rpc.AccessKeyView(ctx, nearrpc.FinalityFinal(), accID, thePubK)
	if err != nil {
		log.Fatal("failed to query access key: ", err)
	}

	nonce := accessKey.Nonce

	// Query latest block
	blockRes, err := rpc.BlockDetails(ctx, nearrpc.FinalityFinal())
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
		PublicKey:  key.WrapED25519(pubKey),
		Nonce:      nonce + 1,
		ReceiverID: targetAccID,
		BlockHash:  blockHash,
		Actions: []action.Action{
			action.NewFunctionCall("increment", []byte(`{"some_complex_args":{"yup":123.456}}`), types.DefaultFunctionCallGas, types.NEARToYocto(0)),
		},
	}

	// Sign the transaction
	signedTxn, err := transaction.NewSignedTransaction(txn, privKey)
	if err != nil {
		log.Fatal("failed to create signed txn: ", err)
	}

	stxnSerialized, err := borsh.Serialize(signedTxn)
	if err != nil {
		log.Fatal("failed to serialize txn: ", err)
	}

	// Try to verify the signature
	sigHash := signedTxn.Hash()
	if !ed25519.Verify(pubKey, sigHash[:], signedTxn.Signature[1:]) {
		log.Fatal("failed to verify payload")
	}

	// Try to parse txn
	var txn2 transaction.Transaction
	if err := borsh.Deserialize(&txn2, signedTxn.SerializedTransaction); err != nil {
		log.Fatal("failed to deserialize txn: ", err)
	}

	stxnBlob := base64.StdEncoding.EncodeToString(stxnSerialized)
	res, err := rpc.TransactionSendAwait(ctx, stxnBlob)
	if err != nil {
		log.Fatal("failed to do txn: ", err)
	}

	spew.Dump(res)
}
