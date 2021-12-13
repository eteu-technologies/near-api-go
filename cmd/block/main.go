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
)

func main() {
	app := &cli.App{
		Name:  "block",
		Usage: "View latest or specified block info",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "network",
				Usage:   "NEAR network",
				Value:   "testnet",
				EnvVars: []string{"NEAR_ENV"},
			},
			&cli.StringFlag{
				Name: "block",
				Usage: "Block hash",
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
	network, ok := config.Networks[networkID]
	if !ok {
		return fmt.Errorf("unknown network '%s'", networkID)
	}

	rpc, err := client.NewClient(network.NodeURL)
	if err != nil {
		return fmt.Errorf("failed to create rpc client: %w", err)
	}

	characteristic := block.FinalityFinal()
	if v := cctx.String("block"); v != "" {
		characteristic = block.BlockHashRaw(v)
	}

	blockDetailsResp, err := rpc.BlockDetails(context.Background(), characteristic)
	if err != nil {
		return fmt.Errorf("failed to query latest block info: %w", err)
	}

	spew.Dump(blockDetailsResp)

	return
}
