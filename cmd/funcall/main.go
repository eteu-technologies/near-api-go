package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/eteu-technologies/near-api-go/pkg/client"
	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/config"
	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/action"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
)

func main() {
	app := &cli.App{
		Name:  "funcall",
		Usage: "Calls function on a smart contract",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "account",
				Usage: "Account id",
			},
			&cli.StringFlag{
				Name:     "target",
				Required: true,
				Usage:    "Account id whose smart contract to call",
			},
			&cli.StringFlag{
				Name:  "mode",
				Usage: "Call mode, either 'view' or 'change'",
				Value: "view",
			},
			&cli.StringFlag{
				Name:  "deposit",
				Usage: "Amount of NEAR to deposit",
			},
			&cli.Uint64Flag{
				Name:  "gas",
				Usage: "Amount of gas to attach for this transaction",
				Value: types.DefaultFunctionCallGas,
			},
			&cli.StringFlag{
				Name:     "method",
				Usage:    "Method to call on specified contract",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "args",
				Usage: "Arguments to pass for specified method. Accepts both JSON and Base64 payload",
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
	network, ok := config.Networks[cctx.String("network")]
	if !ok {
		return fmt.Errorf("unknown network '%s'", cctx.String("network"))
	}

	rpc, err := client.NewClient(network.NodeURL)
	if err != nil {
		return fmt.Errorf("failed to create rpc client: %w", err)
	}

	log.Printf("near network: %s", rpc.NetworkAddr())

	var args []byte = nil
	if a := cctx.String("args"); cctx.IsSet("args") {
		args = []byte(a)
	}

	switch cctx.String("mode") {
	case "view":
		if err := viewFunction(cctx, rpc, args); err != nil {
			return fmt.Errorf("failed to call view function: %w", err)
		}
	case "change":
		if !cctx.IsSet("account") {
			return fmt.Errorf("--account is required for change function call")
		}

		keyPair, err := resolveCredentials(network.NetworkID, cctx.String("account"))
		if err != nil {
			return fmt.Errorf("failed to load private key: %w", err)
		}

		if err := changeFunction(cctx, rpc, keyPair, network, args); err != nil {
			return fmt.Errorf("failed to call change function: %w", err)
		}
	default:
		return fmt.Errorf("either 'change' or 'view' is accepted, you supplied '%s'", cctx.String("mode"))
	}

	return
}

func viewFunction(cctx *cli.Context, rpc client.Client, args []byte) (err error) {
	res, err := rpc.ContractViewCallFunction(cctx.Context, cctx.String("target"), cctx.String("method"), base64.StdEncoding.EncodeToString(args), block.FinalityFinal())
	if err != nil {
		return
	}

	if l := res.Logs; len(l) > 0 {
		log.Println("logs:")
		for _, line := range l {
			log.Printf("- %s", line)
		}
	}

	log.Println("result:")
	if len(res.Result) == 0 {
		fmt.Println("(empty)")
		return
	}

	fmt.Printf("%s", hex.Dump(res.Result))

	return
}

func changeFunction(cctx *cli.Context, rpc client.Client, keyPair key.KeyPair, network config.NetworkInfo, args []byte) (err error) {
	var deposit types.Balance = types.NEARToYocto(0)
	var gas types.Gas = cctx.Uint64("gas")

	if cctx.IsSet("deposit") {
		depositValue := cctx.String("deposit")
		deposit, err = types.BalanceFromString(depositValue)
		if err != nil {
			return fmt.Errorf("failed to parse amount '%s': %w", depositValue, err)
		}
	}

	// Make a transaction
	res, err := rpc.TransactionSendAwait(
		cctx.Context,
		cctx.String("account"),
		cctx.String("target"),
		[]action.Action{
			action.NewFunctionCall(cctx.String("method"), args, gas, deposit),
		},
		client.WithLatestBlock(),
		client.WithKeyPair(keyPair),
	)
	if err != nil {
		return fmt.Errorf("failed to do txn: %w", err)
	}

	// Try to get logs
	type LogEntry struct {
		Executor types.AccountID
		Lines    []string
	}
	logEntries := map[hash.CryptoHash]*LogEntry{}
	for _, receipt := range res.ReceiptsOutcome {
		if len(receipt.Outcome.Logs) == 0 {
			continue
		}

		entry, ok := logEntries[receipt.ID]
		if !ok {
			entry = &LogEntry{
				Executor: receipt.Outcome.ExecutorID,
				Lines:    []string{},
			}
			logEntries[receipt.ID] = entry
		}

		entry.Lines = append(entry.Lines, receipt.Outcome.Logs...)
	}

	if len(logEntries) > 0 {
		oneEntry := len(logEntries) == 1

		log.Println("logs:")
		for _, receipt := range res.ReceiptsOutcome {
			logEntry, ok := logEntries[receipt.ID]
			if !ok {
				continue
			}

			for _, line := range logEntry.Lines {
				if oneEntry {
					log.Printf("- %s", line)
				} else {
					log.Printf("- (%s / %s) %s", receipt.ID, logEntry.Executor, line)
				}
			}
		}
	}

	log.Printf("tx id: %s/transactions/%s", network.ExplorerURL, res.Transaction.Hash)
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
