package types

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"

	uint128 "github.com/eteu-technologies/golang-uint128"
	"github.com/shopspring/decimal"
)

var (
	tenPower24     = uint128.From64(uint64(math.Pow10(12))).Mul64(uint64(math.Pow10(12)))
	zeroNEAR       = Balance(uint128.From64(0))
	dTenPower24, _ = decimal.NewFromString(tenPower24.String())
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

func scaleToYocto(amount decimal.Decimal) (r *big.Int) {
	// Multiply base using the supplied float
	amount = amount.Mul(dTenPower24)

	// Convert it to big.Int
	return amount.BigInt()
}

func BalanceFromFloat(f float64) (bal Balance) {
	bal = Balance(uint128.FromBig(scaleToYocto(decimal.NewFromFloat(f))))
	return
}

func BalanceFromString(s string) (bal Balance, err error) {
	amount, err := decimal.NewFromString(s)
	if err != nil {
		return
	}

	bal = Balance(uint128.FromBig(scaleToYocto(amount)))
	return
}
