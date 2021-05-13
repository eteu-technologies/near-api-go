package main

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	nearrpc "github.com/eteu-technologies/near-rpc-go"
	"github.com/eteu-technologies/near-rpc-go/key"
	"github.com/eteu-technologies/near-rpc-go/shim"
	"github.com/mr-tron/base58"
	borsh "github.com/near/borsh-go"

	"github.com/eteu-technologies/near-api-go/types"
	"github.com/eteu-technologies/near-api-go/types/action"
	"github.com/eteu-technologies/near-api-go/types/hash"
)

var (
	accID       = "node0"
	secretKey   = "ed25519:3D4YudUQRE39Lc4JHghuB5WM8kbgDDa34mnrEP5DdTApVH81af7e2dWgNPEaiQfdJnZq1CNPp5im4Rg5b733oiMP"
	targetAccID = "node1"
	// accID       = "mikroskeem.testnet"
	// secretKey   = os.Getenv("NEAR_PRIV_KEY")
	// targetAccID = "meeksorkim.testnet"
)

func mustKey(key *key.PublicKey, err error) *key.PublicKey {
	if err != nil {
		panic(err)
	}
	return key
}

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
	_ = os.Getenv

	privKey, pubKey, err := loadPrivKey(secretKey)
	if err != nil {
		log.Fatal("failed to load private key: ", err)
	}

	thePubK := key.WrapED25519PubKey(pubKey)

	addr := "http://127.0.0.1:3030"
	//addr := "https://rpc.testnet.near.org"

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

	var blockDetails struct {
		Author types.AccountID `json:"author"`
		Header struct {
			Height                uint64            `json:"height"`
			EpochID               hash.CryptoHash   `json:"epoch_id"`
			NextEpochID           hash.CryptoHash   `json:"next_epoch_id"`
			Hash                  hash.CryptoHash   `json:"hash"`
			PrevHash              hash.CryptoHash   `json:"prev_hash"`
			PrevStateRoot         hash.CryptoHash   `json:"prev_state_root"`
			ChunkReceiptsRoot     hash.CryptoHash   `json:"chunk_receipts_root"`
			ChunkHeadersRoot      hash.CryptoHash   `json:"chunk_headers_root"`
			ChunkTxRoot           hash.CryptoHash   `json:"chunk_tx_root"`
			OutcomeRoot           hash.CryptoHash   `json:"outcome_root"`
			ChunksIncluded        uint64            `json:"chunks_included"`
			ChallengesRoot        string            `json:"challenges_root"`
			Timestamp             uint64            `json:"timestamp"`         // milliseconds
			TimestampNanosec      json.RawMessage   `json:"timestamp_nanosec"` // nanoseconds, uint128
			RandomValue           string            `json:"random_value"`
			ValidatorProposals    []json.RawMessage `json:"validator_proposals"` // TODO: unknown type, check rust code
			ChunkMask             []bool            `json:"chunk_mask"`
			GasPrice              json.RawMessage   `json:"gas_price"`         // TODO: types.Gas deserializing from string
			RentPaid              json.RawMessage   `json:"rent_paid"`         // TODO: unknown type, check rust code
			ValidatorReward       json.RawMessage   `json:"validator_reward"`  // TODO: unknown type, check rust code
			TotalSupply           json.RawMessage   `json:"total_supply"`      // TODO: unknown type, check rust code
			ChallengesResult      []json.RawMessage `json:"challenges_result"` // TODO: unknown type, check rust code
			LastFinalBlock        hash.CryptoHash   `json:"last_final_block"`
			LastDSFinalBlock      hash.CryptoHash   `json:"last_ds_final_block"`
			NextBPHash            hash.CryptoHash   `json:"next_bp_hash"`
			BlockMerkleRoot       hash.CryptoHash   `json:"block_merkle_root"`
			Approvals             []string          `json:"approvals"` // TODO: array of nullable ed25519 signatures
			Signature             string            `json:"signature"` // TODO: ed25519 signature
			LatestProtocolVersion uint64            `json:"latest_protocol_version"`
		} `json:"header"`
		Chunks []struct {
			/*
			   "chunk_hash": "EBM2qg5cGr47EjMPtH88uvmXHDHqmWPzKaQadbWhdw22",
			   "prev_block_hash": "2yUTTubrv1gJhTUVnHXh66JG3qxStBqySoN6wzRzgdVD",
			   "outcome_root": "11111111111111111111111111111111",
			   "prev_state_root": "HqWDq3f5HJuWnsTfwZS6jdAUqDjGFSTvjhb846vV27dx",
			   "encoded_merkle_root": "9zYue7drR1rhfzEEoc4WUXzaYRnRNihvRoGt1BgK7Lkk",
			   "encoded_length": 8,
			   "height_created": 17821130,
			   "height_included": 17821130,
			   "shard_id": 0,
			   "gas_used": 0,
			   "gas_limit": 1000000000000000,
			   "rent_paid": "0",
			   "validator_reward": "0",
			   "balance_burnt": "0",
			   "outgoing_receipts_root": "H4Rd6SGeEBTbxkitsCdzfu9xL9HtZ2eHoPCQXUeZ6bW4",
			   "tx_root": "11111111111111111111111111111111",
			   "validator_proposals": [],
			   "signature": "ed25519:4iPgpYAcPztAvnRHjfpegN37Rd8dTJKCjSd1gKAPLDaLcHUySJHjexMSSfC5iJVy28vqF9VB4psz13x2nt92cbR7"
			*/
		} `json:"chunks"`
	}

	if blockRes.Error != nil {
		log.Fatal("failed to query latest block: ", string(*blockRes.Error))
	} else if err := json.Unmarshal(blockRes.Result, &blockDetails); err != nil {
		log.Fatal("failed to deserialize block details response: ", err)
	}

	blockHash := blockDetails.Header.Hash
	log.Println("latest block hash: ", blockHash)

	// Create a transaction
	bal := big.Int{}
	bal.SetUint64(uint64(10000000000000000000))

	txn := types.Transaction{
		SignerID:   accID,
		PublicKey:  types.PublicKeyFromED25519Key(pubKey),
		Nonce:      nonce + 1,
		ReceiverID: targetAccID,
		BlockHash:  blockHash,
		Actions: []action.Action{
			action.NewTransfer(bal),
		},
	}

	// Sign the transaction
	signedTxn, err := types.NewSignedTransaction(txn, privKey)
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
	var txn2 types.Transaction
	if err := borsh.Deserialize(&txn2, signedTxn.SerializedTransaction); err != nil {
		log.Fatal("failed to deserialize txn: ", err)
	}

	stxnBlob := base64.StdEncoding.EncodeToString(stxnSerialized)
	res, err := rpc.TransactionSendAwait(ctx, stxnBlob)
	if err != nil {
		log.Fatal("failed to do txn: ", err)
	}

	type Status struct {
		SuccessValue     string          `json:"SuccessValue"`
		SuccessReceiptID string          `json:"SuccessReceiptId"`
		Failure          json.RawMessage `json:"Failure"` // TODO
	}

	type Outcome struct {
		Logs        []string      `json:"logs"`        // TODO: verify type
		ReceiptIDs  []interface{} `json:"receipt_ids"` // TODO: unknown type
		GasBurnt    uint64        `json:"gas_burnt"`
		TokensBurnt string        `json:"tokens_burnt"` // TODO: u128
		ExecutorID  string        `json:"executor_id"`  // TODO: account id
		Status      Status        `json:"status"`
	}

	type Receipt struct {
		Proof     []interface{}   `json:"proof"` // TODO: unknown type
		BlockHash hash.CryptoHash `json:"block_hash"`
		ID        hash.CryptoHash `json:"id"`
		Outcome   Outcome         `json:"outcome"`
	}

	var txnRes struct {
		Status      Status `json:"status"`
		Transaction struct {
			SignerID   string                                             `json:"signer_id"`
			PublicKey  string                                             `json:"public_key"`
			Nonce      uint64                                             `json:"nonce"`
			ReceiverID string                                             `json:"receiver_id"`
			Actions    []json.RawMessage/*types.Action*/ `json:"actions"` // TODO: types.Action
			Signature  string                                             `json:"signature"`
			Hash       string                                             `json:"hash"`
		} `json:"transaction"`
		TransactionOutcome Receipt   `json:"transaction_outcome"`
		ReceiptsOutcome    []Receipt `json:"receipts_outcome"`
	}

	if err := json.Unmarshal(res.Result, &txnRes); err != nil {
		log.Fatal("failed to parse txn result: ", err)
	}

	spew.Dump(txnRes)
}
