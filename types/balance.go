package types

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"

	"github.com/eteu-technologies/golang-uint128"
)

var (
	tenPower24 = uint128.From64(uint64(math.Pow10(12))).Mul64(uint64(math.Pow10(12)))
	zeroNEAR   = Balance(uint128.From64(0))
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

// Convenience funcs
func (bal Balance) Div64(div uint64) Balance {
	return Balance(uint128.Uint128(bal).Div64(div))
}

// TODO
func NEARToYocto(near uint64) Balance {
	if near == 0 {
		return zeroNEAR
	}

	return Balance(uint128.From64(near).Mul(tenPower24))
}

// TODO
func YoctoToNEAR(yocto Balance) uint64 {
	div := uint128.Uint128(yocto).Div(tenPower24)
	if h := div.Hi; h != 0 {
		panic(fmt.Errorf("yocto div failed, remaining: %d", h))
	}

	return div.Lo
}
