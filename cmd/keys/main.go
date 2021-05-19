package main

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/eteu-technologies/near-api-go/pkg/client"
	"github.com/eteu-technologies/near-api-go/pkg/client/block"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
)

var (
	accID     = "mllnd.testnet"
	accessKey = "ed25519:nryP2kFtavumu57ThhRP6Qcg5Vy4i52H3Wp2qy4vPY8"
)

func main() {
	addr := "https://rpc.testnet.near.org"

	rpc, err := client.NewClient(addr)
	if err != nil {
		log.Fatal("failed to create rpc client: ", err)
	}

	log.Printf("near network: %s", rpc.NetworkAddr())

	ctx := context.Background()

	pubKey, err := key.NewBase58PublicKey(accessKey)
	if err != nil {
		log.Fatal("failed to parse access pubkey")
	}

	accessKeyViewResp, err := rpc.AccessKeyView(ctx, accID, pubKey, block.FinalityFinal())
	if err != nil {
		log.Fatal("failed to query access key list: ", err)
	}

	spew.Dump(accessKeyViewResp)
}
