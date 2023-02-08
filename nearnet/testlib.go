package nearnet

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/eteu-technologies/near-api-go/pkg/client"
	"github.com/eteu-technologies/near-api-go/pkg/config"
	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/action"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
)

var (
	chainAccountFn client.TransactionOpt
)

func main() {
	if err := entrypoint(); err != nil {
		panic(err)
	}
}

func entrypoint() (err error) {
	ctx := context.Background()

	var rpc client.Client
	cfg := config.Networks["testnet"].NodeURL
	if rpc, err = client.NewClient(cfg); err != nil {
		return
	}

	// TODO: load key from node directory
	var chainKeyPair key.KeyPair
	if chainKeyPair, err = key.GenerateKeyPair(key.KeyTypeED25519, rand.Reader); err != nil {
		return
	}

	chainAccountFn = client.WithKeyPair(chainKeyPair)

	// TODO: generate random valid name for users
	accountA := "integtest-a-.near"
	accountB := "integtest-b-.near"

	var runnables []Runnable = []Runnable{
		CreateAccount("", accountA, types.NEARToYocto(10)),
		CreateAccount(accountA, accountB, types.NEARToYocto(1)),
		AssertThat("accountB has less NEAR than accountA", func(ctx context.Context, rpc client.Client) (err error) {
			// TODO!
			return errors.New("not implemented")
		}),
	}

	_ = rpc
	_ = ctx
	_ = runnables

	return
}

type Runnable func(ctx context.Context, rpc client.Client) error

func CreateAccount(creator, name types.AccountID, balance types.Balance) Runnable {
	return	func(ctx context.Context, rpc client.Client) (err error) {
		_, err = rpc.TransactionSendAwait(ctx, creator, name, []action.Action{
			action.NewTransfer(balance),
		}, chainAccountFn, client.WithLatestBlock())
		return
		}
}

func AssertThat(what string, fn Runnable /* reusing type because lazy */) Runnable {
	return func(ctx context.Context, rpc client.Client) (err error) {
		if err = fn(ctx, rpc); err != nil {
			return fmt.Errorf("Assertion %s failed! %w", what, err)
		}
		return
	}
}
