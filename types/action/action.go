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
	OrdCreateAccount uint8 = iota
	OrdDeployContract
	OrdFunctionCall
	OrdTransfer
	OrdStake
	OrdAddKey
	OrdDeleteKey
	OrdDeleteAccount
)

func (a *Action) PrepaidGas() types.Gas {
	switch uint8(a.Enum) {
	case OrdFunctionCall:
		return a.FunctionCall.Gas
	default:
		return 0
	}
}

func (a *Action) DepositBalance() types.Balance {
	switch uint8(a.Enum) {
	case OrdFunctionCall:
		return types.Balance(a.FunctionCall.Deposit)
	case OrdTransfer:
		return types.Balance(a.Transfer.Deposit)
	default:
		return types.Balance(uint128.Zero)
	}
}

func (a *Action) UnderlyingValue() interface{} {
	switch uint8(a.Enum) {
	case OrdCreateAccount:
		return a.CreateAccount
	case OrdDeployContract:
		return a.DeployContract
	case OrdFunctionCall:
		return a.FunctionCall
	case OrdTransfer:
		return a.Transfer
	case OrdStake:
		return a.Stake
	case OrdAddKey:
		return a.AddKey
	case OrdDeleteKey:
		return a.DeleteKey
	case OrdDeleteAccount:
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
		Enum:          borsh.Enum(OrdCreateAccount),
		CreateAccount: ActionCreateAccount{},
	}
}

type ActionDeployContract struct {
	// TODO
}

// TODO
func NewDeployContract() Action {
	return Action{
		Enum:           borsh.Enum(OrdDeployContract),
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
		Enum: borsh.Enum(OrdFunctionCall),
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
		Enum: borsh.Enum(OrdTransfer),
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
		Enum:  borsh.Enum(OrdStake),
		Stake: ActionStake{},
	}
}

type ActionAddKey struct {
	// TODO
}

// TODO
func NewAddKey() Action {
	return Action{
		Enum:   borsh.Enum(OrdAddKey),
		AddKey: ActionAddKey{},
	}
}

type ActionDeleteKey struct {
	// TODO
}

// TODO
func NewDeleteKey() Action {
	return Action{
		Enum:      borsh.Enum(OrdDeleteKey),
		DeleteKey: ActionDeleteKey{},
	}
}

type ActionDeleteAccount struct {
	// TODO
}

// TODO
func NewDeleteAccount() Action {
	return Action{
		Enum:          borsh.Enum(OrdDeleteAccount),
		DeleteAccount: ActionDeleteAccount{},
	}
}
