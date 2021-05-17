package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/client/block"
	"github.com/eteu-technologies/near-api-go/jsonrpc"
)

type Client struct {
	RPCClient jsonrpc.JSONRPCClient
}

func NewClient(networkAddr string) (client Client, err error) {
	client.RPCClient, err = jsonrpc.NewClient(networkAddr)
	if err != nil {
		return
	}

	return
}

func (c *Client) doRPC(ctx context.Context, method string, block block.BlockCharacteristic, params interface{}) (res jsonrpc.JSONRPCResponse, err error) {
	if block != nil {
		if mapv, ok := params.(map[string]interface{}); ok {
			block(mapv)
		}
	}

	res, err = c.RPCClient.CallRPC(ctx, method, params)
	return
}
