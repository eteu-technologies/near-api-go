package action

import (
	"github.com/eteu-technologies/borsh-go"
	"lukechampine.com/uint128"

	"github.com/eteu-technologies/near-api-go/types"
)

type Action struct {
	Enum borsh.Enum `borsh_enum:"true"`

	CreateAccount  ActionCreateAccount
	DeployContract ActionDeployContract
	FunctionCall   ActionFunctionCall
	Transfer       ActionTransfer
	Stake          ActionStake
	AddKey         ActionAddKey
	DeleteKey      ActionDeleteKey
	DeleteAccount  ActionDeleteAccount
}

const (
	ordCreateAccount uint8 = iota
	ordDeployContract
	ordFunctionCall
	ordTransfer
	ordStake
	ordAddKey
	ordDeleteKey
	ordDeleteAccount
)

func (a *Action) PrepaidGas() types.Gas {
	switch uint8(a.Enum) {
	case ordFunctionCall:
		return a.FunctionCall.Gas
	default:
		return 0
	}
}

func (a *Action) DepositBalance() types.Balance {
	switch uint8(a.Enum) {
	case ordFunctionCall:
		return types.Balance(a.FunctionCall.Deposit)
	case ordTransfer:
		return types.Balance(a.Transfer.Deposit)
	default:
		return types.Balance(uint128.Zero)
	}
}

func (a *Action) UnderlyingValue() interface{} {
	switch uint8(a.Enum) {
	case ordCreateAccount:
		return a.CreateAccount
	case ordDeployContract:
		return a.DeployContract
	case ordFunctionCall:
		return a.FunctionCall
	case ordTransfer:
		return a.Transfer
	case ordStake:
		return a.Stake
	case ordAddKey:
		return a.AddKey
	case ordDeleteKey:
		return a.DeleteKey
	case ordDeleteAccount:
		return a.DeleteAccount
	}

	panic("unreachable")
}

type ActionCreateAccount struct {
	// TODO
}

// TODO
func NewCreateAccount() Action {
	return Action{
		Enum:          borsh.Enum(ordCreateAccount),
		CreateAccount: ActionCreateAccount{},
	}
}

type ActionDeployContract struct {
	// TODO
}

// TODO
func NewDeployContract() Action {
	return Action{
		Enum:           borsh.Enum(ordDeployContract),
		DeployContract: ActionDeployContract{},
	}
}

type ActionFunctionCall struct {
	MethodName string
	Args       string // Base64 string
	Gas        types.Gas
	Deposit    uint128.Uint128
}

func NewFunctionCall(methodName string, args string, gas types.Gas, deposit uint128.Uint128) Action {
	return Action{
		Enum: borsh.Enum(ordFunctionCall),
		FunctionCall: ActionFunctionCall{
			MethodName: methodName,
			Args:       args,
			Gas:        gas,
			Deposit:    deposit,
		},
	}
}

type ActionTransfer struct {
	Deposit uint128.Uint128
}

func NewTransfer(deposit uint128.Uint128) Action {
	return Action{
		Enum: borsh.Enum(ordTransfer),
		Transfer: ActionTransfer{
			Deposit: deposit,
		},
	}
}

type ActionStake struct {
	// TODO
}

// TODO
func NewStake() Action {
	return Action{
		Enum:  borsh.Enum(ordStake),
		Stake: ActionStake{},
	}
}

type ActionAddKey struct {
	// TODO
}

// TODO
func NewAddKey() Action {
	return Action{
		Enum:   borsh.Enum(ordAddKey),
		AddKey: ActionAddKey{},
	}
}

type ActionDeleteKey struct {
	// TODO
}

// TODO
func NewDeleteKey() Action {
	return Action{
		Enum:      borsh.Enum(ordDeleteKey),
		DeleteKey: ActionDeleteKey{},
	}
}

type ActionDeleteAccount struct {
	// TODO
}

// TODO
func NewDeleteAccount() Action {
	return Action{
		Enum:          borsh.Enum(ordDeleteAccount),
		DeleteAccount: ActionDeleteAccount{},
	}
}
