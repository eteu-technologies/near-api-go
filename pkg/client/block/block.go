package block

import (
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
)

// BlockCharacteristic is a function type aiding with specifying a block
type BlockCharacteristic func(map[string]interface{})

// FinalityOptimistic specifies the latest block recorded on the node that responded to your query (<1 second delay after the transaction is submitted)
func FinalityOptimistic() BlockCharacteristic {
	return func(params map[string]interface{}) {
		params["finality"] = "optimistic"
	}
}

// FinalityFinal specifies a block that has been validated on at least 66% of the nodes in the network (usually takes 2 blocks / approx. 2 second delay)
func FinalityFinal() BlockCharacteristic {
	return func(params map[string]interface{}) {
		params["finality"] = "final"
	}
}

// BlockID specifies a block id/height
func BlockID(blockID uint) BlockCharacteristic {
	return func(params map[string]interface{}) {
		params["block_id"] = blockID
	}
}

// BlockHash specifies a block hash
func BlockHash(blockHash hash.CryptoHash) BlockCharacteristic {
	return func(params map[string]interface{}) {
		params["block_id"] = blockHash
	}
}

// BlockHashRaw is a variant of `BlockHash` function accepting a raw block hash (string)
func BlockHashRaw(blockHash string) BlockCharacteristic {
	return func(params map[string]interface{}) {
		params["block_id"] = blockHash
	}
}
