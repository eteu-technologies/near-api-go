package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/eteu-technologies/near-api-go/pkg/client"
	"github.com/eteu-technologies/near-api-go/pkg/config"
	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/action"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
)

func main() {
	app := &cli.App{
		Name:  "transfer",
		Usage: "Transfer NEAR between accounts",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "from",
				Required: true,
				Usage:    "Sender account id",
			},
			&cli.StringFlag{
				Name:     "to",
				Aliases:  []string{"recipient"},
				Required: true,
				Usage:    "Recipient account id",
			},
			&cli.StringFlag{
				Name:     "amount",
				Required: true,
				Usage:    "Amount of NEAR to send",
			},
			&cli.StringFlag{
				Name:    "network",
				Usage:   "NEAR network",
				Value:   "testnet",
				EnvVars: []string{"NEAR_ENV"},
			},
		},
		Action: entrypoint,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func entrypoint(cctx *cli.Context) (err error) {
	networkID := cctx.String("network")
	senderID := cctx.String("from")
	recipientID := cctx.String("to")
	amountValue := cctx.String("amount")

	var amount types.Balance

	amount, err = types.BalanceFromString(amountValue)
	if err != nil {
		return fmt.Errorf("failed to parse amount '%s': %w", amountValue, err)
	}

	network, ok := config.Networks[networkID]
	if !ok {
		return fmt.Errorf("unknown network '%s'", networkID)
	}

	keyPair, err := resolveCredentials(networkID, senderID)
	if err != nil {
		return fmt.Errorf("failed to load private key: %w", err)
	}

	rpc, err := client.NewClient(network.NodeURL)
	if err != nil {
		return fmt.Errorf("failed to create rpc client: %w", err)
	}

	log.Printf("near network: %s", rpc.NetworkAddr())

	ctx := client.ContextWithKeyPair(context.Background(), keyPair)
	res, err := rpc.TransactionSendAwait(
		ctx, senderID, recipientID,
		[]action.Action{
			action.NewTransfer(amount),
		},
		client.WithLatestBlock(),
	)
	if err != nil {
		return fmt.Errorf("failed to do txn: %w", err)
	}

	log.Printf("tx url: %s/transactions/%s", network.ExplorerURL, res.Transaction.Hash)
	return
}

func resolveCredentials(networkName string, id types.AccountID) (kp key.KeyPair, err error) {
	var creds struct {
		AccountID  types.AccountID     `json:"account_id"`
		PublicKey  key.Base58PublicKey `json:"public_key"`
		PrivateKey key.KeyPair         `json:"private_key"`
	}

	var home string
	home, err = os.UserHomeDir()
	if err != nil {
		return
	}

	credsFile := filepath.Join(home, ".near-credentials", networkName, fmt.Sprintf("%s.json", id))

	var cf *os.File
	if cf, err = os.Open(credsFile); err != nil {
		return
	}
	defer cf.Close()

	if err = json.NewDecoder(cf).Decode(&creds); err != nil {
		return
	}

	if creds.PublicKey.String() != creds.PrivateKey.PublicKey.String() {
		err = fmt.Errorf("inconsistent public key, %s != %s", creds.PublicKey.String(), creds.PrivateKey.PublicKey.String())
		return
	}
	kp = creds.PrivateKey

	return
}
