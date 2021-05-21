package client

import "github.com/eteu-technologies/near-api-go/pkg/types"

type GasPrice struct {
	GasPrice types.Balance `json:"gas_price"`
}
