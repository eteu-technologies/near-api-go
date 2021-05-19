package action

import (
	"encoding/json"
	"fmt"

	"github.com/eteu-technologies/borsh-go"
	uint128 "github.com/eteu-technologies/golang-uint128"

	"github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/eteu-technologies/near-api-go/pkg/types/key"
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

var (
	ordMappings = map[string]uint8{
		"CreateAccount":  ordCreateAccount,
		"DeployContract": ordDeployContract,
		"FunctionCall":   ordFunctionCall,
		"Transfer":       ordTransfer,
		"Stake":          ordStake,
		"AddKey":         ordAddKey,
		"DeleteKey":      ordDeleteKey,
		"DeleteAccount":  ordDeleteAccount,
	}
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
		return a.FunctionCall.Deposit
	case ordTransfer:
		return a.Transfer.Deposit
	default:
		return types.Balance(uint128.Zero)
	}
}

func (a *Action) UnderlyingValue() interface{} {
	switch uint8(a.Enum) {
	case ordCreateAccount:
		return &a.CreateAccount
	case ordDeployContract:
		return &a.DeployContract
	case ordFunctionCall:
		return &a.FunctionCall
	case ordTransfer:
		return &a.Transfer
	case ordStake:
		return &a.Stake
	case ordAddKey:
		return &a.AddKey
	case ordDeleteKey:
		return &a.DeleteKey
	case ordDeleteAccount:
		return &a.DeleteAccount
	}

	panic("unreachable")
}

func (a Action) String() string {
	ul := a.UnderlyingValue()
	if u, ok := ul.(interface{ String() string }); ok {
		return fmt.Sprintf("Action{%s}", u.String())
	}

	return fmt.Sprintf("Action{%#v}", ul)
}

func (a *Action) UnmarshalJSON(b []byte) error {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}

	if l := len(obj); l > 1 {
		return fmt.Errorf("action object contains invalid amount of keys (expected: 1, got: %d)", l)
	}

	var firstKey string
	for k := range obj {
		firstKey = k
		break
	}

	ord := ordMappings[firstKey]
	*a = Action{Enum: borsh.Enum(ord)}
	ul := a.UnderlyingValue()

	if err := json.Unmarshal(obj[firstKey], ul); err != nil {
		return err
	}

	return nil
}

type ActionCreateAccount struct {
}

// Create an (sub)account using a transaction `receiver_id` as an ID for a new account
func NewCreateAccount() Action {
	return Action{
		Enum:          borsh.Enum(ordCreateAccount),
		CreateAccount: ActionCreateAccount{},
	}
}

type ActionDeployContract struct {
	Code []byte
}

func NewDeployContract(code []byte) Action {
	return Action{
		Enum: borsh.Enum(ordDeployContract),
		DeployContract: ActionDeployContract{
			Code: code,
		},
	}
}

type ActionFunctionCall struct {
	MethodName string
	Args       []byte
	Gas        types.Gas
	Deposit    types.Balance
}

func (f ActionFunctionCall) String() string {
	return fmt.Sprintf("FunctionCall{MethodName: %s, Args: %s, Gas: %d, Deposit: %s}", f.MethodName, f.Args, f.Gas, f.Deposit)
}

func NewFunctionCall(methodName string, args []byte, gas types.Gas, deposit types.Balance) Action {
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
	Deposit types.Balance
}

func (t ActionTransfer) String() string {
	return fmt.Sprintf("Transfer{Deposit: %s}", t.Deposit)
}

func NewTransfer(deposit types.Balance) Action {
	return Action{
		Enum: borsh.Enum(ordTransfer),
		Transfer: ActionTransfer{
			Deposit: deposit,
		},
	}
}

type ActionStake struct {
	// Amount of tokens to stake.
	Stake types.Balance
	// Validator key which will be used to sign transactions on behalf of singer_id
	PublicKey key.PublicKey
}

func NewStake(stake types.Balance, publicKey key.PublicKey) Action {
	return Action{
		Enum: borsh.Enum(ordStake),
		Stake: ActionStake{
			Stake:     stake,
			PublicKey: publicKey,
		},
	}
}

type ActionAddKey struct {
	PublicKey key.PublicKey
	AccessKey struct {
		Nonce      types.Nonce
		Permission AccessKeyPermission
	}
}

func NewAddKey(publicKey key.PublicKey, nonce types.Nonce, permission AccessKeyPermission) Action {
	return Action{
		Enum:   borsh.Enum(ordAddKey),
		AddKey: ActionAddKey{},
	}
}

type ActionDeleteKey struct {
	PublicKey key.PublicKey
}

func NewDeleteKey(publicKey key.PublicKey) Action {
	return Action{
		Enum: borsh.Enum(ordDeleteKey),
		DeleteKey: ActionDeleteKey{
			PublicKey: publicKey,
		},
	}
}

type ActionDeleteAccount struct {
	BeneficiaryID types.AccountID
}

func NewDeleteAccount(beneficiaryID types.AccountID) Action {
	return Action{
		Enum: borsh.Enum(ordDeleteAccount),
		DeleteAccount: ActionDeleteAccount{
			BeneficiaryID: beneficiaryID,
		},
	}
}
