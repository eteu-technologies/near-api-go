package types

import (
	"fmt"
	"math"

	"lukechampine.com/uint128"
)

var (
	tenPower24 = uint128.From64(uint64(math.Pow10(12))).Mul64(uint64(math.Pow10(12)))
)

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
