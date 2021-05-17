package types

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"

	"lukechampine.com/uint128"
)

var (
	tenPower24 = uint128.From64(uint64(math.Pow10(12))).Mul64(uint64(math.Pow10(12)))
)

// Balance holds amount of yoctoNEAR
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

func (bal Balance) MarshalJSON() ([]byte, error) {
	return json.Marshal(bal.String())
}

func (bal Balance) String() string {
	return uint128.Uint128(bal).String()
}

// TODO
func NEARToYocto(near uint64) uint128.Uint128 {
	return uint128.From64(near).Mul(tenPower24)
}

// TODO
func YoctoToNEAR(yocto uint128.Uint128) uint64 {
	div := yocto.Div(tenPower24)
	if h := div.Hi; h != 0 {
		panic(fmt.Errorf("yocto div failed, remaining: %d", h))
	}

	return div.Lo
}
