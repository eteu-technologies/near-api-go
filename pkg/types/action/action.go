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

	simpleActions = map[string]bool{
		"CreateAccount": true,
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

func (a *Action) UnmarshalJSON(b []byte) (err error) {
	var obj map[string]json.RawMessage

	// actions can be either strings, or objects, so try deserializing into string first
	var actionType string
	if len(b) > 0 && b[0] == '"' {
		if err = json.Unmarshal(b, &actionType); err != nil {
			return
		}

		if _, ok := simpleActions[actionType]; !ok {
			err = fmt.Errorf("Action '%s' had no body", actionType)
			return
		}

		obj = map[string]json.RawMessage{
			actionType: json.RawMessage(`{}`),
		}
	} else {
		if err = json.Unmarshal(b, &obj); err != nil {
			return
		}
	}

	if l := len(obj); l > 1 {
		err = fmt.Errorf("action object contains invalid amount of keys (expected: 1, got: %d)", l)
		return
	}

	for k := range obj {
		actionType = k
		break
	}

	ord := ordMappings[actionType]
	*a = Action{Enum: borsh.Enum(ord)}
	ul := a.UnderlyingValue()

	if err = json.Unmarshal(obj[actionType], ul); err != nil {
		return
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
	Code []byte `json:"code"`
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
	MethodName string        `json:"method_name"`
	Args       []byte        `json:"args"`
	Gas        types.Gas     `json:"gas"`
	Deposit    types.Balance `json:"deposit"`
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
	Deposit types.Balance `json:"deposit"`
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
	Stake types.Balance `json:"stake"`
	// Validator key which will be used to sign transactions on behalf of singer_id
	PublicKey key.PublicKey `json:"public_key"`
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
	PublicKey key.PublicKey `json:"public_key"`
	AccessKey struct {
		Nonce      types.Nonce         `json:"nonce"`
		Permission AccessKeyPermission `json:"permission"`
	} `json:"access_key"`
}

func NewAddKey(publicKey key.PublicKey, nonce types.Nonce, permission AccessKeyPermission) Action {
	return Action{
		Enum:   borsh.Enum(ordAddKey),
		AddKey: ActionAddKey{},
	}
}

type ActionDeleteKey struct {
	PublicKey key.PublicKey `json:"public_key"`
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
	BeneficiaryID types.AccountID `json:"beneficiary_id"`
}

func NewDeleteAccount(beneficiaryID types.AccountID) Action {
	return Action{
		Enum: borsh.Enum(ordDeleteAccount),
		DeleteAccount: ActionDeleteAccount{
			BeneficiaryID: beneficiaryID,
		},
	}
}
