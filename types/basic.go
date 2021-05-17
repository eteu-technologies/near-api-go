package types

import (
	"encoding/json"
	"fmt"
	"math/big"

	"lukechampine.com/uint128"
)

// Account identifier. Provides access to user's state.
type AccountID string

// Gas is a type for storing amounts of gas.
type Gas uint64

type Balance uint128.Uint128

func (bal *Balance) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	val := big.Int{}
	if _, ok := val.SetString(s, 10); !ok {
		return fmt.Errorf("unable to parse '%s'", s)
	}

	*bal = Balance(uint128.FromBig(&val))

	return nil
}

func (bal Balance) String() string {
	return uint128.Uint128(bal).String()
}

// Nonce for transactions.
type Nonce uint64
