package main

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/eteu-technologies/near-api-go/pkg/client"
	"github.com/eteu-technologies/near-api-go/pkg/client/block"
)

func main() {
	addr := "https://rpc.testnet.near.org"

	rpc, err := client.NewClient(addr)
	if err != nil {
		log.Fatal("failed to create rpc client: ", err)
	}

	log.Printf("near network: %s", rpc.NetworkAddr())

	ctx := context.Background()

	blockDetailsResp, err := rpc.BlockDetails(ctx, block.FinalityFinal())
	if err != nil {
		log.Fatal("failed to query block details: ", err)
	}

	spew.Dump(blockDetailsResp)
}
