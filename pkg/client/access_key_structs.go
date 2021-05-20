package client

import (
	"encoding/json"
	"fmt"

	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/hash"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
)

type AccessKey struct {
	Nonce         types.Nonce     `json:"nonce"`
	RawPermission json.RawMessage `json:"permission"`

	// Permission holds parsed access key permission info
	Permission AccessKeyPermission `json:"-"`
}

func (ak *AccessKey) UnmarshalJSON(b []byte) (err error) {
	// Unmarshal into inline struct to avoid recursion
	var tmp struct {
		Nonce         types.Nonce     `json:"nonce"`
		RawPermission json.RawMessage `json:"permission"`
	}
	if err = json.Unmarshal(b, &tmp); err != nil {
		return
	}

	*ak = AccessKey{
		Nonce:         tmp.Nonce,
		RawPermission: tmp.RawPermission,
	}
	err = ak.Permission.UnmarshalJSON(tmp.RawPermission)

	return
}

// AccessKeyPermission holds info whether access key is a FullAccess, or FunctionCall key
type AccessKeyPermission struct {
	FullAccess   bool                   `json:"-"`
	FunctionCall FunctionCallPermission `json:"-"`
}

func (akp *AccessKeyPermission) UnmarshalJSON(b []byte) (err error) {
	*akp = AccessKeyPermission{}

	// Option 1: "FullAccess"
	var s string
	if err = json.Unmarshal(b, &s); err == nil {
		switch s {
		case "FullAccess":
			akp.FullAccess = true
			return
		default:
			return fmt.Errorf("'%s' is neither object or 'FullAccess'", s)
		}
	} else if jerr, ok := err.(*json.UnmarshalTypeError); ok && jerr.Value != "object" {
		// If trying to unmarshal object into string, then continue. Otherwise return here
		return
	}

	// Option 2: Function call
	var perm struct {
		FunctionCall FunctionCallPermission `json:"FunctionCall"`
	}
	println("parsing object: " + string(b))
	err = json.Unmarshal(b, &perm)
	akp.FunctionCall = perm.FunctionCall

	return
}

// FunctionCallPermission represents a function call permission
type FunctionCallPermission struct {
	// Allowance for this function call (default 0.25 NEAR). Can be absent.
	Allowance *types.Balance `json:"allowance"`
	// ReceiverID holds the contract the key is allowed to call methods on
	ReceiverID types.AccountID `json:"receiver_id"`
	// MethodNames hold which functions are allowed to call. Can be empty (all methods are allowed)
	MethodNames []string `json:"method_names"`
}

type AccessKeyView struct {
	AccessKey
	BlockHeight uint64          `json:"block_height"`
	BlockHash   hash.CryptoHash `json:"block_hash"`
}

type AccessKeyViewInfo struct {
	PublicKey key.Base58PublicKey `json:"public_key"`
	AccessKey AccessKey           `json:"access_key"`
}

type AccessKeyList struct {
	Keys []AccessKeyViewInfo `json:"keys"`
}
