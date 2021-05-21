package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/eteu-technologies/near-api-go/pkg/client"
	"github.com/eteu-technologies/near-api-go/pkg/config"
)

func main() {
	app := &cli.App{
		Name:  "genesis",
		Usage: "Gets genesis config for the network",
		Flags: []cli.Flag{
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
	network, ok := config.Networks[networkID]
	if !ok {
		return fmt.Errorf("unknown network '%s'", networkID)
	}

	rpc, err := client.NewClient(network.NodeURL)
	if err != nil {
		return fmt.Errorf("failed to create rpc client: %w", err)
	}

	genesisConfig, err := rpc.GenesisConfig(cctx.Context)
	if err != nil {
		return fmt.Errorf("failed to query genesis config: %w", err)
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	_ = encoder.Encode(genesisConfig)

	return
}
