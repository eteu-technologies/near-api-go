package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/types/key"
)

type rpcContext int

const (
	clientCtx = rpcContext(iota)
	keyPairCtx
)

func ContextWithKeyPair(ctx context.Context, keyPair key.KeyPair) context.Context {
	kp := keyPair
	return context.WithValue(ctx, keyPairCtx, &kp)
}

func getKeyPair(ctx context.Context) *key.KeyPair {
	v, ok := ctx.Value(keyPairCtx).(*key.KeyPair)
	if ok {
		return v
	}

	return nil
}
