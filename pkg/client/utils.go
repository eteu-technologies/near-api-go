package client

import "github.com/eteu-technologies/near-api-go/pkg/client/block"

// HACK
func blockIDArrayParams(block block.BlockCharacteristic) []interface{} {
	params := []interface{}{nil}
	p := map[string]interface{}{}

	block(p)
	if v, ok := p["block_id"]; ok {
		params[0] = v
	}

	return params
}
