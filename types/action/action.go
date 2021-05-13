package action

import (
	"math/big"

	"github.com/near/borsh-go"
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

func (a *Action) PrepaidGas() uint64 {
	switch uint8(a.Enum) {
	case OrdFunctionCall:
		return a.FunctionCall.Gas
	default:
		return 0
	}
}

func (a *Action) DepositBalance() big.Int {
	switch uint8(a.Enum) {
	case OrdFunctionCall:
		return a.FunctionCall.Deposit
	case OrdTransfer:
		return a.Transfer.Deposit
	default:
		return *big.NewInt(0)
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
func NewActionCreateAccount() Action {
	return Action{
		Enum:          borsh.Enum(OrdCreateAccount),
		CreateAccount: ActionCreateAccount{},
	}
}

type ActionDeployContract struct {
	// TODO
}

// TODO
func NewActionDeployContract() Action {
	return Action{
		Enum:           borsh.Enum(OrdDeployContract),
		DeployContract: ActionDeployContract{},
	}
}

type ActionFunctionCall struct {
	MethodName string
	Args       string // Base64 string
	Gas        uint64
	Deposit    big.Int
}

func NewActionFunctionCall(methodName string, args string, gas uint64, deposit big.Int) Action {
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
	Deposit big.Int
}

func NewActionTransfer(deposit big.Int) Action {
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
func NewActionStake() Action {
	return Action{
		Enum:  borsh.Enum(OrdStake),
		Stake: ActionStake{},
	}
}

type ActionAddKey struct {
	// TODO
}

// TODO
func NewActionAddKey() Action {
	return Action{
		Enum:   borsh.Enum(OrdAddKey),
		AddKey: ActionAddKey{},
	}
}

type ActionDeleteKey struct {
	// TODO
}

// TODO
func NewActionDeleteKey() Action {
	return Action{
		Enum:      borsh.Enum(OrdDeleteKey),
		DeleteKey: ActionDeleteKey{},
	}
}

type ActionDeleteAccount struct {
	// TODO
}

// TODO
func NewActionDeleteAccount() Action {
	return Action{
		Enum:          borsh.Enum(OrdDeleteAccount),
		DeleteAccount: ActionDeleteAccount{},
	}
}
