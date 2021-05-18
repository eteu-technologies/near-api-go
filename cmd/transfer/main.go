package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	borsh "github.com/eteu-technologies/borsh-go"

	"github.com/eteu-technologies/near-api-go/client"
	"github.com/eteu-technologies/near-api-go/client/block"
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
		Author types.AccountID `json:"author"`
		Header struct {
			Height                uint64           `json:"height"`
			EpochID               hash.CryptoHash  `json:"epoch_id"`
			NextEpochID           hash.CryptoHash  `json:"next_epoch_id"`
			Hash                  hash.CryptoHash  `json:"hash"`
			PrevHash              hash.CryptoHash  `json:"prev_hash"`
			PrevStateRoot         hash.CryptoHash  `json:"prev_state_root"`
			ChunkReceiptsRoot     hash.CryptoHash  `json:"chunk_receipts_root"`
			ChunkHeadersRoot      hash.CryptoHash  `json:"chunk_headers_root"`
			ChunkTxRoot           hash.CryptoHash  `json:"chunk_tx_root"`
			OutcomeRoot           hash.CryptoHash  `json:"outcome_root"`
			ChunksIncluded        uint64           `json:"chunks_included"`
			ChallengesRoot        hash.CryptoHash  `json:"challenges_root"`
			Timestamp             uint64           `json:"timestamp"`         // milliseconds
			TimestampNanosec      string           `json:"timestamp_nanosec"` // nanoseconds, uint128
			RandomValue           hash.CryptoHash  `json:"random_value"`
			ValidatorProposals    []ValidatorStake `json:"validator_proposals"`
			ChunkMask             []bool           `json:"chunk_mask"`
			GasPrice              types.Balance    `json:"gas_price"`
			RentPaid              types.Balance    `json:"rent_paid"`        // NOTE: deprecated - 2021-05-14
			ValidatorReward       types.Balance    `json:"validator_reward"` // NOTE: deprecated - 2021-05-14
			TotalSupply           types.Balance    `json:"total_supply"`
			ChallengesResult      ChallengesResult `json:"challenges_result"`
			LastFinalBlock        hash.CryptoHash  `json:"last_final_block"`
			LastDSFinalBlock      hash.CryptoHash  `json:"last_ds_final_block"`
			NextBPHash            hash.CryptoHash  `json:"next_bp_hash"`
			BlockMerkleRoot       hash.CryptoHash  `json:"block_merkle_root"`
			Approvals             []string         `json:"approvals"` // TODO: array of nullable ed25519 signatures
			Signature             string           `json:"signature"` // TODO: ed25519 signature
			LatestProtocolVersion uint64           `json:"latest_protocol_version"`
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

	// Create and sign the transaction
	signedTxn, err := transaction.NewSignedTransaction(keyPair, transaction.Transaction{
		SignerID:   accID,
		PublicKey:  keyPair.PublicKey.ToPublicKey(),
		Nonce:      nonce + 1,
		ReceiverID: targetAccID,
		BlockHash:  blockHash,
		Actions: []action.Action{
			action.NewTransfer(types.NEARToYocto(1).Div64(1000)),
		},
	})
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

	if err := json.Unmarshal(res.Result, &txnRes); err != nil {
		log.Fatal("failed to parse txn result: ", err)
	}

	spew.Dump(txnRes)
}
