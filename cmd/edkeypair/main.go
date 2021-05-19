package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"

	"github.com/eteu-technologies/near-api-go/pkg/types/key"
)

func main() {
	keyPair, err := key.GenerateKeyPair(key.KeyTypeED25519, rand.Reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate keypair: %s\n", err)
		os.Exit(1)
	}

	pub := keyPair.PublicKey

	_ = json.NewEncoder(os.Stdout).Encode(struct {
		AccountID  string              `json:"account_id"`
		PublicKey  key.Base58PublicKey `json:"public_key"`
		PrivateKey string              `json:"private_key"`
	}{pub.ToPublicKey().Hash(), pub, keyPair.PrivateEncoded()})
}
