package block

import (
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
)

type BlockCharacteristic func(map[string]interface{})

// Uses the latest block recorded on the node that responded to your query (<1 second delay after the transaction is submitted)
func FinalityOptimistic() BlockCharacteristic {
	return func(params map[string]interface{}) {
		params["finality"] = "optimistic"
	}
}

// Block that has been validated on at least 66% of the nodes in the network (usually takes 2 blocks / approx. 2 second delay)
func FinalityFinal() BlockCharacteristic {
	return func(params map[string]interface{}) {
		params["finality"] = "final"
	}
}

// Block height
func BlockID(blockID uint) BlockCharacteristic {
	return func(params map[string]interface{}) {
		params["block_id"] = blockID
	}
}

// Block hash
func BlockHash(blockHash hash.CryptoHash) BlockCharacteristic {
	return func(params map[string]interface{}) {
		params["block_id"] = blockHash
	}
}

// Variant of `BlockHash`, but accepting a raw block hash
func BlockHashRaw(blockHash string) BlockCharacteristic {
	return func(params map[string]interface{}) {
		params["block_id"] = blockHash
	}
}
