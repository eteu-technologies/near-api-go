package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"

	"github.com/eteu-technologies/near-api-go/pkg/client"
	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/config"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
)

func main() {
	app := &cli.App{
		Name:  "keys",
		Usage: "Display access keys attached to an account",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "account",
				Required: true,
				Usage:    "Account id",
			},
			&cli.StringFlag{
				Name:  "key",
				Usage: "Specific key to query. Otherwise shows all access keys",
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

	ctx := context.Background()
	if rawKey := cctx.String("key"); cctx.IsSet("key") {
		pubKey, err := key.NewBase58PublicKey(rawKey)
		if err != nil {
			return fmt.Errorf("failed to parse access pubkey: %w", err)
		}

		accessKeyViewResp, err := rpc.AccessKeyView(ctx, cctx.String("account"), pubKey, block.FinalityFinal())
		if err != nil {
			return fmt.Errorf("failed to query access key: %w", err)
		}

		spew.Dump(accessKeyViewResp)
	} else {
		accessKeyViewListResp, err := rpc.AccessKeyViewList(ctx, cctx.String("account"), block.FinalityFinal())
		if err != nil {
			return fmt.Errorf("failed to query access key list: %w", err)
		}

		spew.Dump(accessKeyViewListResp)
	}

	return
}
