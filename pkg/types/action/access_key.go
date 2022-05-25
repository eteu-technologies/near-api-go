package action

import (
	"encoding/json"
	"fmt"

	"github.com/eteu-technologies/borsh-go"
	"github.com/eteu-technologies/near-api-go/pkg/types"
)

type AccessKeyPermission struct {
	Enum borsh.Enum `borsh_enum:"true"`

	FunctionCallPermission AccessKeyFunctionCallPermission
	FullAccessPermission   struct{}
}

type fullAccessPermissionWrapper struct {
	FunctionCall AccessKeyFunctionCallPermission `json:"FunctionCall"`
}

func NewFunctionCallPermission(allowance types.Balance, receiverID types.AccountID, methodNames []string) AccessKeyPermission {
	return AccessKeyPermission{
		Enum: borsh.Enum(0),
		FunctionCallPermission: AccessKeyFunctionCallPermission{
			Allowance:   &allowance,
			ReceiverID:  receiverID,
			MethodNames: methodNames,
		},
	}
}

func NewFunctionCallUnlimitedAllowancePermission(receiverID types.AccountID, methodNames []string) AccessKeyPermission {
	return AccessKeyPermission{
		Enum: borsh.Enum(0),
		FunctionCallPermission: AccessKeyFunctionCallPermission{
			Allowance:   nil,
			ReceiverID:  receiverID,
			MethodNames: methodNames,
		},
	}
}

func NewFullAccessPermission() AccessKeyPermission {
	return AccessKeyPermission{
		Enum: borsh.Enum(1),
	}
}

func (a AccessKeyPermission) MarshalJSON() (b []byte, err error) {
	if a.IsFullAccess() {
		b = []byte(`"FullAccess"`)
		return
	}

	var v fullAccessPermissionWrapper
	v.FunctionCall = a.FunctionCallPermission

	b, err = json.Marshal(&v)
	return
}

func (a *AccessKeyPermission) UnmarshalJSON(b []byte) (err error) {
	if len(b) > 0 && b[0] == '{' {
		var permission fullAccessPermissionWrapper
		if err = json.Unmarshal(b, &permission); err != nil {
			return
		}

		*a = AccessKeyPermission{
			Enum:                   borsh.Enum(0),
			FunctionCallPermission: permission.FunctionCall,
		}
		return
	}

	var value string
	if err = json.Unmarshal(b, &value); err != nil {
		return
	}

	if value == "FullAccess" {
		*a = NewFullAccessPermission()
		return
	}

	err = fmt.Errorf("unknown permission '%s'", value)
	return
}

func (a AccessKeyPermission) String() string {
	var value string = "FullAccess"
	if a.IsFunctionCall() {
		value = a.FunctionCallPermission.String()
	}
	return fmt.Sprintf("AccessKeyPermission{%s}", value)
}

func (a *AccessKeyPermission) IsFunctionCall() bool {
	return a.Enum == 0
}

func (a *AccessKeyPermission) IsFullAccess() bool {
	return a.Enum == 1
}

type AccessKeyFunctionCallPermission struct {
	Allowance   *types.Balance  `json:"allowance"`
	ReceiverID  types.AccountID `json:"receiver_id"`
	MethodNames []string        `json:"method_names"`
}

func (a AccessKeyFunctionCallPermission) String() string {
	return fmt.Sprintf("AccessKeyFunctionCallPermission{Allowance=%v, ReceiverID=%v, MethodNames=%v}", a.Allowance, a.ReceiverID, a.MethodNames)
}
